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

		monthString := strconv.Itoa(int(employeeBirthdayDate.Month()))
		if int(employeeBirthdayDate.Month()) < 10 {
			monthString = "0" + strconv.Itoa(int(employeeBirthdayDate.Month()))
		}
		dayString := strconv.Itoa(int(employeeBirthdayDate.Day()))
		if int(employeeBirthdayDate.Day()) < 10 {
			dayString = "0" + strconv.Itoa(int(employeeBirthdayDate.Day()))
		}

		bdayString := strconv.Itoa(today.Year()) + "-" + monthString + "-" + dayString
		employeeBirthday,_ := time.Parse("2006-01-02", bdayString)

		hireDate,_ := time.Parse("2006-01-02", employee.HireDate)
		monthHireString := strconv.Itoa(int(hireDate.Month()))
		if int(hireDate.Month()) < 10 {
			monthHireString = "0" + strconv.Itoa(int(hireDate.Month()))
		}
		dayHireString := strconv.Itoa(int(hireDate.Day()))
		if int(hireDate.Day()) < 10 {
			dayHireString = "0" + strconv.Itoa(int(hireDate.Day()))
		}
		hireString := strconv.Itoa(today.Year()) + "-" + monthHireString + "-" + dayHireString
		employeeHireDate,_ := time.Parse("2006-01-02", hireString)

		if (employeeBirthday.YearDay() >= today.YearDay()) && (employeeBirthday.YearDay() <= endOfWeek.YearDay()) {
			birthdays = append(birthdays, employee)
		}
		if (hireDate.Year() < today.Year()) && (employeeHireDate.YearDay() >= today.YearDay()) && (employeeHireDate.YearDay() <= endOfWeek.YearDay()) {
			 anniversaries = append(anniversaries, employee)
		 }
	}

	celebrations := Celebrations{birthdays, anniversaries}

	return celebrations
}
