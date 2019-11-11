package main

import (
	"log"
	"os"
	"math/rand"
	"time"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/garcijo/robonona/mattermost"
	wr "github.com/mroth/weightedrand"
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

	//Get employees directory
	employees,_ := bambooApi.GetDirectory()
	//Take directory and get extra information for each employee (birthday, anniversary)
	employeeData,_ := bambooApi.GetEmployeeData(employees.Employees)
	//Filter only the employees with celebrations within the next week
	celebrations := mattermost.FilterCelebrations(employeeData)
	//Randomly select one more person for a fake birthday
	rand.Seed(time.Now().UnixNano())

	fakeBday := employeeData[rand.Intn(len(employeeData))]
	fakeBdayName := fmt.Sprintf("@%s.%s", fakeBday.FirstName, fakeBday.LastName)
	fakeBdayString := fmt.Sprintf(":shocked_pikachu: And *finally*, the happiest of all birthdays to %s ! :shocked_pikachu:", fakeBdayName)

	//Determine if humanity will perish this week or not
	c := wr.NewChooser(
		wr.Choice{Item: "true", Weight: 1},
		wr.Choice{Item: "false", Weight: 9},
	)
	destroy := c.Pick().(string)
	destroyHumans := ""
	if destroy == "true" {
		destroyHumans = ":evilparrot: On Thursday I will destroy all humans :evilparrot: \n"
	}

	bdayString := mattermost.ParseBirthdays(celebrations.Birthdays)
	anniString := mattermost.ParseAnniversaries(celebrations.Anniversaries)
	celebrationsString := ":robot: Beep Boop :robot:" + "\n" + bdayString + anniString + destroyHumans + fakeBdayString + "\n" + ":robot: Boop Beep :robot:"
	bot := mattermost.GetBotUser(*api)

	mattermost.MessageMembers(*api, channelName, teamName, bot, celebrationsString)
}
