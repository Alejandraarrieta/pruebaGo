package gateway

import(
"github.com/pruebaapi/general/interno/database"
"github.com/pruebaapi/general/reviews/models"
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