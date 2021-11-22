package services

import (
	"adapter-in-go/ports"
	"fmt"
)

type Client struct {}

func (c *Client) InsertSquareUsbInComputer(com ports.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertInSquarePort()
}