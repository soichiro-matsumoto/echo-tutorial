package models

import "time"

// BaseModel ...
type BaseModel struct {
	LastUpdate time.Time `json:"last_update" gorm:"culumn:last_update"`
}

// BeforeSave ...
func (m *BaseModel) BeforeSave() {
	m.LastUpdate = time.Now()
}

// BeforeUpdate ...
func (m *BaseModel) BeforeUpdate() {
	m.LastUpdate = time.Now()
}
