package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input hexadecimal string
	inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// Convert hexadecimal to Base64
	base64, err := HexToBase64(inputHex)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if base64 != expectedOutput {
		panic("Output does not match expected output")
	}

	// Output results
	fmt.Printf("Hex: %s\n", inputHex)
	fmt.Printf("Base64 = %s", base64)
}

// HexToBase64 converts a hexadecimal string to Base64.
func HexToBase64(input string) (string, error) {
	// Convert hexadecimal to binary
	binary, err := HexToBinary(input)
	if err != nil {
		return "", err
	}

	// Convert binary to Base64
	base64Res, err := BinaryToBase64(binary)
	if err != nil {
		return "", err
	}

	return base64Res, nil
}

// HexToBinary converts a hexadecimal string to binary.
func HexToBinary(hex string) (string, error) {
	var result string

	for _, char := range hex {
		var hexValue uint8

		// Determine the decimal value of the hexadecimal character
		switch {
		case '0' <= char && char <= '9':
			hexValue = uint8(char - '0')
		case 'a' <= char && char <= 'f':
			hexValue = uint8(char - 'a' + 10)
		case 'A' <= char && char <= 'F':
			hexValue = uint8(char - 'A' + 10)
		default:
			return "", fmt.Errorf("invalid hexadecimal digit: %c", char)
		}

		// Convert hexadecimal value to 4-bit binary representation
		binaryString := fmt.Sprintf("%04b", hexValue)
		result += binaryString
	}

	return result, nil
}

// BinaryToBase64 converts a binary string to Base64.
func BinaryToBase64(input string) (string, error) {
	const (
		groupSize   = 6
		base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	)

	// Pad the binary string to make its length a multiple of 6
	for len(input)%6 != 0 {
		input += "0"
	}

	var base64Str strings.Builder

	// Process binary string in chunks of 6 bits
	for i := 0; i < len(input); i += groupSize {
		binaryChunk := input[i : i+groupSize]

		var decimalValue int
		for _, bit := range binaryChunk {
			//decimalValue*2 shifts all bits in decimalValue left by one position
			//int(bit-'0') - converts the ASCII character '0' or '1' to its corresponding integer value (0 or 1)
			decimalValue = decimalValue*2 + int(bit-'0')
		}

		base64Str.WriteByte(base64Chars[decimalValue])
	}

	// Ensure the length of the Base64 string is a multiple of 4 by padding '=' characters if needed
	result := base64Str.String()
	for len(result)%4 != 0 {
		result += "="
	}

	return result, nil
}
