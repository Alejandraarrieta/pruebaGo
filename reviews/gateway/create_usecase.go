package gateway
import(
	"github.com/Alejandraarrieta/pruebaGo/tree/main/reviews/models"
)
func(g *AlumnoGtw) AddAlumno(obj *models.CreateAlumno) (string, error){
	return g.Add(obj)
}