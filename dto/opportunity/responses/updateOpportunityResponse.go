package dto

import (
	"time"

	"github.com/GuilhermeVozniak/go-opportunities/schemas"
)

type UpdateOpportunityResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}

func (*UpdateOpportunityResponse) FromModel(opportunity *schemas.Opportunity) *UpdateOpportunityResponse {
	return &UpdateOpportunityResponse{
		ID:        opportunity.ID,
		CreatedAt: opportunity.CreatedAt,
		UpdatedAt: opportunity.UpdatedAt,
		DeletedAt: opportunity.DeletedAt.Time,
		Role:      opportunity.Role,
		Company:   opportunity.Company,
		Location:  opportunity.Location,
		Remote:    opportunity.Remote,
		Link:      opportunity.Link,
		Salary:    opportunity.Salary,
	}
}
