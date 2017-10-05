package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
	"admin/routes"
	"admin/tools"
)

func main(){
	var staticDir string

	flag.StringVar(&staticDir, "dir", "./static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := routes.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	var myIP = tools.GetOutboundIP()
	fmt.Println("Web interface is ENABLE on",myIP,"at 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
