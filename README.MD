# Metaprint

Metaprint is a tool that lets you give up on your custom scripts to display your computer infos on softwares like i3 / tmux

## Installing
Just download the latest release in the Github release tab and move it to /usr/bin/metaprint

## Building
Building this requires Golang installed (I'm using 1.16 but I think anything > 1.13 should work, not tried it though).

```sh
$ make install
```

## Usage
First, you need to create a config file. It should be located at `$HOME/.config/metaprint/config.yml`.

You can also use the config located at `$HOME/.config/metaprint/$HOSTNAME.yml` if you have multiple computers and you want different setup on each.

To find out how module works, checkout the [config reference](#Config)

Then you can call them with the following command:

```sh
$ metaprint <module> <name> [custom args]
# Example:
$ metaprint datetime tmux
```

## Config
Each module can take a prefix and a suffix:
```yml
module_name:
  your_custom_name:
    prefix: 
    suffix: Go
```

### Module list & usage
- [Battery](wiki/battery.md)
- [Date](wiki/date.md)
- [IP](wiki/ip.md)
- [Music](wiki/music.md)
- [PulseAudio](wiki/pulseaudio.md)
- [Ram](wiki/ram.md)
- [Storage](wiki/storage.md)
- [Temperature](wiki/temperature.md)
- [Uptime](wiki/uptime.md)
- [Custom](wiki/custom.md)


- [Full config example](wiki/example_config.yml)

### Future plans
This project is open for contribution. If you know Go you can definitely create your own module and submit a PR !

Here are the upcoming features:
- fixing the ram (currently used = used + shared + temp/cache or IDK something like that)
- CPU usage
- Percentage stats should have multiple prefix (like for battery percentage, changing the icon)
- Screen brightness
- pacman available updates

Could happen in the future:
- Weather module

Far future, maybe in a parallel universe:
- Windows / OSX support
