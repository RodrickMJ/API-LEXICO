package domain

type Cliente struct {
	ClaveCliente string `json:"clave_cliente"`
	Nombre       string `json:"nombre"`
	Correo       string `json:"correo"`
	Telefono     string `json:"telefono"`
}
