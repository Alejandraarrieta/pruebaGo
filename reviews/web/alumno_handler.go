package web

import(
	"encoding/json"
	"github.com/Alejandraarrieta/pruebaGo/interno/database"
    "github.com/Alejandraarrieta/pruebaGo/reviews/gateway"
    "github.com/Alejandraarrieta/pruebaGo/reviews/models"
	"net/http"
)

func (h *AlumnoHandler) AddAlumnoHandler(w http.ResponseWriter, banana *http.Request) {
	obj := params(banana)
	res, err := h.AddAlumno(obj)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(res))
}

func params(request *http.Request) *models.CreateAlumno{
	var obj models.CreateAlumno

	_ = json.NewDecoder(request.Body).Decode(&obj)

	return &obj
}

type AlumnoHandler struct{
	gateway.AlumnoGateway
}

func NewAlumnoHandler(mongo *database.Mongo) *AlumnoHandler {
	return &AlumnoHandler{gateway.NewAlumnoGateway(mongo)}
}