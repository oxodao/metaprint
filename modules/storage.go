package modules

import (
	"github.com/oxodao/metaprint/utils"
	"strings"
	"syscall"
)

// Inspired/stolen from https://topic.alibabacloud.com/a/golang-get-hard-disk-partition-remaining-space-size_1_38_30919404.html

type Storage struct {
	Prefix     string
	Suffix     string
	Format     string
	MountPoint string `yaml:"mount_point"`
	Rounding   int
	Unit       string
}

func (s Storage) Print(args []string) string {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(s.MountPoint, &fs)
	if err != nil {
		return "-"
	}

	if len(s.Unit) == 0 {
		s.Unit = "Gb"
	}

	size := float64(fs.Blocks * uint64(fs.Bsize))
	avail := float64(fs.Bavail * uint64(fs.Bsize))
	used := size - avail

	size = utils.GetInUnit(size, s.Unit)
	avail = utils.GetInUnit(avail, s.Unit)
	used = utils.GetInUnit(used, s.Unit)

	sizeStr := utils.GetRoundedFloat(size, s.Rounding)
	availStr := utils.GetRoundedFloat(avail, s.Rounding)
	usedStr := utils.GetRoundedFloat(used, s.Rounding)

	str := s.Format

	str = strings.ReplaceAll(str, "%free%", availStr)
	str = strings.ReplaceAll(str, "%used%", usedStr)
	str = strings.ReplaceAll(str, "%total%", sizeStr)

	return str
}

func (s Storage) GetPrefix() string {
	return s.Prefix
}

func (s Storage) GetSuffix() string {
	return s.Suffix
}
