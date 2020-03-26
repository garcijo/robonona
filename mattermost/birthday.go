package mattermost

import (
  "fmt"
  "math/rand"
  "time"
  "strconv"
)


func ParseBirthdays(employees []Employee) (birthdayString string) {
  	for _, employee := range employees {
		employeeName := fmt.Sprintf("@%s", employee.MattermostUsername)
		if len(employeeName) == 0 {
			continue
		}

		employeeBirthdayDate,_ := time.Parse("2006-01-02", employee.DateOfBirth)
		monthString := strconv.Itoa(int(employeeBirthdayDate.Month()))
		if int(employeeBirthdayDate.Month()) < 10 {
			monthString = "0" + strconv.Itoa(int(employeeBirthdayDate.Month()))
		}
		dayString := strconv.Itoa(int(employeeBirthdayDate.Day()))
		if int(employeeBirthdayDate.Day()) < 10 {
			dayString = "0" + strconv.Itoa(int(employeeBirthdayDate.Day()))
		}
		bdayString := strconv.Itoa(time.Now().Year()) + "-" + monthString + "-" + dayString
		employeeBirthday,_ := time.Parse("2006-01-02", bdayString)

		weekdayString := employeeBirthday.Weekday().String()
		WeekDays[weekdayString] = append(WeekDays[weekdayString], employee)
  	}

  	for day, employees := range WeekDays {
  		if len(employees) > 0 {
			birthdayString += CompileDay(day, employees) + "\n"
		}
	}

  	return
}

func CompileDay(day string, employees []Employee) (birthdayString string) {

	var names string
	lastIdx := len(employees) - 1
	for idx, employee := range employees {
		employeeName := fmt.Sprintf("@%s", employee.MattermostUsername)

		if len(employees) == 1 {
			names += employeeName
		} else if idx == lastIdx {
			names += " and " + employeeName
		} else {
			names += employeeName + " "
		}
	}
	birthdayString = StringifyBirthday(names, day)

	return
}

func StringifyBirthday(name, date string) (birthdayString string) {
  birthdayEmojis := GetBirthdayEmojis()

  var emoji1 string
  var emoji2 string

  rand.Seed(time.Now().UnixNano())
  emoji1 = birthdayEmojis[rand.Intn(len(birthdayEmojis))]
  rand.Seed(time.Now().UnixNano())
  emoji2 = birthdayEmojis[rand.Intn(len(birthdayEmojis))]

  birthdayStrings := []string{
      fmt.Sprintf("%s On %s let's wish Happy Birthday to %s %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s %s is %s 's birthday! Woohoo! %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s Time to party with %s on %s %s", emoji1, name, date, emoji2),
      fmt.Sprintf("%s I heard it's %s 's birthday on %s %s",  emoji1, name, date, emoji2),
      fmt.Sprintf("%s On %s we can celebrate with %s! %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s %s requires birthday celebrations for %s! %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s If it's %s we should party with %s %s", emoji1, date, name, emoji2)}

  birthdayString = birthdayStrings[rand.Intn(len(birthdayStrings))]

  return
}
