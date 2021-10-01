package modules

import (
	bh "github.com/oxodao/brightnesshandler/brightnesshandler"
	"github.com/oxodao/metaprint/utils"
)

type Backlight struct {
	Prefix string
	Suffix string
	Device string
	Format string
}

func (b Backlight) Print(args []string) string {
	brightness, err := bh.New(&bh.Config{Device: b.Device})
	if err != nil {
		return err.Error()
	}

	if len(b.Format) == 0 {
		b.Format = "%percentage%%"
	}

	return utils.ReplaceVariables(b.Format, map[string]interface{}{
		"value":      brightness.CurrentValue,
		"max_value":  brightness.MaxValue,
		"percentage": brightness.GetCurrent(),
	})
}

func (b Backlight) GetPrefix() string {
	return b.Prefix
}

func (b Backlight) GetSuffix() string {
	return b.Suffix
}
