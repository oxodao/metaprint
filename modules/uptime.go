package modules

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/oxodao/metaprint/utils"
	"math"
)

type Uptime struct {
	Prefix          string
	Suffix          string
	Format          string
	NoMinutesFormat string `yaml:"no_minutes_format"`
	NoHoursFormat   string `yaml:"no_hours_format"`
	TwoDigitHours   bool   `yaml:"two_digit_hours"`
}

func pad(number float64, condition bool) string {
	padded := fmt.Sprintf("%.0f", number)
	if condition {
		if number < 10 {
			padded = "0" + padded
		}
	}
	return padded
}

func (u Uptime) Print(args []string) string {
	ut, err := linuxproc.ReadUptime("/proc/uptime")
	if err != nil {
		return "-"
	}

	hours := math.Floor(ut.Total / 60 / 60)
	minutes := math.Floor((ut.Total - (hours * 60 * 60)) / 60)
	seconds := math.Floor(ut.Total - (hours * 60 * 60) - (minutes * 60))

	format := u.Format

	if hours == 0 {
		format = u.NoHoursFormat

		if minutes == 0 {
			format = u.NoMinutesFormat
		}
	}

	return utils.ReplaceVariables(format, map[string]interface{}{
		"hours": pad(hours, u.TwoDigitHours),
		"minutes": pad(minutes, true),
		"seconds": pad(seconds, true),
	})
}

func (u Uptime) GetPrefix() string {
	return u.Prefix
}

func (u Uptime) GetSuffix() string {
	return u.Suffix
}
