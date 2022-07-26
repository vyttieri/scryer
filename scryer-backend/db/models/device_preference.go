package models

import (
	"scryer-backend/db"
)

type DevicePreference struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"userId" gorm:"index"`
	DeviceID string `json:"deviceId"`
	SortPosition uint `json:"sortPosition"`
	Visible bool `json:"visible"`
	Icon string `json:"icon"`
}

func (devicePreference *DevicePreference) Save(dbContext *db.DbContext) error {
	return dbContext.Connection.Save(&devicePreference).Error
}
