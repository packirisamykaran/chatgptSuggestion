package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

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
	Result string `json:"result"`
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getSuggestion", getSuggestions).Methods("POST", "OPTIONS")

	return r
}

func getSuggestions(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var body Data

	_ = json.NewDecoder(r.Body).Decode(&body)

	fmt.Printf("%+v", body.Result)
	// query := string(querydata)
	// querySplit := strings.Split(query, " ")
	// query = strings.Join(querySplit, "%20")

	data := Data{
		Result: "hello",
	}

	json.NewEncoder(w).Encode(data)

	// url := "http://suggestqueries.google.com/complete/search?client=firefox&q=" + query

	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer resp.Body.Close()
	// //We Read the response body on the line below.
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var arr []interface{}

	// err = json.Unmarshal(body, &arr)

	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// var sug []string

	// for _, v := range arr[1].([]interface{}) { // use type assertion to loop over []interface{}

	// 	val := v.(string)
	// 	sug = append(sug, val)
	// }

	// fmt.Println(sug)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Context-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")

}
