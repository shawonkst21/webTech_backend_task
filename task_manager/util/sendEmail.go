package util

import "fmt"

func SendResetEmail(email, token string) {
	link := fmt.Sprintf("http://yourfrontend.com/reset-password?token=%s", token)
	fmt.Printf("Send email to %s with link: %s\n", email, link)
}
