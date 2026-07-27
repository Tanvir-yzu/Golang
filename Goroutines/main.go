// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printMessage(text string) {
// 	for i := 1; i <= 11; i++ {
// 		fmt.Println(text, "-", i)	
// 		time.Sleep(500 * time.Millisecond) // ৫০০ মিলিসেকেন্ড অপেক্ষা করা
// 	}
// }

// func main() {
// 	// ১. সাধারণ ফাংশন কল (Sequential)
// 	// printMessage("Hello")

// 	// ২. গোরুটিন দিয়ে কল করা (শুধু সামনে 'go' কি-ওয়ার্ড বসিয়ে দিতে হবে!)
// 	go printMessage("Goroutine Task")

// 	// মেইন ফাংশন নিজের কাজ করছে
// 	fmt.Println("Main function is running...")

// 	// ৩. একটু অপেক্ষা করছি যাতে গোরুটিন তার কাজ শেষ করার সুযোগ পায়
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Main function finished. Program exits.")
// }

package main

import "fmt"

// একটি ফাংশন যা স্লাইসের যোগফল বের করে চ্যানেলে পাঠিয়ে দেবে
func calculateSum(numbers []int, ch chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	if len(numbers) == 0 {
		ch <- 0
		return
	}
	
	// 'ch <- sum' মানে হলো চ্যানেলের ভেতরে sum ভ্যালু পাঠিয়ে দেওয়া (Send)
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// ১. একটি ইন্টিজার টাইপের চ্যানেল তৈরি করলাম (make keyword দিয়ে)
	ch := make(chan int)

	// ২. প্রথম ৫টি সংখ্যার যোগ করার জন্য গোরুটিন দিলাম
	go calculateSum(nums[:5], ch) // {1, 2, 3, 4, 5}
		// ৪. বাকি ৫টি সংখ্যার জন্য আরেকটি গোরুটিন চালাই
	go calculateSum(nums[5:], ch) // {6, 7, 8, 9, 10}

	// ৩. চ্যানেল থেকে ফলাফল রিসিভ করছি (<-ch)
	// এখানে প্রোগ্রাম থেমে থাকবে যতক্ষণ না গোরুটিন চ্যানেলে ডেটা পাঠায়!
	result1 := <-ch
	result2 := <-ch

	fmt.Println("First Result received from channel:", result1)
	fmt.Println("Second Result received from channel:", result2)


	//এবার মডিফাই করে দেখো—যদি একসাথে দুটি গোরুটিন চালানো হয় (দুটি আলাদা চ্যানেল বা একই চ্যানেলে পরপর দুটি <-ch দিয়ে), তবে ফলাফল কীভাবে আরেকটি গোরুটিন চালাই হবে
	fmt.Println("Total sum:", result1+result2)

}