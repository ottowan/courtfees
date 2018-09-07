package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	hc "./havecapital"
)

//Todo store variable
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

//Todos to store variable
type Todos []Todo

func main() {
	//router := mux.NewRouter()
	//router.HandleFunc("/havecapital", GetHaveCapital).Methods("GET")

	http.HandleFunc("/", index)
	http.HandleFunc("/example", example)
	http.HandleFunc("/courtfees", courtfees)
	http.HandleFunc("/courtfees/{type}", courtfeesArbitation)

	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//main project
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Parinya Bumrungchoo")
}

func example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

type Fee struct {
	FeeType  string  `json:"feeType"`
	FeePrice float64 `json:"feePrice"`
}

func courtfees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	result := hc.CaluculateHaveCapitalOver50m(10000.00)

	todos := Fee{FeeType: "อนุญาโต", FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func courtfeesArbitation(w http.ResponseWriter, r *http.Request) {
	r.GetBody()
}

// type ResponseCommands struct {
// 	key   string
// 	value string
// }

// const jsonStream = `
// 	{"ID": 1, "Title": "title 1", "Done": false}
// 	{"ID": 2, "Title": "title 2", "Done": false}

// type Task struct {
// 	ID    int64
// 	Title string
// 	Done  bool
// }

// type Todo struct {
// 	Name      string
// 	Completed bool
// 	Due       time.Time
// }
// type Todos []Todo

// type FeePrice struct {
// 	feePrice string
// }
// type FeePrices []FeePrice

// func GetHaveCapital(w http.ResponseWriter, r *http.Request) {
// 	result := hc.CaluculateHaveCapitalOver50m(10000.00)
// 	fmt.Println(result)
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	feePrices := FeePrices{
// 		FeePrice{feePrice: "3.14"},
// 	}

// 	if err := json.NewEncoder(w).Encode(feePrices); err != nil {
// 		panic(err)
// 	}

// }

// func Examples(w http.ResponseWriter, r *http.Request) {
// 	// result := hc.CaluculateHaveCapitalOver50m(10000.00)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	todos := Todos{
// 		Todo{Name: "Write presentation"},
// 		Todo{Name: "Host meetup"},
// 	}

// 	if err := json.NewEncoder(w).Encode(todos); err != nil {
// 		panic(err)
// 	}

// }
