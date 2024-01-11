package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func stringConcatenation() {
	string1 := "Go"
	string2 := "is awesome"
	result := string1 + " " + string2
	fmt.Println(result)
}

func stringConcatenationWithSprintf() {
	string1 := "Go"
	string2 := "is awesome"
	result := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(result)
}

func stringLengthComparison() {
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))
}

func stringComparison(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
	} else {
		fmt.Printf("%s and %s are not equal\n", str1, str2)
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)

	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)

	fmt.Printf("\n\n")
	charsAndBytePosition("Señor")

	fmt.Printf("\n")
	stringConcatenation()
	stringConcatenationWithSprintf()

	fmt.Printf("\n")
	stringLengthComparison()

	fmt.Printf("\n")
	stringComparison("Go", "Go")
	stringComparison("hello", "world")
}
