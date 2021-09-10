package pulse

import (
	"fmt"
	"github.com/lawl/pulseaudio"
)

type PulseAudio struct {
	client *pulseaudio.Client
}

func Connect() (*PulseAudio, error) {
	client, err := pulseaudio.NewClient()
	if err != nil {
		return nil, err
	}

	return &PulseAudio{client: client}, nil
}

func (pa PulseAudio) FindFirstMatchingSink(names []string) (*pulseaudio.Sink, bool, error) {
	if len(names) == 0 {
		si, err := pa.client.ServerInfo()
		if err != nil {
			return nil, false, err
		}

		names = []string{si.DefaultSink}
	}

	sinks, err := pa.client.Sinks()
	if err != nil {
		return nil, false, err
	}

	for _, name := range names {
		for _, sink := range sinks {
			if name == sink.Name {
				return &sink, true, nil
			}
		}
	}
	return nil, false, nil
}
func (pa PulseAudio) FindFirstMatchingSource(names []string) (*pulseaudio.Source, bool, error) {
	if len(names) == 0 {
		si, err := pa.client.ServerInfo()
		if err != nil {
			return nil, false, err
		}

		names = []string{si.DefaultSource}
	}

	sources, err := pa.client.Sources()
	if err != nil {
		return nil, false, err
	}

	for _, name := range names {
		for _, source := range sources {
			if name == source.Name {
				return &source, true, nil
			}
		}
	}
	return nil, false, nil
}

//#region Volume, inspired by https://gist.github.com/jasonwhite/1df6ee4b5039358701d2
func average(values []uint32) uint32 {
	var sum float32 = 0
	for _, v := range values {
		sum += float32(v)
	}

	return uint32(sum/float32(len(values)))
}

func (pa PulseAudio) FindVolumeSink(sink *pulseaudio.Sink) uint32 {
	avg := average(sink.Cvolume)
	percentage := float32(avg) / float32(sink.BaseVolume)

	return uint32(percentage*100)
}

func (pa PulseAudio) FindVolumeSource(source *pulseaudio.Source) uint32 {
	avg := average(source.Cvolume)
	percentage := float32(avg) / float32(source.BaseVolume)

	return uint32(percentage*100)
}
//#endregion

func PrintInfos() {
	pa, err := Connect()
	if err != nil {
		fmt.Println("Could not connect to PulseAudio: ", err)
		return
	}

	sinks, err := pa.client.Sinks()
	if err != err {
		fmt.Println("Could not get sinks list: ", err)
		return
	}

	fmt.Println("Sinks: ")
	for _, sink := range sinks {
		fmt.Println("\t- " + sink.Name)
	}

	sources, err := pa.client.Sources()
	if err != err {
		fmt.Println("Could not get sources list: ", err)
		return
	}

	fmt.Println("Sources: ")
	for _, source := range sources {
		fmt.Println("\t- " + source.Name)
	}
}