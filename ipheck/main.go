package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"net/http"
	"strings"
)

type DataWS2 struct {
	Domain       string `json:"domain"`
	Isp          string `json:"isp"`
	Mobile_brand string `json:"mobile_brand"`
}

var ip string
var tmpl, _ = template.ParseFiles("sait/index.html")
var Brand string

func main() {
	HandleFunc()
}
func HandleFunc() {
	http.HandleFunc("/", process)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func process(w http.ResponseWriter, r *http.Request) {

	ip = r.FormValue("ip")
	if (ip == "") || (Brand == "") {
	}
	tmpl.Execute(w, Brand)

	if ip != "" {

		key := ""
		MakeRequest(ip, key)
		fmt.Println(Brand)
	} else {
		Brand = ""
	}
}

func MakeRequest(ip string, key string) {

	resp, err := http.Get("https://api.ip2location.com/v2/?ip=" + ip + "&key=" + key + "&package=WS2&format=json")

	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	var data *DataWS2

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if strings.Contains(data.Isp, "MTS") {
		Brand = "MTS"
	} else if strings.Contains(data.Isp, "Vimpelcom") {
		Brand = "Beeline"
	} else if strings.Contains(data.Isp, "MegaFon") {
		Brand = "MegaFon"
	} else if strings.Contains(data.Isp, "2") {
		Brand = "Tele2"
	} else if strings.Contains(data.Isp, "Tinkoff") {
		Brand = "Tinkoff Mobile"
	} else if strings.Contains(data.Isp, "Ekaterinburg") {
		Brand = "MOTIV"
	} else if strings.Contains(data.Isp, "Win") {
		Brand = "Win Mobile"
	} else {
		MakeRequest2(ip, key)
	}

}

func MakeRequest2(ip string, key string) {

	resp, err := http.Get("https://api.ip2location.com/v2/?ip=" + ip + "&key=" + key + "&package=WS19&format=json")

	if err != nil {
		fmt.Println(err.Error())

	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	var data *DataWS2

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if strings.Contains(data.Domain, "MTS") || strings.Contains(data.Mobile_brand, "MTS") {
		Brand = "MTS"
	} else if strings.Contains(data.Domain, "vimpelcom") || strings.Contains(data.Mobile_brand, "Beeline") {
		Brand = "Beeline"
	} else if strings.Contains(data.Domain, "MegaFon") || strings.Contains(data.Mobile_brand, "MegaFon") {
		Brand = "MegaFon"
	} else if strings.Contains(data.Domain, "tele2") || strings.Contains(data.Mobile_brand, "Tele2") {
		Brand = "Tele2"
	} else if strings.Contains(data.Domain, "tinkoff") || strings.Contains(data.Mobile_brand, "Tinkoff") {
		Brand = "Tinkoff Mobile"
	} else if strings.Contains(data.Domain, "motiv") || strings.Contains(data.Mobile_brand, "MOTIV") {
		Brand = "MOTIV"
	} else if strings.Contains(data.Domain, "Win") || strings.Contains(data.Mobile_brand, "Win") {
		Brand = "Win Mobile"
	} else {
		Brand = "Не найдено. Ваши данные, домен: " + data.Domain + ", Оператор: " + data.Mobile_brand + ", организация: " + data.Isp

	}
}
