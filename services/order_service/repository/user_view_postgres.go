package repository

import "database/sql"

type UserViewPostgres struct {
	db *sql.DB
}

func NewUserViewPostgres(db *sql.DB) *UserViewPostgres {
	return &UserViewPostgres{db: db}
}

func (r *UserViewPostgres) Insert(userID int64) error {
	_, err := r.db.Exec(
		`INSERT INTO user_view (user_id)
		 VALUES ($1)
		 ON CONFLICT (user_id) DO NOTHING`,
		userID,
	)
	return err
}
func (r *UserViewPostgres) Exists(userID int64) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		`SELECT EXISTS (SELECT 1 FROM user_view WHERE user_id = $1)`,
		userID,
	).Scan(&exists)

	return exists, err
}
