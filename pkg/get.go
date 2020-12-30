package pkg

import (
	"encoding/json"
	"guitar/mongo"
	"net/http"
)

func GetVideos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	videos := mongo.GetAssets()
	json.NewEncoder(w).Encode(videos)
}
