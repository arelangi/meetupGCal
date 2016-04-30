meetupGCal
----------

This Go library provides some useful functions to extract info from the meetup.com
and interfaces with Google calendar API to add events.

In order to get the tokens for accessing the Calendar, follow the steps here.
https://developers.google.com/google-apps/calendar/quickstart/go#prerequisites

Apart from this, you will also need to provide a config file of the following format

````json
{
	"calendarId":"This is the id of the calendar you want to update. The id can be found in the settings",
	"meetupKey": "The meetup API key",
	"groupsFile":"https://raw.githubusercontent.com/logzer0/dfw-tech-calendar/master/groups.csv -- sample file",
	"zip": "zip code. The program looks for groups in a radius of 50mi"
}
````
