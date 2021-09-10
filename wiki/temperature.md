# Temperature
The temperature module is a bit special as it relies on the lm_sensors software, so you need to have it installed.

You need to be a bit familiar with the JsonPath syntax.

As it relies on lm_sensors you can also use this module to get your fan speed for example.

Let's try to add my GPU temperature monitor:

```shell
$ sensors -j 
{
   "k10temp-pci-00c3":{
      "Adapter": "PCI adapter",
      "Tctl":{
         "temp1_input": 53.125
      },
      "Tdie":{
         "temp2_input": 53.125
      }
   },
   "amdgpu-pci-2600":{
      "Adapter": "PCI adapter",
      "vddgfx":{
         "in0_input": 0.875
      },
      "fan1":{
         "fan1_input": 730.000,
         "fan1_min": 0.000,
         "fan1_max": 3200.000
      },
      "edge":{
         "temp1_input": 53.000,
         "temp1_crit": 94.000,
         "temp1_crit_hyst": -273.150
      },
      "power1":{
         "power1_average": 11.018,
         "power1_cap": 150.000
      }
   }
}
```

The corresponding value I want is `$.amdgpu-pci-2600.edge.temp1_input`

Let's add it to the config:

```yaml
temperature:
  i3_gpu:
    prefix: 🔥🖥️
    json_path: $.amdgpu-pci-2600.edge.temp1_input
    suffix: °C
```