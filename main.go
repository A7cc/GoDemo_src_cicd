package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func execHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")
	// 危险示例：用户输入直接进入系统命令
	result, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "%s", result)
}
func main() {

	var password = "123456"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello ci/cd"))
	})
	http.HandleFunc("/ping", execHandler)
	fmt.Println(password)
	http.ListenAndServe(":8080", nil)
}
