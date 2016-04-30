package meetupGCal

//Group is the data type for a meetup group
type Group struct {
	Score                float64 `json:"score"`
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Link                 string  `json:"link"`
	Urlname              string  `json:"urlname"`
	Description          string  `json:"description"`
	Created              int64   `json:"created"`
	City                 string  `json:"city"`
	Country              string  `json:"country"`
	LocalizedCountryName string  `json:"localized_country_name"`
	State                string  `json:"state"`
	JoinMode             string  `json:"join_mode"`
	Visibility           string  `json:"visibility"`
	Lat                  float64 `json:"lat"`
	Lon                  float64 `json:"lon"`
	Members              int     `json:"members"`
	Organizer            struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Bio   string `json:"bio"`
		Photo struct {
			ID          int    `json:"id"`
			HighresLink string `json:"highres_link"`
			PhotoLink   string `json:"photo_link"`
			ThumbLink   string `json:"thumb_link"`
		} `json:"photo"`
	} `json:"organizer"`
	Who        string `json:"who"`
	GroupPhoto struct {
		ID          int    `json:"id"`
		HighresLink string `json:"highres_link"`
		PhotoLink   string `json:"photo_link"`
		ThumbLink   string `json:"thumb_link"`
	} `json:"group_photo"`
	Timezone  string `json:"timezone"`
	NextEvent struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		YesRsvpCount int    `json:"yes_rsvp_count"`
		Time         int64  `json:"time"`
		UtcOffset    int    `json:"utc_offset"`
	} `json:"next_event"`
	Category struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Shortname string `json:"shortname"`
		SortName  string `json:"sort_name"`
	} `json:"category"`
	Photos []struct {
		ID          int    `json:"id"`
		HighresLink string `json:"highres_link"`
		PhotoLink   string `json:"photo_link"`
		ThumbLink   string `json:"thumb_link"`
	} `json:"photos"`
}

//Event is the data type for a meetup event
type Event struct {
	Created       int64  `json:"created"`
	Duration      int    `json:"duration"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	RsvpLimit     int    `json:"rsvp_limit"`
	Status        string `json:"status"`
	Time          int64  `json:"time"`
	Updated       int64  `json:"updated"`
	UtcOffset     int    `json:"utc_offset"`
	WaitlistCount int    `json:"waitlist_count"`
	YesRsvpCount  int    `json:"yes_rsvp_count"`
	Group         struct {
		Created  int64   `json:"created"`
		Name     string  `json:"name"`
		ID       int     `json:"id"`
		JoinMode string  `json:"join_mode"`
		Lat      float64 `json:"lat"`
		Lon      float64 `json:"lon"`
		Urlname  string  `json:"urlname"`
		Who      string  `json:"who"`
	} `json:"group"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Venue       struct {
		ID                   int     `json:"id"`
		Name                 string  `json:"name"`
		Lat                  float64 `json:"lat"`
		Lon                  float64 `json:"lon"`
		Repinned             bool    `json:"repinned"`
		Address1             string  `json:"address_1"`
		Address2             string  `json:"address_2"`
		Address3             string  `json:"address_3"`
		City                 string  `json:"city"`
		Country              string  `json:"country"`
		LocalizedCountryName string  `json:"localized_country_name"`
	} `json:"venue"`
	Visibility string `json:"visibility"`
}

//json config to provide the data
type Config struct {
	CalendarId   string `json:"calendarId"`
	MeetupKey    string `json:"meetupKey"`
	GroupsFile   string `json:"groupsFile"`
	LookupEvents string `json:"lookup"`
	Zip          string `json:"zip"`
}
