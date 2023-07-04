package client

import (
	"net/http"
	"testing"

	"github.com/renatospaka/http-mock/helpers"
)

func TestAPIGetEmployees(t *testing.T) {
	resp := &ListModel{
		Status: "success",
		Data:   []EmployeeModel{
			{
				ID:             "1",
				EmployeeName:   "Tiger Nixon",
				EmployeeSalary: "320800",
				EmployeeAge:    "61",
				ProfileImage:   "",
			},
		},
	}

	srv := helpers.HttpMock("/api/vi/employees", http.StatusOK, resp)
	defer srv.Close()

	api := API{URL: srv.URL}
	employees, err := api.GetEmployee()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if employees.Status != "success" {
		t.Error("expected status success got:", employees.Status)
	}
	if len(employees.Data) != 1 {
		t.Error("expected 1 data got", len(employees.Data))
	}
}
