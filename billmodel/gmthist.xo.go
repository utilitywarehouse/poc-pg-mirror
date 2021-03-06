// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gmthist represents a row from 'equinox.gmthist'.
type Gmthist struct {
	Gmthfiledate   pq.NullTime    `json:"gmthfiledate"`   // gmthfiledate
	Gmthfilename   sql.NullString `json:"gmthfilename"`   // gmthfilename
	Gmthgref       sql.NullString `json:"gmthgref"`       // gmthgref
	Gmthdateloaded pq.NullTime    `json:"gmthdateloaded"` // gmthdateloaded
	Gmthaction     sql.NullString `json:"gmthaction"`     // gmthaction
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec
}

func AllGmthist(db XODB, callback func(x Gmthist) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gmthist `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		g := Gmthist{}

		// scan
		err = q.Scan(&g.Gmthfiledate, &g.Gmthfilename, &g.Gmthgref, &g.Gmthdateloaded, &g.Gmthaction, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(g) {
			return nil
		}
	}

	return nil
}

// GmthistByEquinoxLrn retrieves a row from 'equinox.gmthist' as a Gmthist.
//
// Generated from index 'gmthist_pkey'.
func GmthistByEquinoxLrn(db XODB, equinoxLrn int64) (*Gmthist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gmthist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gmthist{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gmthfiledate, &g.Gmthfilename, &g.Gmthgref, &g.Gmthdateloaded, &g.Gmthaction, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
