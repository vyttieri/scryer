package models

type DevicePreference struct {
	ID uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id" gorm:"foreign_key"`
	DeviceID string `json:"device_id"`
	SortPosition uint `json:"sort_position"`
	Visible bool `json:"visible"`
	Icon string `json:"icon"`
}
