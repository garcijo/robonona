package mattermost

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-server/model"
)

// ClientV4 is just mattermost's model.Client4
type ClientV4 model.Client4

// NewMatterMostClient returns a NewAPIV4Client after logging in with the provided username and password
func NewMatterMostClient(url string, username string, password string) *model.Client4 {
	api := model.NewAPIv4Client(url)

	api.Login(username, password)

	return api
}

// GetActiveChannelMembers retrieves a list of active members in a given channel for the specified teamName
func GetActiveChannelMembers(m model.Client4, teamName string, channelName string) model.UserSlice {
	team, resp := m.GetTeamByName(teamName, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", team)

	channel, resp := m.GetChannelByName(channelName, team.Id, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", channel)

	members, resp := m.GetUsersInChannel(channel.Id, 0, 100, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}

	slice := model.UserSlice(members)
	return slice.FilterByActive(true)
}

// GetBotUser gets information about the user this program is running as.
func GetBotUser(m model.Client4) *model.User {
	user, resp := m.GetMe("")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}

	return user
}

// GetBotUser gets information about the user this program is running as.
func GetMattermostUsernames(m model.Client4, employeeData []Employee) (employees []Employee) {
	for _, employee := range employeeData {
		search := &model.UserSearch{
			Term:          employee.Email,
			AllowInactive: false,
			Limit:         1,
		}
		users, _ := m.SearchUsers(search)
		if len(users) > 0 {
			employee.MattermostUsername = users[0].Username
			employees = append(employees, employee)
		}
	}

	return employees
}

// GetMattermostUsername gets a username from the employee's email address
func GetMattermostUsername(m model.Client4, employeeData Employee) (employee Employee) {
	search := &model.UserSearch{
		Term:          employeeData.Email,
		AllowInactive: false,
		Limit:         1,
	}
	users, _ := m.SearchUsers(search)
	employeeData.MattermostUsername = users[0].Username
	employee = employeeData

	return employee
}

// MessageMembers sends a message via mattermost to each set of pairs
func MessageMembers(m model.Client4, channelName string, teamName string, botUser *model.User, birthdayString string) {
	team, _ := m.GetTeamByName(teamName, "")
	channel, resp := m.GetChannelByName(channelName, team.Id, "")

	fmt.Printf("Channel: %v", channel)
	fmt.Printf("Received response: %v", resp)

	post := &model.Post{
		ChannelId: channel.Id,
		UserId:    botUser.Id,
		Message:   birthdayString,
	}
	// 	fmt.Printf("%v", post)
	_, resp = m.CreatePost(post)
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
}
