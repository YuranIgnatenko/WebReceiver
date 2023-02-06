package main

import (
	"WebReceiver/receiver"
)

func main() {
	conf := receiver.NewConfig("receiver/main.json")
	logger := receiver.NewLogger(conf.PathLog)
	logger.INFO("Start Receiver")
	receiver.NewReceicerBox(conf, logger).RunReceiver()
}
