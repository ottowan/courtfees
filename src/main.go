package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	hc "./havecapital"
	mg "./mortgage"
	nc "./noncapital"
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

	// http.HandleFunc("/", index)
	// http.HandleFunc("/example", example)
	http.HandleFunc("/havecapital", havecapital)
	http.HandleFunc("/nonecapital", havecapital)
	http.HandleFunc("/mortgage", mortgage)
	http.HandleFunc("/arbitration", arbitration)
	//http.HandleFunc("/courtfees/{type}", courtfeesArbitation)

	fmt.Printf("GOLANG server @port 9090")
	err := http.ListenAndServeTLS(":9090", "./cert/server.crt", "./cert/server.key", nil)
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

func havecapital(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	feeCapital := r.Form.Get("feeCapital")

	//fmt.Printf("%s", feeCapital)

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	result := hc.CalculateHaveCapital(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func noncapital(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	feeCapital := r.Form.Get("feeCapital")

	//fmt.Printf("%s", feeCapital)

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	result := nc.CalculateNonCapital("", feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func mortgage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	feeCapital := r.Form.Get("feeCapital")

	//fmt.Printf("%s", feeCapital)

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	result := mg.CalculateMortgage(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func arbitration(w http.ResponseWriter, r *http.Request) {
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
