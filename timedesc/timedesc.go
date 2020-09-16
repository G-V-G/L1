package timedesc

import (
	"net/http"
	"time"
	"encoding/json"
)

type timestamp struct {
	Time string `json:"time"`
}

// HandleTime for "/time"
func HandleTime (res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	timestampRes := timestamp{currentTime}
	res.Header().Set("Content-Type", "application/json")
	if timeJSON, err := json.MarshalIndent(&timestampRes, "", " "); err != nil {
		res.Write([]byte("{}"))
	} else {
		res.Write(timeJSON)
	}
}