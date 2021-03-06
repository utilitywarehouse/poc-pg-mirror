// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Portcomm represents a row from 'equinox.portcomm'.
type Portcomm struct {
	Portcomment  sql.NullInt64  `json:"portcomment"`  // portcomment
	Portcommdate pq.NullTime    `json:"portcommdate"` // portcommdate
	Portcommby   sql.NullString `json:"portcommby"`   // portcommby
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec
}

func AllPortcomm(db XODB, callback func(x Portcomm) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`portcomment, portcommdate, portcommby, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.portcomm `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Portcomm{}

		// scan
		err = q.Scan(&p.Portcomment, &p.Portcommdate, &p.Portcommby, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PortcommByEquinoxLrn retrieves a row from 'equinox.portcomm' as a Portcomm.
//
// Generated from index 'portcomm_pkey'.
func PortcommByEquinoxLrn(db XODB, equinoxLrn int64) (*Portcomm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`portcomment, portcommdate, portcommby, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.portcomm ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Portcomm{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Portcomment, &p.Portcommdate, &p.Portcommby, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
