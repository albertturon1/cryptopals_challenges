package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main2() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	res, err := HexToBase64WithUtils(input)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hex: %s\n", input)
	fmt.Printf("Base64 = %s", res)
}

func HexToBase64WithUtils(input string) (string, error) {
	// decode the hexadecimal string into a byte slice ([]byte).
	hexBytes, err := hex.DecodeString(input)

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	result := make([]byte, base64.StdEncoding.EncodedLen(len(hexBytes)))
	base64.StdEncoding.Encode(result, hexBytes)

	return string(result), nil
}
