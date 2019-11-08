package main

import (
	"fmt"
	"os"

	"github.com/garcijo/robonona/mattermost"
)

func main() {
	serverURL := os.Getenv("ROBONONA_MATTERMOST_URL")

	botUserName := os.Getenv("ROBONONA_USERNAME")
	botPassword := os.Getenv("ROBONONA_PASSWORD")

	teamName := os.Getenv("ROBONONA_TEAM_NAME")
	channelName := os.Getenv("ROBONONA_CHANNEL_NAME")

	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)

	members := mattermost.GetActiveChannelMembers(*api, teamName, channelName)
	fmt.Printf("%+v\n", members)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(members), channelName, teamName)
	bot := mattermost.GetBotUser(*api)
	pairs := mattermost.SplitIntoPairs(members, bot.Id)

// 	mattermost.MessageMembers(*api, pairs, bot)
}
