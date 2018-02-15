// main.go

package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "strconv"
  
)


type archivo struct {
  ID      int    `json:"id"`
  Nombre   string `json:"title"`
  Formato string `json:"content"`
}


var articleList = []archivo{
  archivo{ID: 1, Nombre: "Documento Word", Formato: ".docx"},
  archivo{ID: 2, Nombre: "Caciòn", Formato: ".mp4"},
  archivo{ID: 3, Nombre: "Documento Excel", Formato: ".xls"},
}

func getAllArticles() []archivo{
  return articleList
}

// Return a list of all the articles
func getArticleById(articles []archivo, id int) archivo{

  emergency := archivo{ID: 0, Nombre: "no", Formato: "no"}
  for _, a := range articles{
    if (a.ID == id) {
      fmt.Println(a.ID, a.Nombre, a.Formato)
      return a

    }
  }
  return emergency
}
func converToId(value string) (int, error) {
  intValue,err := strconv.Atoi(value)

  return intValue, err 
}

func addArchivo(id int, nombre, formato string ){
  
}

var router *gin.Engine

func main() {

  router = gin.Default()
  router.LoadHTMLGlob("templates/*")
  article := getAllArticles()
  // INDEX ROUTE
 router.GET("/archivos", func(c *gin.Context){
      c.HTML(http.StatusOK, "archivos.html",gin.H{
        "title": "Pagina Inicio",
        "payload": article,
      },   
    )
    fmt.Println("SOMEONE REQUEST /archivo DIRECTORY")

    })


 //NEW ROUTE 
  router.GET("/archivo/new", func(c *gin.Context){
      c.HTML(http.StatusOK, "nuevoArchivo.html",gin.H{
        "title": "Pagina Inicio",
      },   
    )
    fmt.Println("nuevo archivo ")

    })

  // CREATE ROUTE
  router.POST("/create", func(c *gin.Context){
      id,_ := strconv.Atoi(c.PostForm("ID"))
      nombre := c.PostForm("Nombre")
      formato:= c.PostForm("Formato")
      addArchivo(id,nombre,formato)

      

      c.HTML(http.StatusOK, "archivos.html",gin.H{
        "title": "Pagina Inicio",
      },   
    )
    fmt.Println("you have reached post route!")
    fmt.Println(id,nombre,formato)

    })



  //SHOW ROUTE

    router.GET("/article/:id", func (c *gin.Context) {

      
      intValue,_ := strconv.Atoi(c.Param("id"))
      article := getArticleById(article,intValue)
      fmt.Println(article)
      c.HTML(http.StatusOK,"show.html",
            gin.H{
              "title":   article.Nombre,
              "payload": article,
            },
          )


      
    })
    // EDIT ROUTE
      router.GET("/article/:id/edit", func (c *gin.Context) {

      
      intValue,_ := strconv.Atoi(c.Param("id"))
      article := getArticleById(article,intValue)
      c.HTML(http.StatusOK,"editarArchivo.html",
            gin.H{
              
              "payload": article,
            },
          )


      
    })

  // trabaja en puerto 8080 por default
  router.Run()

}