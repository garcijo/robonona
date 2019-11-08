package main

import (
	"fmt"
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
	employees, bambooError := bambooApi.GetDirectory()

	testEmployee, bambooError := bambooApi.GetEmployee(employees.Employees[0].Id)
	fmt.Printf("%+v\n", testEmployee)
	fmt.Printf("%+v\n", bambooError)

	members := mattermost.GetActiveChannelMembers(*api, teamName, channelName)
	fmt.Printf("%+v\n", members)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(members), channelName, teamName)
	// 	bot := mattermost.GetBotUser(*api)
	// 	pairs := mattermost.SplitIntoPairs(members, bot.Id)

	// 	mattermost.MessageMembers(*api, pairs, bot)
}
