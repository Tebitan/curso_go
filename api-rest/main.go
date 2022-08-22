package main

import (
	"log"
	"net/http"
)

func main(){
	//router
	router := NewRouter();
	//server
	server := http.ListenAndServe(":8080",router);
	log.Fatal(server);


	/*
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola soy GO");
	});
	server := http.ListenAndServe(":8080",nil)
	log.Fatal(server);*/

} 
