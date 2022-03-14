package gateway

import(
	"github.com/Alejandraarrieta/pruebaGo/tree/main/interno/database"
	"github.com/Alejandraarrieta/pruebaGo/tree/main/reviews/models"
)
type AlumnoGateway interface {
	AddAlumno(obj *models.CreateAlumno) (string, error)
}

type AlumnoGtw struct {
	AlumnoStorage
}

func NewAlumnoGateway(mongoClient *database.Mongo) AlumnoGateway{
	return &AlumnoGtw{&AlumnoStg{mongoClient}}
}