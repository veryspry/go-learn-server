package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var baseURL = "https://hacker-news.firebaseio.com/v0"

// Format a URL to make a request to the hacker news api
func fmtURL(s string) string {
	return baseURL + "/item/" + s + ".json?print=pretty"
}

// Make an http GET request and format the response body
func fmtRes(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

// Filter anything that doesn't have type field == "story"
func filterStories(ids []string) []map[string]interface{} {
	var s []map[string]interface{}

	for i := 0; i < len(ids); i++ {
		url := fmtURL(ids[i])
		body := fmtRes(url)
		byt := []byte(body)
		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		if dat["type"] == "story" {
			s = append(s, dat)
		}
		fmt.Println("finished number:", i)
		fmt.Println("data:", body)
	}
	return s
}

// Get all of the new news item ids
func getNewsItems() []string {
	url := baseURL + "/newstories.json"
	body := fmtRes(url)
	storyIds := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(string(body)), "["), "]")
	idSlice := strings.Split(storyIds, ",")
	return idSlice
}

// GET all stories
// Should return a json response
func getStories(w http.ResponseWriter, r *http.Request) {
	// get all of the story id
	ids := getNewsItems()
	// filter out anything that isn't a story
	// Not sure this is going to be a necessary step
	stories := filterStories(ids)

	var s []map[string]interface{}

	for i := 0; i < len(stories); i++ {
		s = append(s, stories[i])
		// fmt.Println("stories", byt)
	}
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(s)
	byt, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	w.Write(byt)
}

// GET a single story by id read from url parameter
// Should return a json response
func getStory(w http.ResponseWriter, r *http.Request) {
	// read the id
	id := mux.Vars(r)["id"]
	// format the url
	storyURL := fmtURL(id)
	// Make request and get response body
	data := fmtRes(storyURL)
	fmt.Println("data:", data)
	w.Write([]byte(data))
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/news", getStories)
	rtr.HandleFunc("/story/{id}", getStory)
	http.Handle("/", rtr)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
