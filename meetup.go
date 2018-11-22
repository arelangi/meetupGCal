package meetupGCal

import (
	"fmt"
	"strings"
	"time"

	"gitlab.logzero.in/arelangi/mlog"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func AddEventToGCal(event *calendar.Event, secretConf []byte) {
	ctx := context.Background()

	gConfig, err := google.ConfigFromJSON(secretConf, calendar.CalendarScope)
	if err != nil {
		mlog.Error(fmt.Sprintf("Unable to parse client secret file to gConfig: %v", err.Error()))
	}
	client := getClient(ctx, gConfig)

	srv, err := calendar.New(client)
	if err != nil {
		mlog.Error(fmt.Sprintf("Unable to retrieve calendar Client %v", err.Error()))
	}
	gCalInsertedEvent, err := srv.Events.Insert(config.CalendarId, event).Do()
	if err != nil {
		if !strings.Contains(err.Error(), "duplicate") {
			mlog.Error(fmt.Sprintf("Unable to create event. %v\n", err.Error()))
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
