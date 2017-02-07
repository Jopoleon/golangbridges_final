package main

import (
	"encoding/json"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2"
)

var mongoUrl = "mongodb://egor_m:qwer1234@ds135029.mlab.com:35029/bridges_mlab"

type resBody struct {
	Waterlevel string `json:"waterlevel" bson:"waterlevel"`
	Time       string `json:"time" bson:"time"`
	ShipHight  string `json:"shipHight" bson:"shipHight"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index1.ejs")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index1.ejs", nil)
}

func ScraperHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Reqest from: ", r.Host, r.URL.Path)
	url := "http://spun.fkpkzs.ru/Level/Gorny"
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	time_p := (doc.Find("#waterleveltable td.timestampvalue").First().Text())
	waterlevel_p := (doc.Find("#waterleveltable td.value").First().Text())
	log.Println("Watrelevel = ", waterlevel_p)
	log.Println("Time = ", time_p)
	if r.Body == nil {
		http.Error(w, "Nil body in request", 400)
		return
	}

	r.ParseForm()
	shipHight_p := r.FormValue("value")
	log.Println("Ship hight = ", shipHight_p)
	mapShip := map[string]string{
		"waterlevel": waterlevel_p,
		"time":       time_p,
		"shipHight":  shipHight_p,
	}

	log.Println("Data BEFORE marshaling json: ", mapShip)
	data, err := json.Marshal(mapShip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}

	shipDataCollection := session.DB("bridges_mlab").C("bridgeRequests")
	err = shipDataCollection.Insert(
		&resBody{
			Waterlevel: waterlevel_p,
			Time:       time_p,
			ShipHight:  shipHight_p})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server started on port: ", port)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/scrape", ScraperHandler)

	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.ListenAndServe(":"+port, nil)
}
