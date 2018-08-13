package havecapital
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	WritRate float64 `json:"writrate"`
}

const url = "http://localhost:3000/writrate"
const feeCapitalConst = 50000000.00
const feeTypeConst = "มีทุนทรัพย์"
const failFeePriceConst = 0.00

func CalculateCapitalEqual50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00

	if feeType == feeTypeConst && feeCapital == feeCapitalConst {
		return feePrice
	}

	return failFeePriceConst

}

func CalculateCapitalLess50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00

	if feeType == feeTypeConst && feeCapital < feeCapitalConst {
		return (feeCapital * 2) / 100
	}

	return failFeePriceConst

}


func CalculateCapitalOver50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00

	if feeType == feeTypeConst && feeCapital > feeCapitalConst {
		return feePrice + (((feeCapital - feeCapitalConst) * 0.1) / 100)
	}

	return failFeePriceConst

}

func CalculateWritRate(moo int, subDistrict string, district string, province string) float64 {

	values := map[string]string{"moo": strconv.Itoa(moo), "subDistrict": subDistrict, "district": district, "province": province}
	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var datas Data
	json.Unmarshal(body, &datas)

	// i, err := strconv.ParseFloat(datas[0].WritRate, 64)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	return datas.WritRate
}
