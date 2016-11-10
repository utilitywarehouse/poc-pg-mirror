// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Vcheader represents a row from 'equinox.vcheader'.
type Vcheader struct {
	Vchdealer      sql.NullString  `json:"vchdealer"`      // vchdealer
	Vchclawbackbal sql.NullFloat64 `json:"vchclawbackbal"` // vchclawbackbal
	EquinoxPrn     sql.NullInt64   `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec
}

// VcheaderByEquinoxLrn retrieves a row from 'equinox.vcheader' as a Vcheader.
//
// Generated from index 'vcheader_pkey'.
func VcheaderByEquinoxLrn(db XODB, equinoxLrn int64) (*Vcheader, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`vchdealer, vchclawbackbal, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.vcheader ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	v := Vcheader{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&v.Vchdealer, &v.Vchclawbackbal, &v.EquinoxPrn, &v.EquinoxLrn, &v.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
