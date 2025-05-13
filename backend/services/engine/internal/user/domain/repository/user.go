package repository

import (
	"database/sql"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/adapters"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/model"
)

type Repository interface {
	// Create(firstname, lastname, email, password string) error
	// Create(firstname, lastname, email, password string) error
	// Create(email string) error
	// Create(u *model.User) error
	Create(u *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	// GetByUID(uid string) (*types.Account, error)
	// GetByEmail(email string) (*types.Account, error)
	// GetByUsername(username string) (*types.Account, error)
	// GetByUsernameAndID(username string, id uint) (map[string]interface{}, error)
	// ListByUserName(username string, id uint) ([]map[string]interface{}, error)
	// IsUserNameExist(username string, id uint) (bool, error)

	// GetOrCreateInfoByID(id uint) (*types.AccountInfo, error)
	// GetInfoByID(id uint) (*types.AccountInfo, error)
	// GetInfoByReferralCode(code string) (*types.AccountInfo, error)
	// GetSelfieImage(infos []types.AccountInfoDocumentDocs) (*types.GetSelfieImage, error)
	// CreateOrUpdateInfoMeta(id uint, infoMeta types.AccountInfoMeta) error

	// UpdateStatus(id uint, status types.AccountStatusType) error
	// UpdateUsername(id uint, username string) error
	// UpdatePassword(id uint, psw string) error
	// Deactivate(id uint) error

	// CreateToken(tx *gorm.DB, at model.AccountToken) error
}

func NewRepository(db *sql.DB) Repository {
	return adapters.NewUserRepository(db)
}
