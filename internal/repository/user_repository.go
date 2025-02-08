package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"technodom/models"
)

type UserRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewUserRepository(db *sql.DB, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password, role, registration_date)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	r.logger.Info("Creating user", slog.String("username", user.Username))

	err := r.db.QueryRow(query, user.Username, user.Email, user.Password, user.Role, user.RegistrationDate).Scan(&user.ID)
	if err != nil {
		r.logger.Error("Error creating user", slog.String("error", err.Error()))
		return fmt.Errorf("unable to create user: %v", err)
	}

	r.logger.Info("User created", slog.Int("id", user.ID))
	return nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, username, email, password, role, registration_date FROM users WHERE id = $1`
	r.logger.Info("Getting user by ID", slog.Int("id", id))

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.RegistrationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("User not found", slog.Int("id", id))
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		r.logger.Error("Error getting user", slog.String("error", err.Error()))
		return nil, fmt.Errorf("unable to get user: %v", err)
	}

	r.logger.Info("User found", slog.Int("id", user.ID))
	return user, nil
}

// UpdateUser обновляет данные пользователя.
func (r *UserRepository) UpdateUser(updatedUser *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4, registration_date = $5
			  WHERE id = $6`
	r.logger.Info("Updating user", slog.Int("id", updatedUser.ID))

	_, err := r.db.Exec(query, updatedUser.Username, updatedUser.Email, updatedUser.Password, updatedUser.Role,
		updatedUser.RegistrationDate, updatedUser.ID)
	if err != nil {
		r.logger.Error("Error updating user", slog.String("error", err.Error()))
		return fmt.Errorf("unable to update user: %v", err)
	}

	r.logger.Info("User updated", slog.Int("id", updatedUser.ID))
	return nil
}

// DeleteUser удаляет пользователя по ID.
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	r.logger.Info("Deleting user", slog.Int("id", id))

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Error("Error deleting user", slog.String("error", err.Error()))
		return fmt.Errorf("unable to delete user: %v", err)
	}

	r.logger.Info("User deleted", slog.Int("id", id))
	return nil
}

// GetAllUsers возвращает всех пользователей.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, username, email, password, role, registration_date FROM users`
	r.logger.Info("Getting all users")

	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Error("Error getting users", slog.String("error", err.Error()))
		return nil, fmt.Errorf("unable to get users: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.RegistrationDate)
		if err != nil {
			r.logger.Error("Error scanning user", slog.String("error", err.Error()))
			return nil, fmt.Errorf("unable to scan user: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error iterating rows", slog.String("error", err.Error()))
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	r.logger.Info("Retrieved users", slog.Int("count", len(users)))
	return users, nil
}
