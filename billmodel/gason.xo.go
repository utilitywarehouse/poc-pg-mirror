// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gason represents a row from 'equinox.gason'.
type Gason struct {
	Gonmpr         sql.NullString `json:"gonmpr"`         // gonmpr
	Gonmprlocc     sql.NullString `json:"gonmprlocc"`     // gonmprlocc
	Gonmprlocw     sql.NullString `json:"gonmprlocw"`     // gonmprlocw
	Gonmprlocnotes sql.NullInt64  `json:"gonmprlocnotes"` // gonmprlocnotes
	Gonmpraccess   sql.NullInt64  `json:"gonmpraccess"`   // gonmpraccess
	Gonc1          sql.NullString `json:"gonc1"`          // gonc1
	Gonc2          sql.NullString `json:"gonc2"`          // gonc2
	Gonc3          sql.NullString `json:"gonc3"`          // gonc3
	Gonn1          sql.NullInt64  `json:"gonn1"`          // gonn1
	Gonn2          sql.NullInt64  `json:"gonn2"`          // gonn2
	Gonn3          sql.NullInt64  `json:"gonn3"`          // gonn3
	Gond1          pq.NullTime    `json:"gond1"`          // gond1
	Gond2          pq.NullTime    `json:"gond2"`          // gond2
	Gond3          pq.NullTime    `json:"gond3"`          // gond3
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec
}

// GasonByEquinoxLrn retrieves a row from 'equinox.gason' as a Gason.
//
// Generated from index 'gason_pkey'.
func GasonByEquinoxLrn(db XODB, equinoxLrn int64) (*Gason, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gonmpr, gonmprlocc, gonmprlocw, gonmprlocnotes, gonmpraccess, gonc1, gonc2, gonc3, gonn1, gonn2, gonn3, gond1, gond2, gond3, equinox_lrn, equinox_sec ` +
		`FROM equinox.gason ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gason{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gonmpr, &g.Gonmprlocc, &g.Gonmprlocw, &g.Gonmprlocnotes, &g.Gonmpraccess, &g.Gonc1, &g.Gonc2, &g.Gonc3, &g.Gonn1, &g.Gonn2, &g.Gonn3, &g.Gond1, &g.Gond2, &g.Gond3, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
