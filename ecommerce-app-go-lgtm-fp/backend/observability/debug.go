package observability

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func StartDebugServer(port string) {
	go func() {
		log.Printf("[DEBUG] pprof listening on :%s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Printf("[DEBUG] pprof server error: %v", err)
		}
	}()
}
