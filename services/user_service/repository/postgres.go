package repository

import (
	"user_service/domain"
	"database/sql"
	"fmt"
	"time"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) UserRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByEmail(email string) (*domain.User, error) {
	query := `
		SELECT 
			u.id, u.full_name, u.email, u.password, u.created_at,
			p.first_name, p.last_name, p.birth_date, p.address,
			r.role
		FROM users u
		LEFT JOIN profiles p ON u.id = p.user_id
		LEFT JOIN roles r ON u.id = r.user_id
		WHERE u.email = $1
		LIMIT 1
	`

	user, err := r.scanUser(query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	return user, nil
}

func (r *postgresRepository) GetByID(id int64) (*domain.User, error) {
	query := `
		SELECT 
			u.id, u.full_name, u.email, u.password, u.created_at,
			p.first_name, p.last_name, p.birth_date, p.address,
			r.role
		FROM users u
		LEFT JOIN profiles p ON u.id = p.user_id
		LEFT JOIN roles r ON u.id = r.user_id
		WHERE u.id = $1
		LIMIT 1
	`

	user, err := r.scanUser(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}

	return user, nil
}

func (r *postgresRepository) Create(user *domain.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	userQuery := `
		INSERT INTO users (full_name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err = tx.QueryRow(
		userQuery,
		user.FullName,
		user.Email,
		user.Password,
		user.CreatedAt,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	profileQuery := `
		INSERT INTO profiles (user_id, first_name, last_name, birth_date, address)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			first_name = EXCLUDED.first_name,
			last_name = EXCLUDED.last_name,
			birth_date = EXCLUDED.birth_date,
			address = EXCLUDED.address
		`

	_, err = tx.Exec(
		profileQuery,
		user.ID,
		user.Profile.FirstName,
		user.Profile.LastName,
		user.Profile.BirthDate,
		user.Profile.Address,
	)

	if err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	roleQuery := `
		INSERT INTO roles (user_id, role)
		VALUES ($1, $2)
		ON CONFLICT (user_id, role) DO NOTHING
	`

	_, err = tx.Exec(roleQuery, user.ID, string(user.Role))
	if err != nil {
		return fmt.Errorf("failed to assign role: %w", err)
	}

	return tx.Commit()
}

func (r *postgresRepository) GetAll() ([]*domain.User, error) {
	query := `
		SELECT 
			u.id, u.full_name, u.email, u.password, u.created_at,
			p.first_name, p.last_name, p.birth_date, p.address,
			r.role
		FROM users u
		LEFT JOIN profiles p ON u.id = p.user_id
		LEFT JOIN roles r ON u.id = r.user_id
		ORDER BY u.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user, err := r.scanUserFromRows(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}

func (r *postgresRepository) GetUserWithProfile(userID int64) (*domain.User, error) {
	return r.GetByID(userID)
}

func (r *postgresRepository) EmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}
	
	return exists, nil
}

func (r *postgresRepository) scanUser(query string, args ...interface{}) (*domain.User, error) {
	var user domain.User
	var firstName, lastName, address sql.NullString
	var birthDate sql.NullTime
	var role sql.NullString

	err := r.db.QueryRow(query, args...).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&firstName,
		&lastName,
		&birthDate,
		&address,
		&role,
	)

	if err != nil {
		return nil, err
	}

	user.Profile = domain.Profile{
		FirstName: firstName.String,
		LastName:  lastName.String,
		Address:   address.String,
	}

	if birthDate.Valid {
		user.Profile.BirthDate = birthDate.Time
	}

	if role.Valid {
		user.Role = domain.Role(role.String)
	}

	return &user, nil
}

func (r *postgresRepository) scanUserFromRows(rows *sql.Rows) (*domain.User, error) {
	var user domain.User
	var firstName, lastName, address sql.NullString
	var birthDate sql.NullTime
	var role sql.NullString

	err := rows.Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&firstName,
		&lastName,
		&birthDate,
		&address,
		&role,
	)

	if err != nil {
		return nil, err
	}

	user.Profile = domain.Profile{
		FirstName: firstName.String,
		LastName:  lastName.String,
		Address:   address.String,
	}

	if birthDate.Valid {
		user.Profile.BirthDate = birthDate.Time
	}

	if role.Valid {
		user.Role = domain.Role(role.String)
	}

	return &user, nil
}
