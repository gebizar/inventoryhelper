package main

import (
	"fmt"
	"inventoryhelper/servers/gateway/character"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Hello")
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/character", character.CharacterHandler)
	log.Printf("server is listening on %s...", ":80")
	// 80 is standard for http 443 is standard for https
	log.Fatal(http.ListenAndServe(":80", mux))

}
