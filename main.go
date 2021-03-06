package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Files []FitFile `json:"activityFiles"`
}

type FitFile struct {
	UserId             string `json:"userId"`
	UserAccessToken    string `json:"userAccessToken"`
	FileType           string `json:"fileType"`
	CallBackUrl        string `json:"callbackURL"`
	StartTimeInSeconds int32  `json:"startTimeInSeconds"`
	Manual             bool   `json:"manual"`
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// Client to wrap around real client and allow mocks
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func main() {
	file := FitFile{
		UserId:             "adf_xdsd",
		UserAccessToken:    "dfdzf",
		FileType:           "FIT",
		CallBackUrl:        "http://localhost:8080/fit/",
		StartTimeInSeconds: 2323232,
		Manual:             false,
	}

	file2 := FitFile{
		UserId:             "adf_xdsd",
		UserAccessToken:    "dfdzf",
		FileType:           "FIT",
		CallBackUrl:        "http://localhost:8080/fit2/",
		StartTimeInSeconds: 2323292,
		Manual:             false,
	}

	file3 := FitFile{
		UserId:             "adf_xdsd",
		UserAccessToken:    "dfdzf",
		FileType:           "FIT",
		CallBackUrl:        "http://localhost:8080/fit3/",
		StartTimeInSeconds: 2323292,
		Manual:             false,
	}

	file4 := FitFile{
		UserId:             "adf_xdsd",
		UserAccessToken:    "dfdzf",
		FileType:           "FIT",
		CallBackUrl:        "http://localhost:8080/fit4/?id=12",
		StartTimeInSeconds: 2323292,
		Manual:             false,
	}

	message := Message{[]FitFile{file, file2, file3, file4}}

	requestBody, err := json.Marshal(message)
	fmt.Printf("%s\n", requestBody)

	if err != nil {
		log.Fatalln(err)
		return
	}
	address := "http://localhost:8000/rowers/garmin/files/"
	req, err := http.NewRequest("POST", address, bytes.NewBuffer([]byte(requestBody)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Response status code %v\n", resp.StatusCode)
		return
	}

	fmt.Println("Ping OK")

	http.HandleFunc("/fit/", serveFit)
	http.HandleFunc("/fit2/", serveFit2)
	http.HandleFunc("/fit3/", serveFit3)
	http.HandleFunc("/fit4/", serveFit4)

	http.ListenAndServe(":8080", nil)
}

func serveFit(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "5206944119.fit")
}

func serveFit2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "3x250m.fit")
}

func serveFit3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "davidfit.fit")
}

func serveFit4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "FR-GAR-2020-07-11.fit")
}
