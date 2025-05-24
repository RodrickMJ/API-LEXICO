package mysql

import (
	"api_go/internal/domain"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepo struct {
	db *sql.DB
}

func NewMySQLRepo(db *sql.DB) *MySQLRepo {
	schema := `
	CREATE TABLE IF NOT EXISTS clientes (
		clave_cliente VARCHAR(50) PRIMARY KEY,
		nombre VARCHAR(100),
		correo VARCHAR(100),
		telefono VARCHAR(20)
	);`
	db.Exec(schema)
	return &MySQLRepo{db}
}

func (r *MySQLRepo) Crear(c domain.Cliente) error {
	_, err := r.db.Exec("INSERT INTO clientes (clave_cliente, nombre, correo, telefono) VALUES (?, ?, ?, ?)",
		c.ClaveCliente, c.Nombre, c.Correo, c.Telefono)
	return err
}

func (r *MySQLRepo) Listar() ([]domain.Cliente, error) {
	rows, err := r.db.Query("SELECT clave_cliente, nombre, correo, telefono FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []domain.Cliente
	for rows.Next() {
		var c domain.Cliente
		rows.Scan(&c.ClaveCliente, &c.Nombre, &c.Correo, &c.Telefono)
		clientes = append(clientes, c)
	}
	return clientes, nil
}

func (r *MySQLRepo) Buscar(clave string) (domain.Cliente, error) {
	var c domain.Cliente
	err := r.db.QueryRow("SELECT clave_cliente, nombre, correo, telefono FROM clientes WHERE clave_cliente = ?", clave).
		Scan(&c.ClaveCliente, &c.Nombre, &c.Correo, &c.Telefono)
	return c, err
}

func (r *MySQLRepo) Actualizar(c domain.Cliente) error {
	_, err := r.db.Exec("UPDATE clientes SET nombre = ?, correo = ?, telefono = ? WHERE clave_cliente = ?",
		c.Nombre, c.Correo, c.Telefono, c.ClaveCliente)
	return err
}

func (r *MySQLRepo) Eliminar(clave string) error {
	_, err := r.db.Exec("DELETE FROM clientes WHERE clave_cliente = ?", clave)
	return err
}
