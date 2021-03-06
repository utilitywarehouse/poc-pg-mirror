// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Hol2014o represents a row from 'equinox.hol2014o'.
type Hol2014o struct {
	Hol14opartnerid sql.NullString `json:"hol14opartnerid"` // hol14opartnerid
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

func AllHol2014o(db XODB, callback func(x Hol2014o) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`hol14opartnerid, equinox_lrn, equinox_sec ` +
		`FROM equinox.hol2014o `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		h := Hol2014o{}

		// scan
		err = q.Scan(&h.Hol14opartnerid, &h.EquinoxLrn, &h.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(h) {
			return nil
		}
	}

	return nil
}

// Hol2014oByEquinoxLrn retrieves a row from 'equinox.hol2014o' as a Hol2014o.
//
// Generated from index 'hol2014o_pkey'.
func Hol2014oByEquinoxLrn(db XODB, equinoxLrn int64) (*Hol2014o, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`hol14opartnerid, equinox_lrn, equinox_sec ` +
		`FROM equinox.hol2014o ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	h := Hol2014o{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&h.Hol14opartnerid, &h.EquinoxLrn, &h.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
