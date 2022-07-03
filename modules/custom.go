package modules

import (
	"os/exec"
	"strings"

	"github.com/oxodao/metaprint/utils"
)

type Custom struct {
	Prefix  string
	Suffix  string
	Command string
        Format string
}

func (c Custom) Print(args []string) string {
	out, err := exec.Command("bash", "-c", c.Command).Output()
	if err != nil {
		return ""
	}

	return utils.ReplaceVariables(c.Format, map[string]interface{}{
          "output": strings.Trim(string(out), " \r\n\t"),
        })
	// return strings.Trim(string(out), " \n\t")
}

func (c Custom) GetPrefix() string {
	return c.Prefix
}

func (c Custom) GetSuffix() string {
	return c.Suffix
}
