package commits

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

// Person comitter
type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Date string `json:"date"`
}

// Commit in repo story
type Commit struct {
	Author Person `json:"author"`
	Committer Person `json:"committer"`
	Msg string `json:"message"`
}

// CommitWrapper from pasring
type CommitWrapper struct {
	Branch string `json:"sha"`
	Commit Commit `json:"commit"`
}

// HandleCommits for "/commits"
func HandleCommits(res http.ResponseWriter, req *http.Request) {
	comRes, err := http.Get("https://api.github.com/repos/Andrew1407/test_repo/commits")
	if (err != nil) {
		http.NotFound(res, req)
		return
	}
	defer comRes.Body.Close()
	res.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(comRes.Body)
	if (err != nil) {
		res.Write([]byte("{}"))
		return
	}
	var commits []CommitWrapper
	err = json.Unmarshal(body, &commits)
	if (err != nil) {
		res.Write([]byte("{}"))
		return
	}
	commitsJSON, err := json.MarshalIndent(&commits, "", " ")
	res.Header().Set("Content-Type", "application/json")
	if (err != nil) {
		res.Write([]byte("{}"))
		return
	}
	res.Write(commitsJSON)
}
