package gateway
import(
	"github.com/pruebaapi/general/reviews/models"
)
func(g *AlumnoGtw) AddAlumno(obj *models.CreateAlumno) (string, error){
	return g.Add(obj)
}