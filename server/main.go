package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	_ "modernc.org/sqlite"
)

var db *sql.DB
var jwtSecret = []byte("your-256-bit-secret-key-here-must-be-long-enough-for-security")

// TokenBlacklist stores revoked JWT tokens
var tokenBlacklist = make(map[string]time.Time)

// cleanupInterval is how often we clean up expired tokens from the blacklist
const cleanupInterval = 10 * time.Minute

// User struct for database operations
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Exclude from JSON response
}

// RegisterRequest struct
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse struct
type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
	UserID  int    `json:"user_id"`
}

// Claims struct for JWT
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// TodoRequest struct
type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Todo struct
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}

// Email validation regex
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func initDatabase() {
	var err error
	db, err = sql.Open("sqlite", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create todos table with user_id foreign key
	createTodosTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		user_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	_, err = db.Exec(createTodosTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected and tables created successfully!")
}

// ValidateEmail checks if email is valid
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// ValidatePassword checks password strength
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return fmt.Errorf("password must contain at least one number")
	}
	return nil
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "todo-server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT token and returns user ID
func ValidateJWT(tokenString string) (int, error) {
	// Check if token is blacklisted
	if _, exists := tokenBlacklist[tokenString]; exists {
		return 0, fmt.Errorf("token has been revoked")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserID, nil
}

// BlacklistToken adds a token to the blacklist with its expiration time
func BlacklistToken(tokenString string, expiresAt time.Time) {
	tokenBlacklist[tokenString] = expiresAt
}

// CleanupBlacklist removes expired tokens from the blacklist
func CleanupBlacklist() {
	for {
		time.Sleep(cleanupInterval)
		now := time.Now()
		for token, expiresAt := range tokenBlacklist {
			if now.After(expiresAt) {
				delete(tokenBlacklist, token)
			}
		}
	}
}

// AuthMiddleware is a middleware to authenticate requests
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		userID, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store user ID in context for later use
		r.Header.Set("X-User-ID", strconv.Itoa(userID))
		next(w, r)
	}
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate email
	if !ValidateEmail(req.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Validate password
	err = ValidatePassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if email already exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create user
	result, err := db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", req.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User registered successfully",
		"user_id": userID,
	})
}

// LogoutHandler handles user logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
		return
	}

	tokenString := parts[1]

	// Parse token to get expiration time
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Add token to blacklist
	BlacklistToken(tokenString, claims.ExpiresAt.Time)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Logout successful",
	})
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !ValidateEmail(req.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Get user from database
	var user User
	err = db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", req.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Verify password
	if !CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Token:   token,
		UserID:  user.ID,
	})
}

// createTodoHandler handles creating a todo
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	var newTodo TodoRequest
	err = json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO todos (title, description, user_id) VALUES (?, ?, ?)"
	result, err := db.Exec(query, newTodo.Title, newTodo.Description, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todoID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Todo saved to database successfully!",
		"id":      todoID,
	})
}

// getTodosHandler handles getting all todos for a user
func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, title, description FROM todos WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.UserID = userID
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// getTodoHandler handles getting a single todo
func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Path[len("/api/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo Todo
	query := "SELECT id, title, description FROM todos WHERE id = ? AND user_id = ?"
	err = db.QueryRow(query, id, userID).Scan(&todo.ID, &todo.Title, &todo.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todo.UserID = userID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// updateTodoHandler handles updating a todo
func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Path[len("/api/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo TodoRequest
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE todos SET title = ?, description = ? WHERE id = ? AND user_id = ?"
	result, err := db.Exec(query, updatedTodo.Title, updatedTodo.Description, id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Todo updated successfully!",
	})
}

// deleteTodoHandler handles deleting a todo
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Path[len("/api/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM todos WHERE id = ? AND user_id = ?"
	result, err := db.Exec(query, id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Todo deleted successfully!",
	})
}

func main() {
	initDatabase()
	defer db.Close()

	// Start token blacklist cleanup goroutine
	go CleanupBlacklist()

	// Auth routes (public)
	http.HandleFunc("/api/auth/register", RegisterHandler)
	http.HandleFunc("/api/auth/login", LoginHandler)
	http.HandleFunc("/api/auth/logout", AuthMiddleware(LogoutHandler))

	// Todo routes (protected)
	http.HandleFunc("/api/todos", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodosHandler(w, r)
		case http.MethodPost:
			createTodoHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/api/todos/", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodoHandler(w, r)
		case http.MethodPut:
			updateTodoHandler(w, r)
		case http.MethodDelete:
			deleteTodoHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	fmt.Println("Server is running on port 8080 with SQLite and Authentication...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
