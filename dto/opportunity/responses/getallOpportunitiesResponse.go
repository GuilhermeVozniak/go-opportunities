package dto

import (
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
)

type GetallOpportunityResponse struct {
	ID       uint   `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (*GetallOpportunityResponse) FromModel(opportunityArr *[]schemas.Opportunity) *[]GetallOpportunityResponse {
	var opportunities = []GetallOpportunityResponse{}

	for _, opportunity := range *opportunityArr {
		opportunities = append(opportunities, GetallOpportunityResponse{
			ID:       opportunity.ID,
			Role:     opportunity.Role,
			Company:  opportunity.Company,
			Location: opportunity.Location,
			Remote:   opportunity.Remote,
			Link:     opportunity.Link,
			Salary:   opportunity.Salary,
		})
	}

	return &opportunities
}
