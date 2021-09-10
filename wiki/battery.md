# Battery module
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