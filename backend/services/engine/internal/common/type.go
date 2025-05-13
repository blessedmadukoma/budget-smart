package common

//go:generate cp -r ../../VERSION ./VERSION

// go:embed VERSION
// var F embed.FS

type QueueName string

type TaskName string

const (
	NotificationProcessor TaskName = "NotificationProcessor"
	// SendPaymentProcessor         TaskName = "SendPaymentProcessor"
	// AccountVerificationProcessor TaskName = "AccountVerificationProcessor"
	CreateCardProcessor TaskName = "CreateCardProcessor"
	// LinkCardProcessor            TaskName = "LinkCardProcessor"
	// CreditCardProcessor          TaskName = "CreditCardProcessor"
	// WithdrawCardProcessor        TaskName = "WithdrawCardProcessor"
	// ResendCardPinProcessor       TaskName = "ResendCardPinProcessor"
)

// queues
const (
	AccountQueue  QueueName = "Account"
	PaymentQueue  QueueName = "Payment"
	ScheduleQueue QueueName = "Schedule"
	DefaultQueue  QueueName = "Default"
)

// func readVersion(fs embed.FS) ([]byte, error) {
// 	data, err := fs.ReadFile("VERSION")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func GetVersion() string {
// 	v := "0.1.0"

// 	f, err := readVersion(F)
// 	if err != nil {
// 		return v
// 	}

// 	v = strings.TrimSpace(string(f))
// 	return v
// }
