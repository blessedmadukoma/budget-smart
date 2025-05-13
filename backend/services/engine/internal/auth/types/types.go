package types

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AccountStatusType string

// Simplified status keys
var AccountStatusPendingKey = uint(1)
var AccountStatusActiveKey = uint(2)
var AccountStatusRejectedKey = uint(3)
var AccountStatusRestrictedKey = uint(4)
var AccountStatusLockedKey = uint(5)

// Simplified account status types
var AccountStatus = struct {
	PENDING    AccountStatusType
	ACTIVE     AccountStatusType
	REJECTED   AccountStatusType
	RESTRICTED AccountStatusType
	LOCKED     AccountStatusType
}{
	PENDING:    "PENDING",
	ACTIVE:     "ACTIVE",
	REJECTED:   "REJECTED",
	RESTRICTED: "RESTRICTED",
	LOCKED:     "LOCKED",
}

// Returns a map of account status types to their corresponding keys
func AccountStatuses() map[AccountStatusType]uint {
	m := map[AccountStatusType]uint{
		AccountStatus.PENDING:    AccountStatusPendingKey,
		AccountStatus.ACTIVE:     AccountStatusActiveKey,
		AccountStatus.REJECTED:   AccountStatusRejectedKey,
		AccountStatus.RESTRICTED: AccountStatusRestrictedKey,
		AccountStatus.LOCKED:     AccountStatusLockedKey,
	}
	return m
}

// GetAccountStatusKey returns the key for an account status type
// If the status type is not found, it returns the PENDING key as default
func GetAccountStatusKey(t AccountStatusType) uint {
	if v, found := AccountStatuses()[t]; found {
		return v
	}
	return AccountStatusPendingKey
}

// GetAccountStatusName returns the account status type for a given key
// If the key is not found, it returns PENDING as default
func GetAccountStatusName(val uint) AccountStatusType {
	for k, v := range AccountStatuses() {
		if v == val {
			return k
		}
	}
	return AccountStatus.PENDING
}
