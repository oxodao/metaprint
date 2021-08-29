package modules

import (
	"strings"

	"github.com/oxodao/metaprint/utils"
	"github.com/pbnjay/memory"
)

type Ram struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
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

	str = strings.ReplaceAll(str, "%used%", utils.GetRoundedFloat(used, r.Rounding))
	str = strings.ReplaceAll(str, "%free%", utils.GetRoundedFloat(free, r.Rounding))
	str = strings.ReplaceAll(str, "%total%", utils.GetRoundedFloat(total, r.Rounding))
	str = strings.ReplaceAll(str, "%percentage%", utils.GetRoundedFloat(percentage, r.Rounding))
	str = strings.ReplaceAll(str, "%percentage_free%", utils.GetRoundedFloat(percentageFree, r.Rounding))

	str = strings.ReplaceAll(str, "%unit%", r.Unit)

	return str
}

func (r Ram) GetPrefix() string {
	return r.Prefix
}

func (r Ram) GetSuffix() string {
	return r.Suffix
}
