package main

import (
	"net/http"
	"./base"
	"./timedesc"
	"./commits"
)

func main() {
	http.HandleFunc("/", base.HandleMain)
	http.HandleFunc("/time", timedesc.HandleTime)
	http.HandleFunc("/commits", commits.HandleCommits)

	http.ListenAndServe(":8795", nil)
}
