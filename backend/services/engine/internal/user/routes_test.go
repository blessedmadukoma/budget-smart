package user

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/blessedmadukoma/ecom/types"
// 	"github.com/gorilla/mux"
// )

// func TestUserServiceHandlers(t *testing.T) {
// 	userStore := &mockUserStore{}

// 	handler := NewHandler(userStore)

// 	t.Run("should register the user", func(t *testing.T) {
// 		payload := types.RegisterUserPayload{
// 			FirstName: "first",
// 			LastName:  "last",
// 			Email:     "b@m.com",
// 			Password:  "asd;lkj",
// 		}

// 		marshalled, _ := json.Marshal(payload)

// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		recorder := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(recorder, req)

// 		if recorder.Code != http.StatusCreated {
// 			t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
// 		}
// 	})

// 	t.Run("should fail if the request body is nil", func(t *testing.T) {
// 		req, err := http.NewRequest(http.MethodPost, "/register", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		recorder := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(recorder, req)

// 		if recorder.Code != http.StatusBadRequest {
// 			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
// 		}
// 	})

// 	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
// 		payload := types.RegisterUserPayload{
// 			FirstName: "first",
// 			LastName:  "last",
// 			// Email:     "",
// 			Password: "asd",
// 		}

// 		marshalled, _ := json.Marshal(payload)

// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		recorder := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(recorder, req)

// 		if recorder.Code != http.StatusBadRequest {
// 			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
// 		}
// 	})

// 	// should fail if the user already exists

// 	//should fail if password hashing fails
// }

// type mockUserStore struct {
// 	GetUserByEmailFunc func(email string) (*types.User, error)
// }

// func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
// 	if m.GetUserByEmailFunc != nil {
// 		return m.GetUserByEmailFunc(email)
// 	}
// 	return nil, nil
// }

// func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
// 	return nil, nil
// }

// func (m *mockUserStore) CreateUser(user types.User) error {
// 	return nil
// }
