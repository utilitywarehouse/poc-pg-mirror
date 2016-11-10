// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Repevent represents a row from 'equinox.repevent'.
type Repevent struct {
	Repname      sql.NullString  `json:"repname"`      // repname
	Repowner     sql.NullString  `json:"repowner"`     // repowner
	Repcat       sql.NullString  `json:"repcat"`       // repcat
	Repkey       sql.NullString  `json:"repkey"`       // repkey
	Repnavname   sql.NullString  `json:"repnavname"`   // repnavname
	Repdescr     sql.NullString  `json:"repdescr"`     // repdescr
	Repid        sql.NullFloat64 `json:"repid"`        // repid
	Repfreq      sql.NullFloat64 `json:"repfreq"`      // repfreq
	Repfirstused pq.NullTime     `json:"repfirstused"` // repfirstused
	Replastused  pq.NullTime     `json:"replastused"`  // replastused
	Repoutpath   sql.NullString  `json:"repoutpath"`   // repoutpath
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec
}

// RepeventByEquinoxLrn retrieves a row from 'equinox.repevent' as a Repevent.
//
// Generated from index 'repevent_pkey'.
func RepeventByEquinoxLrn(db XODB, equinoxLrn int64) (*Repevent, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`repname, repowner, repcat, repkey, repnavname, repdescr, repid, repfreq, repfirstused, replastused, repoutpath, equinox_lrn, equinox_sec ` +
		`FROM equinox.repevent ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	r := Repevent{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&r.Repname, &r.Repowner, &r.Repcat, &r.Repkey, &r.Repnavname, &r.Repdescr, &r.Repid, &r.Repfreq, &r.Repfirstused, &r.Replastused, &r.Repoutpath, &r.EquinoxLrn, &r.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
