package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Hello, World!\n")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// サーバー起動
	log.Println("サーバーを起動します。ポート：8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("サーバー起動に失敗しました: %v", err)
	}
}
