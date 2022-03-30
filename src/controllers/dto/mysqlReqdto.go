package dto

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TestSlackMessage struct {
	IntegrationName    string
	HookUrl            string
	SenderName         string
	ChannelName        string
	MessageTitle       string
	TagsSentWithAlerts string
}
