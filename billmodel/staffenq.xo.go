// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Staffenq represents a row from 'equinox.staffenq'.
type Staffenq struct {
	StaffEnqcode sql.NullString `json:"staff_enqcode"` // staff_enqcode
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`   // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`   // equinox_sec
}

// StaffenqByEquinoxLrn retrieves a row from 'equinox.staffenq' as a Staffenq.
//
// Generated from index 'staffenq_pkey'.
func StaffenqByEquinoxLrn(db XODB, equinoxLrn int64) (*Staffenq, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`staff_enqcode, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.staffenq ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Staffenq{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.StaffEnqcode, &s.EquinoxPrn, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
