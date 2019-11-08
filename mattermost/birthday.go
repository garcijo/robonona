package birthday

import (
  "fmt"
  "math/rand"
  "time"
)

func parseBirthdays(employees []Employee) birthdayString {
  birthdayString := ""

  for _, employee := range employees {
    employeeName = "@%s.%s, employee.FirstName, employee.LastName"
    birthdayString += stringifyBirthday(employeeName, employee.DateOfBirth) + "\n"
  }

  return
}

func stringifyBirthday(name, date string) (birthdayString string) {

  birthdayEmojis := []string{
    ":partyparrot:",
    ":dance:",
    ":partyparrot:",
    ":dancer",
    ":airdancer:",
    "banana_dance:",
    ":dancing_women",
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
    ":birthday:"}

  rand.Seed(time.Now().Unix())

  emoji1 := birthdayEmojis[rand.Intn(len(birthdayEmojis))]
  emoji2 := birthdayEmojis[rand.Intn(len(birthdayEmojis))]

  birthdayStrings := []string{
      fmt.Sprintf("%s On %s let's wish Happy Birthday to %s %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s %s is %s's birthday! Woohoo! %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s Time to party with %s on %s %s", emoji1, name, date, emoji2),
      fmt.Sprintf("%s I heard it's %s's birthday on %s %s",  emoji1, name, date, emoji2),
      fmt.Sprintf("%s On %s we can celebrate with %s! %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s %s requires birthday celebrations for %s! %s", emoji1, date, name, emoji2),
      fmt.Sprintf("%s If it's %s we should party with %s %s", emoji1, date, name, emoji2)}

  birthdayString = birthdayStrings[rand.Intn(len(birthdayStrings))]

  return
}
