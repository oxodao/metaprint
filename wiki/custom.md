# Custom module
This one aims at adding the custom script we promised to remove. Yeah... If there is no module for what you want, you can keep the advantage of using metaprint to run a script, letting it adding prefix/suffix and still use metaprint everywhere instead of calling directly your script

An example with weather, and quantity of outdated packages for Archlinux based distros:

```yml
custom:
  temperature:
    prefix: ☁️
    command: curl wttr.in/Paris?format=%t
  packages:
    prefix: ⚙️
    command: sudo pacman -Syy > /dev/null && pacman -Qu | wc -l
```
