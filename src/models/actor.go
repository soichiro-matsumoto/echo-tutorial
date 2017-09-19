package models

// Actor ...
type Actor struct {
	ActorID   int    `json:"actor_id" gorm:"primary_key; column:actor_id"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	BaseModel
}

// TableName ...
func (m *Actor) TableName() string {
	return "actor"
}
