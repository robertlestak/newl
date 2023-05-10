package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
}

func usage() {
	u := `Usage: %s [file] [opt:out file]
If file is not specified, stdin is used.
If out file is not specified, stdout is used.
`
	fmt.Fprintf(os.Stderr, u, os.Args[0])
}

func newline(in io.Reader, out io.Writer) error {
	// replace all newlines with literal "\n"
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Fprintf(out, "%s\\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) > 3 {
		usage()
		os.Exit(1)
	}
	var in io.Reader
	var out io.Writer
	var inPath string
	var outPath string
	if len(os.Args) == 1 {
		in = os.Stdin
		out = os.Stdout
	} else if len(os.Args) == 2 {
		inPath = os.Args[1]
		inFile, err := os.Open(inPath)
		if err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
		in = inFile
		out = os.Stdout
	} else if len(os.Args) == 3 {
		inPath = os.Args[1]
		inFile, err := os.Open(inPath)
		if err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
		in = inFile
		outPath = os.Args[2]
		outFile, err := os.Create(outPath)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
		out = outFile
	}
	if err := newline(in, out); err != nil {
		log.Fatal(err)
	}
}
