package mysql

import (
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"
	"database/sql"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (s *orderRepo) GetOrders() (*[]models.Order, error) {
	query := `
		SELECT 
			order_id, customer_name, ordered_at
		FROM orders
		ORDER BY order_id ASC
	`

	// proses query data
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	// matikan koneksi
	defer rows.Close()

	var orders []models.Order

	// proses scan data
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerName,
			&order.OrderedAt,
			// &order.Grade,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return &orders, nil
}



func (s *orderRepo) CreateOrder(order *models.Order) error {
	query := `
		INSERT INTO orders (customer_name,ordered_at)
		VALUES ($1, $2, $3)
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.CustomerName, order.OrderedAt)

	return err
}



func (s *orderRepo) UpdateOrder(order *models.Order) ( error) {
	query := `
		UPDATE orders
		SET 
		customer_name = $1
		ordered_at = $2
		
		WHERE order_id = $3;
	`

	// proses query data
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	// matikan koneksi
	defer stmt.Close()

	// proses scan data
	_, err = stmt.Exec(order.CustomerName, order.OrderedAt, order.OrderID)

	return nil
}


func (s *orderRepo) DeleteOrder(id *string) (error) {
	query := `
		DELETE FROM orders
		
		WHERE order_id = $1;
	`

	// proses query data
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	// matikan koneksi
	defer stmt.Close()



	// proses scan data
	_, err = stmt.Exec(id)

	return  nil
}
