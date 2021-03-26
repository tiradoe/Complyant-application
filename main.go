package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Applicant struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type Application struct {
	ApplicantToken string `json:"applicant_token"`
	ListingSlug    string `json:"listing_slug"`
	GithubUsername string `json:"github_username"`
	Phone          string `json:"phone"`
	Address1       string `json:"address1"`
	Address2       string `json:"address2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Zip            string `json:"zip"`
	Cover          string `json:"cover"`
	Resume         string `json:"resume"`
	ResumeUrl      string `json:"resume_url"`
	SocialMedia    string `json:"social_media"`
	Projects       string `json:"projects"`
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "submit" {
		submitApplication()
	} else {
		createApplicantAccount()
	}

}

func createApplicantAccount() {
	applicant := Applicant{
		Email:     "rony@starkindustries.com",
		FirstName: "Tony",
		LastName:  "Stark",
	}

	applicant_json, err := json.Marshal(applicant)
	errorCheck(err)
	sendRequest(applicant_json, "http://cars.complyant.co/api/applicant")
}

func submitApplication() {
	application := Application{
		ApplicantToken: "abcdedfg",
		ListingSlug:    "listing-slug-here",
		GithubUsername: "WizardRockstarCodeNinja",
		Phone:          "(555)-123-4567",
		Address1:       "1234 Street St",
		Address2:       "Suite 500",
		City:           "Denver",
		State:          "Colorado",
		Zip:            "12345",
		Cover:          "cover letter goes here",
		Resume:         "resume goes here",
		ResumeUrl:      "resume.com",
		SocialMedia:    "twitter",
		Projects:       "projects",
	}

	application_json, err := json.Marshal(application)
	errorCheck(err)

	sendRequest(application_json, "http://cars.complyant.co/api/application")
}

func sendRequest(json_data []byte, url string) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	errorCheck(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorCheck(err)

	log.Println(string(body))
}

func errorCheck(err error) {
	if err != nil {
		log.Println("Oh, get a job?  Just get a job? Why don’t I strap on my job helmet and squeeze down into a job cannon and fire off into Jobland, where jobs grow on jobbies!")
		log.Fatal(err)
	}
}
