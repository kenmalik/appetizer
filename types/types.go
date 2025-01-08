package types

type Application struct {
	Company     string
	Position    string
	Location    string
	DatePosted  string
	DateApplied string
	Url         string
	Notes       string
	Status      string
}

type ReadApplication struct {
	Id          int
	Company     string
	Position    string
	Location    *string
	DatePosted  *string
	DateApplied *string
	Url         *string
	Notes       *string
	Status      string
}

func NewApplication(ra ReadApplication) Application {
	var application Application
	application.Status = ra.Status

	if ra.Location == nil {
		application.Location = ""
	} else {
		application.Location = *ra.Location
	}
	if ra.DatePosted == nil {
		application.DatePosted = ""
	} else {
		application.DatePosted = *ra.DatePosted
	}
	if ra.DateApplied == nil {
		application.DateApplied = ""
	} else {
		application.DateApplied = *ra.DateApplied
	}
	if ra.Url == nil {
		application.Url = ""
	} else {
		application.Url = *ra.Url
	}
	if ra.Notes == nil {
		application.Notes = ""
	} else {
		application.Notes = *ra.Notes
	}

	return application
}

