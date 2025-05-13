package user

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/blessedmadukoma/ecom/config"
// 	"github.com/blessedmadukoma/ecom/service/auth"
// 	"github.com/blessedmadukoma/ecom/types"
// 	"github.com/blessedmadukoma/ecom/utils"
// 	"github.com/gorilla/mux"
// )

// type Handler struct {
// 	store types.UserStore
// }

// func NewHandler(store types.UserStore) *Handler {
// 	return &Handler{
// 		store: store,
// 	}
// }

// func (h *Handler) RegisterRoutes(router *mux.Router) {
// 	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
// 	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
// }

// func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
// 	// get JSON payload
// 	var payload types.LoginUserPayload

// 	if err := utils.ParseJSON(r, &payload); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	if err := utils.ValidateLoginUserPayload(&payload); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	// check if user exists
// 	user, err := h.store.GetUserByEmail(payload.Email)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found -> maybe invalid email or password: %w", err))
// 		// utils.WriteError(w, http.StatusBadRequest, utils.WrapError(utils.ErrExists, "user"))
// 		return
// 	}

// 	// check password match
// 	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
// 		utils.WriteError(w, http.StatusBadRequest, utils.ErrWrongPassword)
// 		return
// 	}

// 	secret := []byte(config.Envs.JWTSecret)
// 	token, err := auth.CreateJWTToken(secret, user.ID)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	utils.WriteJSON(w, http.StatusOK, "user logged in successfully", map[string]string{"token": token})
// }

// func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
// 	// get JSON payload
// 	var payload types.RegisterUserPayload

// 	if err := utils.ParseJSON(r, &payload); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	if err := utils.ValidateRegisterUserPayload(&payload); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	// check if user exists
// 	user, err := h.store.GetUserByEmail(payload.Email)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err) // Handle actual database errors
// 		return
// 	}
// 	if user != nil {
// 		utils.WriteError(w, http.StatusBadRequest, utils.WrapError(utils.ErrExists, "user"))
// 		return
// 	}

// 	hashedPassword, err := auth.HashPassword(payload.Password)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	// create new user
// 	err = h.store.CreateUser(types.User{
// 		FirstName: payload.FirstName,
// 		LastName:  payload.LastName,
// 		Email:     payload.Email,
// 		Password:  hashedPassword,
// 	})
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, utils.WrapError(utils.ErrInternalServer, "unable to create user"))
// 		return
// 	}

// 	utils.WriteJSON(w, http.StatusCreated, "user registered successfully", nil)
// }
