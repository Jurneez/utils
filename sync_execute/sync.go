package syncexecute

import (
	"sync"
)

func executeFunc(f func(), wg *sync.WaitGroup) {
	defer wg.Done()
	f()
}

// f并发次数n
func SyncExecute(f func(), n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go executeFunc(f, &wg)
	}
	wg.Wait()
}
