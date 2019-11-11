package mattermost

import (
	"time"
)

type Celebrations struct {
	Birthdays  []Employee
	Anniversaries  []Employee
}

func FilterCelebrations(employees []Employee) Celebrations {
	today := time.Now()
	//Take current date and add 6 days to cover the whole week (Monday - Sunday)
	endOfWeek := today.AddDate(0, 0, 6)

	birthdays := []Employee{}
	anniversaries := []Employee{}
	for _, employee := range employees {
		employeeBirthday,_ := time.Parse("2006-01-02", employee.DateOfBirth)
		employeeHireDate,_ := time.Parse("2006-01-02", employee.HireDate)

		if !(employeeBirthday.YearDay() < today.YearDay()) && !(employeeBirthday.YearDay() > endOfWeek.YearDay()) {
			birthdays = append(birthdays, employee)
		}
		if !(employeeHireDate.YearDay() < today.YearDay()) && !(employeeHireDate.YearDay() > endOfWeek.YearDay()) {
			 anniversaries = append(anniversaries, employee)
		 }
	}

	celebrations := Celebrations{birthdays, anniversaries}

	return celebrations
}
