package main

import "testing"

func TestHexToBase64(t *testing.T) {
	inputs := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			expectedOutput: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
		{
			input:          "1F4A",
			expectedOutput: "H0o=",
		},
		{
			input:          "B7E9",
			expectedOutput: "t+k=",
		},
	}

	for _, item := range inputs {
		binary, err := HexToBase64(item.input)

		if err != nil {
			t.Fatal("Error:", err)
		}

		if binary != item.expectedOutput {
			t.Errorf("Expected: %s, Got: %s", item.expectedOutput, binary)
		}
	}
}

func TestHexToBinary(t *testing.T) {
	inputs := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			expectedOutput: "010010010010011101101101001000000110101101101001011011000110110001101001011011100110011100100000011110010110111101110101011100100010000001100010011100100110000101101001011011100010000001101100011010010110101101100101001000000110000100100000011100000110111101101001011100110110111101101110011011110111010101110011001000000110110101110101011100110110100001110010011011110110111101101101",
		},
		{
			input:          "1F4A",
			expectedOutput: "0001111101001010",
		},
		{
			input:          "B7E9",
			expectedOutput: "1011011111101001",
		},
	}

	for _, item := range inputs {
		binary, err := HexToBinary(item.input)

		if err != nil {
			t.Fatal("Error:", err)
		}

		if binary != item.expectedOutput {
			t.Errorf("Expected: %s, Got: %s", item.expectedOutput, binary)
		}
	}

}

func TestBinaryToBase64(t *testing.T) {
	inputs := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "010010010010011101101101001000000110101101101001011011000110110001101001011011100110011100100000011110010110111101110101011100100010000001100010011100100110000101101001011011100010000001101100011010010110101101100101001000000110000100100000011100000110111101101001011100110110111101101110011011110111010101110011001000000110110101110101011100110110100001110010011011110110111101101101",
			expectedOutput: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
		{
			input:          "0001111101001010",
			expectedOutput: "H0o=",
		},
		{
			input:          "1101011010110110101101011011010110101101",
			expectedOutput: "1ra1ta0=",
		},
		{
			input:          "1",
			expectedOutput: "g===",
		},
		//TODO: add test with invalid binary input
	}

	for _, item := range inputs {
		binary, err := BinaryToBase64(item.input)

		if err != nil {
			t.Fatal("Error:", err)
		}

		if binary != item.expectedOutput {
			t.Errorf("Input = %s\nExpected: %s, Got: %s\n\n", item.input, item.expectedOutput, binary)
		}
	}

}

func TestHexToBase64WithUtils(t *testing.T) {
	inputs := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			expectedOutput: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
		{
			input:          "1F4A",
			expectedOutput: "H0o=",
		},
		{
			input:          "B7E9",
			expectedOutput: "t+k=",
		},
	}

	for _, item := range inputs {
		binary, err := HexToBase64(item.input)

		if err != nil {
			t.Fatal("Error:", err)
		}

		if binary != item.expectedOutput {
			t.Errorf("Expected: %s, Got: %s", item.expectedOutput, binary)
		}
	}
}
