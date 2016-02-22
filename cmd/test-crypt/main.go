package main

import (
	"fmt"

	"github.com/stevenwilkin/coconut/crypt" 
)

func main() {
	plaintext := []byte("I am he as you are he as you are me and we are all together")
	fmt.Printf("plaintext: %s\n", plaintext)

	ciphertext := crypt.Encrypt(plaintext)
	fmt.Printf("ciphertext: %q\n", ciphertext)

	decrypted := crypt.Decrypt(ciphertext)
	fmt.Printf("decrypted: %s\n", decrypted)
}
