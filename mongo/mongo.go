package mongo

import (
	"context"
	"fmt"
	muxgo "github.com/muxinc/mux-go"
	"go.mongodb.org/mongo-driver/bson"
	"guitar/entities"
	"log"
)

func SaveAsset(asset muxgo.Asset) {
	var assetData entities.AssetData
	assetData.ID = asset.Id
	assetData.Status = asset.Status
	pIds := make([]entities.PlaybackID,0)
	for _, id := range asset.PlaybackIds {
		pid := entities.PlaybackID{
			Policy: id.Policy,
			ID:     id.Id,
		}
		pIds = append(pIds,pid)
	}
	assetData.PlaybackIDS = pIds
	assetData.Test = asset.Test
	assetData.CreatedAt = asset.CreatedAt
	assetData.MP4Support = asset.Mp4Support
	assetData.MasterAccess = asset.MasterAccess
	collection := entities.MongoClient.Database("test").Collection("videos")
	collection.InsertOne(context.TODO(),assetData)
}

func GetAssets() []entities.AssetData {
	collection := entities.MongoClient.Database("test").Collection("videos")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cursor.Close(context.TODO())
	videos := make([]entities.AssetData,0)
	for cursor.Next(context.TODO()) {
		var video entities.AssetData
		if err = cursor.Decode(&video); err != nil {
			log.Fatal(err)
			return nil
		}
		videos = append(videos,video)
		fmt.Println(video)
	}
	return videos
}
