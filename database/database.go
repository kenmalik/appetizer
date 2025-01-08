package database

import "database/sql"

type Application struct {
	Id            int
	Company       string
	PositionTitle string
	Location      string
	DatePosted    string
	DateApplied   string
	Url           string
	Notes         string
	Status        int
}

type ReadApplication struct {
	Id            int
	Company       string
	PositionTitle string
	Location      *string
	DatePosted    *string
	DateApplied   *string
	Url           *string
	Notes         *string
	Status        int
}

type ApplicationModel struct {
	DB *sql.DB
}

func newApplication(ra ReadApplication) Application {
	var application Application
	application.Id = ra.Id
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
	rows, err := m.DB.Query("SELECT * FROM applications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []Application

	for rows.Next() {
		var application ReadApplication
		err = rows.Scan(
			&application.Id,
			&application.Company,
			&application.PositionTitle,
			&application.Location,
			&application.DatePosted,
			&application.DateApplied,
			&application.Url,
			&application.Notes,
			&application.Status,
		)
		if err != nil {
			return nil, err
		}

		applications = append(applications, newApplication(application))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}
