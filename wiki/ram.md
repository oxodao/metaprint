# Ram module
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