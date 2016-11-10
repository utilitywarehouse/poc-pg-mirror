// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Chpact represents a row from 'equinox.chpact'.
type Chpact struct {
	Chpactid         sql.NullString `json:"chpactid"`         // chpactid
	Chpactstatus     sql.NullInt64  `json:"chpactstatus"`     // chpactstatus
	Chpactcreateddt  pq.NullTime    `json:"chpactcreateddt"`  // chpactcreateddt
	Chpactcreatedtm  pq.NullTime    `json:"chpactcreatedtm"`  // chpactcreatedtm
	Chpacttype       sql.NullString `json:"chpacttype"`       // chpacttype
	Chpactescto      sql.NullString `json:"chpactescto"`      // chpactescto
	Chpactnotes      sql.NullInt64  `json:"chpactnotes"`      // chpactnotes
	Chpactreminder   sql.NullInt64  `json:"chpactreminder"`   // chpactreminder
	Chpactremsched   sql.NullInt64  `json:"chpactremsched"`   // chpactremsched
	Chpactremtype    sql.NullInt64  `json:"chpactremtype"`    // chpactremtype
	Chpactremdt      pq.NullTime    `json:"chpactremdt"`      // chpactremdt
	Chpactremtm      pq.NullTime    `json:"chpactremtm"`      // chpactremtm
	Chpactremmsg     sql.NullInt64  `json:"chpactremmsg"`     // chpactremmsg
	Chpactcreatedby  sql.NullString `json:"chpactcreatedby"`  // chpactcreatedby
	Chpactcompletedt pq.NullTime    `json:"chpactcompletedt"` // chpactcompletedt
	Chpactcompletetm pq.NullTime    `json:"chpactcompletetm"` // chpactcompletetm
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

// ChpactByEquinoxLrn retrieves a row from 'equinox.chpact' as a Chpact.
//
// Generated from index 'chpact_pkey'.
func ChpactByEquinoxLrn(db XODB, equinoxLrn int64) (*Chpact, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`chpactid, chpactstatus, chpactcreateddt, chpactcreatedtm, chpacttype, chpactescto, chpactnotes, chpactreminder, chpactremsched, chpactremtype, chpactremdt, chpactremtm, chpactremmsg, chpactcreatedby, chpactcompletedt, chpactcompletetm, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.chpact ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Chpact{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Chpactid, &c.Chpactstatus, &c.Chpactcreateddt, &c.Chpactcreatedtm, &c.Chpacttype, &c.Chpactescto, &c.Chpactnotes, &c.Chpactreminder, &c.Chpactremsched, &c.Chpactremtype, &c.Chpactremdt, &c.Chpactremtm, &c.Chpactremmsg, &c.Chpactcreatedby, &c.Chpactcompletedt, &c.Chpactcompletetm, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
