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
	Fields []Field `json:"fields"`
	Employees []Employee `json:"employees"`
}

type Employee struct {
	Id   string    `json:"id,attr"`
	DisplayName string `json:"displayName,attr"`
	FirstName string `json:"firstName,attr"`
	LastName string `json:"lastName,attr"`
	DateOfBirth string `json:"dateOfBirth,attr"`
	HireDate string `json:"hireDate,attr"`
}

type Field struct {
	Id   string    `json:"id,attr"`
	Type string `json:"type,attr"`
	Name string `json:"name,attr"`
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
	log.Printf("%v", req)
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

// Get employee directory
func (b Bamboo) GetEmployee(id string) (employee Employee, err error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/v1/employees/%s?fields=firstName,lastName,displayName,dateOfBirth,hireDate", b.base, b.subdomain, id), nil)
	log.Printf("%v", req)
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



// Extract body from the HTTP response
func body(resp *http.Response) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

