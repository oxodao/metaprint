# IP module
The IP module lets you display a local IP address.

It takes an interface array which will be scanned in the same order.

As of today, it only supports IPv4. IPv6 support may be coming later (PR are welcome ;)).

This will also ignore non-existing interface, so if you want to put first your Android's RNDIS it won't mess your setup when you are not connected through your smartphone.

It also takes a `no_ip` string that is displayed when no address IP can be found.

```yml
ip:
  your_custom_name:
    interface: ['enp37s0']
    no_ip: No address
```
