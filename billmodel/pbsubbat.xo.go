// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Pbsubbat represents a row from 'equinox.pbsubbat'.
type Pbsubbat struct {
	Pbsubbatchref    sql.NullString  `json:"pbsubbatchref"`    // pbsubbatchref
	Pbsubbatchdate   pq.NullTime     `json:"pbsubbatchdate"`   // pbsubbatchdate
	Pbsubbatchamount sql.NullFloat64 `json:"pbsubbatchamount"` // pbsubbatchamount
	Pbsubbatchtype   sql.NullString  `json:"pbsubbatchtype"`   // pbsubbatchtype
	Pbsubbatchbalanc sql.NullFloat64 `json:"pbsubbatchbalanc"` // pbsubbatchbalanc
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// PbsubbatByEquinoxLrn retrieves a row from 'equinox.pbsubbat' as a Pbsubbat.
//
// Generated from index 'pbsubbat_pkey'.
func PbsubbatByEquinoxLrn(db XODB, equinoxLrn int64) (*Pbsubbat, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pbsubbatchref, pbsubbatchdate, pbsubbatchamount, pbsubbatchtype, pbsubbatchbalanc, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pbsubbat ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pbsubbat{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pbsubbatchref, &p.Pbsubbatchdate, &p.Pbsubbatchamount, &p.Pbsubbatchtype, &p.Pbsubbatchbalanc, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
