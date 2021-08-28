package modules

import (
	"os/exec"
	"strings"
)

type Custom struct {
	Prefix  string
	Suffix  string
	Command string
}

func (c Custom) Print(args []string) string {
	out, err := exec.Command("bash", "-c", c.Command).Output()
	if err != nil {
		return ""
	}

	return strings.Trim(string(out), " \n\t")
}

func (c Custom) GetPrefix() string {
	return c.Prefix
}

func (c Custom) GetSuffix() string {
	return c.Suffix
}
