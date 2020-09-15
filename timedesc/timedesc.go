package timedesc

import (
	"net/http"
	"time"
	"encoding/json"
)

// Timestamp response
type Timestamp struct {
	Time string `json:"time"`
}

// HandleTime for "/time"
func HandleTime (res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	timeFormated := currentTime.Format(time.RFC3339)
	timestampRes := Timestamp{timeFormated}
	timeJSON, err := json.MarshalIndent(timestampRes, "", " ")
	res.Header().Set("Content-Type", "application/json")
	if (err != nil) {
		res.Write([]byte("{}"))
		return
	}
	res.Write(timeJSON)
}