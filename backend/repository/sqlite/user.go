package sqlite

import (
	"fmt"
	"time"

	"backend/domain/models"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user in the database
func (db *DB) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	query := `
		INSERT INTO users (
			id, username, email, password_hash,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = db.Exec(query,
		user.ID,
		user.Username,
		user.Email,
		string(hashedPassword),
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

// UpdateUser updates an existing user in the database
func (db *DB) UpdateUser(user *models.User) error {
	query := `
		UPDATE users 
		SET username = ?, email = ?, updated_at = ?
		WHERE id = ?
	`

	user.UpdatedAt = time.Now()
	_, err := db.Exec(query,
		user.Username,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

// UpdatePassword updates a user's password
func (db *DB) UpdatePassword(userID, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	query := `
		UPDATE users 
		SET password_hash = ?, updated_at = ?
		WHERE id = ?
	`

	_, err = db.Exec(query, string(hashedPassword), time.Now(), userID)
	if err != nil {
		return fmt.Errorf("error updating password: %v", err)
	}

	return nil
}

// GetUserByID retrieves a user by their ID
func (db *DB) GetUserByID(id string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	user := &models.User{}
	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email address
func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	user := &models.User{}
	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting user by email: %v", err)
	}

	return user, nil
}

// GetUserByUsername retrieves a user by their username
func (db *DB) GetUserByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE username = ?
	`

	user := &models.User{}
	err := db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting user by username: %v", err)
	}

	return user, nil
}

// DeleteUser deletes a user from the database
func (db *DB) DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}

// ValidatePassword checks if the provided password matches the stored hash
func (db *DB) ValidatePassword(userID, password string) (bool, error) {
	var hashedPassword string
	query := "SELECT password_hash FROM users WHERE id = ?"

	err := db.QueryRow(query, userID).Scan(&hashedPassword)
	if err != nil {
		return false, fmt.Errorf("error getting password hash: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, fmt.Errorf("error comparing passwords: %v", err)
	}

	return true, nil
}
