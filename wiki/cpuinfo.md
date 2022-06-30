# Cpu Info module
Cpu info displays info about num cpus, cores, physical cpus, average speed, per-cpu cores/speed

- Total cpus `%cpus%`
- Total cores `%cores%`
- Total physical cpus `%pcpus%`
- Average current GHz clockspeed `%avgghz%`
- Per CPU cores and mhz `%cpuN_cores%` `%cpuN_mhz%` where `N` is the cpu num 0 indexed. 

```yml
cpuinfo:
  tmux:
    format: " #[bold,fg=colour0]%cpus% #[nobold] #[bold,fg=colour8 dim]%avgghz%㎓"
    rounding: 2
```
