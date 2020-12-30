package entities

import muxgo "github.com/muxinc/mux-go"

type AssetData struct {
	Test bool `bson:"test",json:"test"`
	Status string `bson:"status",json:"status"`
	PlaybackIDS []PlaybackID `bson:"playbackIDS",json:"playbackIDS"`
	MP4Support string `bson:"mp4support",json:"mp4support"`
	MasterAccess string `bson:"masterAccess",json:"masterAccess"`
	ID string `bson:"id",json:"id"`
	CreatedAt string `bson:"created_at",json:"created_at"`
}

type PlaybackID struct {
	Policy muxgo.PlaybackPolicy `bson:"policy",json:"policy"`
	ID string `bson:"playbackID",json:"playbackID"`
}

