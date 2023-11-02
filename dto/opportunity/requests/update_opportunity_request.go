package dto

import (
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/helper"
)

type UpdateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (v *UpdateOpportunityRequest) Validate() error {
	// if any filed is provided validation is truthy
	if v.Role != "" || v.Company != "" || v.Location != "" || v.Link != "" || v.Remote != nil || v.Salary > 0 {
		return nil
	}
	return helper.ErrEmptyRequestBody()
}

func (v *UpdateOpportunityRequest) ToModel(opportunity *schemas.Opportunity) schemas.Opportunity {
	if v.Role != "" {
		opportunity.Role = v.Role
	}
	if v.Company != "" {
		opportunity.Company = v.Company
	}
	if v.Location != "" {
		opportunity.Location = v.Location
	}
	if v.Remote != nil {
		opportunity.Remote = *v.Remote
	}
	if v.Link != "" {
		opportunity.Link = v.Link
	}
	if v.Salary > 0 {
		opportunity.Salary = v.Salary
	}
	return *opportunity
}
