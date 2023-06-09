package main

import (
	"basededatos/backend"
	"log"
	"net/http"
	"os"
)

func main() {
	path := "../mamuro-email/dist"

	if len(os.Args) < 3 {
		log.Printf("\x1b[31;1mError: Missing arguments <port>\x1b[0m\n")
		return
	} else if len(os.Args) > 4 {
		if os.Args[3] == "-d" && os.Args[4] == "" {
			log.Printf("\x1b[31;1mError: Missing path to files e.g ./mamuro -p 8080 -d ../mamuro-email/dist\x1b[0m\n")
			return
		} else if os.Args[3] == "-d" && os.Args[4] != "" {
			path = os.Args[4]
			// Verify that the path exists
			if _, err := os.Stat(os.Args[4]); os.IsNotExist(err) {
				log.Printf("\x1b[31;1mError: Path does not exist\x1b[0m\n")
				// Print paths that exist in the current directory
				log.Printf("\x1b[33;1mPaths that exist in the current directory:\x1b[0m\n")
				file, err := os.Open(".")
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				list, _ := file.Readdirnames(0) // 0 to read all files and folders
				for _, name := range list {
					log.Printf(name)
				}
				return
			}
		}
	}

	port := os.Args[2]

	if os.Args[1] != "-p" && os.Args[2] == "" {
		log.Printf("\x1b[31;1mError: Missing port number e.g ./mamuro -p 8080\x1b[0m\n")
		return
	}

	go backend.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpDir := http.Dir(path)
		http.FileServer(httpDir).ServeHTTP(w, r)
	})

	log.Printf("\x1b[32;1mServer listening on port %s...\x1b[0m\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
