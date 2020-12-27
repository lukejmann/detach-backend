package main

import (
	"log"
	"sync"

	"github.com/lukejmann/detach2-backend/timing-server/timing"
)

func main() {
	tM, err := timing.NewTimingManager()
	if err != nil {
		log.Fatal("NewTimingManager:", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	tM.Start()
	wg.Wait()
}
