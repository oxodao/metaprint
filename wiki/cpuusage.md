# Cpu Usage module
Cpu usage displays a constantly updating cpu usage in percentage every second with the formatting provided in config.

- CPU Usage (all cpus) `%pusage%`
- Interrupts `%intr%`
- Context Switches `%ctxt%`
- Total Processes `%procs%`
- Running Processes `%running%`
- Blocked Processes `%blocked%` 

```yml
cpuusage:
  tmux:
    format: " #{p-3:#{l:%pusage%}}%#[bold,fg=colour#{?#{e|<:%pusage%,70},2,1}] "
    rounding: 0
```
