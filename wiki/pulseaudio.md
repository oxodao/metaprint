# Pulseaudio module

The pulseaudio module will let you display the percentage of your sound inputs or outputs.

You need to get the name of your source (input) or sink (output). Conveniently, metaprint features a custom command that find them easily:

```shell
$ metaprint pulseaudio-infos
Sinks: 
        - alsa_output.usb-C-Media_Electronics_Inc._Arozzi_Sfera_Pro_Microphone-00.analog-stereo
        - alsa_output.pci-0000_28_00.3.analog-stereo
Sources: 
        - alsa_output.usb-C-Media_Electronics_Inc._Arozzi_Sfera_Pro_Microphone-00.analog-stereo.monitor
        - alsa_input.usb-C-Media_Electronics_Inc._Arozzi_Sfera_Pro_Microphone-00.analog-stereo
        - alsa_output.pci-0000_28_00.3.analog-stereo.monitor
```

The `names` parameter is optional, if none supplied, it will take the default one selected in pulseaudio.
The names are tested one by one. The first one working is chosen.

The `not_found_format` will be displayed when none of the given names is found (i.e. you unplugged everything)

The `muted_format` will be used instead of the `format` when the source/sink is muted.

The percentage is calculated depending on the Base volume, not on the real "100%" you can find in pavucontrol for example.

Config reference:
```yaml
pulseaudio:
  output:
    type: sink
    format: "%percentage%%"
    not_found_format: "-"
    muted_format: "🔇 %percentage%%"
    names:
        - alsa_output.usb-C-Media_Electronics_Inc._Arozzi_Sfera_Pro_Microphone-00.analog-stereo
        - alsa_output.pci-0000_28_00.3.analog-stereo
  input:
    type: source
    format: "%percentage%%"
```