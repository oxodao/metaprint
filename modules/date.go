package modules

import (
	"time"
)

type Date struct {
	Prefix string
	Suffix string
	Format string
}

func (d Date) Print(args []string) string {
	return time.Now().Format(d.Format)
}

func (d Date) GetPrefix() string {
	return d.Prefix
}

func (d Date) GetSuffix() string {
	return d.Suffix
}
