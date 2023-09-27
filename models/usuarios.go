package models

type usuario struct {
	Id_usuario      int    `json:"id_usuario"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	SegundoApellido string `json:"segundo_apellido"`
	Email           string `json:"email"`
	Rut             string `json:"rut"`
	Fono            string `json:"fono"`
}

var Usuarios = []usuario{
	{
		Id_usuario:      1,
		Nombre:          "John",
		Apellido:        "Doe",
		SegundoApellido: "Smith",
		Email:           "john@example.com",
		Rut:             "123456789",
		Fono:            "555-123-456",
	},
	{
		Id_usuario:      2,
		Nombre:          "Jane",
		Apellido:        "Smith",
		SegundoApellido: "Doe",
		Email:           "jane@example.com",
		Rut:             "987654321",
		Fono:            "555-987-654",
	},
}
