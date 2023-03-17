package main

import (
	"asso-api/config"
	"strings"
)

func main() {
	ctx := config.Connexion()
	defer config.Client.Disconnect(ctx)
	filenameWithExt := Download()

	segments := strings.Split(filenameWithExt, ".")
	fileName = segments[0]

	Unzip(filenameWithExt, fileName)
	Upload(fileName)
}
