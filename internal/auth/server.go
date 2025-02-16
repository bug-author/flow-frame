package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func StartAuthServer(port int, authChan chan string) {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code != "" {
			fmt.Fprintf(w, "Authorization successful! You can close this window.")
			authChan <- code
		} else {
			http.Error(w, "Authorization code not found", http.StatusBadRequest)
		}
	})

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		log.Printf("Listening on http://localhost:%d/callback", port)

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	go func() {
		time.Sleep(time.Minute * 2)
		log.Println("Shutting down auth server due to timeout.")
		_ = server.Close()
	}()

}
