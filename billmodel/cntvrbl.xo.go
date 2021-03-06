// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Cntvrbl represents a row from 'equinox.cntvrbls'.
type Cntvrbl struct {
	Cntvvariablename sql.NullString `json:"cntvvariablename"` // cntvvariablename
	Cntvvariabledata sql.NullInt64  `json:"cntvvariabledata"` // cntvvariabledata
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllCntvrbl(db XODB, callback func(x Cntvrbl) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cntvvariablename, cntvvariabledata, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cntvrbls `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cntvrbl{}

		// scan
		err = q.Scan(&c.Cntvvariablename, &c.Cntvvariabledata, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CntvrblByEquinoxLrn retrieves a row from 'equinox.cntvrbls' as a Cntvrbl.
//
// Generated from index 'cntvrbls_pkey'.
func CntvrblByEquinoxLrn(db XODB, equinoxLrn int64) (*Cntvrbl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cntvvariablename, cntvvariabledata, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cntvrbls ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cntvrbl{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cntvvariablename, &c.Cntvvariabledata, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
