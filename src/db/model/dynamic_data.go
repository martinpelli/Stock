package model

type Product struct {
	Id           int32  `json:"id"`
	Nombre       string `json:"nombre"`
	IdContenedor int32  `json:"id_contenedor"`
	Cantidad     int16  `json:"cantidad"`
}
