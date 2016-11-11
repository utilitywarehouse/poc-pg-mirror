// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Smartcom represents a row from 'equinox.smartcom'.
type Smartcom struct {
	Smcid      sql.NullString `json:"smcid"`       // smcid
	Smcdate    pq.NullTime    `json:"smcdate"`     // smcdate
	Smcsentby  sql.NullString `json:"smcsentby"`   // smcsentby
	Smctype    sql.NullInt64  `json:"smctype"`     // smctype
	EquinoxPrn sql.NullInt64  `json:"equinox_prn"` // equinox_prn
	EquinoxLrn int64          `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64  `json:"equinox_sec"` // equinox_sec
}

func AllSmartcom(db XODB, callback func(x Smartcom) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`smcid, smcdate, smcsentby, smctype, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.smartcom `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		s := Smartcom{}

		// scan
		err = q.Scan(&s.Smcid, &s.Smcdate, &s.Smcsentby, &s.Smctype, &s.EquinoxPrn, &s.EquinoxLrn, &s.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(s) {
			return nil
		}
	}

	return nil
}

// SmartcomByEquinoxLrn retrieves a row from 'equinox.smartcom' as a Smartcom.
//
// Generated from index 'smartcom_pkey'.
func SmartcomByEquinoxLrn(db XODB, equinoxLrn int64) (*Smartcom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`smcid, smcdate, smcsentby, smctype, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.smartcom ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Smartcom{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Smcid, &s.Smcdate, &s.Smcsentby, &s.Smctype, &s.EquinoxPrn, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
