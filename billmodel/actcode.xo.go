// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Actcode represents a row from 'equinox.actcodes'.
type Actcode struct {
	Actcode    sql.NullString  `json:"actcode"`     // actcode
	Actdesc    sql.NullString  `json:"actdesc"`     // actdesc
	Actsparec1 sql.NullString  `json:"actsparec1"`  // actsparec1
	Actsparen1 sql.NullFloat64 `json:"actsparen1"`  // actsparen1
	Actspared1 pq.NullTime     `json:"actspared1"`  // actspared1
	EquinoxLrn int64           `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64   `json:"equinox_sec"` // equinox_sec
}

func AllActcode(db XODB, callback func(x Actcode) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`actcode, actdesc, actsparec1, actsparen1, actspared1, equinox_lrn, equinox_sec ` +
		`FROM equinox.actcodes `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		a := Actcode{}

		// scan
		err = q.Scan(&a.Actcode, &a.Actdesc, &a.Actsparec1, &a.Actsparen1, &a.Actspared1, &a.EquinoxLrn, &a.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(a) {
			return nil
		}
	}

	return nil
}

// ActcodeByEquinoxLrn retrieves a row from 'equinox.actcodes' as a Actcode.
//
// Generated from index 'actcodes_pkey'.
func ActcodeByEquinoxLrn(db XODB, equinoxLrn int64) (*Actcode, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`actcode, actdesc, actsparec1, actsparen1, actspared1, equinox_lrn, equinox_sec ` +
		`FROM equinox.actcodes ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Actcode{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Actcode, &a.Actdesc, &a.Actsparec1, &a.Actsparen1, &a.Actspared1, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
