package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Event struct {
	Base

	Name       string `json:"name"`
	URL        string `json:"url"`
	VisitorID  string `json:"visitor_id"`
	CampaignID string `json:"campaign_id"`

	Payload JSONB `gorm:"type:jsonb;default:'{}'"`
	Meta    JSONB `gorm:"type:jsonb;default:'{}'"`
}

type CreateEvent struct {
	Name       string `json:"name" binding:"required"`
	URL        string `json:"url" binding:"required"`
	VisitorID  string `json:"visitor_id" binding:"required"`
	CampaignID string `json:"campaign_id" binding:"required"`

	Payload JSONB `json:"payload"`
	Meta    JSONB `json:"meta"`
}

type JSONB map[string]interface{}

func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
