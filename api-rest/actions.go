package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	//conexion a mongoDB
	session, err :=mgo.Dial("mongodb://localhost")
	//validamos si hay error 
	if err !=nil{
		panic(err);
	}

	return session;
}

//indicamos la BD y la coleccion(tabla)
var collectionMovie = getSession().DB("curso_go").C("movies")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola soy GO");
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	//objeto de respuesta
	var results []Movie;
	//Instaciamos session BD y listamos
	err := collectionMovie.Find(nil).All(&results);
	if(err !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(results);
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	//captura parametro
	params := mux.Vars(r);
	movieId := params["id"];
	//objeto de respuesta
	result := Movie{};
	//Instaciamos session BD y buscamos por id
	err := collectionMovie.FindId(bson.ObjectIdHex(movieId)).One(&result);
	if(err !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", err)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(result);
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	//capturamos el BODY de la request
	decoder := json.NewDecoder(r.Body);
	//variable que captura el objeto del BODY de la request
	var movie_data Movie;
	//deceodifica y asigna a la variable movie_data
	err := decoder.Decode(&movie_data);
	//validamos si existe error
	if(err !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", err)
		return
	}
	//cerramos la lectura de la request 
	defer r.Body.Close();
	//indicamos la BD y la coleccion(tabla) . Insertamos
	erro := collectionMovie.Insert(movie_data);
	if(erro !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", erro)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movie_data);
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	//captura parametro
	params := mux.Vars(r);
	movieId := params["id"];
	//capturamos el BODY de la request
	decoder := json.NewDecoder(r.Body);
	//variable que captura el objeto del BODY de la request
	var movie_data Movie;
	//deceodifica y asigna a la variable movie_data
	err := decoder.Decode(&movie_data);
	//validamos si existe error
	if(err !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", err)
		return
	}
	//cerramos la lectura de la request 
	defer r.Body.Close();
	//Instaciamos session BD y modificamos
	erro := collectionMovie.UpdateId(bson.ObjectIdHex(movieId),movie_data)
	if(erro !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", erro)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movie_data);
}

func MovieDelete(w http.ResponseWriter, r *http.Request) {
	//captura parametro
	params := mux.Vars(r);
	movieId := params["id"];
	//objeto de respuesta
	result := Movie{Name:movieId };
	//Instaciamos session BD y buscamos por id
	err := collectionMovie.RemoveId(bson.ObjectIdHex(movieId));
	if(err !=nil){
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%v", err)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(result);
}