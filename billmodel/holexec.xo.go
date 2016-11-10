// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Holexec represents a row from 'equinox.holexec'.
type Holexec struct {
	Holexecmonth     sql.NullString `json:"holexecmonth"`     // holexecmonth
	Holexecpartnerid sql.NullString `json:"holexecpartnerid"` // holexecpartnerid
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

// HolexecByEquinoxLrn retrieves a row from 'equinox.holexec' as a Holexec.
//
// Generated from index 'holexec_pkey'.
func HolexecByEquinoxLrn(db XODB, equinoxLrn int64) (*Holexec, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`holexecmonth, holexecpartnerid, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.holexec ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	h := Holexec{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&h.Holexecmonth, &h.Holexecpartnerid, &h.EquinoxPrn, &h.EquinoxLrn, &h.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
