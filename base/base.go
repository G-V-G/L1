package base

import "net/http"

// HandleMain for "/"
func HandleMain (res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Lab 1: Gomilko-Vasylynenko-Golovko"))
}