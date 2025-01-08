package database

import (
	"database/sql"
	"fmt"
)

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

type ApplicationModel struct {
	DB *sql.DB
}

func newApplication(ra ReadApplication) Application {
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

func (m ApplicationModel) All() ([]Application, error) {
	rows, err := m.DB.Query("SELECT company, position, location, date_posted, date_applied, url, notes, status FROM applications LEFT JOIN statuses ON applications.status_id = statuses.id")
	if err != nil {
		return nil, fmt.Errorf("Error querying database - %v", err)
  }
	defer rows.Close()

	var applications []Application

	for rows.Next() {
		var application ReadApplication
		err = rows.Scan(
			&application.Company,
			&application.Position,
			&application.Location,
			&application.DatePosted,
			&application.DateApplied,
			&application.Url,
			&application.Notes,
			&application.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row - %v", err)
		}

		applications = append(applications, newApplication(application))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}
