package mattermost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bamboo struct {
	subdomain, key string
	debug          bool
	client         *http.Client
	base           string
}

type Directory struct {
	Fields    []Field    `json:"fields"`
	Employees []Employee `json:"employees"`
}

type Employee struct {
	Id                 string `json:"id,attr"`
	DisplayName        string `json:"displayName,attr"`
	FirstName          string `json:"firstName,attr"`
	LastName           string `json:"lastName,attr"`
	PreferredName      string `json:"preferredName,attr"`
	DateOfBirth        string `json:"dateOfBirth,attr"`
	HireDate           string `json:"hireDate,attr"`
	Email              string `json:"workEmail,attr"`
	MattermostUsername string `json:"mattermostUsername,attr"`
}

type Field struct {
	Id   string `json:"id,attr"`
	Type string `json:"type,attr"`
	Name string `json:"name,attr"`
}

var WeekDays = map[string][]Employee{
	"Monday":    {},
	"Tuesday":   {},
	"Wednesday": {},
	"Thursday":  {},
	"Friday":    {},
	"Saturday":  {},
	"Sunday":    {},
}

// Start using the API here -
func BambooHR(subdomain, key string) Bamboo {
	return Bamboo{subdomain, key, false, &http.Client{}, "https://api.bamboohr.com/api/gateway.php"}
}

// Enable/disable debug mode. When debug mode is enabled,
// you will get additional logging showing the HTTP requests
// and responses
func (b *Bamboo) Debug(d bool) {
	b.debug = d
}

// Configure a custom HTTP client (e.g. to configure a proxy server)
func (b *Bamboo) Client(client *http.Client) {
	b.client = client
}

// Get employee directory
func (b Bamboo) GetDirectory() (dir Directory, err error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/v1/employees/directory", b.base, b.subdomain), nil)
	if err != nil {
		return
	}

	req.SetBasicAuth(b.key, "x")
	req.Header.Set("Accept", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := body(resp)
	if err != nil {
		return
	}
	if b.debug {
		log.Printf("Got response %s: %s", resp.Status, data)
	}

	err = json.Unmarshal(data, &dir)

	return
}

// Get employee. Will return names, birthday, and hire date
func (b Bamboo) GetEmployee(id string) (employee Employee) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/v1/employees/%s?fields=firstName,lastName,displayName,preferredName,dateOfBirth,hireDate", b.base, b.subdomain, id), nil)
	// 	log.Printf("%v", req)
	if err != nil {
		return
	}

	req.SetBasicAuth(b.key, "x")
	req.Header.Set("Accept", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := body(resp)
	if err != nil {
		return
	}
	if b.debug {
		log.Printf("Got response %s: %s", resp.Status, data)
	}

	err = json.Unmarshal(data, &employee)

	return
}

// Get employee data. Will return names, birthday, and hire date
func (b Bamboo) GetEmployeeData(employees []Employee) (employeeData []Employee, err error) {
	for _, employee := range employees {
		emp := b.GetEmployee(employee.Id)
		emp.Email = employee.Email
		employeeData = append(employeeData, emp)
	}

	return
}

// Extract body from the HTTP response
func body(resp *http.Response) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
