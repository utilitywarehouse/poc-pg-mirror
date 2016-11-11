// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gmprompt represents a row from 'equinox.gmprompt'.
type Gmprompt struct {
	Gmpfrequency     sql.NullInt64 `json:"gmpfrequency"`     // gmpfrequency
	Gmplastactiond   pq.NullTime   `json:"gmplastactiond"`   // gmplastactiond
	Gmplastaction    sql.NullInt64 `json:"gmplastaction"`    // gmplastaction
	Gmpfreqsetbycust sql.NullInt64 `json:"gmpfreqsetbycust"` // gmpfreqsetbycust
	EquinoxPrn       sql.NullInt64 `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64         `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64 `json:"equinox_sec"`      // equinox_sec
}

func AllGmprompt(db XODB, callback func(x Gmprompt) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gmpfrequency, gmplastactiond, gmplastaction, gmpfreqsetbycust, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gmprompt `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		g := Gmprompt{}

		// scan
		err = q.Scan(&g.Gmpfrequency, &g.Gmplastactiond, &g.Gmplastaction, &g.Gmpfreqsetbycust, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(g) {
			return nil
		}
	}

	return nil
}

// GmpromptByEquinoxLrn retrieves a row from 'equinox.gmprompt' as a Gmprompt.
//
// Generated from index 'gmprompt_pkey'.
func GmpromptByEquinoxLrn(db XODB, equinoxLrn int64) (*Gmprompt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gmpfrequency, gmplastactiond, gmplastaction, gmpfreqsetbycust, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gmprompt ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gmprompt{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gmpfrequency, &g.Gmplastactiond, &g.Gmplastaction, &g.Gmpfreqsetbycust, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
