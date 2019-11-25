package app

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

func Run() {
	err := rpio.Open()

	if err != nil {
		fmt.Println("Error while getting access to the GPIO memory")
	}
}
