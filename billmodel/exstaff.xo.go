// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Exstaff represents a row from 'equinox.exstaff'.
type Exstaff struct {
	Exstafid       sql.NullInt64   `json:"exstafid"`       // exstafid
	Exstaffname    sql.NullString  `json:"exstaffname"`    // exstaffname
	Exstaffreplace sql.NullString  `json:"exstaffreplace"` // exstaffreplace
	Exstaffsparec1 sql.NullString  `json:"exstaffsparec1"` // exstaffsparec1
	Exstaffsparen1 sql.NullFloat64 `json:"exstaffsparen1"` // exstaffsparen1
	Exstaffspared1 pq.NullTime     `json:"exstaffspared1"` // exstaffspared1
	EquinoxPrn     sql.NullInt64   `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec
}

func AllExstaff(db XODB, callback func(x Exstaff) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`exstafid, exstaffname, exstaffreplace, exstaffsparec1, exstaffsparen1, exstaffspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.exstaff `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Exstaff{}

		// scan
		err = q.Scan(&e.Exstafid, &e.Exstaffname, &e.Exstaffreplace, &e.Exstaffsparec1, &e.Exstaffsparen1, &e.Exstaffspared1, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// ExstaffByEquinoxLrn retrieves a row from 'equinox.exstaff' as a Exstaff.
//
// Generated from index 'exstaff_pkey'.
func ExstaffByEquinoxLrn(db XODB, equinoxLrn int64) (*Exstaff, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`exstafid, exstaffname, exstaffreplace, exstaffsparec1, exstaffsparen1, exstaffspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.exstaff ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Exstaff{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Exstafid, &e.Exstaffname, &e.Exstaffreplace, &e.Exstaffsparec1, &e.Exstaffsparen1, &e.Exstaffspared1, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
