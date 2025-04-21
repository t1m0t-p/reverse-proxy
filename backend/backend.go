package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func displayMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[backend] Requête reçue à : %s, URL : %s\n", time.Now().Format(time.RFC3339), r.URL.Path)
	fmt.Fprint(w, "Réponse du serveur backend")
}

func main() {
	http.HandleFunc("/", displayMessage)
	fmt.Println("Serveur backend démarré sur :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
