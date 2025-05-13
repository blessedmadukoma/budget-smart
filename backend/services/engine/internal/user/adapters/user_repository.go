package adapters

import (
	"database/sql"
	"fmt"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

//	func (r UserRepository) Create(email string) error {
//		return nil
//	}
func (r UserRepository) Create(user *model.User) error {
	r.GetByEmail(user.Email)

	fmt.Println("User details:", user.FirstName, user.LastName, user.Email, user.Password)

	return nil
}

// 	// store in db
// 	fmt.Println("User details:", payload.FirstName, payload.LastName, payload.Email, payload.Password)

// 	// Generate a unique ID for the user
// 	uid := uuid.New().String()

// 	// Hash password if using local authentication
// 	var hashedPassword string
// 	var err error

// 	if payload.AuthProvider == "local" {
// 		hashedBytes, err := password.HashPassword(payload.Password)
// 		if err != nil {
// 			return fmt.Errorf("failed to hash password: %w", err)
// 		}
// 		hashedPassword = string(hashedBytes)
// 	}

// 	// Begin transaction
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %w", err)
// 	}
// 	defer tx.Rollback()

// 	// Prepare the query based on auth provider
// 	var query string
// 	var args []interface{}

// 	if payload.AuthProvider == "local" {
// 		query = `
// 			INSERT INTO users (
// 				uid, first_name, last_name, email, password,
// 				status, auth_provider, created_at, updated_at
// 			) VALUES (
// 				$1, $2, $3, $4, $5, $6, $7, $8, $9
// 			) RETURNING uid`

// 		args = []interface{}{
// 			uid,
// 			strings.TrimSpace(payload.FirstName),
// 			strings.TrimSpace(payload.LastName),
// 			strings.TrimSpace(payload.Email),
// 			hashedPassword,
// 			"PENDING",
// 			"local",
// 			time.Now(),
// 			time.Now(),
// 		}
// 	} else if payload.AuthProvider == "google" {
// 		// Query for Google authentication
// 		query = `
// 			INSERT INTO users (
// 				uid, first_name, last_name, email,
// 				status, auth_provider, google_id, google_token,
// 				created_at, updated_at
// 			) VALUES (
// 				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
// 			) RETURNING uid`

// 		args = []interface{}{
// 			uid,
// 			strings.TrimSpace(payload.FirstName),
// 			strings.TrimSpace(payload.LastName),
// 			strings.TrimSpace(payload.Email),
// 			"ACCEPTED",
// 			"google",
// 			payload.GoogleID,
// 			payload.GoogleToken,
// 			time.Now(),
// 			time.Now(),
// 		}
// 	} else {
// 		return fmt.Errorf("unsupported auth provider: %s", payload.AuthProvider)
// 	}

// 	// Execute the query
// 	var returnedUID string
// 	err = tx.QueryRow(query, args...).Scan(&returnedUID)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			// Check for unique constraint violation (email already exists)
// 			if pqErr.Code == "23505" {
// 				return fmt.Errorf("user with this email already exists")
// 			}
// 		}
// 		return fmt.Errorf("failed to create user: %w", err)
// 	}

// 	// Commit the transaction
// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("failed to commit transaction: %w", err)
// 	}

// 	return nil
// }

func (r UserRepository) GetByID(id uint) (*model.User, error) {
	return nil, nil
}

func (r UserRepository) GetByEmail(email string) (*model.User, error) {
	return nil, nil
}
