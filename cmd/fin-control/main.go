package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/MuriLovin/fin-control/internal/router"
)

func main() {
	router := router.InitRouter()
	err := http.ListenAndServe(":8090", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed with error: %s\n", err)
		return
	}

	if err != nil {
		fmt.Printf("Server closed with error: %s\n", err)
		os.Exit(1)
		return
	}
}
