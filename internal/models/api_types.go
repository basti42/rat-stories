package models

import (
	"github.com/google/uuid"
)

type NewStory struct {
	Assignee           *uuid.UUID            `json:"assignee,omitempty"`
	Team               *uuid.UUID            `json:"team,omitempty"`
	Project            *uuid.UUID            `json:"project,omitempty"`
	Feature            *uuid.UUID            `json:"feature,omitempty"`
	Status             string                `json:"status"`
	Title              string                `json:"title"`
	Description        string                `json:"description"`
	AcceptanceCriteria []AcceptanceCriterium `json:"acceptance_criteria,omitempty"`
}

type NewAcceptanceCriterium struct {
	Description string `json:"description"`
}

type NewStoryHistory struct {
	FromStatus string `json:"from_status"`
	ToStatus   string `json:"to_status"`
}
