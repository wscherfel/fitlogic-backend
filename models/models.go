package models

import (
	"github.com/jinzhu/gorm"
	"github.com/wscherfel/fitlogic-backend/common"
)

const (
	RoleAdmin = 1
	RoleManager = 2
	RoleUser = 3

	ImpactInsignificant = 0.05
	ImpactSmall = 0.1
	ImpactMedium = 0.2
	ImpactBig = 0.4
	ImpactExtraordinary = 0.8
)

// @dao
type User struct {
	gorm.Model

	Name string `valid:"required"`
	Email string `valid:"email,required" gorm:"unique"`
	Password string `valid:"required"`
	Role int `valid:"required"`
	Skills string

	Projects []Project `gorm:"many2many:user_projects;"`

	Risks []Risk
}

// @dao
type Project struct {
	gorm.Model

	Start common.JSONTime
	End common.JSONTime
	IsFinished bool
	ManagerID uint

	Name string
	Description string

	Users []User `gorm:"many2many:user_projects;"`

	Risks []Risk `gorm:"many2many:risk_projects;"`
}

// @dao
type Risk struct {
	gorm.Model

	Value float64
	Cost int
	Probability float64
	Risk float64

	Name string
	Description string
	Category string
	Threat string
	Status string
	Trigger string
	Impact float64

	Start common.JSONTime
	End common.JSONTime

	UserID uint

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
