// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Event represents a row from 'equinox.events'.
type Event struct {
	Eventid      sql.NullInt64  `json:"eventid"`      // eventid
	Evname       sql.NullString `json:"evname"`       // evname
	Evfrequency  sql.NullString `json:"evfrequency"`  // evfrequency
	Evstart      pq.NullTime    `json:"evstart"`      // evstart
	Evend        pq.NullTime    `json:"evend"`        // evend
	Evtimeofday  pq.NullTime    `json:"evtimeofday"`  // evtimeofday
	Evdayofweek  sql.NullString `json:"evdayofweek"`  // evdayofweek
	Evdayofmonth sql.NullInt64  `json:"evdayofmonth"` // evdayofmonth
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec
}

func AllEvent(db XODB, callback func(x Event) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`eventid, evname, evfrequency, evstart, evend, evtimeofday, evdayofweek, evdayofmonth, equinox_lrn, equinox_sec ` +
		`FROM equinox.events `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Event{}

		// scan
		err = q.Scan(&e.Eventid, &e.Evname, &e.Evfrequency, &e.Evstart, &e.Evend, &e.Evtimeofday, &e.Evdayofweek, &e.Evdayofmonth, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// EventByEquinoxLrn retrieves a row from 'equinox.events' as a Event.
//
// Generated from index 'events_pkey'.
func EventByEquinoxLrn(db XODB, equinoxLrn int64) (*Event, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`eventid, evname, evfrequency, evstart, evend, evtimeofday, evdayofweek, evdayofmonth, equinox_lrn, equinox_sec ` +
		`FROM equinox.events ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Event{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Eventid, &e.Evname, &e.Evfrequency, &e.Evstart, &e.Evend, &e.Evtimeofday, &e.Evdayofweek, &e.Evdayofmonth, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
