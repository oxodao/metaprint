package modules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pbnjay/memory"
)

type Ram struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
}

func ramReplaceVar(input string, toReplace string, with float64, rounding int) string {
	return strings.ReplaceAll(
		input,
		toReplace,
		fmt.Sprintf("%."+strconv.Itoa(rounding)+"f", with),
	)
}

func (r Ram) Print(args []string) string {
	str := r.Format

	free := float64(memory.FreeMemory())
	total := float64(memory.TotalMemory())
	used := total - free
	percentage := used / total * 100
	percentageFree := free / total * 100

	var divisor float64 = 1

	switch strings.ToLower(r.Unit) {
	case "go":
		divisor = 1000000000
	case "mo":
		divisor = 1000000
	case "ko":
		divisor = 1000
	}

	free /= divisor
	total /= divisor
	used /= divisor

	str = ramReplaceVar(str, "%used%", used, r.Rounding)
	str = ramReplaceVar(str, "%free%", free, r.Rounding)
	str = ramReplaceVar(str, "%total%", total, r.Rounding)
	str = ramReplaceVar(str, "%percentage%", percentage, r.Rounding)
	str = ramReplaceVar(str, "%percentage_free%", percentageFree, r.Rounding)

	str = strings.ReplaceAll(str, "%unit%", r.Unit)

	return str
}

func (r Ram) GetPrefix() string {
	return r.Prefix
}

func (r Ram) GetSuffix() string {
	return r.Suffix
}
