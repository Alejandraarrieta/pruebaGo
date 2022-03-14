package main
import(

//	_ "github.com/go-sql-driver/mysql"
//	"github.com/golang-migrate/migrate"
//	migration "github.com/golang-migrate/migrate/database/mysql"
//	_ "github.com/golang-migrate/migrate/source/file"
//	"github.com/tomiok/course-phones-review/gadgets/smartphones/web"
"github.com/Alejandraarrieta/pruebaGo/tree/main/interno/database"
"github.com/Alejandraarrieta/pruebaGo/tree/main/interno/logs"
	reviews "github.com/Alejandraarrieta/pruebaGo/tree/main/reviews/web"
//"github.com/Alejandraarrieta/pruebaGo/tree/main/reviews/web"



)

func main(){

	mongoClient := database.NweMongoClient("localhost")
	alumnoHandler := models.NewAlumnoHandler(mongoClient) 

	server := NewServer(mux)
	server.Run()
}