package main

import (
	"asso-api/internal/config"
	"asso-api/internal/importer"
	"os"
	"strings"
)

func main() {
	ctx := config.Connexion()
	defer config.Client.Disconnect(ctx)
	filenameWithExt := importer.Download()

	segments := strings.Split(filenameWithExt, ".")
	fileName := segments[0]
	defer os.RemoveAll("data")

	importer.Unzip(filenameWithExt, fileName)
	importer.Upload(fileName)
}
