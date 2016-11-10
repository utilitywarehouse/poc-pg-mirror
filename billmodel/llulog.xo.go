// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Llulog represents a row from 'equinox.llulog'.
type Llulog struct {
	Llucsource    sql.NullString `json:"llucsource"`    // llucsource
	Lluccli       sql.NullString `json:"lluccli"`       // lluccli
	Lluccheckdate pq.NullTime    `json:"lluccheckdate"` // lluccheckdate
	Llucchecktime pq.NullTime    `json:"llucchecktime"` // llucchecktime
	Llucstatus    sql.NullString `json:"llucstatus"`    // llucstatus
	Llucexchange  sql.NullString `json:"llucexchange"`  // llucexchange
	Llucsparec1   sql.NullString `json:"llucsparec1"`   // llucsparec1
	Llucsparec2   sql.NullString `json:"llucsparec2"`   // llucsparec2
	Llucsparec3   sql.NullString `json:"llucsparec3"`   // llucsparec3
	Llucspared1   pq.NullTime    `json:"llucspared1"`   // llucspared1
	Llucspared2   pq.NullTime    `json:"llucspared2"`   // llucspared2
	Llucspared3   pq.NullTime    `json:"llucspared3"`   // llucspared3
	Llucsparen1   sql.NullInt64  `json:"llucsparen1"`   // llucsparen1
	Llucsparen2   sql.NullInt64  `json:"llucsparen2"`   // llucsparen2
	Llucsparen3   sql.NullInt64  `json:"llucsparen3"`   // llucsparen3
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec
}

// LlulogByEquinoxLrn retrieves a row from 'equinox.llulog' as a Llulog.
//
// Generated from index 'llulog_pkey'.
func LlulogByEquinoxLrn(db XODB, equinoxLrn int64) (*Llulog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`llucsource, lluccli, lluccheckdate, llucchecktime, llucstatus, llucexchange, llucsparec1, llucsparec2, llucsparec3, llucspared1, llucspared2, llucspared3, llucsparen1, llucsparen2, llucsparen3, equinox_lrn, equinox_sec ` +
		`FROM equinox.llulog ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	l := Llulog{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&l.Llucsource, &l.Lluccli, &l.Lluccheckdate, &l.Llucchecktime, &l.Llucstatus, &l.Llucexchange, &l.Llucsparec1, &l.Llucsparec2, &l.Llucsparec3, &l.Llucspared1, &l.Llucspared2, &l.Llucspared3, &l.Llucsparen1, &l.Llucsparen2, &l.Llucsparen3, &l.EquinoxLrn, &l.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
