package mattermost

import (
	"fmt"
	"math/rand"
	"time"
)

func StringifyRobot() (robotString string) {
	robotEmojis := GetRobotEmojis()

	var emoji1 string
	var emoji2 string

	rand.Seed(time.Now().UnixNano())
	emoji1 = robotEmojis[rand.Intn(len(robotEmojis))]
	rand.Seed(time.Now().UnixNano())
	emoji2 = robotEmojis[rand.Intn(len(robotEmojis))]

	robotStrings := []string{
		fmt.Sprintf("%s Beep Boop %s", emoji1, emoji2),
		fmt.Sprintf("%s Boop %s", emoji1, emoji2),
		fmt.Sprintf("%s Beep %s", emoji1, emoji2),
		fmt.Sprintf("%s Boop Beep %s", emoji1, emoji2)}

		robotString = robotStrings[rand.Intn(len(robotStrings))]

	return
}
