// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gappmet represents a row from 'equinox.gappmet'.
type Gappmet struct {
	Gammsn          sql.NullString  `json:"gammsn"`          // gammsn
	Gaminstalldate  pq.NullTime     `json:"gaminstalldate"`  // gaminstalldate
	Gamremovedate   pq.NullTime     `json:"gamremovedate"`   // gamremovedate
	Gamcf           sql.NullFloat64 `json:"gamcf"`           // gamcf
	Gamiorm         sql.NullString  `json:"gamiorm"`         // gamiorm
	Gamrf           sql.NullFloat64 `json:"gamrf"`           // gamrf
	Gamreadingunits sql.NullInt64   `json:"gamreadingunits"` // gamreadingunits
	Gamdials        sql.NullInt64   `json:"gamdials"`        // gamdials
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllGappmet(db XODB, callback func(x Gappmet) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gammsn, gaminstalldate, gamremovedate, gamcf, gamiorm, gamrf, gamreadingunits, gamdials, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gappmet `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		g := Gappmet{}

		// scan
		err = q.Scan(&g.Gammsn, &g.Gaminstalldate, &g.Gamremovedate, &g.Gamcf, &g.Gamiorm, &g.Gamrf, &g.Gamreadingunits, &g.Gamdials, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(g) {
			return nil
		}
	}

	return nil
}

// GappmetByEquinoxLrn retrieves a row from 'equinox.gappmet' as a Gappmet.
//
// Generated from index 'gappmet_pkey'.
func GappmetByEquinoxLrn(db XODB, equinoxLrn int64) (*Gappmet, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gammsn, gaminstalldate, gamremovedate, gamcf, gamiorm, gamrf, gamreadingunits, gamdials, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gappmet ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gappmet{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gammsn, &g.Gaminstalldate, &g.Gamremovedate, &g.Gamcf, &g.Gamiorm, &g.Gamrf, &g.Gamreadingunits, &g.Gamdials, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
