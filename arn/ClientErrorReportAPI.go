package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

var (
	_ api.Newable = (*ClientErrorReport)(nil)
)

// Create sets the data for a new report with data we received from the API request.
func (report *ClientErrorReport) Create(ctx aero.Context) error {
	data, err := ctx.Request().Body().JSONObject()

	if err != nil {
		return err
	}

	// Get user
	user := GetUserFromContext(ctx)

	// Create report
	report.ID = GenerateID("ClientErrorReport")
	report.Message = data["message"].(string)
	report.Stack = data["stack"].(string)
	report.FileName = data["fileName"].(string)
	report.LineNumber = int(data["lineNumber"].(float64))
	report.ColumnNumber = int(data["columnNumber"].(float64))
	report.Created = DateTimeUTC()

	if user != nil {
		report.CreatedBy = user.ID
	}

	report.Save()
	return nil
}

// Save saves the client error report in the database.
func (report *ClientErrorReport) Save() {
	DB.Set("ClientErrorReport", report.ID, report)
}

// Authorize returns an error if the given API request is not authorized.
func (report *ClientErrorReport) Authorize(ctx aero.Context, action string) error {
	if action == "create" {
		return nil
	}

	return errors.New("Action " + action + " not allowed")
}
