package main

import (
	"sync"
)

func main(){
	var rwmutex *sync.RWMutex
	rwmutex = new(sync.RWMutex)
	rwmutex.RUnlock()
}