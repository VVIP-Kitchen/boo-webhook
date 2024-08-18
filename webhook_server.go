package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

// PushEvent represents a GitHub push event
type PushEvent struct {
	// Ref is the reference of the push event (e.g., "refs/heads/main")
	Ref string `json:"ref"`
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Starting webhook server on :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var pushEvent PushEvent
	err := json.NewDecoder(r.Body).Decode(&pushEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the push event is for the "main" branch
	if pushEvent.Ref == "refs/heads/main" {
		// Trigger the deployment in a separate goroutine
		go deploy()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deployment started"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// deploy runs the deployment script
func deploy() {
	// Create a command to run the deployment script
	cmd := exec.Command("/home/ifkash/deploy.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Deployment failed: %v\n", err)
		log.Printf("Output: %s\n", output)
	} else {
		log.Printf("Deployment successful. Output: %s\n", output)
	}
}