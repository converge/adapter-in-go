package domain

import "fmt"

type Mac struct {}

func (m Mac) InsertInSquarePort() {
	fmt.Println("insert square port into mac machine")
}