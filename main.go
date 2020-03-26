package main

import (
	"encoding/json"
	"fmt"
	"github.com/garcijo/robonona/mattermost"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
	holidaysURL := "https://canada-holidays.ca/api/v1/provinces/ON"

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

	resp, err := http.Get(holidaysURL)
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	var generic2 map[string]interface{}
	generic2 = generic["provinces"]
	var generic3 map[string]interface{}
	generic3 = generic["holidays"]

	//	today := time.Now()
	//Take current date and add 6 days to cover the whole week (Monday - Sunday)
	//	endOfWeek := today.AddDate(0, 0, 6)

	// map[province:map[id:ON nameEn:Ontario nameFr:Ontario holidays:[map[id:1 date:2020-01-01 nameEn:New Year’s Day nameFr:Jour de l’An federal:1] map[id:4 date:2020-02-17 nameEn:Family Day nameFr:Fête de la famille federal:0] map[federal:1 id:7 date:2020-04-10 nameEn:Good Friday nameFr:Vendredi saint] map[id:11 date:2020-05-18 nameEn:Victoria Day nameFr:Fête de la Reine federal:1] map[nameFr:Fête du Canada federal:1 id:15 date:2020-07-01 nameEn:Canada Day] map[id:24 date:2020-09-07 nameEn:Labour Day nameFr:Fête du travail federal:1] map[date:2020-10-12 nameEn:Thanksgiving nameFr:Action de grâce federal:1 id:25] map[federal:1 id:27 date:2020-12-25 nameEn:Christmas Day nameFr:Noël] map[federal:1 id:28 date:2020-12-28 nameEn:Boxing Day nameFr:Lendemain de Noël]] nextHoliday:map[nameEn:Good Friday nameFr:Vendredi saint federal:1 id:7 date:2020-04-10]]]

	for _, holiday := range generic3 {
		fmt.Println(holiday)
		/*
				employeeBirthdayDate,_ := time.Parse("2006-01-02", employee.DateOfBirth)

				monthString := strconv.Itoa(int(employeeBirthdayDate.Month()))
				if int(employeeBirthdayDate.Month()) < 10 {
					monthString = "0" + strconv.Itoa(int(employeeBirthdayDate.Month()))
				}
				dayString := strconv.Itoa(int(employeeBirthdayDate.Day()))
				if int(employeeBirthdayDate.Day()) < 10 {
					dayString = "0" + strconv.Itoa(int(employeeBirthdayDate.Day()))
				}

				bdayString := strconv.Itoa(today.Year()) + "-" + monthString + "-" + dayString
				employeeBirthday,_ := time.Parse("2006-01-02", bdayString)

				hireDate,_ := time.Parse("2006-01-02", employee.HireDate)
				monthHireString := strconv.Itoa(int(hireDate.Month()))
				if int(hireDate.Month()) < 10 {
					monthHireString = "0" + strconv.Itoa(int(hireDate.Month()))
				}
				dayHireString := strconv.Itoa(int(hireDate.Day()))
				if int(hireDate.Day()) < 10 {
					dayHireString = "0" + strconv.Itoa(int(hireDate.Day()))
				}
				hireString := strconv.Itoa(today.Year()) + "-" + monthHireString + "-" + dayHireString
				employeeHireDate,_ := time.Parse("2006-01-02", hireString)

				if (employeeBirthday.YearDay() >= today.YearDay()) && (employeeBirthday.YearDay() <= endOfWeek.YearDay()) {
					birthdays = append(birthdays, employee)
				}
				if (hireDate.Year() < today.Year()) && (employeeHireDate.YearDay() >= today.YearDay()) && (employeeHireDate.YearDay() <= endOfWeek.YearDay()) {
					 anniversaries = append(anniversaries, employee)
				 }
			celebrations := Celebrations{birthdays, anniversaries}
		*/
	}

	celebrationsString := ":robot: Beep Boop :robot:" + "\n" + anniString + bdayString + "\n" + ":robot: Boop Beep :robot:"

	//Define bot account
	bot := mattermost.GetBotUser(*api)
	//Trigger message to specified channel
	mattermost.MessageMembers(*api, channelName, teamName, bot, celebrationsString)
}
