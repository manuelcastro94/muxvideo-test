package pkg

import (
	"fmt"
	muxgo "github.com/muxinc/mux-go"
	"guitar/entities"
	"guitar/mongo"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request){
	asset, err := entities.MuxVideoClient.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			muxgo.InputSettings{
				Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
			},
		},
		PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC},
	})
	// Check everything was good, and output the playback URL
	if err == nil {
		mongo.SaveAsset(asset.Data)
		fmt.Printf("Playback URL: https://stream.mux.com/%s.m3u8 \n", asset.Data.PlaybackIds[0].Id)
	} else {
		fmt.Printf("Oh no, there was an error: %s \n", err)
	}
}