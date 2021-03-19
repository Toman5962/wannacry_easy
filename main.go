package main

import "wannacry_easy/utils"

func main() {
	// path
	path := "test/a.txt"
	// key
	key := []byte("1234567890123456")
	// encrypt
	utils.Encrypt(path, key)
}
