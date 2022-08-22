package main

type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

//asi se haria lo mas paresido al encapsulamiento
//Ejemplo de Encapsulamiento
/*
movie := new(Movie);
movie.setName("CUALQUIER_COSA");
*/

func (this *Movie) setName(name string) {
	this.Name = name
}

type Movies []Movie