package app

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"sync"
	"time"
)

func Run() {
	err := rpio.Open()

	if err != nil {
		fmt.Println("Error while getting access to the GPIO memory")
		os.Exit(1)
	}

	firstPump := Pump{14, true}
	secondPump := Pump{23, false}
	thirdPump := Pump{24, false}

	// Test: Start first pump for 5 sec, and both sec and third for 10 sec

	wg := new(sync.WaitGroup)

	defer rpio.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		firstPump.StartFlow()

		time.Sleep(time.Second * 5)

		firstPump.StopFlow()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		secondPump.StartFlow()
		thirdPump.StartFlow()

		time.Sleep(time.Second * 10)

		secondPump.StopFlow()
		thirdPump.StopFlow()
	}()
	wg.Wait()
}
