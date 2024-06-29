package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pin/tftp/v3"
)

var (
	rootDir = ""
)

func pathSanitizer(path string) (string, error) {
	// Prevent directory traversal by removing leading slashes
	for strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	// Prevent access to parent directories by throwing an error if the path contains ".."
	if strings.Contains(path, "..") {
		return path, fmt.Errorf("invalid path, should not contains '..'")
	}

	return rootDir + "/" + path, nil
}

func readHandler(filename string, rf io.ReaderFrom) error {
	sanitizedFilename, err := pathSanitizer(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}

	fmt.Println("Will read file ", sanitizedFilename)

	file, err := os.Open(sanitizedFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)

	return nil
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	listenAddress := flag.String("tftp.listen-address", ":69", "Address to listen on for TFTP requests")
	timeoutSeconds := flag.Int("tftp.timeout-seconds", 5, "Timeout in seconds for TFTP requests")
	rd := flag.String("tftp.root-dir", wd, "Root directory for the TFTP server")

	flag.Parse()
	rootDir = *rd

	fmt.Println("Starting TFTP server on", *listenAddress, "with root directory", rootDir, "and timeout", *timeoutSeconds, "seconds.")

	s := tftp.NewServer(readHandler, nil)
	s.SetTimeout(time.Duration(*timeoutSeconds) * time.Second)
	err = s.ListenAndServe(*listenAddress)
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
