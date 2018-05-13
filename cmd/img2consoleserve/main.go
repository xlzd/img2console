package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/xlzd/img2console"
)

var (
	defaultCols           = 30
	defaultImageString, _ = img2console.ConvertFromPath("assets/demo.jpg", defaultCols)
)

func main() {
	http.HandleFunc("/", convert)
	log.Fatal(http.ListenAndServe(":1926", nil))
}

func convert(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("User-Agent"), "curl") {
		http.Redirect(w, r, "https://github.com/xlzd/img2console", http.StatusFound)
		return
	}
	finalResp := defaultImageString

	if err := r.ParseForm(); err == nil {
		if url := r.Form.Get("url"); url != "" {
			cols, err := strconv.Atoi(r.Form.Get("cols"))
			if err != nil || cols >= 160 || cols <= 5 {
				cols = defaultCols
			}
			finalResp, err = img2console.ConvertFromURL(url, cols)
			if err != nil {
				finalResp = defaultImageString
			}
		}
	}

	flusher, _ := w.(http.Flusher)
	fmt.Fprintf(w, finalResp)
	flusher.Flush()
}
