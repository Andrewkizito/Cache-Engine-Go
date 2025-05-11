package main

import (
	"cache-engine/utils"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Map default func
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		action := r.URL.Query().Get("action")
		key := r.URL.Query().Get("key")
		contentType := r.Header.Get("Content-Type")

		fmt.Printf("Content-Type: %s", contentType)

		if method == "PUT" {
			if action == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Action is required")
				return
			}

			switch action {
			// Handle read
			case "get":
				entry, err := utils.ReadCachEntry(key)

				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "%s", err.Error())
					return
				}

				w.Header().Set("Content-Type", entry.ContentType)
				fmt.Fprintf(w, "%s", string(entry.Raw))
			// Handle adding new entries
			case "set":
				bodyBytes, readBodyErr := io.ReadAll(r.Body)

				if readBodyErr != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Failed to read body - %s", readBodyErr.Error())
					return
				}

				id, err := utils.AddCacheEntry(key, utils.RawCacheEntry{
					ContentType: contentType,
					Raw:         bodyBytes,
				})

				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "%s", err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"id": "%s"}`, id)
				return
			default:
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "Action Not Allowed - [set, get]")
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Not Allowed")
		}
	}))

	// Serve and listen on port 9000
	log.Println("Server is running on http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
