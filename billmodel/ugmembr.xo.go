// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Ugmembr represents a row from 'equinox.ugmembr'.
type Ugmembr struct {
	Ugmuserid  sql.NullString `json:"ugmuserid"`   // ugmuserid
	Ugmgroupid sql.NullString `json:"ugmgroupid"`  // ugmgroupid
	EquinoxPrn sql.NullInt64  `json:"equinox_prn"` // equinox_prn
	EquinoxLrn int64          `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64  `json:"equinox_sec"` // equinox_sec
}

// UgmembrByEquinoxLrn retrieves a row from 'equinox.ugmembr' as a Ugmembr.
//
// Generated from index 'ugmembr_pkey'.
func UgmembrByEquinoxLrn(db XODB, equinoxLrn int64) (*Ugmembr, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ugmuserid, ugmgroupid, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ugmembr ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	u := Ugmembr{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&u.Ugmuserid, &u.Ugmgroupid, &u.EquinoxPrn, &u.EquinoxLrn, &u.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
