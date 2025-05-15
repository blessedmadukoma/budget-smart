package jwt

import (
	"context"
)

func GetToken(ctx context.Context) string {
	// if ip, ok := ctx.Value("clientIP").(string); ok {
	// 	return ip
	// }

	if token, ok := ctx.Value("auth_token").(string); ok {
		return token
		// Parse the token to get user information
		// claims, err := jwt.Parse(token, h.config.JWTSecret)
		// if err != nil {
		// 	return "unknown"
		// }
		// // Assuming claims has a field "UserID"
		// if userID, ok := claims["user_id"].(string); ok {
		// 	return userID
		// }
		// return "unknown"
	}

	return "unknown"
}
