package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	arb "./arbitration"

	pvAppealOrSupreme "./provincialcourt/appealorsupreme"
	pvCompensation "./provincialcourt/compensation"
	pvHaveCapital "./provincialcourt/havecapital"
	pvMortgage "./provincialcourt/mortgage"
	pvNonAndHaveCapital "./provincialcourt/nonandhavecapital"
	pvNonCapital "./provincialcourt/noncapital"

	kwAppealOrSupreme "./kwaengcourt/appealorsupreme"
	kwCompensation "./kwaengcourt/compensation"
	kwHaveCapital "./kwaengcourt/havecapital"
	kwMortgage "./kwaengcourt/mortgage"
	kwNonAndHaveCapital "./kwaengcourt/nonandhavecapital"
	kwNonCapital "./kwaengcourt/noncapital"
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

	http.HandleFunc("/arbitration", arbitration)

	http.HandleFunc("/provincialcourt/havecapital", ProvincialCourtHaveCapital)
	http.HandleFunc("/provincialcourt/noncapital", ProvincialCourtNonCapital)
	http.HandleFunc("/provincialcourt/mortgage", ProvincialCourtMortgage)
	http.HandleFunc("/provincialcourt/appealorsupreme", ProvincialCourtAppealOrSupreme)
	http.HandleFunc("/provincialcourt/compensation", ProvincialCourtCompensation)
	http.HandleFunc("/provincialcourt/nonandhavecapital", ProvincialCourtNonAndHaveCapital)

	http.HandleFunc("/kwaengcourt/havecapital", KwaengCourtHaveCapital)
	http.HandleFunc("/kwaengcourt/noncapital", KwaengCourtNonCapital)
	http.HandleFunc("/kwaengcourt/mortgage", KwaengCourtMortgage)
	http.HandleFunc("/kwaengcourt/appealorsupreme", KwaengCourtAppealOrSupreme)
	http.HandleFunc("/kwaengcourt/compensation", KwaengCourtCompensation)
	http.HandleFunc("/kwaengcourt/nonandhavecapital", KwaengCourtNonAndHaveCapital)

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

type Fee struct {
	FeeType  string  `json:"feeType"`
	FeePrice float64 `json:"feePrice"`
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

func ProvincialCourtHaveCapital(w http.ResponseWriter, r *http.Request) {

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
	result := pvHaveCapital.CalculateHaveCapital(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ProvincialCourtNonAndHaveCapital(w http.ResponseWriter, r *http.Request) {

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
	result := pvNonAndHaveCapital.CalculateNonAndHaveCapital(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ProvincialCourtNonCapital(w http.ResponseWriter, r *http.Request) {

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
	fmt.Printf("%s", feeCapital)

	feeCapitalFloat, err := strconv.ParseFloat(feeCapital, 64)
	result := pvNonCapital.CalculateNonCapital(section, feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ProvincialCourtMortgage(w http.ResponseWriter, r *http.Request) {

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
	result := pvMortgage.CalculateMortgage(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ProvincialCourtAppealOrSupreme(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", feeCapital)
	result := pvAppealOrSupreme.CalculateAppealOrSupreme()
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ProvincialCourtCompensation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", feeCapital)
	result := pvCompensation.CalculateCompensation()
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func KwaengCourtHaveCapital(w http.ResponseWriter, r *http.Request) {

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
	result := kwHaveCapital.CalculateHaveCapital(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func KwaengCourtNonCapital(w http.ResponseWriter, r *http.Request) {

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
	result := kwNonCapital.CalculateNonCapital(section, feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func KwaengCourtMortgage(w http.ResponseWriter, r *http.Request) {

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
	result := kwMortgage.CalculateMortgage(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func KwaengCourtAppealOrSupreme(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", feeCapital)
	result := kwAppealOrSupreme.CalculateAppealOrSupreme()
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func KwaengCourtCompensation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", feeCapital)
	result := kwCompensation.CalculateCompensation()
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func KwaengCourtNonAndHaveCapital(w http.ResponseWriter, r *http.Request) {

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
	result := kwNonAndHaveCapital.CalculateNonAndHaveCapital(feeCapitalFloat)
	todos := Fee{FeePrice: result}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
