package main

import (
	"log"
	"os"
	"math/rand"
	"time"
	"fmt"

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

	//Get employees directory
	employees,_ := bambooApi.GetDirectory()
	//Take directory and get extra information for each employee (birthday, anniversary)
	employeeData,_ := bambooApi.GetEmployeeData(employees.Employees)
	//Filter only the employees with celebrations within the next week
	celebrations := mattermost.FilterCelebrations(employeeData)

	//Randomly select one more person for a fake birthday
	rand.Seed(time.Now().UnixNano())
	fakeBday := employeeData[rand.Intn(len(employeeData))]
	fakeBdayUser := mattermost.GetMattermostUsername(*api, fakeBday)
	fakeBdayName := fmt.Sprintf("@%s", fakeBdayUser.MattermostUsername)
	fakeBdayString := fmt.Sprintf(":shocked_pikachu: And *finally*, the happiest of all birthdays to %s ! :shocked_pikachu:", fakeBdayName)

	birthdays := mattermost.GetMattermostUsernames(*api, celebrations.Birthdays)
	bdayString := mattermost.ParseBirthdays(birthdays)
	anniversaries := mattermost.GetMattermostUsernames(*api, celebrations.Anniversaries)
	anniString := mattermost.ParseAnniversaries(anniversaries)
	celebrationsString := ":robot: Beep Boop :robot:" + "\n" + anniString + bdayString + fakeBdayString + "\n" + ":robot: Boop Beep :robot:"

	//Define bot account
	bot := mattermost.GetBotUser(*api)
	//Trigger message to specified channel
	mattermost.MessageMembers(*api, channelName, teamName, bot, celebrationsString)
}
