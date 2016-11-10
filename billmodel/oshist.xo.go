// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Oshist represents a row from 'equinox.oshist'.
type Oshist struct {
	Oshdateout   pq.NullTime    `json:"oshdateout"`   // oshdateout
	Oshtimeout   pq.NullTime    `json:"oshtimeout"`   // oshtimeout
	Oshaccountno sql.NullString `json:"oshaccountno"` // oshaccountno
	Oshlocation  sql.NullString `json:"oshlocation"`  // oshlocation
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec
}

// OshistByEquinoxLrn retrieves a row from 'equinox.oshist' as a Oshist.
//
// Generated from index 'oshist_pkey'.
func OshistByEquinoxLrn(db XODB, equinoxLrn int64) (*Oshist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`oshdateout, oshtimeout, oshaccountno, oshlocation, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.oshist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	o := Oshist{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&o.Oshdateout, &o.Oshtimeout, &o.Oshaccountno, &o.Oshlocation, &o.EquinoxPrn, &o.EquinoxLrn, &o.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
