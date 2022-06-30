# Hackspeed
This module monitors the /dev/input events for your system in an attempt to track your 

- keys per second `%keys%`
- shorcuts per second `%shorties%`
- words per minute `%wpm%`
- spaces per minute (fake wpm) `%fakewpm%`

```yml
hackspeed:
  tmux:
    format: " %keys%/kps וּ %shorties%/s  %wpm%/wpm"
    unit: ps
    rounding: 1
```
