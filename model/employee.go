package model

type Response struct {
	Status string         `json:"status"`
	Data   []ListEmployee `json:"data"`
}

type ListEmployee struct {
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
}
