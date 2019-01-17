package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var testString = "yeah, buddy"
var baseURL = "https://hacker-news.firebaseio.com/v0"

func fmtURL(s string) string {
	return baseURL + "/item/" + s + ".json?print=pretty"
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	fmt.Println("http.Request is here:", r.URL)
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func formatNums(w http.ResponseWriter, r *http.Request) {
	nums := r.URL.Path
	nums = strings.TrimPrefix(nums, "/addnums/")
	w.Write([]byte(nums))
}

func printShit() {
	var s = "hello"
	fmt.Println(s)
	s = "goodbye"
	fmt.Println(s)
	var p = &s
	*p = "yaassss"
	fmt.Println(s)
	s = "whoa"
	fmt.Println(s)
}

// Get and format http response body
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

// Get all of the new news items
func getNewsItems() []string {
	url := baseURL + "/newstories.json"
	body := fmtRes(url)
	storyIds := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(string(body)), "["), "]")
	idSlice := strings.Split(storyIds, ",")
	return idSlice
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
	}
	// fmt.Println(dat["type"])
	return s
}

// Get all stories
func getNews(w http.ResponseWriter, r *http.Request) {

	ids := getNewsItems()
	stories := filterStories(ids)

	fmt.Println("stories", stories[0])
}

func getStory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	storyURL := fmtURL(id)
	data := fmtRes(storyURL)

	fmt.Println("vars", data)
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/writename", sayHello)
	rtr.HandleFunc("/addnums/*", formatNums)
	rtr.HandleFunc("/news", getNews)
	rtr.HandleFunc("/story/{id}", getStory)
	http.Handle("/", rtr)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	// http.Handle("/", http.FileServer(http.Dir("./public")))
	// http.HandleFunc("/writename", sayHello)
	// http.HandleFunc("/addnums/*", formatNums)
	// http.HandleFunc("/news", getNews)
	// http.HandleFunc("/story*", getStory)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }
}
