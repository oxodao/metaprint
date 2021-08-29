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

First, you need to create a config file. It should be located at $HOME/.config/metaprint/config.yml. To find out how module works, checkout the [config reference](#Config)

Then you can call them with the following command:

```sh
$ metaprint <module> <name> [custom args]
# Example:
$ metaprint date tmux
```

## Config

There are not much modules for now, but more will come soon.

Each module can take a prefix and a suffix:
```yml
module_name:
  your_custom_name:
    prefix: 
    suffix: Go
```

### Battery
Available variables for the format: 
  - `%percentage%`: The battery percentage
  - `%status_text%`: charging_text or discharging_text from the config
  - `%cycles%`: The amount of cycles of the battery

```yml
battery:
  your_custom_name:
    battery_name: BAT0 # Defaults to BAT0, found at /sys/class/power_supply/ 
    charging_text: 🔌
    discharging_text: 🔋
    format: "%status_text% %percentage%%"
```

### Date
The date module takes a golang date format
```yml
datetime:
  your_custom_name:
    prefix: 
    format: 01/02/06 <b>15:04:05</b>
```

### Ram
The ram module takes a format, a unit and the amount of number after the comma to display.

The format can take multiple params: `%used%`, `%free%`, `%total%`, `%percentage%`, `%percentage_free%`.

The unit can be either `Go`, `Mo`, `Ko`

```yml
ram:
  your_custom_name:
    format: "%used% / %total% %unit%"
    unit: Go
    rounding: 2
```

### IP
The IP module lets you display a local IP address.

It takes an interface array which will be scanned in the same order. As of today, it only supports IPv4. IPv6 support may be coming later (PR are welcome ;)). This will also ignore non-existing interface, so if you want to put first your Android's RNDIS it won't mess your setup when you are not connected through your smartphone.

It also takes a `no_ip` string that is displayed when no address IP can be found.

```yml
ip:
  your_custom_name:
    interface: ['enp37s0']
    no_ip: No address
```

### Music
The music module accesses the mpris data through dbus so you need a mpris compatible player to go with this (Most of them are: Firefox, Spotify, VLC, ...)

You can limit / order the players allowed the same way you'd do with the IP module, the string required is (@TODO: Find on what I sort and explain how to get it here).

You can put any mpris data as variables, here I only use the title and the artist.

```yml
music:
  tmux:
    format: "%xesam:title% - %xesam:artist%"
    players: ['Mozilla Firefox', 'Spotify'] # Limit to these players, first found is preferred. @TODO Still need to explain how to get those names, IDK
    no_player: none # This is the text displayed when there is no player running
    trim_at: 15 # each variable will be trimmed at 15 chars
    trim_all: 50 # the max response (excluding pre/suffix) will be 50 chars
```
### Custom
This one aims at adding the custom script we promised to remove. Yeah... If there is no module for what you want, you can keep the advantage of using metaprint to run a script, letting it adding prefix/suffix and still use metaprint everywhere instead of calling directly your script

```yml
custom:
  my_custom_script:
    prefix: ☁️
    command: curl wttr.in/Paris?format=%t
```

### Full config example
```yml
battery:
  i3:
    charging_text: 🔌
    discharging_text: 🔋
    format: "%status_text% %percentage%%"

datetime:
  i3:
    prefix: 
    format: 01/02/06 <b>15:04:05</b>
  tmux:
    format: 15:04:05

ram:
  i3:
    format: "%percentage%"
    unit: Go
    rounding: 2
  tmux:
    format: "%used% / %total% %unit% - %free%"
    unit: Go
    rounding: 2

ip:
  tmux:
    interface: ['enp37s0']
    no_ip: No address

music:
  tmux:
    format: "%xesam:title% - %xesam:artist%"
    no_player: none
    trim_at: 15 

custom:
  my_custom_script:
    prefix: ☁️
    command: curl wttr.in/Paris?format=%t
```

### Future plans
This project is open for contribution. If you know Go you can definetly create your own module and submit a PR !

Here are the upcoming features:
- include / multiple config file / Including files based on the machine hostname
- fixing the ram (currently used = used + shared + temp/cache or IDK something like that)
- CPU usage
- Temperature sensors module
- Percentage stats should have multiple prefix (like for battery percentage, changing the icon)
- Volume (Input / Output)
- Screen brightness

Could happen in the future:
- Weather module
