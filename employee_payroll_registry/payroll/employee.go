package payroll

import "errors"

type Employee struct {
	EmpID      int
	Name       string
	Department string
	BaseSalary float64
	Role       string
}

var employees = make(map[int]Employee)

func AddEmployee(e Employee) error {

	if e.Department == "" {
		return errors.New("department cannot be empty")
	}

	if e.BaseSalary <= 0 {
		return errors.New("salary must be greater than 0")
	}

	if _, exists := employees[e.EmpID]; exists {
		return errors.New("employee ID already exists")
	}

	employees[e.EmpID] = e

	return nil
}

func FindByID(id int) (*Employee, error) {

	e, exists := employees[id]

	if !exists {
		return nil, errors.New("employee not found")
	}

	return &e, nil
}

func PromoteEmployee(id int, newRole string, salaryBump float64) error {

	e, exists := employees[id]

	if !exists {
		return errors.New("employee not found")
	}

	if newRole == "" {
		return errors.New("role cannot be empty")
	}

	if salaryBump <= 0 {
		return errors.New("salary bump must be greater than 0")
	}

	e.Role = newRole
	e.BaseSalary = e.BaseSalary + salaryBump

	employees[id] = e

	return nil
}

func TerminateEmployee(id int) error {

	if _, exists := employees[id]; !exists {
		return errors.New("employee not found")
	}

	delete(employees, id)

	return nil
}

func GetAllEmployees() map[int]Employee {
	return employees
}

func SetEmployees(data map[int]Employee) {
	employees = data
}