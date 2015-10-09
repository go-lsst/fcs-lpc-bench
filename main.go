package main

import (
	"flag"

	"github.com/go-lsst/ncs"
	"github.com/go-lsst/ncs/drivers/canbus"
	"github.com/go-lsst/ncs/drivers/hd2001"
)

var (
	mock = flag.Bool("mock", false, "switch to use mock-canbus")
	port = flag.Int("port", 50000, "port for c-wrapper")
)

func main() {

	flag.Parse()

	newBus := canbus.New
	if *mock {
		newBus = canbus.NewMock
	}

	app, err := ncs.New(
		"lpc",
		newBus(
			"canbus", *port,
			canbus.NewADC("ai814", "c7c80499", 0x1),
			canbus.NewDAC("ao412", "c7c60327"),
		),
		canbus.NewLED("led", "canbus"),
		hd2001.New("hpt", "canbus"),
	)
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
