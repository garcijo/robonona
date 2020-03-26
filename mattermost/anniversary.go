package mattermost

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var NumberToWord = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func Convert1to9(n int) (w string) {
	if n < 20 {
		w = NumberToWord[n]
		return
	}
	r := n % 10
	if r == 0 {
		w = NumberToWord[n]
	} else {
		w = NumberToWord[n-r] + "-" + NumberToWord[r]
	}
	return
}

func ParseAnniversaries(employees []Employee) (anniversaryString string) {
	for _, employee := range employees {
		employeeName := fmt.Sprintf("@%s", employee.MattermostUsername)
		if len(employeeName) == 0 {
			continue
		}

		anniversary, _ := time.Parse("2006-01-02", employee.HireDate)
		monthHireString := strconv.Itoa(int(anniversary.Month()))
		if int(anniversary.Month()) < 10 {
			monthHireString = "0" + strconv.Itoa(int(anniversary.Month()))
		}
		dayHireString := strconv.Itoa(int(anniversary.Day()))
		if int(anniversary.Day()) < 10 {
			dayHireString = "0" + strconv.Itoa(int(anniversary.Day()))
		}
		hireString := strconv.Itoa(time.Now().Year()) + "-" + monthHireString + "-" + dayHireString
		employeeAnniversary, _ := time.Parse("2006-01-02", hireString)
		today := time.Now().Year()
		anniversaryNumber := today - anniversary.Year()
		numberString := Convert1to9(anniversaryNumber)
		anniversaryString += StringifyAnniversary(employeeName, employeeAnniversary.Weekday().String(), numberString) + "\n"
	}

	return
}

func StringifyAnniversary(name, date string, year string) (anniversaryString string) {

	anniversaryEmojis := GetAnniversaryEmojis()

	var emoji1 string
	var emoji2 string

	rand.Seed(time.Now().UnixNano())
	emoji1 = anniversaryEmojis[rand.Intn(len(anniversaryEmojis))]
	rand.Seed(time.Now().UnixNano())
	emoji2 = anniversaryEmojis[rand.Intn(len(anniversaryEmojis))]

	anniversaryStrings := []string{
		fmt.Sprintf("%s It's %s's :%s: anniversary at Tulip on %s! %s", emoji1, name, year, date, emoji2),
		fmt.Sprintf("%s On %s, %s has been here for :%s: years! Wow! %s", emoji1, date, name, year, emoji2),
		fmt.Sprintf("%s WOO! It's been :%s: years since %s started here on %s! %s", emoji1, year, name, date, emoji2),
		fmt.Sprintf("%s Can you believe it? %s is having their :%s: Tulip-day on %s! %s", emoji1, name, year, date, emoji2)}

	anniversaryString = anniversaryStrings[rand.Intn(len(anniversaryStrings))]

	return
}
