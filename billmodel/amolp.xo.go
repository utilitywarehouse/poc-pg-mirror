// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Amolp represents a row from 'equinox.amolp'.
type Amolp struct {
	Amolpdatefrom    pq.NullTime     `json:"amolpdatefrom"`    // amolpdatefrom
	Amolpdateto      pq.NullTime     `json:"amolpdateto"`      // amolpdateto
	Amolppercent     sql.NullFloat64 `json:"amolppercent"`     // amolppercent
	Amolppercentearn sql.NullFloat64 `json:"amolppercentearn"` // amolppercentearn
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllAmolp(db XODB, callback func(x Amolp) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`amolpdatefrom, amolpdateto, amolppercent, amolppercentearn, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.amolp `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		a := Amolp{}

		// scan
		err = q.Scan(&a.Amolpdatefrom, &a.Amolpdateto, &a.Amolppercent, &a.Amolppercentearn, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(a) {
			return nil
		}
	}

	return nil
}

// AmolpByEquinoxLrn retrieves a row from 'equinox.amolp' as a Amolp.
//
// Generated from index 'amolp_pkey'.
func AmolpByEquinoxLrn(db XODB, equinoxLrn int64) (*Amolp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`amolpdatefrom, amolpdateto, amolppercent, amolppercentearn, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.amolp ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Amolp{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Amolpdatefrom, &a.Amolpdateto, &a.Amolppercent, &a.Amolppercentearn, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
