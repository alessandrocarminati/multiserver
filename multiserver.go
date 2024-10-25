package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pin/tftp"
)

const (
	tftpPort    = 69
	httpPort    = 8080
	servePath   = "./files"
//	defaultFile = "flasher.img"
)

func main() {
	go startTFTP()
	startHTTP()
}

func startTFTP() {
	server := tftp.NewServer(
		func(filename string, rf io.ReaderFrom) error {
			fmt.Printf("TFTP Request: %s\n", filename)

			filePath := "./files/" + filename

			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = rf.ReadFrom(file)
			return err
		},
		func(filename string, wt io.WriterTo) error {
			return fmt.Errorf("Write operation not supported")
		},
	)

	err := server.ListenAndServe("0.0.0.0:" + fmt.Sprint(tftpPort))
	if err != nil {
		fmt.Printf("Error starting TFTP server: %s\n", err)
	}
}

func startHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("HTTP Request: %s\n", r.URL.Path)

		filePath := servePath + r.URL.Path
		if strings.HasSuffix(filePath, "/") {
			filePath += "index.html"
		}

		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error opening file: %s", err), http.StatusNotFound)
			return
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting file info: %s", err), http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, filePath, info.ModTime(), file)
	})

	fmt.Printf("HTTP server listening on port %d...\n", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
	if err != nil {
		fmt.Printf("Error starting HTTP server: %s\n", err)
	}
}
