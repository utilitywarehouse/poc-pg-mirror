// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Greendl represents a row from 'equinox.greendl'.
type Greendl struct {
	Gdplanid         sql.NullString  `json:"gdplanid"`         // gdplanid
	Gdmpan           sql.NullString  `json:"gdmpan"`           // gdmpan
	Gdstartdate      pq.NullTime     `json:"gdstartdate"`      // gdstartdate
	Gdenddate        pq.NullTime     `json:"gdenddate"`        // gdenddate
	Gdannualgassave  sql.NullFloat64 `json:"gdannualgassave"`  // gdannualgassave
	Gdannualelecsave sql.NullFloat64 `json:"gdannualelecsave"` // gdannualelecsave
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// GreendlByEquinoxLrn retrieves a row from 'equinox.greendl' as a Greendl.
//
// Generated from index 'greendl_pkey'.
func GreendlByEquinoxLrn(db XODB, equinoxLrn int64) (*Greendl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gdplanid, gdmpan, gdstartdate, gdenddate, gdannualgassave, gdannualelecsave, equinox_lrn, equinox_sec ` +
		`FROM equinox.greendl ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Greendl{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gdplanid, &g.Gdmpan, &g.Gdstartdate, &g.Gdenddate, &g.Gdannualgassave, &g.Gdannualelecsave, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
