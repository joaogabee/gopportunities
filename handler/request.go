package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type RequestOpening struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *RequestOpening) Validate() error {
	if r.Role == "" && r.Remote == nil && r.Location == "" && r.Company == "" && r.Salary < 0 {
		return fmt.Errorf("request body is empty!")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary < 1 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpening struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpening) Validate() error {
	if r.Role != "" || r.Link != "" || r.Location != "" || r.Company != "" || r.Remote != nil || r.Salary < 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
