package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	RoleAdmin = 1
	RoleManager = 2
	RoleUser = 3
)

// @dao
type User struct {
	gorm.Model

	Name string
	Email string `valid:"email" gorm:"unique"`
	Password string
	Role int
	Skills string

	Projects []Project `gorm:"many2many:user_projects;"`
}

// @dao
type Project struct {
	gorm.Model

	Start time.Time
	End time.Time
	IsFinished bool

	Name string
	Description string

	Users []User `gorm:"many2many:user_projects;"`

	Risks []Risk `gorm:"many2many:risk_projects;"`
}

// @dao
type Risk struct {
	gorm.Model

	Value int
	Cost int
	Probability float64

	Name string
	Description string
	Category string
	Threat string
	Status string
	Trigger string

	Start time.Time
	End time.Time

	Owner int

	Projects []Project `gorm:"many2many:risk_projects;"`

	CounterMeasures []CounterMeasure `gorm:"many2many:risk_counter_measures;"`
}

// @dao
type CounterMeasure struct {
	gorm.Model

	Name string
	Description string
	Cost int

	Risks []Risk `gorm:"many2many:risk_counter_measures;"`
}
