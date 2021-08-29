package modules

import (
	"fmt"
	"math"
	"strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Uptime struct {
	Prefix          string
	Suffix          string
	Format          string
	NoMinutesFormat string `yaml:"no_minutes_format"`
	NoHoursFormat   string `yaml:"no_hours_format"`
	TwoDigitHours   bool   `yaml:"two_digit_hours"`
}

func (u Uptime) Print(args []string) string {
	ut, err := linuxproc.ReadUptime("/proc/uptime")
	if err != nil {
		return "-"
	}

	hours := math.Floor(ut.Total / 60 / 60)
	minutes := math.Floor((ut.Total - (hours * 60 * 60)) / 60)
	seconds := math.Floor(ut.Total - (hours * 60 * 60) - (minutes * 60))

	str := u.Format

	if hours == 0 {
		str = u.NoHoursFormat

		if minutes == 0 {
			str = u.NoMinutesFormat
		}
	}

	tmpHours := fmt.Sprintf("%.0f", hours)
	if u.TwoDigitHours {
		if hours < 10 {
			tmpHours = "0" + tmpHours
		}
	}

	str = strings.ReplaceAll(
		str,
		"%hours%",
		tmpHours,
	)

	tmpMinutes := fmt.Sprintf("%.0f", minutes)
	if minutes < 10 {
		tmpMinutes = "0" + tmpMinutes
	}

	str = strings.ReplaceAll(
		str,
		"%minutes%",
		tmpMinutes,
	)

	tmpSeconds := fmt.Sprintf("%.0f", seconds)
	if seconds < 10 {
		tmpSeconds = "0" + tmpSeconds
	}

	str = strings.ReplaceAll(
		str,
		"%seconds%",
		tmpSeconds,
	)

	return str
}

func (u Uptime) GetPrefix() string {
	return u.Prefix
}

func (u Uptime) GetSuffix() string {
	return u.Suffix
}
