package printer

import (
	"fmt"

	"github.com/Stern-Ritter/go/hw02_fix_app/types"
)

func PrintStaff(employees []types.Employee) {
	for _, employee := range employees {
		fmt.Printf("User ID: %d; Age: %d; Name: %s; Department ID: %d; \n",
			employee.UserID, employee.Age, employee.Name, employee.DepartmentID)
	}
}
