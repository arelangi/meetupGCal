package meetupGCal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"fmt"
)

var (
	config         Config
	ConfigFilePath string
)

const (
	baseURL = "https://api.meetup.com/"
)

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

func UpdateCalendar() {
	baseURL := "https://api.meetup.com/"
	eventURLParams := "/events?&photo-host=public&page=" + config.LookupEvents + "&key="
	var meetupGroups []Group
	var err error

	if err = getConfig(ConfigFilePath); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if meetupGroups, err = getTechGroupsInDallas(); err != nil {
		log.Println(err)
	}

	for _, group := range meetupGroups {
		time.Sleep(time.Second) //Delay introduced to be under meetup api rate limits
		var nextEvents []Event
		eventURL := baseURL + group.Urlname + eventURLParams + config.MeetupKey

		resp, err := Call(eventURL)
		if err != nil {
			log.Println(err)
			continue
		}

		err = json.Unmarshal(resp, &nextEvents)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("Group: ", group.Urlname," Number of Events:",len(nextEvents)," Lookup:",config.LookupEvents, "url:",eventURL)
		for _, eachEvent := range nextEvents {
			AddEventToGCal(ConvertMeetupEventToGCalEvent(group, eachEvent))
		}
	}
}

func Call(url string) (resp []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
