package mattermost

import (
  "fmt"
  "math/rand"
  "time"
  "strconv"
)

func ParseBirthdays(employees []Employee) (birthdayString string) {
  	for _, employee := range employees {
  	firstName := employee.FirstName
	if (employee.PreferredName != "") {
		firstName = employee.PreferredName
	}
		employeeName := fmt.Sprintf("@%s", employee.MattermostUsername)
		if (employeeName == "") {
			continue
		}

		birthday,_ := time.Parse("2006-01-02", employee.DateOfBirth)
		//Build new string with current year and anniversary month and day
		bdayString := strconv.Itoa(time.Now().Year()) + "-" + strconv.Itoa(int(birthday.Month())) + "-" + strconv.Itoa(birthday.Day())
		employeeBirthday,_ := time.Parse("2006-01-02", bdayString)

		birthdayString += StringifyBirthday(employeeName, employeeBirthday.Weekday().String()) + "\n"
  	}

  	return
}

func StringifyBirthday(name, date string) (birthdayString string) {
  birthdayEmojis := []string{
    ":partyparrot:",
    ":dance:",
    ":partyparrot:",
    ":dancer:",
    ":airdancer:",
    ":banana_dance:",
    ":dancing_women:",
    ":gopherdance:",
    ":party_dead:",
    ":ultrafastparrot:",
    ":partydinosaur:",
    ":partygopher:",
    ":partyshark:",
    ":congapartyparrot:",
    ":megaman_party:",
    ":gift:",
    ":fireworks:",
    ":matrixparrot:",
    ":birthday:",
    ":moneybag:",
    ":pizza:",
    ":penguin_dance:",
    ":danghoul:",
    ":happydance:",
    ":dancing_corgi:",
    ":dancing_men:",
    ":shufflepartyparrot:",
    ":aussiereversecongaparrot:",
    ":sassyparrot:",
    ":portalparrot:",
    ":shocked_pikachu:",
    ":pikachu:"}

  rand.Seed(time.Now().UnixNano())
  var emoji1 string
  var emoji2 string

  emoji1 = birthdayEmojis[rand.Intn(len(birthdayEmojis))]
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
