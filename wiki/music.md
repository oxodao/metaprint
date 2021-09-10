# Music module
The music module accesses the mpris data through dbus, so you need a mpris compatible player to go with this (Most of them are: Firefox, Spotify, VLC, ...)

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