// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Chpaudit represents a row from 'equinox.chpaudit'.
type Chpaudit struct {
	Chpauddt     pq.NullTime    `json:"chpauddt"`     // chpauddt
	Chpaudtm     pq.NullTime    `json:"chpaudtm"`     // chpaudtm
	Chpauduser   sql.NullString `json:"chpauduser"`   // chpauduser
	Chpaudaction sql.NullInt64  `json:"chpaudaction"` // chpaudaction
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec
}

// ChpauditByEquinoxLrn retrieves a row from 'equinox.chpaudit' as a Chpaudit.
//
// Generated from index 'chpaudit_pkey'.
func ChpauditByEquinoxLrn(db XODB, equinoxLrn int64) (*Chpaudit, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`chpauddt, chpaudtm, chpauduser, chpaudaction, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.chpaudit ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Chpaudit{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Chpauddt, &c.Chpaudtm, &c.Chpauduser, &c.Chpaudaction, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
