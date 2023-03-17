package main

import (
	"asso-api/config"
	"asso-api/model"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func insert(collection *mongo.Collection, dirName string, fileName string, c chan int) {
	file, err := os.Open(filepath.Join(dirName, fileName))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	idx := 0
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		association := model.Association{
			IDRna:           rec[0],
			Name:            strings.ToValidUTF8(rec[8], ""),
			DateCreation:    rec[4],
			DatePublication: rec[5],
			Object:          rec[9],
			ObjetSocial1:    rec[10],
			ObjetSocial2:    rec[11],
			Address:         rec[12],
			ZipCode:         rec[15],
			City:            rec[16],
		}
		// do something with read line
		if idx > 0 {
			_, err := collection.InsertOne(context.Background(), association)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		idx += 1
	}
	// records, _ := reader.ReadAll()
	// for line := range records {
	// 	fmt.Printf(": %+v\n", line)
	// }
	// fmt.Printf("records")
	c <- 1
}

func Upload(fileName string) {
	collection := config.Client.Database("asso").Collection("asso")

	mod := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "name", Value: bsonx.String("text")}},
		Options: options.Index().SetDefaultLanguage("french"),
	}
	collection.Indexes().CreateOne(context.Background(), mod)

	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int, len(files))
	fmt.Printf("Upload started\n")

	start := time.Now()
	for _, file := range files {
		go insert(collection, fileName, file.Name(), c)
	}

	for i := 0; i < len(files); i++ {
		<-c // wait for one task to complete
	}
	elapsed := time.Since(start)
	fmt.Printf("Upload finished in :%v\n", elapsed)
}
