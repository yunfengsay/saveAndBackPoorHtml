/**
touch index.html
mkdir back
go run server.go
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type DocumentContent struct {
	Content string `json:"content"`
}
type Res struct {
	Success bool   `json:"success"`
	msg     string `json:"msg"`
}

func main() {
	http.HandleFunc("/api/save", save)
	fmt.Println("监听 8809 端口中")
	http.ListenAndServe(":8809", nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	var d DocumentContent
	var res Res
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	exec.Command("bash", "-c", "echo d > ./sss.html")
	currentTimeStr := strconv.FormatInt(time.Now().Unix(), 10)
	exec.Command("bash", "-c", "cp ./index.html  ./back/index."+currentTimeStr+".html").Output()
	exec.Command("bash", "-c", "echo "+d.Content+" > ./index.html").Output()
	res.Success = true
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
