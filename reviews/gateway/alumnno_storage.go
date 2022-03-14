package gateway

//Para hacer las consultas

import(
	"context"
	"github.com/Alejandraarrieta/pruebaGo/tree/main/interno/database"
	"github.com/Alejandraarrieta/pruebaGo/tree/main/interno/logs"
	"github.com/Alejandraarrieta/pruebaGo/tree/main/reviews/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AlumnoStorage interface {
	Add(obj *models.CreateAlumno) (string, error)
}

type AlumnoStg struct{
	*database.Mongo
}

func (s *AlumnoStg) Add(obj *CreateAlumno) (string, error){
	coll := s.Client.Database("alumnoDB").Collection("alumnos")

	res, err := coll.InsertOne(context.Background(),
		bson.D{
			{"dni", obj.Dni},
			{"nombre", obj.Nombre},
			{"apellido", obj.Apellido},
			{"carrera", obj.Carrera},
		})

	if err != nil {
		logs.Log().Error("cannot insert review")
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)

	return id.String(), nil
}
