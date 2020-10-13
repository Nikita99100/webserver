package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var groups []Group
	file, _ := os.Open("groups.json")
	defer file.Close()
	json.NewDecoder(file).Decode(&groups)
	fmt.Println(groups[1].Name)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, groups)
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "не понятно")
	})
	http.ListenAndServe(":80", nil)
}

type Group struct {
	Name        string `json:"group_name"`
	Description string `json:"group_description"`
	ID          int    `json:"group_id"`
	ParentID    int    `json:"parent_id"`
}
