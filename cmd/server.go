package cmd

import (
	"github.com/parth105/simple-http/internal/httpserver"
)

func WikiServer(port string) {
	httpserver.StartServer(port)
}
