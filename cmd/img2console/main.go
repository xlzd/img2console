package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xlzd/img2console"
)

type config struct {
	path   string
	url    string
	cols   int
	output string
}

func parseConfig() config {
	pathPtr := flag.String("path", "", "image path to convert")
	urlPtr := flag.String("url", "", "image url to convert")
	colsPtr := flag.Int("cols", 30, "convert cols")
	outputPtr := flag.String("output", "", "save to file, if not set, output to stdout")
	flag.Parse()

	return config{
		path:   strings.TrimSpace(*pathPtr),
		url:    strings.TrimSpace(*urlPtr),
		cols:   *colsPtr,
		output: strings.TrimSpace(*outputPtr),
	}
}

func main() {
	cfg := parseConfig()

	if cfg.path == "" && cfg.url == "" {
		fmt.Println("Fatal error: no input source to convert!")
		os.Exit(2)
	}

	var (
		consoleImg string
		err        error
	)
	if cfg.path != "" {
		consoleImg, err = img2console.ConvertFromPath(cfg.path, cfg.cols)
	} else {
		consoleImg, err = img2console.ConvertFromURL(cfg.url, cfg.cols)
	}
	if err != nil {
		log.Fatal(err)
	}

	var writer *bufio.Writer
	if cfg.output != "" {
		writer = bufio.NewWriter(nil)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	writer.WriteString(consoleImg)
	writer.Flush()
}
