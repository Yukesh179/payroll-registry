package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"employee-payroll-registry/payroll"
)

const dataFile = "employees.csv"

// Save all employee data to CSV file
func saveEmployees() error {

	file, err := os.Create(dataFile)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// CSV header
	writer.Write([]string{
		"EmpID",
		"Name",
		"Department",
		"BaseSalary",
		"Role",
	})

	employees := payroll.GetAllEmployees()

	for _, e := range employees {

		record := []string{
			strconv.Itoa(e.EmpID),
			e.Name,
			e.Department,
			fmt.Sprintf("%.2f", e.BaseSalary),
			e.Role,
		}

		writer.Write(record)
	}

	return writer.Error()
}

// Load employee data from CSV file
func loadEmployees() error {

	file, err := os.Open(dataFile)

	if os.IsNotExist(err) {
		// File does not exist yet.
		// It will be created when first employee is added.
		return nil
	}

	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return err
	}

	employees := make(map[int]payroll.Employee)

	// Start from 1 because row 0 is the header
	for i := 1; i < len(records); i++ {

		record := records[i]

		if len(record) < 5 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		salary, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		employee := payroll.Employee{
			EmpID:      id,
			Name:       record[1],
			Department: record[2],
			BaseSalary: salary,
			Role:       record[4],
		}

		employees[id] = employee
	}

	payroll.SetEmployees(employees)

	return nil
}

func main() {

	// Load previously saved employee data
	err := loadEmployees()

	if err != nil {
		fmt.Println("Error loading employees:", err)
	}

	for {

		fmt.Println("\n========== EMPLOYEE PAYROLL REGISTRY ==========")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Find Employee")
		fmt.Println("3. Promote Employee")
		fmt.Println("4. Terminate Employee")
		fmt.Println("5. Display All Employees")
		fmt.Println("6. Exit")
		fmt.Println("===============================================")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			var e payroll.Employee

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&e.EmpID)

			fmt.Print("Enter Name: ")
			fmt.Scan(&e.Name)

			fmt.Print("Enter Department: ")
			fmt.Scan(&e.Department)

			fmt.Print("Enter Base Salary: ")
			fmt.Scan(&e.BaseSalary)

			fmt.Print("Enter Role: ")
			fmt.Scan(&e.Role)

			err := payroll.AddEmployee(e)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				err := saveEmployees()

				if err != nil {
					fmt.Println("Error saving data:", err)
				} else {
					fmt.Println("Employee added successfully!")
					fmt.Println("Data saved to employees.csv")
				}
			}

		// --------------------------------
		// FIND EMPLOYEE
		// --------------------------------

		case 2:

			var id int

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			e, err := payroll.FindByID(id)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				fmt.Println("\nEmployee Details")
				fmt.Println("-------------------------")
				fmt.Println("ID         :", e.EmpID)
				fmt.Println("Name       :", e.Name)
				fmt.Println("Department :", e.Department)
				fmt.Printf("Salary     : %.2f\n", e.BaseSalary)
				fmt.Println("Role       :", e.Role)
			}

		// --------------------------------
		// PROMOTE EMPLOYEE
		// --------------------------------

		case 3:

			var id int
			var newRole string
			var salaryBump float64

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter New Role: ")
			fmt.Scan(&newRole)

			fmt.Print("Enter Salary Bump: ")
			fmt.Scan(&salaryBump)

			err := payroll.PromoteEmployee(
				id,
				newRole,
				salaryBump,
			)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				err := saveEmployees()

				if err != nil {
					fmt.Println("Error saving data:", err)
				} else {
					fmt.Println("Employee promoted successfully!")
					fmt.Println("Data updated in employees.csv")
				}
			}
		case 4:

			var id int

			fmt.Print("Enter Employee ID to terminate: ")
			fmt.Scan(&id)

			err := payroll.TerminateEmployee(id)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				err := saveEmployees()

				if err != nil {
					fmt.Println("Error saving data:", err)
				} else {
					fmt.Println("Employee terminated successfully!")
					fmt.Println("Data updated in employees.csv")
				}
			}
		case 5:

			fmt.Println("\n========== ALL EMPLOYEES ==========")

			employees := payroll.GetAllEmployees()

			if len(employees) == 0 {

				fmt.Println("No employees found.")

			} else {

				for _, e := range employees {

					fmt.Printf(
						"ID: %d | Name: %s | Department: %s | Salary: %.2f | Role: %s\n",
						e.EmpID,
						e.Name,
						e.Department,
						e.BaseSalary,
						e.Role,
					)
				}
			}

		case 6:

			fmt.Println("\nEmployee data is saved in employees.csv")
			fmt.Println("Exiting Payroll Registry...")
			return

		default:

			fmt.Println("Invalid choice. Please try again.")
		}
	}
}