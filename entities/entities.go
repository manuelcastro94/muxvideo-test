package entities

import (
	muxgo "github.com/muxinc/mux-go"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ServerConfig *ServerConfigModel
	MongoClient *mongo.Client
	MuxVideoClient *muxgo.APIClient
)

type ServerConfigModel struct {
	Port   string `yaml:"port"`
	MongoHost   string `yaml:"mongodb"`
	VideoCollection  string `yaml:"collection"`
}

func InitMongoClient(receivedMongoClient *mongo.Client) {
	MongoClient = receivedMongoClient
}

func InitMuxVideoClient(receivedMuxVideoClient *muxgo.APIClient) {
	MuxVideoClient = receivedMuxVideoClient
}
