package modules

import (
	"encoding/json"
	"fmt"
	_ "strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/oxodao/metaprint/utils"
)

type CpuInfo struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
}

func (r CpuInfo) Print(args []string) string {
  cpuinfo := make(map[string]interface{})
	cpu, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		return "-"
	}
  
  cpuinfo["cores"] = cpu.NumCore()
  cpuinfo["cpus"] = cpu.NumCPU()
  cpuinfo["pcpus"] = cpu.NumPhysicalCPU()

  mhzavg := 0.0

  for n, proc := range cpu.Processors {
      pi, ok := json.Marshal(proc)
      if ok != nil {
        // fmt.Println(fmt.Errorf("json marshall err %s", ok))
        return ""
      }
      var x map[string]interface{}
      _ = json.Unmarshal(pi, &x)
    for _, p := range []string{"cores","mhz"} {
        val, _ := x[p].(float64)
      if (p == "mhz"){
        mhzavg = (mhzavg + val)
      }
      cpukey := fmt.Sprintf("cpu%d_%s", n, p)
      cpuinfo[cpukey] = val
    }
  }
  mhzavg = (mhzavg / float64(cpu.NumCPU()))

  cpuinfo["avgghz"] = utils.GetRoundedFloat((mhzavg / 1024.0), r.Rounding)

	return utils.ReplaceVariables(r.Format, cpuinfo)
}

func (r CpuInfo) GetPrefix() string {
	return r.Prefix
}

func (r CpuInfo) GetSuffix() string {
	return r.Suffix
}
