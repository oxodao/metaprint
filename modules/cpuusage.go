package modules

import (
	"encoding/json"
	"fmt"
	_ "strings"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/oxodao/metaprint/utils"
)

type CpuUsage struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
}

func (r CpuUsage) Print(args []string) string {
	cpuinfo := make(map[string]interface{})

	totalTime, idleTime := r.getAvg()
	initTotaltime, initIdletime := r.getAvg()
	count := int64(0)
	cpuUsage := float64(0.0)
	for {
		time.Sleep(1 * time.Second)
		nextTotaltime, nextIdletime := r.getAvg()

		if count > 1 {
			deltaTotalTime := nextTotaltime - totalTime
			deltaIdleTime := nextIdletime - idleTime
			cpuUsage = (1.0 - float64(deltaIdleTime)/float64(deltaTotalTime)) * 100.0
		}
    
		cpu, err := linuxproc.ReadStat("/proc/stat")
		if err != nil {
			break
		}

		cpuinfo["ctxt"] = cpu.ContextSwitches
		cpuinfo["procs"] = cpu.Processes
		cpuinfo["running"] = cpu.ProcsRunning
		cpuinfo["blocked"] = cpu.ProcsBlocked
		cpuinfo["intr"] = cpu.Interrupts
		cpuinfo["pusage"] = utils.GetRoundedFloat(cpuUsage, r.Rounding)

		fmt.Println(utils.ReplaceVariables(r.Format, cpuinfo))
		totalTime, idleTime = r.getAvg()

		count++
	}
  
	cpuinfo["initIdleTime"] = initIdletime
	cpuinfo["initTotalTime"] = initTotaltime

	return utils.ReplaceVariables(r.Format, cpuinfo)
}

func (r CpuUsage) getAvg() (float64,float64) {
	cpu, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		fmt.Println(fmt.Errorf("cpu read error %s", err))
		return 0, 0
	}
	cpustatall, ok := json.Marshal(cpu.CPUStatAll)
	if ok != nil {
		fmt.Println(fmt.Errorf("json marshall err %s", ok))
		return 0, 0
	}
	var x map[string]interface{}
	_ = json.Unmarshal(cpustatall, &x)

	idletime := float64(0)
	cputime := float64(0)
	for n, proc := range x {
    if n == "id" {
      continue
    }
		if n == "idle" {
			idletime = proc.(float64)
		}
		cputime = cputime + proc.(float64)
	}
	return cputime, idletime
}

func (r CpuUsage) GetPrefix() string {
	return r.Prefix
}

func (r CpuUsage) GetSuffix() string {
	return r.Suffix
}
