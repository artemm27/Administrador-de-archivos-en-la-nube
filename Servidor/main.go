// main.go

package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "strconv"
  _ "github.com/go-sql-driver/mysql" //my sql driver
  "database/sql"
)


type archivo struct {
  ID      int    `json:"id"`
  Nombre   string `json:"title"`
  Formato string `json:"content"`
}

//Función para verificar errores de mySQL
func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}


/*var articleList = []archivo{
  archivo{ID: 1, Nombre: "Documento Word", Formato: ".docx"},
  archivo{ID: 2, Nombre: "Caciòn", Formato: ".mp4"},
  archivo{ID: 3, Nombre: "Documento Excel", Formato: ".xls"},
}*/

var articleList = make([]archivo,0)
//Obtiene los archivos del slice.
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

// Convierte un string a un Int
func converToId(value string) (int, error) {
  intValue,err := strconv.Atoi(value)

  return intValue, err 
}
// Agrega un archivo a nuestro slice de archivos
func addArchivo(id int, nombre, formato string ){

  articleList = append(articleList, archivo{ID: id, Nombre: nombre, Formato: formato})
  fmt.Println(articleList)

}
// Modifica un Archivo dado un Id, nombre, y un Formato
func modificarArchivo(articles []archivo ,id int, nombre, formato string ){
  fmt.Println("you have reached me")
  articless := getAllArticles()
  fmt.Println("all articles", articles)
   for i, a := range articles{
    if (a.ID == id) {
      fmt.Println("no modiciado", a)
      articles[i].Nombre = nombre
      articles[i].Formato = formato
      fmt.Println("modiciado", a)
  fmt.Println("again articles", articless)
    }
  }

}

var router *gin.Engine

func main() {

  db, err := sql.Open("mysql", "root:1718@/GoFiles")
  checkErr(err)

  router = gin.Default()
  router.LoadHTMLGlob("templates/*")
  
  // INDEX ROUTE
 router.GET("/archivos", func(c *gin.Context){
      articles := getAllArticles()
      fmt.Println(articles)
      c.HTML(http.StatusOK, "archivos.html",gin.H{
        "title": "Pagina Inicio",
        "payload": articles,
      },   
    )
    fmt.Println("SOMEONE REQUEST /archivo DIRECTORY")

    })

 //RUTA UNICAMENTE DE PRUEBA.
  router.GET("/ejemplo", func(c *gin.Context){
      articles := getAllArticles()
      fmt.Println(articles)
      c.HTML(http.StatusOK, "ejemplo.html",gin.H{
        "title": "Pagina Inicio",
        "payload": articles,
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

      

      c.Redirect(302, "/archivos")
    fmt.Println("you have reached post route!")
    fmt.Println(id,nombre,formato)

    })



  //SHOW ROUTE

    router.GET("/article/:id", func (c *gin.Context) {

      articles := getAllArticles()
      intValue,_ := strconv.Atoi(c.Param("id"))
      article := getArticleById(articles,intValue)
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

      articles := getAllArticles()
      intValue,_ := strconv.Atoi(c.Param("id"))
      article := getArticleById(articles,intValue)
      c.HTML(http.StatusOK,"editarArchivo.html",
            gin.H{
              
              "payload": article,
            },
          )      
    })
      //RUTA UNICAMENTE DE PRUEBA
      router.PUT("/fruta", func (c *gin.Context) {

      fmt.Println("you have reached pUT REQUEST")
      c.HTML(http.StatusOK,"fruta.html",
            gin.H{
              
              "payload": "hola",
            },
          )      
    })
     //UPDATE ROUTE 

      router.POST("/article/:id", func(c *gin.Context){
            articles := getAllArticles()
            id,_ := strconv.Atoi(c.PostForm("ID"))
            nombre := c.PostForm("Nombre")
            formato:= c.PostForm("Formato")
            fmt.Println("you have reached the PUT ROUTE")
            modificarArchivo(articles,id,nombre,formato)
           c.Redirect(302, "/archivos")
        })


      router.GET("/", func(c *gin.Context){
      c.HTML(http.StatusOK, "login.html",gin.H{
        "title": "Login",
        },   
      )
      fmt.Println("nuevo usuario")

      })

    router.POST("/login", func(c *gin.Context){
        userName := c.PostForm("uname")
        password := c.PostForm("psw")
        fmt.Println("USUARIO")
        fmt.Println(userName)
        fmt.Println(password)
        
        //insertar en la base de datos
        stmt, err := db.Prepare("INSERT Users SET UserId=?, Passw=?, UserPath=?") 
        checkErr(err)
        res, err := stmt.Exec(userName, password, "/home/GoProj/" + userName)
        checkErr(err)
        id, err := res.LastInsertId()
        checkErr(err)
        fmt.Println(id)

        c.Redirect(302, "/archivos")
        //fmt.Println("you have reached post route!")

      })

  // trabaja en puerto 8080 por default
  router.Run()

}