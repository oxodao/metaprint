package modules

import (
	_ "strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/oxodao/metaprint/utils"
)

type LoadAvg struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
}

func (r LoadAvg) Print(args []string) string {
	load, err := linuxproc.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return "-"
	}

	return utils.ReplaceVariables(r.Format, map[string]interface{}{
		"avg1min":  utils.GetRoundedFloat(load.Last1Min, r.Rounding),
		"avg5min":  utils.GetRoundedFloat(load.Last5Min, r.Rounding),
		"avg15min": utils.GetRoundedFloat(load.Last15Min, r.Rounding),
		"running":  load.ProcessRunning,
		"procs":    load.ProcessTotal,
		"unit":     r.Unit,
	})
}

func (r LoadAvg) GetPrefix() string {
	return r.Prefix
}

func (r LoadAvg) GetSuffix() string {
	return r.Suffix
}
