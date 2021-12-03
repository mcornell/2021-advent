package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	return strings.Split(str, "\n")

}
