// Discord Bot用コード
package main

import (
	"fmt"
	"net/http"
)

// http://localhost:8080/ へアクセスしたときの処理
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}