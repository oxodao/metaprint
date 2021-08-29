package modules

import (
	"strings"

	"github.com/oxodao/metaprint/utils"
)

type Battery struct {
	Prefix          string
	Suffix          string
	BatteryName     string `yaml:"battery_name"`
	ChargingText    string `yaml:"charging_text"`
	DischargingText string `yaml:"discharging_text"`
	Rounding        int
	Format          string
}

func (b Battery) Print(args []string) string {
	if len(b.BatteryName) == 0 {
		b.BatteryName = "BAT0"
	}

	path := "/sys/class/power_supply/" + b.BatteryName

	str := b.Format

	if strings.Contains(str, "%percentage%") {
		currentCharge, err := utils.GetFloatFromFile(path + "/charge_now")
		if err != nil {
			return "-"
		}

		maxCharge, err := utils.GetFloatFromFile(path + "/charge_full")
		if err != nil {
			return "-"
		}

		percentage := 100 * (currentCharge / maxCharge)

		str = strings.ReplaceAll(str, "%percentage%", utils.GetRoundedFloat(percentage, b.Rounding))
	}

	if strings.Contains(str, "%status_text%") {
		status, err := utils.GetStringFromFile(path + "/status")
		if err != nil {
			return "-"
		}

		statusText := b.DischargingText
		if status != "Discharging" {
			statusText = b.ChargingText
		}

		str = strings.ReplaceAll(str, "%status_text%", statusText)
	}

	if strings.Contains(str, "%cycles%") {
		cycles, err := utils.GetStringFromFile(path + "/cycle_count")
		if err != nil {
			return "-"
		}

		str = strings.ReplaceAll(str, "%cycles%", cycles)
	}

	return str
}

func (b Battery) GetPrefix() string {
	return b.Prefix
}

func (b Battery) GetSuffix() string {
	return b.Suffix
}
