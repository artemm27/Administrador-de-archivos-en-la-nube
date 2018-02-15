package main 

import (
"fmt"

)

type archivo struct {
  ID      int    `json:"id"`
  Nombre   string `json:"title"`
  Formato string `json:"content"`
}


func main(){

	s := make([]archivo, 0)


	s = append(s, archivo{ID: 1, Nombre: "Documento Word1", Formato: ".docx"})
	s = append(s, archivo{ID: 2, Nombre: "Documento Word2", Formato: ".docx"})
	s = append(s, archivo{ID: 3, Nombre: "Documento Word3", Formato: ".docx"})

	for _, a := range s {
		fmt.Println(a.ID,a.Nombre,a.Formato)
	}




}