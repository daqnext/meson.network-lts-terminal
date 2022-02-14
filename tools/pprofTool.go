package tools

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func RunPprofTool(port string) {
	go func() {
		log.Println(http.ListenAndServe(":"+port, nil))
	}()
}
