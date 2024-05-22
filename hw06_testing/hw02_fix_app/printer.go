package hw02

import (
	"fmt"
)

func PrintStaff(employees []Employee) {
	for _, employee := range employees {
		fmt.Printf("User ID: %d; Age: %d; Name: %s; Department ID: %d; \n",
			employee.UserID, employee.Age, employee.Name, employee.DepartmentID)
	}
}
