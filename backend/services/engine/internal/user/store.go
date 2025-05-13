package user

// import (
// 	"database/sql"

// 	"github.com/blessedmadukoma/ecom/types"
// 	"github.com/blessedmadukoma/ecom/utils"
// )

// type Store struct {
// 	db *sql.DB
// }

// func NewStore(db *sql.DB) *Store {
// 	return &Store{
// 		db: db,
// 	}
// }

// func (s *Store) GetUserByEmail(email string) (*types.User, error) {
// 	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	u := new(types.User)
// 	found := false
// 	for rows.Next() {
// 		u, err = scanRowIntoUser(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 		found = true
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	if !found {
// 		return nil, nil // No user found, return nil user and nil error
// 	}

// 	return u, nil
// }

// func (s *Store) GetUserByID(id int) (*types.User, error) {
// 	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	u := new(types.User)
// 	for rows.Next() {
// 		u, err = scanRowIntoUser(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if u.ID == 0 {
// 		return nil, utils.WrapError(utils.ErrNotFound, "user")
// 	}

// 	return u, nil
// }

// func (s *Store) CreateUser(user types.User) error {
// 	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
// 	user := new(types.User)

// 	err := rows.Scan(
// 		&user.ID,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Email,
// 		&user.Password,
// 		&user.CreatedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
