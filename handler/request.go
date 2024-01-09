package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

// create lead

type CreateLeadRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Client *bool  `json:"client"`
	Age    int64  `json:"age"`
}

func (r *CreateLeadRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Phone == "" {
		return errParamIsRequired("phone", "string")
	}
	if r.Client == nil {
		return errParamIsRequired("client", "bool")
	}
	if r.Age <= 0 {
		return errParamIsRequired("age", "int64")
	}
	return nil
}

// update request
type UpdateLeadRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Client *bool  `json:"client"`
	Age    int64  `json:"age"`
}

func (r *UpdateLeadRequest) Validate() error {
	if r.Name != "" ||
		r.Email != "" ||
		r.Phone != "" ||
		r.Client != nil ||
		r.Age > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provied")
}
