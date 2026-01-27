package tgbot

import "github.com/go-telegram/bot/models"

// Get user's telegram username
func GetUsername(user *models.User) string {
	if user == nil {
		return ""
	}
	if user.Username != "" {
		return "@" + user.Username
	}
	return ""
}
