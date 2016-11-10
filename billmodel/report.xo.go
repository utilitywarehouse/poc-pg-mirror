// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Report represents a row from 'equinox.reports'.
type Report struct {
	ReportName  sql.NullString `json:"report_name"`  // report_name
	ReportDesc  sql.NullString `json:"report_desc"`  // report_desc
	ReportStart pq.NullTime    `json:"report_start"` // report_start
	ReportEnd   pq.NullTime    `json:"report_end"`   // report_end
	ReportDate  pq.NullTime    `json:"report_date"`  // report_date
	ReportBy    sql.NullString `json:"report_by"`    // report_by
	ReportFreq  sql.NullInt64  `json:"report_freq"`  // report_freq
	EquinoxLrn  int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec  sql.NullInt64  `json:"equinox_sec"`  // equinox_sec
}

// ReportByEquinoxLrn retrieves a row from 'equinox.reports' as a Report.
//
// Generated from index 'reports_pkey'.
func ReportByEquinoxLrn(db XODB, equinoxLrn int64) (*Report, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`report_name, report_desc, report_start, report_end, report_date, report_by, report_freq, equinox_lrn, equinox_sec ` +
		`FROM equinox.reports ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	r := Report{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&r.ReportName, &r.ReportDesc, &r.ReportStart, &r.ReportEnd, &r.ReportDate, &r.ReportBy, &r.ReportFreq, &r.EquinoxLrn, &r.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
