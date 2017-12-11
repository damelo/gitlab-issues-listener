package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Issue Issue
type Issue struct {
}

//IssueHandler IssueHandler
func IssueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IssueHandler...")

	body, err := ioutil.ReadAll(r.Body)
	check(err)

	log.Println(string(body))
	var issue Issue

	err = json.Unmarshal(body, &issue)
	check(err)

	//node := r.URL.Path[len("/cisreport/"):]
	//p := extractReportFromFile(node, config.Dir)
	//json.NewEncoder(w).Encode//(p) //write json to
	//extractReportFromFile(node,)

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(p)
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

func main() {
	fmt.Println("Iniciando...")
	//loadConfig("config.yml")
	//fmt.Println("config dir: ", config.Dir)

	//http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/put/issue/", IssueHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)

}
