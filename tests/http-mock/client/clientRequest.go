package client

import (
	"context"
	"net/http"
	"time"

	"github.com/renatospaka/http-mock/helpers"
)

type ListModel struct {
	Status string          `json:"status"`
	Data   []EmployeeModel `json:"data"`
}

type EmployeeModel struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type API struct {
	URL string
}

func (a *API) GetEmployee() (*ListModel, error) {
	employees := &ListModel{}

	to := time.Duration(10)
	opt := &helpers.HttpOptions{
		Ctx:     context.Background(),
		Url:     a.URL + "/api/v1/employees",
		TO:      &to,
		Method:  http.MethodGet,
	}

	_, err := helpers.DoRequest(opt, employees)
	return employees, err
}
