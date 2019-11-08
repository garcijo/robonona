package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/garcijo/robonona/mattermost"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
	log.Print("No .env file found")
    }
}

func main() {
	serverURL := os.Getenv("ROBONONA_MATTERMOST_URL")

	botUserName := os.Getenv("ROBONONA_USERNAME")
	botPassword := os.Getenv("ROBONONA_PASSWORD")

	teamName := os.Getenv("ROBONONA_TEAM_NAME")
	channelName := os.Getenv("ROBONONA_CHANNEL_NAME")

	bambooURL := os.Getenv("ROBONONA_BAMBOO_URL")
	bambooKey := os.Getenv("ROBONONA_BAMBOO_KEY")

	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)

	bambooApi := mattermost.BambooHR(bambooURL, bambooKey)
// 	bambooApi.Debug(true)
	employees,_ := bambooApi.GetDirectory()

	employeeData,_ := bambooApi.GetEmployeeData(employees.Employees[0:50])

	celebrations := mattermost.FilterCelebrations(employeeData)

	bdayString := mattermost.ParseBirthdays(celebrations.Birthdays)
	anniString := mattermost.ParseAnniversaries(celebrations.Anniversaries)
	celebrationsString := ":robot: Beep Boop :robot:" + "\n" + bdayString + anniString + ":robot: Boop Beep :robot:"
	bot := mattermost.GetBotUser(*api)

	mattermost.MessageMembers(*api, channelName, teamName, bot, celebrationsString)
}
