package batch

import (
	"github.com/state303/go-discogs/model"
	"gorm.io/gorm/clause"
)

const (
	id                = "id"
	dataQuality       = "data_quality"
	contactInfo       = "contact_info"
	name              = "name"
	profile           = "profile"
	country           = "country"
	realName          = "real_name"
	title             = "title"
	releasedYear      = "released_year"
	releasedMonth     = "released_month"
	mainReleaseID     = "main_release_id"
	releasedDay       = "released_day"
	listedReleaseDate = "listed_release_date"
	isMaster          = "is_master"
	notes             = "notes"
	status            = "status"
)

var (
	styleConstraint = clause.OnConflict{Columns: getClauseColumns([]string{id, name}), OnConstraint: "style_name_key", DoNothing: true}
	genreConstraint = clause.OnConflict{Columns: getClauseColumns([]string{id, name}), OnConstraint: "genre_name_key", DoNothing: true}
)

func ExtractClause(i interface{}) clause.OnConflict {
	switch i.(type) {
	case *model.Artist:
		return updateOnIdConflict(dataQuality, name, profile, realName)
	case *model.Label:
		return updateOnIdConflict(contactInfo, dataQuality, name, profile)
	case *model.Master:
		return updateOnIdConflict(dataQuality, title, releasedYear)
	case *model.Release:
		return updateOnIdConflict(title, country, dataQuality, releasedYear, releasedMonth, releasedDay, listedReleaseDate, isMaster, notes, status)
	case *model.MasterMainRelease:
		return updateOnIdConflict(mainReleaseID)
	case *model.Style:
		return styleConstraint
	case *model.Genre:
		return genreConstraint
	}
	return clause.OnConflict{DoNothing: true}
}

func getClauseColumns(columns []string) []clause.Column {
	clauseCols := make([]clause.Column, len(columns))
	for i, v := range columns {
		clauseCols[i] = clause.Column{Name: v}
	}
	return clauseCols
}

func updateOnIdConflict(columns ...string) clause.OnConflict {
	return onConflictDoUpdate([]string{id}, columns)
}

func onConflictDoUpdate(conflictColumns []string, updateColumns []string) clause.OnConflict {
	return clause.OnConflict{
		Columns:   getClauseColumns(conflictColumns),
		DoUpdates: clause.AssignmentColumns(updateColumns)}
}
