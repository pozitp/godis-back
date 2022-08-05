package main

import (
	"fmt"
	"net/http"

	"github.com/pozitp/godis-back/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)
	fmt.Println("Server started at port 8000")
	http.ListenAndServe(":8000", nil)
}
