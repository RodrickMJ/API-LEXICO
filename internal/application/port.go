package application

import "api_go/internal/domain"

type ClienteRepository interface {
	Crear(cliente domain.Cliente) error
	Listar() ([]domain.Cliente, error)
	Buscar(clave string) (domain.Cliente, error)
	Actualizar(cliente domain.Cliente) error
	Eliminar(clave string) error
}
