package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("String: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	hash := sha256.Sum256([]byte(input))
	hashStr := hex.EncodeToString(hash[:])

	fmt.Println("\nSHA256 Hash:")
	fmt.Println(hashStr)

	fmt.Print("\nPress Enter to close...")
	reader.ReadString('\n')
}