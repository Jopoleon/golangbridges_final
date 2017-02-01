package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"html/template"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index1.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index1.html", nil)
}
func scraperHandler(w http.ResponseWriter, r *http.Request) {

	url := "http://spun.fkpkzs.ru/Level/Gorny"

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
	fmt.Println("Watrelevel = ", waterlevel_p)
	fmt.Println("Time = ", time_p)
	if r.Body == nil {
		http.Error(w, "Nil body in request", 400)
		return
	}

	r.ParseForm()
	shipHight_p := r.FormValue("value")
	fmt.Println("Ship hight = ", shipHight_p)
	mapShip := map[string]string{
		"waterlevel": waterlevel_p,
		"time":       time_p,
		"shipHight":  shipHight_p,
	}

	fmt.Println("Data BEFORE marshaling json: ", mapShip)
	data, err := json.Marshal(mapShip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Data AFTER marshaling json: ", data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func main() {
	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = ":3000"
	}
	fmt.Println(os.Getenv("PORT"))

	//var port = ":3000"
	fmt.Println("Server started on port: ", port)
	//fmt.Println("Server started on port: 3000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/scrape", scraperHandler)

	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.ListenAndServe(port, nil)
}
