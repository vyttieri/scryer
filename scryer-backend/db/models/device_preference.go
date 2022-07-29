package models

type DevicePreference struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"index"`
	DeviceID string `json:"device_id"`
	SortPosition uint `json:"sort_position"`
	Visible bool `json:"visible"`
	Icon string `json:"icon"`
}
