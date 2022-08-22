package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Lector");

	nuevo_texto := os.Args[1]+"\n"
	//escribir
	fichero, err := os.OpenFile("fichero.txt",os.O_APPEND,0777);
	showErro(err);
	file,err := fichero.WriteString(nuevo_texto);
	showErro(err);
	fmt.Println(file);
	fichero.Close();
	
	

	//leer
	fileArc,errorFichero := ioutil.ReadFile("fichero.txt");
	showErro(errorFichero);
	fmt.Println(string(fileArc));
	

}


func showErro(e error)  {
	if(e != nil){
		panic(e);
	}
	
}