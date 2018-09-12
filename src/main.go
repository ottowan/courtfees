package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	arb "./arbitration"
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

	http.HandleFunc("/havecapital", havecapital)
	http.HandleFunc("/noncapital", noncapital)
	http.HandleFunc("/mortgage", mortgage)
	http.HandleFunc("/arbitration", arbitration)

	fmt.Printf("GOLANG server @port 8080")
	err := http.ListenAndServeTLS(":8080", "./cert/server.crt", "./cert/server-decrypted.key", nil)
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

	section := r.Form.Get("section")
	feeCapital := r.Form.Get("feeCapital")

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	result := nc.CalculateNonCapital(section, feeCapitalFloat)
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

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	result := mg.CalculateMortgage(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func arbitration(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	feeCapital := r.Form.Get("feeCapital")
	amountPerson := r.Form.Get("amountPerson")

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	amountPersonInt, err := strconv.ParseInt(amountPerson, 10, 0)
	result := arb.CalculateArbitration(int(amountPersonInt), feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
