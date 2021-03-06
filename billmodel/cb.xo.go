// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Cb represents a row from 'equinox.cb'.
type Cb struct {
	Cbuniquesys      sql.NullInt64  `json:"cbuniquesys"`      // cbuniquesys
	Cblocation       sql.NullString `json:"cblocation"`       // cblocation
	Cbdescription    sql.NullString `json:"cbdescription"`    // cbdescription
	Cbretailband     sql.NullString `json:"cbretailband"`     // cbretailband
	Cbwholebandbt    sql.NullString `json:"cbwholebandbt"`    // cbwholebandbt
	Cbwholebandcw    sql.NullString `json:"cbwholebandcw"`    // cbwholebandcw
	Cbwholebandcwin  sql.NullString `json:"cbwholebandcwin"`  // cbwholebandcwin
	Cbwholebandenerg sql.NullString `json:"cbwholebandenerg"` // cbwholebandenerg
	Cbwholebandgamma sql.NullString `json:"cbwholebandgamma"` // cbwholebandgamma
	Cbwholebandmpc   sql.NullString `json:"cbwholebandmpc"`   // cbwholebandmpc
	Cbwholebandmpcp  sql.NullString `json:"cbwholebandmpcp"`  // cbwholebandmpcp
	Cbwholebandntl   sql.NullString `json:"cbwholebandntl"`   // cbwholebandntl
	Cbwholebando2    sql.NullString `json:"cbwholebando2"`    // cbwholebando2
	Cbwholebandopal  sql.NullString `json:"cbwholebandopal"`  // cbwholebandopal
	Cbtype           sql.NullInt64  `json:"cbtype"`           // cbtype
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllCb(db XODB, callback func(x Cb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_lrn, equinox_sec ` +
		`FROM equinox.cb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cb{}

		// scan
		err = q.Scan(&c.Cbuniquesys, &c.Cblocation, &c.Cbdescription, &c.Cbretailband, &c.Cbwholebandbt, &c.Cbwholebandcw, &c.Cbwholebandcwin, &c.Cbwholebandenerg, &c.Cbwholebandgamma, &c.Cbwholebandmpc, &c.Cbwholebandmpcp, &c.Cbwholebandntl, &c.Cbwholebando2, &c.Cbwholebandopal, &c.Cbtype, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CbByEquinoxLrn retrieves a row from 'equinox.cb' as a Cb.
//
// Generated from index 'cb_pkey'.
func CbByEquinoxLrn(db XODB, equinoxLrn int64) (*Cb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_lrn, equinox_sec ` +
		`FROM equinox.cb ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cb{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cbuniquesys, &c.Cblocation, &c.Cbdescription, &c.Cbretailband, &c.Cbwholebandbt, &c.Cbwholebandcw, &c.Cbwholebandcwin, &c.Cbwholebandenerg, &c.Cbwholebandgamma, &c.Cbwholebandmpc, &c.Cbwholebandmpcp, &c.Cbwholebandntl, &c.Cbwholebando2, &c.Cbwholebandopal, &c.Cbtype, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
