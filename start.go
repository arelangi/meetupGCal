package meetupGCal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"gitlab.logzero.in/arelangi/mlog"
)

var (
	config Config
)

const (
	baseURL = "https://api.meetup.com/"
)

func init() {
	mlog.SetPrintStackTrace(false)
}

func getConfig(configFile string) (err error) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &config)
	return
}

func getTechGroupsInDallas() (groups []Group, err error) {
	var content []byte
	if content, err = Call(config.GroupsFile); err != nil {
		return
	}
	selectGroups := strings.Split(string(content), "|*|")
	for _, eachSelectedGroup := range selectGroups {
		split := strings.Split(eachSelectedGroup, ",")
		if len(split) > 1 {
			groups = append(groups, Group{Name: split[0], Urlname: split[1], Link: split[2]})
		}
	}
	return
}

func UpdateCalendar(configFilePath, secretFilePath string) {
	var err error
	if err = getConfig(configFilePath); err != nil {
		mlog.Error("Failed to get the config", mlog.Items{"error": err})
		os.Exit(1)
	}

	secretConf, err := ioutil.ReadFile(secretFilePath)
	if err != nil {
		mlog.Fatal(fmt.Sprintf("Unable to read client secret file: %v", err.Error()))
	}

	baseURL := "https://api.meetup.com/"
	eventURLParams := "/events?&photo-host=public&page=" + config.LookupEvents + "&key="
	var meetupGroups []Group

	if meetupGroups, err = getTechGroupsInDallas(); err != nil {
		mlog.Error("Failed to get the tech groups in dallas")
	}

	for _, group := range meetupGroups {
		time.Sleep(time.Second) //Delay introduced to be under meetup api rate limits
		var nextEvents []Event
		eventURL := baseURL + group.Urlname + eventURLParams + config.MeetupKey

		resp, err := Call(eventURL)
		if err != nil {
			mlog.Error("Failed to make a call", mlog.Items{"error": err, "group": group.Name})
			continue
		}

		err = json.Unmarshal(resp, &nextEvents)
		if err != nil {
			mlog.Error("Failed to unmarshal the data", mlog.Items{"error": err, "group": group.Name})
			continue
		}

		for _, eachEvent := range nextEvents {
			AddEventToGCal(ConvertMeetupEventToGCalEvent(group, eachEvent), secretConf)
		}
	}
}

func Call(url string) (resp []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		mlog.Error("Failed to make the call to the url", mlog.Items{"error": err, "url": url})
		return
	}
	defer response.Body.Close()
	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		mlog.Error("Failed to read the body", mlog.Items{"error": err, "url": url})
		return
	}
	return
}
