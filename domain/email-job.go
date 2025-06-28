package domain

import (
	"time"

	"github.com/google/uuid"
)

type EmailJob struct {
	ID                uuid.UUID
	CampaignID        uuid.UUID
	SubscriberListIDs []uuid.UUID
	ErrorMessage      string
	Status            string
	CreatedAt         time.Time
}
