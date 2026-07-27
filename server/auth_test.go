package main

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		want    bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with numbers", "user123@domain.org", true},
		{"valid email with special chars", "user.name+tag@domain.co.uk", true},
		{"invalid email no @", "testexample.com", false},
		{"invalid email empty domain", "test@", false},
		{"invalid email no dot", "test@domain", false},
		{"empty email", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.email); got != tt.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name    string
		password string
		wantErr bool
	}{
		{"valid password", "Password1", false},
		{"too short", "Pass1", true},
		{"no uppercase", "password1", true},
		{"no lowercase", "PASSWORD1", true},
		{"no number", "Password", true},
		{"valid complex", "MyP@ssw0rd!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword(%q) error = %v, wantErr %v", tt.password, err, tt.wantErr)
				return
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	password := "Password123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	if len(hash) == 0 {
		t.Error("HashPassword returned empty string")
	}

	// Verify hash is not the same as password
	if hash == password {
		t.Error("HashPassword returned plain text password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "Password123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{"correct password", password, hash, true},
		{"wrong password", "WrongPassword", hash, false},
		{"empty password", "", hash, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.password, tt.hash); got != tt.want {
				t.Errorf("CheckPasswordHash(%q, %q) = %v, want %v", tt.password, tt.hash, got, tt.want)
			}
		})
	}
}

func TestGenerateAndValidateJWT(t *testing.T) {
	userID := 123

	// Generate token
	token, err := GenerateJWT(userID)
	if err != nil {
		t.Fatalf("GenerateJWT failed: %v", err)
	}

	if len(token) == 0 {
		t.Error("GenerateJWT returned empty token")
	}

	// Validate token
	validatedUserID, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	if validatedUserID != userID {
		t.Errorf("ValidateJWT returned userID %d, want %d", validatedUserID, userID)
	}
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		wantErr     bool
	}{
		{"empty token", "", true},
		{"invalid token", "invalid.token.string", true},
		{"wrong signature", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.invalid-signature-that-wont-verify", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateJWT(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT(%q) error = %v, wantErr %v", tt.token, err, tt.wantErr)
			}
		})
	}
}