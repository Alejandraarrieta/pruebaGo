package main
import(

//	_ "github.com/go-sql-driver/mysql"
//	"github.com/golang-migrate/migrate"
//	migration "github.com/golang-migrate/migrate/database/mysql"
//	_ "github.com/golang-migrate/migrate/source/file"
//	"github.com/tomiok/course-phones-review/gadgets/smartphones/web"
"github.com/pruebaapi/general/interno/database"
"github.com/pruebaapi/general/interno/logs"
//	reviews "github.com/tomiok/course-phones-review/reviews/web"
"/../web/alumno_handler.go"



)

func main(){

	mongoClient := database.NweMongoClient("localhost")
	alumnoHandler := models.NewAlumnoHandler(mongoClient) 

	server := NewServer(mux)
	server.Run()
}