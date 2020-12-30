package main

import (
	"context"
	"flag"
	"fmt"
	muxgo "github.com/muxinc/mux-go"
	"gopkg.in/yaml.v2"
	"guitar/entities"
	"guitar/handler"
	endpoints "guitar/pkg"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MUX_CLIENT_ID = "35de6729-c6f6-451b-9ad7-5d7354e9c415"
	MUX_CLIENT_SECRET = "oPsItqE0YljDzcqDAZ/y0wo3AFBA5tYpTEn5l3zU+rFp13S/ybtkDTgntqs90PbieSQo7rR3LHr"
)

var (
	mongoClient *mongo.Client
)

func main() {

	yamlConfigFile := flag.String("config", "config.yml", "Path to the yaml server configuration")
	yamlFile, err := ioutil.ReadFile(*yamlConfigFile)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	entities.ServerConfig = &entities.ServerConfigModel{}
	err = yaml.Unmarshal(yamlFile, entities.ServerConfig)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/upload", endpoints.Upload)
	mux.HandleFunc("/api/v0/videos", endpoints.GetVideos)

	spa := handler.SpaHandler{StaticPath: "guitar-web/build", IndexPath: "index.html"}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spa.ServeHTTP(w, r)
	})
	var mongoErr error
	clientOptions := options.Client().ApplyURI(entities.ServerConfig.MongoHost)
	mongoClient, mongoErr = mongo.Connect(context.TODO(), clientOptions)
	entities.InitMongoClient(mongoClient)
	if mongoErr != nil {
		fmt.Println(mongoErr)
	}

	// API Client Initialization
	muxVideoClient := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(MUX_CLIENT_ID, MUX_CLIENT_SECRET),
		))

	entities.InitMuxVideoClient(muxVideoClient)

	srv := &http.Server{
		Addr:         ":" + entities.ServerConfig.Port,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}