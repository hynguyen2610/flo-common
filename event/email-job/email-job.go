package event

import (
	"time"

	domain "github.com/hynguyen2610/flo-common/domain" //

	"github.com/google/uuid"
)

type EmailJobAction string

const (
	ActionCreate EmailJobAction = "create"
	ActionRead   EmailJobAction = "read"
	ActionUpdate EmailJobAction = "update"
	ActionDelete EmailJobAction = "delete"
)

type EmailJobMessage struct {
	Version  int            `json:"version"`
	EventID  uuid.UUID      `json:"eventId"`
	Action   EmailJobAction `json:"action"`
	EventAt  time.Time      `json:"eventAt"`
	EmailJob EmailJobDTO    `json:"emailJob"`
}

type EmailJobDTO struct {
	ID                uuid.UUID   `json:"id"`
	CampaignID        uuid.UUID   `json:"campaignId"`
	SubscriberListIDs []uuid.UUID `json:"subscriberListIds"`
	ErrorMessage      string      `json:"errorMessage,omitempty"`
	Status            string      `json:"status"`
	CreatedAt         time.Time   `json:"createdAt"`
}

func ToDTO(job domain.EmailJob) EmailJobDTO {
	return EmailJobDTO{
		ID:                job.ID,
		CampaignID:        job.CampaignID,
		SubscriberListIDs: job.SubscriberListIDs,
		ErrorMessage:      job.ErrorMessage,
		Status:            job.Status,
		CreatedAt:         job.CreatedAt,
	}
}

func FromDTO(dto EmailJobDTO) domain.EmailJob {
	return domain.EmailJob{
		ID:                dto.ID,
		CampaignID:        dto.CampaignID,
		SubscriberListIDs: dto.SubscriberListIDs,
		ErrorMessage:      dto.ErrorMessage,
		Status:            dto.Status,
		CreatedAt:         dto.CreatedAt,
	}
}
