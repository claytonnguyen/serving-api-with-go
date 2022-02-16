package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}


func main() {
	http.HandleFunc("/", ShowcaseBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowcaseBooks(w http.ResponseWriter, r *http.Request) {
	// book := []Book {
	// 	Book{"This is the end of the world", "Joanne Schmidtt"},
	// 	Book{"I can't believe this is happening right now", "Roy Kim"},
	// 	Book{"I have flashbacks to the time when", "Joanne Schmidtt"},
	// 	Book{"TTHis is ludacris", "Joanne Schmidtt"},
	// 	Book{"WHat is the word for the ugliest human on the earth", "Joanne Schmidtt"},
	// }

	b, err := json.Marshal(map[string]interface{}{
		"Book":[]Book{
			Book{"This is the end of the world", "Joanne Schmidtt"},
			Book{"I can't believe this is happening right now", "Roy Kim"},
			Book{"I have flashbacks to the time when", "Joanne Schmidtt"},
			Book{"TTHis is ludacris", "Joanne Schmidtt"},
			Book{"WHat in the world and the earth", "Jane Scott"},
		},
	})
	// b, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}