package repository

import (
	"database/sql"

	"github.com/gingenericssql/model"
)

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id int) (T, error)
	Create(entity T) (T, error)
	Update(id int, entity T) (T, error)
	Delete(id int) error
}
type PostgresRepository[T any] struct {
	DB *sql.DB

	TableName string

	ScanRow  func(*sql.Row) (T, error)
	ScanRows func(*sql.Rows) (T, error)

	Insert   func(*sql.DB, T) (T, error)
	UpdateFn func(*sql.DB, int, T) (T, error)
}

func (r *PostgresRepository[T]) GetAll() ([]T, error) {
	rows, err := r.DB.Query("SELECT * FROM " + r.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []T
	for rows.Next() {
		item, err := r.ScanRows(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}
func (r *PostgresRepository[T]) GetByID(id int) (T, error) {
	return r.ScanRow(r.DB.QueryRow("SELECT * FROM "+r.TableName+" WHERE id=$1", id))
}
func (r *PostgresRepository[T]) Create(entity T) (T, error) {
	return r.Insert(r.DB, entity)
}
func (r *PostgresRepository[T]) Update(id int, entity T) (T, error) {
	return r.UpdateFn(r.DB, id, entity)
}

func (r *PostgresRepository[T]) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM "+r.TableName+" WHERE id=$1", id)
	return err
}

//factory apis -  generic interface and structure implemenation
//now we have implemented with Employee, later you can with other types like products

func NewEmployeeRepo(db *sql.DB) *PostgresRepository[model.Employee] {
	return &PostgresRepository[model.Employee]{
		DB:        db,
		TableName: "employees",

		ScanRow: func(row *sql.Row) (model.Employee, error) {
			var e model.Employee
			err := row.Scan(&e.ID, &e.Name, &e.City)
			return e, err
		},

		ScanRows: func(rows *sql.Rows) (model.Employee, error) {
			var e model.Employee
			err := rows.Scan(&e.ID, &e.Name, &e.City)
			return e, err
		},

		Insert: func(db *sql.DB, e model.Employee) (model.Employee, error) {
			err := db.QueryRow(
				"INSERT INTO employees(name, city) VALUES($1,$2) RETURNING id",
				e.Name, e.City,
			).Scan(&e.ID)
			return e, err
		},

		UpdateFn: func(db *sql.DB, id int, e model.Employee) (model.Employee, error) {
			_, err := db.Exec(
				"UPDATE employees SET name=$1, city=$2 WHERE id=$3",
				e.Name, e.City, id,
			)
			e.ID = id
			return e, err
		},
	}
}
