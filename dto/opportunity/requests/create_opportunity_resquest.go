package dto

import (
	"github.com/GuilhermeVozniak/go-opportunities/helper"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
)

type CreateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (v *CreateOpportunityRequest) Validate() error {
	if v.Role == "" && v.Company == "" && v.Location == "" && v.Link == "" && v.Remote == nil && v.Salary <= 0 {
		return helper.ErrEmptyRequestBody()
	}

	if v == nil {
		return helper.ErrMalformedJSON()
	}

	if v.Role == "" {
		return helper.ErrParamIsRequired("role", "string")
	}

	if v.Company == "" {
		return helper.ErrParamIsRequired("company", "string")
	}

	if v.Location == "" {
		return helper.ErrParamIsRequired("location", "string")
	}

	if v.Link == "" {
		return helper.ErrParamIsRequired("link", "string")
	}

	if v.Remote == nil {
		return helper.ErrParamIsRequired("remote", "bool")
	}

	if v.Salary <= 0 {
		return helper.ErrParamIsRequired("salary", "int64")
	}
	return nil
}

func (v *CreateOpportunityRequest) ToModel() *schemas.Opportunity {
	return &schemas.Opportunity{
		Role:     v.Role,
		Company:  v.Company,
		Location: v.Location,
		Remote:   *v.Remote,
		Link:     v.Link,
		Salary:   v.Salary,
	}
}
