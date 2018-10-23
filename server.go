package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

var appBuildDirPtr = flag.String("dir", "static", "file directory location")
var appLocationPtr = flag.String("loc", "/StaticApp/", "location of app under the domain")

func main() {
	flag.Parse()

	fsHandler := http.FileServer(http.Dir(*appBuildDirPtr))

	http.HandleFunc("/greet", greet)

	appLocation := *appLocationPtr
	http.Handle(appLocation, http.StripPrefix(appLocation, fsHandler))

	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
