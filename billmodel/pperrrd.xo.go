// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Pperrrd represents a row from 'equinox.pperrrd'.
type Pperrrd struct {
	Pperrrdreading   sql.NullInt64  `json:"pperrrdreading"`   // pperrrdreading
	Pperrrddate      pq.NullTime    `json:"pperrrddate"`      // pperrrddate
	Pperrrdtariff    sql.NullInt64  `json:"pperrrdtariff"`    // pperrrdtariff
	Pperrrdtrfdt     pq.NullTime    `json:"pperrrdtrfdt"`     // pperrrdtrfdt
	Pperrrdreg       sql.NullString `json:"pperrrdreg"`       // pperrrdreg
	Pperrrdadded     pq.NullTime    `json:"pperrrdadded"`     // pperrrdadded
	Pperrrdallocated pq.NullTime    `json:"pperrrdallocated"` // pperrrdallocated
	Pperrrdallocto   sql.NullString `json:"pperrrdallocto"`   // pperrrdallocto
	Pperrrdsrc       sql.NullString `json:"pperrrdsrc"`       // pperrrdsrc
	Pperrrdorigref   sql.NullString `json:"pperrrdorigref"`   // pperrrdorigref
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

// PperrrdByEquinoxLrn retrieves a row from 'equinox.pperrrd' as a Pperrrd.
//
// Generated from index 'pperrrd_pkey'.
func PperrrdByEquinoxLrn(db XODB, equinoxLrn int64) (*Pperrrd, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pperrrdreading, pperrrddate, pperrrdtariff, pperrrdtrfdt, pperrrdreg, pperrrdadded, pperrrdallocated, pperrrdallocto, pperrrdsrc, pperrrdorigref, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pperrrd ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pperrrd{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pperrrdreading, &p.Pperrrddate, &p.Pperrrdtariff, &p.Pperrrdtrfdt, &p.Pperrrdreg, &p.Pperrrdadded, &p.Pperrrdallocated, &p.Pperrrdallocto, &p.Pperrrdsrc, &p.Pperrrdorigref, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
