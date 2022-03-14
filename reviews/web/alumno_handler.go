package web

import(
	"encoding/json"
	"github.com/pruebaapi/general/interno/database"
	"github.com/pruebaapi/general/reviews/gateway"
	"github.com/pruebaapi/general/reviews/models"
	"net/http"
)

func (h *AlumnoHandler) AddAlumnoHandler(w http.ResponseWriter, r *http.Request){
	obj := params(r)
	res, err := h.AddAlumno(obj)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _= w.Write([]byte(res))
}

func params(r http.Request) *models.CreateAlumno{
	var obj models.CreateAlumno

	_=json.NewDecoder(r.Body).Decode(&obj)

	return &obj
}
type AlumnoHandler struct{
	gateway.AlumnoGateway
}

func NewAlumnoHandler(mongo *database.Mongo) *AlumnoHandler {
	return &AlumnoHandler{gateway.NewAlumnoGateway(mongo)}
}