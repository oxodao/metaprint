# Load average module
Load averages 1, 5, 15 and number of procs

- Load Averages `%avg1min%`, `%avg5min%`, `%avg15min%`
- Total Processes `%procs%`
- Running Processes `%running%`

```yml
loadavg:
  tmux:
    format: "龍#{p-5:#{l:%avg1min%}} (ﰌ #{p-3:#{l:%running%}})[	#{p-3:#{l:%procs%}}]"
    rounding: 2
```
