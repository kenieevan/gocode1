package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type database map[string]int

func main() {
	db := database{"shoes": 50, "socks": 5}

	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/update", http.HandlerFunc(db.update))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

//need do more err check against the input
//unserstant req.Form , why its value is []string??
func (db database) update(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println(err)
	}

	var item string
	var price int
	var err error
	for k, v := range req.Form {
		fmt.Fprintf(os.Stdout, "Form[%q]= %q\n", k, v)
		if k == "item" {
			item = v[0]
		}

		if k == "price" {
			price, err = strconv.Atoi(v[0])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "invalid price")
			}
		}

		//item doesn't exist
		if _, ok := db[item]; !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such items")
			return
		}

		db[item] = price
	}

}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s :%s\n", item, price)
	}
}
