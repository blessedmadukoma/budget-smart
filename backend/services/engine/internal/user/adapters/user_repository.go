package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/model"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/password"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/uuid"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Create(ctx context.Context, user *model.User) error {
	fmt.Println("User details:", user.FirstName, user.LastName, user.Email, user.Password)

	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Ensure rollback if commit fails
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Prepare query and arguments based on auth provider
	query, args, err := r.prepareUserInsertQuery(*user)
	if err != nil {
		return err
	}

	// Execute the query
	var returnedUID string
	err = tx.QueryRowContext(ctx, query, args...).Scan(&returnedUID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return fmt.Errorf("user with this email already exists")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	query := `
		SELECT 
			id, uid, first_name, last_name, email, password, 
			status, auth_provider, google_id, google_token,
			created_at
		FROM users
		WHERE id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by ID: %w", err)
	}
	defer rows.Close()

	return scanRowsIntoUser(rows)
}

// GetByUID retrieves a user by their UID
func (r *UserRepository) GetByUID(ctx context.Context, uid string) (*model.User, error) {
	query := `
		SELECT 
			id, uid, first_name, last_name, email, password, 
			status, auth_provider, google_id, google_token,
			created_at
		FROM users
		WHERE uid = $1`

	rows, err := r.db.QueryContext(ctx, query, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by UID: %w", err)
	}
	defer rows.Close()

	return scanRowsIntoUser(rows)
}

// GetByEmail retrieves a user by their email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT 
			id, uid, first_name, last_name, email, password, 
			status, auth_provider, google_id, google_token,
			created_at
		FROM users
		WHERE email = $1`

	rows, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}
	defer rows.Close()

	return scanRowsIntoUser(rows)
}

func (r *UserRepository) prepareUserInsertQuery(payload model.User) (string, []interface{}, error) {
	uid := uuid.New()

	now := time.Now()

	// Common fields for all auth providers
	args := []interface{}{
		uid,
		strings.TrimSpace(payload.FirstName),
		strings.TrimSpace(payload.LastName),
		strings.TrimSpace(payload.Email),
	}

	// Base query with common fields
	query := `
		INSERT INTO users (
			uid, first_name, last_name, email,
			%s, -- Placeholder for auth-specific fields
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4,
			%s, -- Placeholder for auth-specific values
			$%d, $%d
		) RETURNING uid`

	var authFields, authValues string
	// var status string
	var nextArgIndex = 5

	switch payload.AuthProvider {
	case "local":
		hashedPassword, err := password.HashPassword(payload.Password)
		if err != nil {
			return "", nil, fmt.Errorf("failed to hash password: %w", err)
		}
		authFields = "password, status, auth_provider"
		authValues = fmt.Sprintf("$%d, $%d, $%d", nextArgIndex, nextArgIndex+1, nextArgIndex+2)
		args = append(args, hashedPassword, "PENDING", "local")
		nextArgIndex += 3

	case "google":
		authFields = "status, auth_provider, google_id, google_token"
		authValues = fmt.Sprintf("$%d, $%d, $%d, $%d", nextArgIndex, nextArgIndex+1, nextArgIndex+2, nextArgIndex+3)
		args = append(args, "ACCEPTED", "google", payload.GoogleID, payload.GoogleToken)
		nextArgIndex += 4

	default:
		return "", nil, fmt.Errorf("unsupported auth provider: %s", payload.AuthProvider)
	}

	// Append created_at and updated_at
	args = append(args, now, now)
	query = fmt.Sprintf(query, authFields, authValues, nextArgIndex, nextArgIndex+1)

	return query, args, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*model.User, error) {
	user := new(model.User)
	found := false
	var googleID, googleToken, authProvider sql.NullString

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.UID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Status,
			&authProvider,
			&googleID,
			&googleToken,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		found = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	if !found {
		return nil, nil // No user found, return nil user and nil error
	}

	// Handle nullable fields
	if authProvider.Valid {
		user.AuthProvider = authProvider.String
	}
	if googleID.Valid {
		user.GoogleID = googleID.String
	}
	if googleToken.Valid {
		user.GoogleToken = googleToken.String
	}

	return user, nil
}
