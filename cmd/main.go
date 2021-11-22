package main

import (
	"adapter-in-go/adapters"
	"adapter-in-go/domain"
	"adapter-in-go/services"
)

func main() {

	client :=services.Client{}

	mac := domain.Mac{}
	client.InsertSquareUsbInComputer(mac)

	windowsMachine := domain.Windows{}
	windowsMachineAdapter := adapters.WindowsAdapter{
		WindowMachine: &windowsMachine,
	}

	client.InsertSquareUsbInComputer(&windowsMachineAdapter)

}