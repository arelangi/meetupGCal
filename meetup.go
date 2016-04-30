package meetupGCal

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func AddEventToGCal(event *calendar.Event) {
	ctx := context.Background()
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	gConfig, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to gConfig: %v", err)
	}
	client := getClient(ctx, gConfig)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}
	gCalInsertedEvent, err := srv.Events.Insert(config.CalendarId, event).Do()
	if err != nil {
		if !strings.Contains(err.Error(), "duplicate") {
			log.Printf("Unable to create event. %v\n", err)
		}
		return
	}
	fmt.Printf("Event created: %s\n", gCalInsertedEvent.HtmlLink)
}

func ConvertMeetupEventToGCalEvent(group Group, event Event) *calendar.Event {
	startTime := time.Unix(0, int64(time.Millisecond)*event.Time)
	endTime := time.Unix(0, int64(time.Millisecond)*(event.Time+int64(event.Duration)))
	if startTime == endTime {
		endTime = startTime.Add(time.Duration(int64(time.Hour) * 3))
	}
	gEvent := &calendar.Event{
		ICalUID:     event.ID,
		Summary:     group.Name + " - " + event.Name,
		Description: group.Link + "\n" + event.Description,
		Location:    event.Venue.Address1 + " " + event.Venue.Address2 + " " + event.Venue.Address3 + " " + event.Venue.City + " " + event.Venue.Country,
		Start: &calendar.EventDateTime{
			DateTime: startTime.Format(time.RFC3339),
			TimeZone: "America/Chicago",
		},
		End: &calendar.EventDateTime{
			DateTime: endTime.Format(time.RFC3339),
			TimeZone: "America/Chicago",
		},
		AnyoneCanAddSelf: true,
		Source: &calendar.EventSource{
			Title: group.Name,
			Url:   event.Link,
		},
	}
	return gEvent
}
