package application

import "api_go/internal/domain"

type ClienteService struct {
	repo ClienteRepository
}

func NewClienteService(repo ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Crear(c domain.Cliente) error        { return s.repo.Crear(c) }
func (s *ClienteService) Listar() ([]domain.Cliente, error)   { return s.repo.Listar() }
func (s *ClienteService) Buscar(clave string) (domain.Cliente, error) {
	return s.repo.Buscar(clave)
}
func (s *ClienteService) Actualizar(c domain.Cliente) error   { return s.repo.Actualizar(c) }
func (s *ClienteService) Eliminar(clave string) error         { return s.repo.Eliminar(clave) }
