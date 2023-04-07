package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Fengshui struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func request(water int, wind int) Fengshui {
	payload, err := json.Marshal(Fengshui{
		Water: water,
		Wind:  wind,
	})
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var fengshui Fengshui
	json.Unmarshal(result, &fengshui)
	return fengshui
}

func checkFengshui(fengshui Fengshui) (string, string) {
	var waterStatus, windStatus string
	if fengshui.Water < 5 {
		waterStatus = "aman"
	} else if fengshui.Water >= 6 && fengshui.Water <= 8 {
		waterStatus = "siaga"
	} else if fengshui.Water > 8 {
		waterStatus = "bahaya"
	}

	if fengshui.Wind < 6 {
		windStatus = "aman"
	} else if fengshui.Wind >= 7 && fengshui.Wind <= 15 {
		windStatus = "siaga"
	} else if fengshui.Wind > 15 {
		windStatus = "bahaya"
	}

	return waterStatus, windStatus
}

func main() {
	log.Println("Start")
	for i := 0; i < 5; i++ {
		time.Sleep(15 * time.Second)
		fengshui := request(rand.Intn(100), rand.Intn(100))
		waterStatus, windStatus := checkFengshui(fengshui)

		data, err := json.MarshalIndent(fengshui, "", "  ")
		if err != nil {
			log.Println(err)
		}
		log.Printf("\n%s\nstatus water : %s\nstatus wind : %s", string(data), waterStatus, windStatus)
	}
	log.Println("End")
}
