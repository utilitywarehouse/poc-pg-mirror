// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Ps71rec represents a row from 'equinox.ps71recs'.
type Ps71rec struct {
	Ps71rejfilename sql.NullString `json:"ps71rejfilename"` // ps71rejfilename
	EquinoxPrn      sql.NullInt64  `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

func AllPs71rec(db XODB, callback func(x Ps71rec) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`ps71rejfilename, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ps71recs `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Ps71rec{}

		// scan
		err = q.Scan(&p.Ps71rejfilename, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// Ps71recByEquinoxLrn retrieves a row from 'equinox.ps71recs' as a Ps71rec.
//
// Generated from index 'ps71recs_pkey'.
func Ps71recByEquinoxLrn(db XODB, equinoxLrn int64) (*Ps71rec, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ps71rejfilename, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ps71recs ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Ps71rec{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Ps71rejfilename, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
