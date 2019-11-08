package mattermost

import (
  "fmt"
  "math/rand"
  "time"
)

var NumberToWord = map[int]string{
      1:  "one",
      2:  "two",
      3:  "three",
      4:  "four",
      5:  "five",
      6:  "six",
      7:  "seven",
      8:  "eight",
      9:  "nine",
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
		employeeName := fmt.Sprintf("@%s.%s", employee.FirstName, employee.LastName)
		employeeAnniversary,_ := time.Parse("2006-01-02", employee.HireDate)
		today := time.Now().Year()
		anniversaryNumber := today - employeeAnniversary.Year()
		numberString := Convert1to9(anniversaryNumber)
		anniversaryString += StringifyAnniversary(employeeName, employeeAnniversary.Weekday().String(), numberString) + "\n"
  	}

  	return
}

func StringifyAnniversary(name, date string, year string) (anniversaryString string) {

  anniversaryEmojis := []string{
    ":tulipio:",
	":raised_hands:",
	":clap:",
	":wave:",
	":open_mouth:",
	":tulip:"}

  rand.Seed(time.Now().UnixNano())
  var emoji1 string
  var emoji2 string

  emoji1 = anniversaryEmojis[rand.Intn(len(anniversaryEmojis))]
  emoji2 = anniversaryEmojis[rand.Intn(len(anniversaryEmojis))]


  anniversaryStrings := []string{
      fmt.Sprintf("%s It's %s's :%s: anniversary at Tulip on %s! %s", emoji1, name, year, date, emoji2),
      fmt.Sprintf("%s On %s, %s has been here for :%s: years! Wow! %s", emoji1, date, name, year, emoji2),
      fmt.Sprintf("%s WOO! It's been :%s: years since %s started here on %s! %s", emoji1, year, name, date, emoji2),
	  fmt.Sprintf("%s Can you believe it? %s is having their :%s: Tulip-day on %s! %s", emoji1, name, year, date, emoji2)}

  anniversaryString = anniversaryStrings[rand.Intn(len(anniversaryStrings))]

  return
}
