package models

import (
	"github.com/jinzhu/gorm"
)

// constants for Roles and levels of Impact
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
// User is a DB model of a user, email is unique in DB
type User struct {
	gorm.Model

	Name string `valid:"required"`
	Email string `valid:"email,required" gorm:"unique"`
	Password string `valid:"required" json:",omitempty"`
	Role int `valid:"required"`
	Skills string
	Status string

	Projects []Project `gorm:"many2many:user_projects;" json:",omitempty"`

	Risks []Risk `json:",omitempty"`
}

// @dao
// Project is a DB model of a Project, name is unique in DB
type Project struct {
	gorm.Model

	Start string
	End string
	IsFinished bool
	ManagerID uint

	Name string `gorm:"unique"`
	Description string

	Users []User `gorm:"many2many:user_projects;" json:",omitempty"`

	Risks []Risk `gorm:"many2many:risk_projects;" json:",omitempty"`
}

// @dao
// Risk is a DB model of a Risk, name is unique in DB
type Risk struct {
	gorm.Model

	Value float64
	Cost int
	Probability float64
	Risk float64

	Name string `gorm:"unique"`
	Description string
	Category string
	Threat string
	Status string
	Trigger string
	Impact float64

	Start string
	End string

	UserID uint

	Projects []Project `gorm:"many2many:risk_projects;" json:",omitempty"`

	CounterMeasureUsed bool
	CounterMeasureCost int
	CounterMeasureDesc string
}

// dao.db.Model(&m).Association("CounterMeasures").Find(&retVal)

// CounterMeasure is a DB model of a countermeasure to risk,
// currently deprecated
// @dao
type CounterMeasure struct {
	gorm.Model

	Name string
	Description string
	Cost int

	Risks []Risk `gorm:"many2many:risk_counter_measures;"`
}
