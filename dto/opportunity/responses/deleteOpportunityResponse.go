package dto

import (
	"fmt"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"time"
)

type DeleteOpportunityResponse struct {
	Message string                `json:"message"`
	Data    DeleteOpportunityData `json:"data"`
}

type DeleteOpportunityData struct {
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

func (*DeleteOpportunityResponse) FromModelToResponse(opportunity *schemas.Opportunity, op string) *DeleteOpportunityResponse {
	return &DeleteOpportunityResponse{
		Message: fmt.Sprintf("operation from handler %s successfull", op),
		Data: DeleteOpportunityData{
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
		},
	}
}
