package importer

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	fileName    string
	fullURLFile string
)

func Download() string {
	currentYearMonth := time.Now().Format("20060102")

	fullURLFile = strings.Join([]string{"https://media.interieur.gouv.fr/rna/rna_import_", currentYearMonth, ".zip"}, "")

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = strings.Join([]string{"data", segments[len(segments)-1]}, "/")

	// Create blank file
	os.MkdirAll("data", 0755)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	defer file.Close()

	return fileName
}
