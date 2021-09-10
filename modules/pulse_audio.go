package modules

import (
	"fmt"
	"github.com/lawl/pulseaudio"
	"github.com/oxodao/metaprint/pulse"
	"github.com/oxodao/metaprint/utils"
	"strings"
)

type PulseAudio struct {
	Prefix string
	Suffix string
	Names  []string
	Type   string
	Format string
	NotFoundFormat string `yaml:"not_found_format"`
	MutedFormat string `yaml:"muted_format"`
}

func (pa PulseAudio) Print(args []string) string {
	itemType := strings.ToLower(pa.Type)

	pAudio, err := pulse.Connect()
	if err != nil {
		return "Could not connect to PulseAudio"
	}

	var sink *pulseaudio.Sink
	var source *pulseaudio.Source
	var found bool

	if itemType == "sink" {
		sink, found, err = pAudio.FindFirstMatchingSink(pa.Names)
	} else if itemType == "source" {
		source, found, err = pAudio.FindFirstMatchingSource(pa.Names)
	} else {
		return "Unknown type"
	}

	if !found {
		if err != nil {
			fmt.Println("Could not fetch " + strings.ToLower(pa.Type) + ": ", err)
			return "-"
		}

		return pa.NotFoundFormat
	}

	var percentage uint32
	if itemType == "sink" {
		percentage = pAudio.FindVolumeSink(sink)
	} else {
		percentage = pAudio.FindVolumeSource(source)
	}

	format := pa.Format
	if (itemType == "sink" && sink.Muted) || (itemType == "source" && source.Muted) {
		format = pa.MutedFormat
	}

	return utils.ReplaceVariables(format, map[string]interface{}{
		"percentage": percentage,
	})
}

func (pa PulseAudio) GetPrefix() string {
	return pa.Prefix
}

func (pa PulseAudio) GetSuffix() string {
	return pa.Suffix
}
