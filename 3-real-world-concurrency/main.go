package main

import (
	"fmt"
	"net/http"
	"time"
)

var visitantsNumber uint64 = 0

func main() {
	//mutex := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//mutex.Lock()
		//visitantsNumber++
		//mutex.Unlock()

		//atomic.AddUint64(&visitantsNumber, 1)

		time.Sleep(300 * time.Millisecond)

		w.Write([]byte(fmt.Sprintf("Visitant number: %d", visitantsNumber)))
	})
	http.ListenAndServe(":3000", nil)
}
