package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Status Data `json:"status"`
}

type Data struct {
	Wind int `json:"wind"`
	Water int `json:"water"`
}




func randWaterWind() (int, int) {
	return rand.Intn(100) + 1, rand.Intn(100) + 1
}

func generateJSON(){ 
	water, wind := randWaterWind()

	data := Response{
		Status: Data{
			Wind: wind,
			Water: water,
		},
	}

	jsonByte, err := json.Marshal(&data)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = ioutil.WriteFile("data.json", jsonByte, 0644)
	fmt.Println(string(jsonByte))

}

func transformStatus(water int, wind int) (rWater string,rWind string){

	if water <= 5{
		rWater = "aman"
	} else if water >= 6 && water <= 8 {
		rWater = "siaga"
	} else {
		rWater = "bahaya"
	}

	if wind <= 6 {
		rWind = "aman"
	} else if wind >= 7 && wind <= 15 {
		rWind = "siaga"
	} else {
		rWind = "bahaya"
	}

	rWater += " (" + strconv.Itoa(water) + " meter)"
	rWind += " (" + strconv.Itoa(wind) + " meter/detik)"

	return
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var res Response

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	json.Unmarshal(byteValue, &res)

	

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	water, wind := transformStatus(res.Status.Water, res.Status.Wind)
	t.Execute(w, map[string]string {
		"Water" : water,
		"Wind" : wind,
	})
}


func main(){
	http.HandleFunc("/", IndexHandler)


	ticker := time.NewTicker(5*time.Second)
	
	go func ()  {
		for {
			select {
			case <- ticker.C:
				generateJSON()
				
			}
		}
	}()


	fmt.Println("starting server at port 4000");

	http.ListenAndServe(":4000", nil)
}

