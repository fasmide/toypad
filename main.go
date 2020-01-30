package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fasmide/toypad/draw"
	"github.com/rakyll/launchpad"
)

const fileStorage string = "state.json"

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		log.Fatalf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()

	fd, err := os.Open(fileStorage)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	logic := draw.NewLogic(pad)

	// if the above file exists
	if !os.IsNotExist(err) {
		log.Printf("Loading state from %s", fileStorage)

		dec := json.NewDecoder(fd)
		err := dec.Decode(logic)
		if err != nil {
			panic(err)
		}
	}

	logic.Render()

	var timer *time.Timer
	var lock sync.Mutex

	ch := pad.Listen()
	for {
		hit := <-ch
		log.Printf("Button pressed at <x=%d, y=%d>", hit.X, hit.Y)

		// Keydown is racing against time.AfterFunc which will save the state
		// if no key press have been detected in the last 5 seconds
		lock.Lock()
		logic.KeyDown(hit.X, hit.Y)

		if timer != nil {
			timer.Reset(time.Second * 5)
		}

		if timer == nil {
			timer = time.AfterFunc(time.Second*5, func() {
				log.Print("Saving state...")
				payload, err := json.Marshal(logic)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(fileStorage, payload, 0644)
				if err != nil {
					panic(err)
				}

				lock.Lock()
				timer = nil
				lock.Unlock()
			})
		}
		lock.Unlock()
	}
}
