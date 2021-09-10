package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GetRoundedFloat(value float64, rounding int) string {
	return fmt.Sprintf("%."+strconv.Itoa(rounding)+"f", value)
}

const (
	B  = float64(1)
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func GetInUnit(val float64, unit string) float64 {
	switch strings.ToLower(unit) {
	case "kb":
		return val / KB
	case "mb":
		return val / MB
	case "gb":
		return val / GB
	}

	return val
}

func ReplaceVariables(str string, variables map[string]interface{}) string {
	for k, v := range variables {
		str = strings.ReplaceAll(str, "%" + k + "%", fmt.Sprintf("%v", v))
	}

	return str
}