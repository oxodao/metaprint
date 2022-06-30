# Storage module
The storage module lets you display your storage infos.

You need to set the `mount_point` value to the mount point of your device.

The format can take multiple params: `%free%`, `%used%`, `%total`

The params `%pfree%`, `%pused%` are the percent of free space vs. total or the percent of used vs. total

The unit can be either `Gb`, `Mb`, `Kb`, `b`

You can also set the rounding.

```yml
storage:
  tmux:
    format: "%free% / %total% (%pused%%)Gb "
    rounding: 1
    unit: Gb # Optional, default to Gb
    mount_point: /
```
