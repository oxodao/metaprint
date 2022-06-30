package modules

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-cmd/cmd"
	// "github.com/gookit/goutil/dump"

	// "github.com/gookit/goutil/dump"
	"github.com/oxodao/metaprint/utils"
)

type HackSpeed struct {
	Prefix   string
	Suffix   string
	Format   string
	Rounding int
	Unit     string
}

func (r HackSpeed) Print(args []string) string {
	cmds := r.streamKeys()
	if len(r.Unit) == 0 {
		r.Unit = "pm"
	}

	for {
		select {
		case is := <-cmds[0].outChan:
			fmt.Println(r.PrintFormatted(is.calculateStats()))
		case <-cmds[0].exitChan:
			fmt.Println("Key input failed")
			os.Exit(2)
		case <-time.After(1 * time.Second):
			continue
		}
	}
}

func (r HackSpeed) PrintFormatted(spaces, keys, shorties, fakewpm, wpm float64) string {
	return utils.ReplaceVariables(r.Format, map[string]interface{}{
		"keys":     utils.GetRoundedFloat(keys, r.Rounding),
		"spaces":   utils.GetRoundedFloat(spaces, r.Rounding),
		"shorties": utils.GetRoundedFloat(shorties, r.Rounding),
		"fakewpm":      utils.GetRoundedFloat(fakewpm, r.Rounding),
		"wpm":      utils.GetRoundedFloat(wpm, r.Rounding),
	})
}

func (r HackSpeed) GetPrefix() string {
	return r.Prefix
}

func (r HackSpeed) GetSuffix() string {

	return r.Suffix
}

func PrintMarked(keyopts *map[string]interface{}) {
	colors := []*color.Color{
		color.New(color.FgHiGreen),
		color.New(color.FgHiBlue),
		color.New(color.FgMagenta),
		color.New(color.FgHiCyan),
		color.New(color.FgHiRed),
		color.New(color.FgRed),
	}

	keys := []string{}
	colors[0].Add(color.Bold).Add(color.Underline)
	colors[2].Add(color.Bold).Add(color.Underline)
	// count := 0
	for tag, keymatch := range *keyopts {
		if str, ok := keymatch.(string); ok {
			tag = colors[0].Sprint(strings.TrimSpace(tag))
			str = colors[2].Sprint(strconv.Quote(strings.TrimSpace(str)))
			t := fmt.Sprintf("%s = %s", tag, str)
			keys = append(keys, t)
		}
		if len(keys)%6 == 0 {
			keys = append(keys, "\n")
		}
	}

	// status := strings.Join(keys, ",")
	// fmt.Printf("%s\n", status)
}

var keyEventKeys = []string{"keysyms", "mods", "unicode", "leds", "level"}
func (is *InputStats) getKeys(s string) (*map[string]interface{}, error) {

	regex := regexp.MustCompile(`(?Ui)(?i:([a-z]+) \[([^\]]+)\][ ]?)`)
	match := regex.FindAllStringSubmatch(s, -1)

	out := make(map[string]interface{}, len(match))
	if is.lastKey["keysyms"] != nil && strings.ToLower(is.lastKey["keysyms"].(string)) == "j" && is.lastKey["mods"] == nil {
		out = is.lastKey
	}
	for i, keymatch := range match {
		_ = i
		if len(keymatch) == 3 {
			out[keymatch[1]] = strings.TrimSpace(keymatch[2])
		}
	}
	return &out, nil
}

type streamCmd struct {
	exe      string
	args     []string
	exitChan chan struct{}
	proc     *cmd.Cmd

	outChan    chan InputStats
	started    time.Time
	lastOutput time.Time
}

func NewCmd(exe string, args []string) (*streamCmd, error) {
	cmdOptions := cmd.Options{
		Buffered:       false,
		LineBufferSize: 0,
		Streaming:      true,
	}

	exitChan := make(chan struct{})
	envCmd := cmd.NewCmdOptions(cmdOptions, exe, args...)
	proc := streamCmd{
		exe:      exe,
		args:     args,
		exitChan: exitChan,
		proc:     envCmd,
	}
	proc.outChan = make(chan InputStats)
	return &proc, nil
}

type InputStats struct {
	windowStart     time.Time
	windowEnd       time.Time
	lastElapsed     time.Duration
	keysPressed     int
	lastAvgs        []float64
	currentAvg      float64
	smoothCount     int
	lastKey         map[string]interface{}
	shortcutPressed int
	spacePressed    int
  keySpaceRatio   float64
}

type InputStat struct {
	Type  string
	Key   string
	Value float64
}

func (is *InputStats) calculateStats() (float64, float64, float64, float64, float64) {
	spaces := is.lastAvgs[1]
	keys := is.lastAvgs[0]
	shorties := is.lastAvgs[2]

  wpm := is.lastAvgs[3] * 60.0
	fakewpm := (spaces) * 60.0

 //  dump.Print(is.lastAvgs)
	// fmt.Printf("\r\nkeys: %0.3f spaces: %0.3f,shorties: %0.3f fwpm: %0.3f wpm: %0.3f\r\n", keys, spaces, shorties, fakewpm, wpm )
	return spaces, keys, shorties, fakewpm, wpm
}

func (is *InputStats) Tick() {
	is.windowEnd = time.Now()
	is.lastElapsed = time.Since(is.windowStart)

  elapse := float64(is.lastElapsed.Seconds())
  spaceRatio := float64(is.spacePressed) / float64(is.keysPressed)
  smoothavg := (is.lastAvgs[3] * (float64(is.smoothCount - 1)))
  accwps := (float64(is.keysPressed) / (float64(is.spacePressed) * (1.0-spaceRatio))) / float64(elapse) 
  smoothavg = (smoothavg + accwps) / float64(is.smoothCount)
  
  if (math.IsNaN(smoothavg) || math.IsInf(smoothavg, 1)){
    smoothavg = 0.00
  }
  // dump.Print(smoothavg, spaceRatio, accwps)
  
  is.lastAvgs[3] = smoothavg

	is.avg(is.lastAvgs[0:1], &is.keysPressed)
	is.avg(is.lastAvgs[1:2], &is.spacePressed)
	is.avg(is.lastAvgs[2:3], &is.shortcutPressed)
  // is.avg(is.lastAvgs[3:4], is.keysPressed / is.spacePressed)

	is.windowStart = time.Now()
}

func (is *InputStats) avg(lAvg []float64, keys *int) {
	tmpAvg := float64(*keys) / float64(is.lastElapsed.Seconds())

	smoothavg := (lAvg[0] * (float64(is.smoothCount - 1.0)))
	smoothavg = (smoothavg + tmpAvg) / float64(is.smoothCount)
	is.currentAvg = smoothavg
	// dump.Print(smoothavg, is.currentAvg, lAvg)
	lAvg[0] = smoothavg
	*keys = 0
}
func (is *InputStats) keyed(keyed map[string]interface{}) {
  keycheck := []string{"keysyms", "mods"}
  somenil := false
  for _, k := range keycheck {
    if keyed[k] == nil {
      somenil = true
    }
  }
	
  if somenil && keyed["keysyms"] != nil && strings.ToLower(keyed["keysyms"].(string)) == "j" {
	is.lastKey = keyed
		return
	}
  
	if len(strings.TrimSpace(keyed["mods"].(string))) > 0 {
		is.shortcutPressed += 1
	} else if spstr, ok := keyed["keysyms"].(string); ok {
		if strings.Contains(spstr, "space") {
			is.spacePressed += 1
		} else {
			is.keysPressed += 1

		}
	}

	is.lastKey = keyed
}
func (envCmd *streamCmd) startKeyCmd() error {
	smoothCount := 4
	is := InputStats{
		windowStart:     time.Now(),
		windowEnd:       time.Now(),
		lastElapsed:     time.Duration(0.0),
		keysPressed:     0,
		spacePressed:    0,
		shortcutPressed: 0,
		lastAvgs:        make([]float64, smoothCount),
		smoothCount:     smoothCount,
	}
	go func() {
		defer close(envCmd.exitChan)

		go func() {
			for {
				select {
				case <-time.After(time.Duration(1) * time.Second):
					if time.Now().Second()%1 == 0 {
						envCmd.outChan <- is
					}
					if time.Now().Second()%5 == 0 {
						is.Tick()
					}
					// fmt.Println("KeyTick!")
				}

			}
		}()
		for envCmd.proc.Stdout != nil || envCmd.proc.Stderr != nil {
			select {
			case line, open := <-envCmd.proc.Stdout:
				if !open {
					envCmd.proc.Stdout = nil
					continue
				}
				key, ok := is.getKeys(line)
				if ok != nil {
					fmt.Println(ok)
					return
				}
				if key != nil {

				}
				is.keyed(*key)

			case line, open := <-envCmd.proc.Stderr:
				if !open {
					envCmd.proc.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	<-envCmd.proc.Start()

	return nil
}

func (r *HackSpeed) streamKeys() []*streamCmd {
	cmds := make([]*streamCmd, 2)

	envCmd, ok := NewCmd("/home/chris/.local/bin/keyev", []string{})
	if ok != nil {
		fmt.Println(ok)
		return nil
	}
	cmds[0] = envCmd
	go func() {
		ok = envCmd.startKeyCmd()
		if ok != nil {
			fmt.Println(ok)
			return
		}
	}()

	return cmds

}
