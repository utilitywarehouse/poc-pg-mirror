// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cscomm represents a row from 'equinox.cscomm'.
type Cscomm struct {
	Cscomment     sql.NullInt64  `json:"cscomment"`     // cscomment
	Cscommdate    pq.NullTime    `json:"cscommdate"`    // cscommdate
	Cscommentered sql.NullString `json:"cscommentered"` // cscommentered
	EquinoxPrn    sql.NullInt64  `json:"equinox_prn"`   // equinox_prn
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec
}

// CscommByEquinoxLrn retrieves a row from 'equinox.cscomm' as a Cscomm.
//
// Generated from index 'cscomm_pkey'.
func CscommByEquinoxLrn(db XODB, equinoxLrn int64) (*Cscomm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cscomment, cscommdate, cscommentered, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cscomm ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cscomm{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cscomment, &c.Cscommdate, &c.Cscommentered, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
