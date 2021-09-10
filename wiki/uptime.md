# Uptime module
This lets you display your uptime.
Available variables for the format:
- `%hours%`
- `%minutes%`
- `%seconds%`

```yml
uptime:
  your_custom_name:
    two_digit_hours: true # Hours < 10 should have a leading zero
    format: "%hours%h%minutes%"
    no_hours_format: "%minutes% minutes"
    no_minutes_format: "%seconds% seconds"
```