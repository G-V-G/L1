package commits

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type person struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Date string `json:"date"`
}

type commit struct {
	Author person `json:"author"`
	Committer person `json:"committer"`
	Msg string `json:"message"`
}

type commitWrapper struct {
	Branch string `json:"sha"`
	Commit commit `json:"commit"`
}

// HandleCommits for "/commits"
func HandleCommits(res http.ResponseWriter, req *http.Request) {
	comRes, err := http.Get("https://api.github.com/repos/G-V-G/l1/commits")
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
	var commits []commitWrapper
	if err := json.Unmarshal(body, &commits); err != nil {
		res.Write([]byte("{}"))
		return
	}
	if commitsJSON, err := json.MarshalIndent(&commits, "", " "); err != nil {
		res.Write([]byte("{}"))
	} else {
		res.Write(commitsJSON)
	}
}
