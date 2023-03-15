package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Suggestion struct {
	Suggestions []Values
}

type Values struct {
	Value string
}

func main() {

	http.HandleFunc("/", serverRunning)
	http.HandleFunc("/getSuggestion", getSuggestions)

	err := http.ListenAndServe(":8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	// getSuggestions("would chatgpt be")
}

func getSuggestions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	querydata, _ := io.ReadAll(r.Body)
	query := string(querydata)
	querySplit := strings.Split(query, " ")
	query = strings.Join(querySplit, "%20")

	url := "http://suggestqueries.google.com/complete/search?client=firefox&q=" + query

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var arr []interface{}

	err = json.Unmarshal(body, &arr)

	if err != nil {
		fmt.Println("error:", err)
	}

	var sug []string

	for _, v := range arr[1].([]interface{}) { // use type assertion to loop over []interface{}

		val := v.(string)
		sug = append(sug, val)
	}

	fmt.Println(sug)

}

func serverRunning(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server is running!\n")
}
