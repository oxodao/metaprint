# Metaprint

Metaprint is a tool that lets you give up on your custom scripts to display your computer infos on softwares like i3 / tmux

![tmux status line 80% metaprint](https://user-images.githubusercontent.com/1828125/176643911-3c604abf-55f9-4732-8595-319cdfe48f95.png)
[tmux config for the preview](wiki/tmux_preview.tmux)

[metaprint config for the preview](wiki/example_config.yml)

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
- [Backlight](wiki/backlight.md)
- [Battery](wiki/battery.md)
- [Date](wiki/date.md)
- [IP](wiki/ip.md)
- [Music](wiki/music.md)
- [PulseAudio](wiki/pulseaudio.md)
- [Ram](wiki/ram.md)
- [Load](wiki/loadavg.md)
- [CPU Usage](wiki/cpuusage.md)
- [CPU Info](wiki/cpuinfo.md)
- [Hackspeed](wiki/hackspeed.md)
- [Storage](wiki/storage.md)
- [Temperature](wiki/temperature.md)
- [Uptime](wiki/uptime.md)
- [Custom](wiki/custom.md)
- [Full config example](wiki/example_config.yml)

### Future plans
This project is open for contribution. If you know Go you can definitely create your own module and submit a PR !

Here are the upcoming features:
- fixing the ram (currently used = used + shared + temp/cache or IDK something like that)
- ~~CPU usage~~
- Percentage stats should have multiple prefix (like for battery percentage, changing the icon)

Could happen in the future:
- Weather module

Far future, maybe in a parallel universe:
- Windows / OSX support

## License

Metaprint
Copyright (C) 2021 Nathan <Oxodao> JANCZEWSKI

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
