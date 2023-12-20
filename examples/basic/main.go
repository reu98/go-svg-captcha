package main

import (
	"fmt"
	"log"

	captcha "github.com/reu98/go-svg-captcha"
)

func main() {
	option := captcha.OptionText{}
	result, err := captcha.CreateByText(option)
	if err != nil {
		log.Fatalln(err)
	}

	// Text
	fmt.Printf("Text: %v", (*result).Text)

	// HTML SVG
	fmt.Printf("SVG: %v", (*result).Data)
}
