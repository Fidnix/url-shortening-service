package utils

import (
	"fmt"
	"regexp"
)

func VerifyUrl(url string) bool {
	re, err := regexp.Compile(`https?://[\w\.-]+\.\w+`)
	if err != nil {
		fmt.Println("Error compiling regexp", err)
		return false
	}

	return re.MatchString(url)
}
