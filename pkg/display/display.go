package display

import (
	"log"
	"sync"
)

type Display struct {
	mutex sync.Mutex
}

func (d *Display) Show(message string) {
	d.mutex.Lock()
	log.Printf("|| %s ||", message)
	d.mutex.Unlock()
}
