package main

import (
	"fmt"
	"net/http"
	// "sync"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()
		atomic.AddUint64(&number, 1)

		// time.Sleep(150 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Voce teve %d acessos nessa página.\n", number)))
	})
	http.ListenAndServe(":3000", nil)
}
