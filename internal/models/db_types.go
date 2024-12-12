package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type Story struct {
	UUID               uuid.UUID             `gorm:"primaryKey" json:"uuid"`
	Type               string                `json:"type"`
	Creator            uuid.UUID             `json:"creator"`
	Assignee           *uuid.UUID            `gorm:"null" json:"assignee"`
	Team               *uuid.UUID            `gorm:"null" json:"team"`
	Project            *uuid.UUID            `gorm:"null" json:"project"`
	Feature            *uuid.UUID            `gorm:"null" json:"feature"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
	ClosedAt           *time.Time            `json:"closed_at"`
	Status             string                `json:"status"`
	Title              string                `json:"title"`
	Description        string                `json:"description"`
	Estimation         *Estimation           `json:"estimation" gorm:"foreignKey:StoryUUID"`
	AcceptanceCriteria []AcceptanceCriterium `json:"acceptance_criteria" gorm:"foreignKey:StoryUUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	History            []StoryHistory        `json:"history" gorm:"foreignKey:StoryUUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments           []Comment             `json:"comments" gorm:"foriegnKey:StoryUUID;contraint:OnUpdate:CASCASE,OnDelete:SET NULL;"`
}

func (Story) TableName() string { return "stories" }

type Estimation struct {
	UUID       uuid.UUID `gorm:"primaryKey" json:"uuid"`
	StoryUUID  uuid.UUID `json:"story_uuid" gorm:"index"`
	Risk       int       `json:"risk"`
	Complexity int       `json:"complexity"`
	Effort     int       `json:"effort"`
}

func (Estimation) TableName() string { return "estimations" }

type AcceptanceCriterium struct {
	UUID        uuid.UUID  `gorm:"primaryKey" json:"uuid"`
	StoryUUID   uuid.UUID  `json:"story_uuid"`
	Description string     `json:"description"`
	AcceptedAt  *time.Time `json:"accepted_at"`
	AcceptedBy  *uuid.UUID `json:"accepted_by"`
}

func (AcceptanceCriterium) TableName() string { return "acceptance_criteria" }

type StoryHistory struct {
	UUID       uuid.UUID `gorm:"primaryKey" json:"uuid"`
	StoryUUID  uuid.UUID `json:"story_uuid"`
	FromStatus string    `json:"from_status"`
	ToStatus   string    `json:"to_status"`
	Date       time.Time `json:"date"`
	By         uuid.UUID `json:"by"`
}

func (StoryHistory) TableName() string { return "story_history" }

type Comment struct {
	gorm.Model
	StoryUUID uuid.UUID `json:"story_uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Text      string    `json:"text"`
}

func (Comment) TableName() string { return "comments" }
