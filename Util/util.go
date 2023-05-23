package util

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
func ChannelFunc() {
	cnp := make(chan func(), 10) // here we create a channel

	for i := 0; i < 4; i++ { // using loop we create 4 goroutines using anonymous functions

		go func() {

			for f := range cnp { //Each goroutine runs in the background and continuously listens for functions on the cnp channel.

				f()

			}

		}()

	}

	cnp <- func() { //we add a function to the cnp channel using the expression cnp <- func() { fmt.Println("HERE1") }. This function simply prints "HERE1" when executed.

		fmt.Println("HERE1")

	}

	fmt.Println("Hello") //After adding the function to the channel, the code continues to execute. It prints "Hello" to the console.
}
