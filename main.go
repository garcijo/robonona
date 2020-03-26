package main

import (
	"encoding/json"
	"fmt"
	"github.com/garcijo/robonona/mattermost"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	//Extract env variables
	serverURL := os.Getenv("ROBONONA_MATTERMOST_URL")
	botUserName := os.Getenv("ROBONONA_USERNAME")
	botPassword := os.Getenv("ROBONONA_PASSWORD")
	teamName := os.Getenv("ROBONONA_TEAM_NAME")
	channelName := os.Getenv("ROBONONA_CHANNEL_NAME")
	bambooURL := os.Getenv("ROBONONA_BAMBOO_URL")
	bambooKey := os.Getenv("ROBONONA_BAMBOO_KEY")

	//Define API Clients
	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)
	bambooApi := mattermost.BambooHR(bambooURL, bambooKey)

	// Get employees directory
	employees, _ := bambooApi.GetDirectory()
	// Take directory and get extra information for each employee (birthday, anniversary)
	employeeData, _ := bambooApi.GetEmployeeData(employees.Employees)
	// Filter only the employees with celebrations within the next week
	celebrations := mattermost.FilterCelebrations(employeeData)

	birthdays := mattermost.GetMattermostUsernames(*api, celebrations.Birthdays)
	bdayString := mattermost.ParseBirthdays(birthdays)
	anniversaries := mattermost.GetMattermostUsernames(*api, celebrations.Anniversaries)
	anniString := mattermost.ParseAnniversaries(anniversaries)
	holidayString := "Holidays this week: "

	// TODO holidaysURL := "https://canada-holidays.ca/api/v1/provinces/ON"
	holidays := `[
         {
            "id":1,
            "date":"2020-01-01",
            "nameEn":"New Year’s Day",
            "nameFr":"Jour de l’An",
            "federal":1
         },
         {
            "id":4,
            "date":"2020-02-17",
            "nameEn":"Family Day",
            "nameFr":"Fête de la famille",
            "federal":0
         },
         {
            "id":7,
            "date":"2020-04-10",
            "nameEn":"Good Friday",
            "nameFr":"Vendredi saint",
            "federal":1
         },
         {
            "id":11,
            "date":"2020-05-18",
            "nameEn":"Victoria Day",
            "nameFr":"Fête de la Reine",
            "federal":1
         },
         {
            "id":15,
            "date":"2020-07-01",
            "nameEn":"Canada Day",
            "nameFr":"Fête du Canada",
            "federal":1
         },
         {
            "id":24,
            "date":"2020-09-07",
            "nameEn":"Labour Day",
            "nameFr":"Fête du travail",
            "federal":1
         },
         {
            "id":25,
            "date":"2020-10-12",
            "nameEn":"Thanksgiving",
            "nameFr":"Action de grâce",
            "federal":1
         },
         {
            "id":27,
            "date":"2020-12-25",
            "nameEn":"Christmas Day",
            "nameFr":"Noël",
            "federal":1
         },
         {
            "id":28,
            "date":"2020-03-27",
            "nameEn":"Boxing Day",
            "nameFr":"Lendemain de Noël",
            "federal":1
         }
      ]`

	// TODO BOXING DAY IS ACTUALLY - "date":"2020-12-28",

	// Declared an empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(holidays), &results)

	for _, result := range results {
		today := time.Now()
		endOfWeek := today.AddDate(0, 0, 6)

		holidayDate, _ := time.Parse("2006-01-02", result["date"].(string))

		if (holidayDate.YearDay() >= today.YearDay()) && (holidayDate.YearDay() <= endOfWeek.YearDay()) {
			holidayString = holidayString + result["nameEn"].(string) + " "
		}
	}

	celebrationsString := ":robot: Beep Boop :robot:" + "\n" + anniString + bdayString + holidayString + "\n" + ":robot: Boop Beep :robot:"

	//Define bot account
	bot := mattermost.GetBotUser(*api)
	//Trigger message to specified channel
	mattermost.MessageMembers(*api, channelName, teamName, bot, celebrationsString)
}
