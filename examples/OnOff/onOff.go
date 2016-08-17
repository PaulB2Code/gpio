package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/PaulB2Code/gpio"
	"github.com/PaulB2Code/gpio/rpi"
)

func main() {
	// set GPIO25 to output mode
	pin33, err := gpio.OpenPin(rpi.GPIO13, gpio.ModeOutput)
	pin35, err := gpio.OpenPin(rpi.GPIO19, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	// turn the led off on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nClearing and unexporting the pin.\n")
			pin33.Clear()
			pin33.Close()
			pin35.Clear()
			pin35.Close()
			os.Exit(0)
		}
	}()

	for {
		fmt.Println("Mise a un de 35")
		pin35.Set()
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("Mise a zero de 35")
		pin35.Clear()
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("Mise a un de 33 ")
		pin33.Set()
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("Mise a zero de 33 ")
		pin33.Clear()
	}
}
