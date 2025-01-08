package database

import (
	"database/sql"
	"fmt"
  
  "github.com/kenmalik/appetizer/types"
)

type ApplicationModel struct {
	DB *sql.DB
}

func (m ApplicationModel) All() ([]types.Application, error) {
	rows, err := m.DB.Query(`SELECT company, position, location, date_posted, date_applied, url, notes, status 
  FROM applications LEFT JOIN statuses
  ON applications.status_id = statuses.id`)
	if err != nil {
		return nil, fmt.Errorf("Error querying database - %v", err)
	}
	defer rows.Close()

	var applications []types.Application

	for rows.Next() {
		var application types.ReadApplication
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

		applications = append(applications, types.NewApplication(application))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}
