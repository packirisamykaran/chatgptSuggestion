package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Suggestion struct {
	Suggestions []Values
}

type Values struct {
	Value string
}

func main() {

	r := Router()

	err := http.ListenAndServe(":8000", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	// getSuggestions("would chatgpt be")
}

type Data struct {
	Query string `json:"query"`
}

type ReturnData struct {
	Suggestions []string `json:"suggestions"`
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getSuggestion", getChatGPTresponse).Methods("POST", "OPTIONS")

	return r
}

func getSuggestions(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var reqbody Data

	_ = json.NewDecoder(r.Body).Decode(&reqbody)

	query := string(reqbody.Query)
	querySplit := strings.Split(query, " ")
	query = strings.Join(querySplit, "%20")
	// fmt.Printf("%+v \n", query)
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

	// fmt.Println(sug)

	data := ReturnData{
		Suggestions: sug,
	}

	json.NewEncoder(w).Encode(data)

}

type ChatApiData struct {
	Model    string `json:model`
	Messages []MessageData
}

type MessageData struct {
	Role    string `json:role`
	Content string `json:content`
}

func getChatGPTresponse(w http.ResponseWriter, r *http.Request) {
	url := "https://chatgpt-api.shn.hk/v1/"
	enableCors(&w)

	var chatQeury ChatApiData = ChatApiData{
		Model:    "gpt-3.5-turbo",
		Messages: []MessageData{{Role: "user", Content: "who is thanos"}},
	}

	chatQueryJson, _ := json.Marshal(chatQeury)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(chatQueryJson))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Context-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")

}
