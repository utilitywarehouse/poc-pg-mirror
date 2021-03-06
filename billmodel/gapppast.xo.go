// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gapppast represents a row from 'equinox.gapppast'.
type Gapppast struct {
	Gappps         pq.NullTime   `json:"gappps"`         // gappps
	Gapppe         pq.NullTime   `json:"gapppe"`         // gapppe
	Gappaqapplied  pq.NullTime   `json:"gappaqapplied"`  // gappaqapplied
	Gappp1aq       sql.NullInt64 `json:"gappp1aq"`       // gappp1aq
	Gappp2aq       sql.NullInt64 `json:"gappp2aq"`       // gappp2aq
	Gappp3aq       sql.NullInt64 `json:"gappp3aq"`       // gappp3aq
	Gapppbilledkwh sql.NullInt64 `json:"gapppbilledkwh"` // gapppbilledkwh
	EquinoxPrn     sql.NullInt64 `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64         `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64 `json:"equinox_sec"`    // equinox_sec
}

func AllGapppast(db XODB, callback func(x Gapppast) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gappps, gapppe, gappaqapplied, gappp1aq, gappp2aq, gappp3aq, gapppbilledkwh, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gapppast `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		g := Gapppast{}

		// scan
		err = q.Scan(&g.Gappps, &g.Gapppe, &g.Gappaqapplied, &g.Gappp1aq, &g.Gappp2aq, &g.Gappp3aq, &g.Gapppbilledkwh, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(g) {
			return nil
		}
	}

	return nil
}

// GapppastByEquinoxLrn retrieves a row from 'equinox.gapppast' as a Gapppast.
//
// Generated from index 'gapppast_pkey'.
func GapppastByEquinoxLrn(db XODB, equinoxLrn int64) (*Gapppast, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gappps, gapppe, gappaqapplied, gappp1aq, gappp2aq, gappp3aq, gapppbilledkwh, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gapppast ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gapppast{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gappps, &g.Gapppe, &g.Gappaqapplied, &g.Gappp1aq, &g.Gappp2aq, &g.Gappp3aq, &g.Gapppbilledkwh, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
