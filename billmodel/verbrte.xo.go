// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Verbrte represents a row from 'equinox.verbrte'.
type Verbrte struct {
	Vrtableno    sql.NullInt64   `json:"vrtableno"`    // vrtableno
	Vrtabletype  sql.NullString  `json:"vrtabletype"`  // vrtabletype
	Vrterm       sql.NullInt64   `json:"vrterm"`       // vrterm
	Vrcommission sql.NullFloat64 `json:"vrcommission"` // vrcommission
	Vrfromdate   pq.NullTime     `json:"vrfromdate"`   // vrfromdate
	Vrtodate     pq.NullTime     `json:"vrtodate"`     // vrtodate
	Vrmultiple   sql.NullInt64   `json:"vrmultiple"`   // vrmultiple
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec
}

func AllVerbrte(db XODB, callback func(x Verbrte) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`vrtableno, vrtabletype, vrterm, vrcommission, vrfromdate, vrtodate, vrmultiple, equinox_lrn, equinox_sec ` +
		`FROM equinox.verbrte `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		v := Verbrte{}

		// scan
		err = q.Scan(&v.Vrtableno, &v.Vrtabletype, &v.Vrterm, &v.Vrcommission, &v.Vrfromdate, &v.Vrtodate, &v.Vrmultiple, &v.EquinoxLrn, &v.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(v) {
			return nil
		}
	}

	return nil
}

// VerbrteByEquinoxLrn retrieves a row from 'equinox.verbrte' as a Verbrte.
//
// Generated from index 'verbrte_pkey'.
func VerbrteByEquinoxLrn(db XODB, equinoxLrn int64) (*Verbrte, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`vrtableno, vrtabletype, vrterm, vrcommission, vrfromdate, vrtodate, vrmultiple, equinox_lrn, equinox_sec ` +
		`FROM equinox.verbrte ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	v := Verbrte{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&v.Vrtableno, &v.Vrtabletype, &v.Vrterm, &v.Vrcommission, &v.Vrfromdate, &v.Vrtodate, &v.Vrmultiple, &v.EquinoxLrn, &v.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
