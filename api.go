package gophercise3

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func JsonToStoryMap(path string) map[string]story {

	stories := map[string]story{}

	jsonData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(jsonData, &stories)
	if err != nil {
		log.Print(string(jsonData))
		log.Fatalln(err)
	}
	return stories
}
func StartServer() {
	mux := http.NewServeMux()

	storiesMap := JsonToStoryMap("stories.json")
	mapHandler := MapHandler(storiesMap, mux)
	http.ListenAndServe(":9019", mapHandler)
}

func readTemplateFromFile(path string) string {
	templ, err := os.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
	}
	return string(templ)
}

func MapHandler(pathsToUrls map[string]story, fallback http.Handler) http.HandlerFunc {

	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.RequestURI)
		log.Println(strings.Split(r.RequestURI, "/")[1])
		log.Println(pathsToUrls[strings.Split(r.RequestURI, "/")[1]])
		w.Header().Set("Content-Type", "text/html")
		t, err := template.ParseFiles("template.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, pathsToUrls[strings.Split(r.RequestURI, "/")[1]])
	}
}

func htmlHandler(html []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {}, nil
}
