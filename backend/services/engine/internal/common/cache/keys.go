package cache

import "fmt"

var JWTTokenKey = func(uid string) string { return fmt.Sprintf("jwt:token:%s", uid) }
var ConfigurationKey = "configurations"
