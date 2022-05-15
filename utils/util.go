package utils

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
	CheckIfError(err)
	return string(content)
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
