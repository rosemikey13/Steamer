package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/getGames", GetGamesHandler)

	http.ListenAndServe(":8069", serveMux)
}
