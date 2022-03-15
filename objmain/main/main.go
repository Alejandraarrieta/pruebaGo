package main
import(

//	_ "github.com/go-sql-driver/mysql"
//	"github.com/golang-migrate/migrate"
//	migration "github.com/golang-migrate/migrate/database/mysql"
//	_ "github.com/golang-migrate/migrate/source/file"
//	"github.com/tomiok/course-phones-review/gadgets/smartphones/web"
//	"github.com/tomiok/course-phones-review/internal/database"
//	"github.com/tomiok/course-phones-review/internal/logs"
//	reviews "github.com/tomiok/course-phones-review/reviews/web"

"github.com/Alejandraarrieta/pruebaGo/interno/logs"
"github.com/Alejandraarrieta/pruebaGo/interno/database"
reviews "github.com/Alejandraarrieta/pruebaGo/reviews/web"


)

func main(){

	_ = logs.InitLogger()

	mongoClient := database.NewMongoClient("localhost")
	alumnoHandler := reviews.NewAlumnoHandler(mongoClient) 

	mux := Routes(alumnoHandler)
	server := NewServer(mux)
	server.Run()
}