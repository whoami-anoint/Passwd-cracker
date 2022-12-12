package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read the password from the user
	fmt.Println("Enter the password to test:")
	reader := bufio.NewReader(os.Stdin)
	password, _ := reader.ReadString('\n')

	// Trim the password and convert it to lowercase
	password = string(password[:len(password)-1])
	password = strings.ToLower(password)

	// Use a brute-force attack to guess the password
	for i := 0; i < len(password); i++ {
		for j := 0; j < len(password); j++ {
			for k := 0; k < len(password); k++ {
				guess := string(password[i]) + string(password[j]) + string(password[k])
				if guess == password {
					fmt.Println("Password cracked:", password)
					return
				}
			}
		}
	}

	// If the password was not found, it is likely strong
	fmt.Println("Password not cracked. It may be strong.")
}
