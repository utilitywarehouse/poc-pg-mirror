// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Smartsvr represents a row from 'equinox.smartsvr'.
type Smartsvr struct {
	Ssvrid          sql.NullInt64  `json:"ssvrid"`          // ssvrid
	Ssvrdtadded     pq.NullTime    `json:"ssvrdtadded"`     // ssvrdtadded
	Ssvrdtprocessed pq.NullTime    `json:"ssvrdtprocessed"` // ssvrdtprocessed
	Ssvrfilename    sql.NullString `json:"ssvrfilename"`    // ssvrfilename
	Ssvrsource      sql.NullString `json:"ssvrsource"`      // ssvrsource
	Ssvraccount     sql.NullString `json:"ssvraccount"`     // ssvraccount
	Ssvrmpxn        sql.NullString `json:"ssvrmpxn"`        // ssvrmpxn
	Ssvrdata1       sql.NullString `json:"ssvrdata1"`       // ssvrdata1
	Ssvrdata2       sql.NullString `json:"ssvrdata2"`       // ssvrdata2
	Ssvrdata3       sql.NullString `json:"ssvrdata3"`       // ssvrdata3
	Ssvrdata4       sql.NullString `json:"ssvrdata4"`       // ssvrdata4
	Ssvrdata5       sql.NullString `json:"ssvrdata5"`       // ssvrdata5
	Ssvrdata6       sql.NullString `json:"ssvrdata6"`       // ssvrdata6
	Ssvrdata7       sql.NullString `json:"ssvrdata7"`       // ssvrdata7
	Ssvrdata8       sql.NullString `json:"ssvrdata8"`       // ssvrdata8
	Ssvrdata9       sql.NullString `json:"ssvrdata9"`       // ssvrdata9
	Ssvrdata10      sql.NullString `json:"ssvrdata10"`      // ssvrdata10
	Ssvrdata11      sql.NullString `json:"ssvrdata11"`      // ssvrdata11
	Ssvrdata12      sql.NullString `json:"ssvrdata12"`      // ssvrdata12
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

// SmartsvrByEquinoxLrn retrieves a row from 'equinox.smartsvr' as a Smartsvr.
//
// Generated from index 'smartsvr_pkey'.
func SmartsvrByEquinoxLrn(db XODB, equinoxLrn int64) (*Smartsvr, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ssvrid, ssvrdtadded, ssvrdtprocessed, ssvrfilename, ssvrsource, ssvraccount, ssvrmpxn, ssvrdata1, ssvrdata2, ssvrdata3, ssvrdata4, ssvrdata5, ssvrdata6, ssvrdata7, ssvrdata8, ssvrdata9, ssvrdata10, ssvrdata11, ssvrdata12, equinox_lrn, equinox_sec ` +
		`FROM equinox.smartsvr ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Smartsvr{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Ssvrid, &s.Ssvrdtadded, &s.Ssvrdtprocessed, &s.Ssvrfilename, &s.Ssvrsource, &s.Ssvraccount, &s.Ssvrmpxn, &s.Ssvrdata1, &s.Ssvrdata2, &s.Ssvrdata3, &s.Ssvrdata4, &s.Ssvrdata5, &s.Ssvrdata6, &s.Ssvrdata7, &s.Ssvrdata8, &s.Ssvrdata9, &s.Ssvrdata10, &s.Ssvrdata11, &s.Ssvrdata12, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
