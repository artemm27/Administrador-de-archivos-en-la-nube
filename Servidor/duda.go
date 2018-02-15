package main 

import (
  "fmt"
 
 
)


type archivo struct {
  ID      int    `json:"id"`
  Nombre   string `json:"title"`
  Formato string `json:"content"`
}

func changeValue( archivos []archivo){
	fmt.Println(archivos)
	archivos[0].ID = 5
	
}
func checkvalue(archivos []archivo){
	fmt.Println(archivos)
}

var articleList = make([]archivo,0)


func main() {


articleList = append(articleList, archivo{ID: 1, Nombre: "Documento Word1", Formato: ".docx"})
articleList = append(articleList, archivo{ID: 2, Nombre: "Documento Word2", Formato: ".docx"})
articleList = append(articleList, archivo{ID: 3, Nombre: "Documento Word3", Formato: ".docx"})



/*fmt.Println(archivos)
fmt.Println(&archivos)
fmt.Println(*&archivos)
*&archivos.ID = 500

fmt.Println(archivos)*/
changeValue(articleList)
fmt.Println(articleList)
checkvalue(articleList)








}