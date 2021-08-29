package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFloatFromFile(file string) (float64, error) {
	f, err := os.Open(file)
	if err != nil {
		return -1, err
	}

	var val float64
	n, err := fmt.Fscanln(f, &val)
	if n == 0 || err != nil {
		return -1, err
	}

	return val, err
}

func GetStringFromFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return strings.Trim(string(data), " \n\t"), err
}
