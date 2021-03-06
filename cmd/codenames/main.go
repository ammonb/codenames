package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jbowens/codenames"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	server := &codenames.Server{
		Server: http.Server{
			Addr: ":80",
		},
	}
	if err := server.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
