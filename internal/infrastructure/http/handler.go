package http

import (
	"api_go/internal/application"
	"api_go/internal/domain"
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	service *application.ClienteService
}

func NewHandler(s *application.ClienteService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/clientes", h.handleClientes)
	mux.HandleFunc("/clientes/", h.handleCliente)
}

func (h *Handler) handleClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		clientes, err := h.service.Listar()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(clientes)

	case http.MethodPost:
		var c domain.Cliente
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Formato de solicitud inválido", http.StatusBadRequest)
			return
		}

		if err := h.service.Crear(c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(c)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clave := strings.TrimPrefix(r.URL.Path, "/clientes/")
	if clave == "" {
		http.Error(w, "Se requiere ID de cliente", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		c, err := h.service.Buscar(clave)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(c)

	case http.MethodPut:
		var c domain.Cliente
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Formato de solicitud inválido", http.StatusBadRequest)
			return
		}

		c.ClaveCliente = clave
		if err := h.service.Actualizar(c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(c)

	case http.MethodDelete:
		if err := h.service.Eliminar(clave); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}