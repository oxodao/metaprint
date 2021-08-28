package modules

import (
	"regexp"
	"strings"

	"github.com/godbus/dbus"
	"github.com/oxodao/metaprint/mpris"
)

var re = regexp.MustCompile("\\%(([a-zA-Z0-9]|:)*)\\%")

type Music struct {
	Prefix   string
	Suffix   string
	Format   string
	NoPlayer string `yaml:"no_player"`
	TrimAt   int    `yaml:"trim_at"`
	TrimAll  int    `yaml:"trim_all"`
	Players  []string
}

type mprisValue = map[string]dbus.Variant

func (m Music) Print(args []string) string {
	dbus, err := dbus.SessionBus()
	if err != nil {
		return ""
	}

	mprisPlayers, err := mpris.List(dbus)
	if err != nil {
		return "Could not list players"
	}

	if len(mprisPlayers) == 0 {
		return m.NoPlayer
	}

	var firstPlayer *mprisValue = nil
	players := map[string]mprisValue{}

	for _, p := range mprisPlayers {
		player := mpris.New(dbus, p)
		metadata := player.GetMetadata()

		players[player.GetIdentity()] = metadata

		if firstPlayer == nil {
			firstPlayer = &metadata
		}
	}

	var toProcessData *mprisValue = nil
	if len(m.Players) > 0 {
		for _, name := range m.Players {
			if data, ok := players[name]; ok {
				toProcessData = &data
				break
			}
		}
	} else {
		toProcessData = firstPlayer
	}

	if toProcessData == nil {
		return "An error occured"
	}

	str := m.Format

	variables := re.FindAllStringSubmatch(str, -1)
	for _, variable := range variables {
		currentVar := variable[0]
		mprisVariableName := currentVar[1 : len(currentVar)-1]

		if val, ok := (*toProcessData)[mprisVariableName]; ok {
			currentVal := ""

			if casted, ok := val.Value().(string); ok {
				currentVal = casted
			} else if casted, ok := val.Value().([]string); ok {
				currentVal = strings.Join(casted, ", ")
			}

			if m.TrimAt > 0 && len(currentVal) > m.TrimAt {
				currentVal = currentVal[:m.TrimAt]
			}

			str = strings.ReplaceAll(str, currentVar, currentVal)
		}
	}

	if m.TrimAll > 0 && len(str) > m.TrimAll {
		str = str[:m.TrimAll]
	}

	return str
}

func (m Music) GetPrefix() string {
	return m.Prefix
}

func (m Music) GetSuffix() string {
	return m.Suffix
}
