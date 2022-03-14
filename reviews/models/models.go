package models

type Alumno struct {
	Dni string
	Nombre string
	Apellido string
	Carrera string
}

type CreateAlumno struct{

	Dni string `json:"dni"`
	Nombre string `json:"nombre`
	Apellido string `json:"apellido`
	Carrera string `json:"carrera`
}
