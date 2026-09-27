package main

import (
	"fmt"

	"employee-payroll-registry/payroll"
)

func main() {
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

			if err := payroll.AddEmployee(e); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully!")
			}

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

			if err := payroll.PromoteEmployee(id, newRole, salaryBump); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee promoted successfully!")
			}

		case 4:
			var id int
			fmt.Print("Enter Employee ID to terminate: ")
			fmt.Scan(&id)

			if err := payroll.TerminateEmployee(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee terminated successfully!")
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
						e.EmpID, e.Name, e.Department, e.BaseSalary, e.Role,
					)
				}
			}

		case 6:
			fmt.Println("Exiting Payroll Registry...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
