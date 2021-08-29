package modules

import (
	"strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/oxodao/metaprint/utils"
)

type Ram struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
}

func (r Ram) Print(args []string) string {
	mem, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return "-"
	}

	available := float64(mem.MemAvailable)
	total := float64(mem.MemTotal)
	used := total - available

	percentage := used / total
	percentageFree := available / total

	var divisor float64 = 1

	switch strings.ToLower(r.Unit) {
	case "go":
		divisor = 1000000
	case "mo":
		divisor = 1000
	case "ko":
		divisor = 1
	}

	available /= divisor
	total /= divisor
	used /= divisor

	str := r.Format

	str = strings.ReplaceAll(str, "%used%", utils.GetRoundedFloat(used, r.Rounding))
	str = strings.ReplaceAll(str, "%free%", utils.GetRoundedFloat(available, r.Rounding))
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
