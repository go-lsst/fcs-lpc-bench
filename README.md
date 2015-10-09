fcs-lpc-bench
=============

A little application using the `github.com/go-lsst/ncs` control system.

## Installation

```sh
sh> go get github.com/go-lsst/fcs-lpc-bench
sh> fcs-lpc-bench -mock
canbus INFO    >>> boot...
canbus INFO    detected node 0x41
canbus INFO    detected node 0x42
canbus INFO    node=canbus.Node{id:65, device:262545, vendor:23, product:587333634, revision:65538, serial:"c7c80499"}
canbus INFO    node=canbus.Node{id:66, device:524689, vendor:23, product:587464706, revision:1, serial:"c7c60327"}
canbus INFO    adc=&canbus.ADC{Base:(*ncs.Base)(0xc82000a220), node:65, serial:"c7c80499", tx:1, bus:(*canbus.busImpl)(0xc82009c000)}
canbus INFO    dac=&canbus.DAC{Base:(*ncs.Base)(0xc82000a260), node:66, serial:"c7c60327", bus:(*canbus.busImpl)(0xc82009c000)}
canbus INFO    >>> boot... [done]
canbus INFO    handle...
sysbus INFO    recv: name=hpt; hygro=46.875%; press=756.25mbar; temp=42.5C
sysbus INFO    recv: name=hpt; hygro=46.875%; press=756.25mbar; temp=42.5C
sysbus INFO    recv: name=hpt; hygro=46.875%; press=756.25mbar; temp=42.5C
sysbus INFO    recv: name=hpt; hygro=46.875%; press=756.25mbar; temp=42.5C
sysbus INFO    recv: name=hpt; hygro=46.875%; press=756.25mbar; temp=42.5C
^C
lpc INFO    stopping app...
canbus INFO    stopping...
canbus INFO    shutdown...
canbus INFO    received 'quit' request...
canbus INFO    handle... [done]
canbus INFO    closing tcp server
sysbus ERROR    received EOF: connection closed
```
