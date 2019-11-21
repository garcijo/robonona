package mattermost

import (
	"time"
	"strconv"
)

type Celebrations struct {
	Birthdays  []Employee
	Anniversaries  []Employee
}

//Takes an array with all employees and filters out all those who don't have a birthday or anniversary within the next week
func FilterCelebrations(employees []Employee) Celebrations {
	today := time.Now()
	//Take current date and add 6 days to cover the whole week (Monday - Sunday)
	endOfWeek := today.AddDate(0, 0, 6)

	birthdays := []Employee{}
	anniversaries := []Employee{}
	for _, employee := range employees {
		employeeBirthdayDate,_ := time.Parse("2006-01-02", employee.DateOfBirth)
		bdayString := strconv.Itoa(today.Year()) + "-" + strconv.Itoa(int(employeeBirthdayDate.Month())) + "-" + strconv.Itoa(employeeBirthdayDate.Day())
		employeeBirthday,_ := time.Parse("2006-01-02", bdayString)

		hireDate,_ := time.Parse("2006-01-02", employee.HireDate)
		hireString := strconv.Itoa(today.Year()) + "-" + strconv.Itoa(int(hireDate.Month())) + "-" + strconv.Itoa(hireDate.Day())
		employeeHireDate,_ := time.Parse("2006-01-02", hireString)

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
