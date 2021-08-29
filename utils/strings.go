package utils

import (
	"fmt"
	"strconv"
)

func GetRoundedFloat(value float64, rounding int) string {
	return fmt.Sprintf("%."+strconv.Itoa(rounding)+"f", value)
}
