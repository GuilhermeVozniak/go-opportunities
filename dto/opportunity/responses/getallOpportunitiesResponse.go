package dto

import (
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
)

type GetallOpportunityResponse struct {
	Message string                `json:"message"`
	Data    []GetallOpportunityData `json:"data"`
}

type GetallOpportunityData struct {
	ID       uint   `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (*GetallOpportunityResponse) FromModelToResponse(opportunityArr *[]schemas.Opportunity, op string) *GetallOpportunityResponse {
	var opportunities = GetallOpportunityResponse{}

	opportunities.Message = "operation from handler " + op + " successfull"
	for _, opportunity := range *opportunityArr {
		opportunities.Data = append(opportunities.Data, GetallOpportunityData{
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
