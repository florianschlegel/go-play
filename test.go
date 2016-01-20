package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/janhalfar/go-play/client"
)

// creating and chan
func getJobs() map[string]string {
	//jobs := make(map[string]string)
	jobs := getJobsUsingLiterals()
	jobs["google"] = "http://google.de/"
	jobs["new heise"] = "http://heise.de"
	delete(jobs, "deleteMe")
	return jobs
}

func getJobsUsingLiterals() map[string]string {
	return map[string]string{
		"google":                  "http://google.com/",
		"deleteMe":                "wrong",
		"the bestbytes home page": "http://bestbytes.de/",
	}
}

func modify(job client.Job) client.Job {
	job.Name += " modified"
	return job
}

func dataPlay() {
	var job client.Job

	job = client.Job{
		Name: "foo",
		URL:  "this is a url",
	}
	job = modify(job)
	fmt.Println("non refernce name   :", job.Name)
}

func modifyReference(job *client.Job) {
	job.Name += " modified"
}

func dataPlayReference() {
	var jobReference *client.Job

	jobReference = &client.Job{
		Name: "foo",
		URL:  "this is a url",
	}
	modifyReference(jobReference)
	fmt.Println("reference name      :", jobReference.Name)
}

type stringInt int

func (si stringInt) String() string {
	return fmt.Sprint("my value is: ", int(si))
}

type Person struct {
	Name    string    `json:"name" xml:"name"`
	Age     int       `json:"age" xml:"age"`
	Friends []*Friend `json:"friends" xml:"friends"`
}

type Friend struct {
	Name    string   `json:"name" xml:"name"`
	Surname string   `json:"surname" xml:"surname"`
	Address *Address `json:"address" xml:"address"`
}

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city" xml:",comment"`
	Addition string `json:"addition,omitempty" xml:"addition,attr"`
}

func main() {

	pers := &Person{
		Name: "Nicola",
		Age:  25,
		Friends: []*Friend{
			&Friend{
				Name:    "Mario",
				Surname: "Rossi",
				Address: &Address{
					Street:   "Via Corta",
					City:     "Gradisca",
					Addition: "App 2",
				},
			},
			&Friend{
				Name:    "Jan",
				Surname: "Halfar",
				Address: &Address{
					Street: "Holzweg 12",
					City:   "München",
				},
			},
		},
	}

	/*
		addr := &Address{
			Street:   "Pienzenauerstrasse 10",
			City:     "Munich",
			Addition: "in the back yard",
		}
	*/

	jsonBytes, _ := json.MarshalIndent(pers, "", "	")
	fmt.Println(string(jsonBytes))

	xmlBytes, _ := xml.MarshalIndent(pers, "", "	")
	fmt.Println(string(xmlBytes))

	/*
		job := client.NewJob("jan", "bestbytes.de")
		fmt.Println(job.GetDescription())
		otherJob := &client.Job{
			Name: "foo",
			URL:  "somewhere else",
		}
		fmt.Println(otherJob.GetDescription())
	*/

	/*
		jobs := getJobs()
		for name, url := range jobs {
			fmt.Println(name)
			client.CallURL(url)
		}
	*/
}
