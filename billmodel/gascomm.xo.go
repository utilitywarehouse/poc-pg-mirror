// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gascomm represents a row from 'equinox.gascomm'.
type Gascomm struct {
	Gcommdate        pq.NullTime     `json:"gcommdate"`        // gcommdate
	Gcommcomment     sql.NullInt64   `json:"gcommcomment"`     // gcommcomment
	Gcommenteredby   sql.NullString  `json:"gcommenteredby"`   // gcommenteredby
	Gcommcommentcode sql.NullString  `json:"gcommcommentcode"` // gcommcommentcode
	Gcommspecific    sql.NullString  `json:"gcommspecific"`    // gcommspecific
	Gcommoutcome     sql.NullString  `json:"gcommoutcome"`     // gcommoutcome
	Gcommsparec1     sql.NullString  `json:"gcommsparec1"`     // gcommsparec1
	Gcommsparen1     sql.NullFloat64 `json:"gcommsparen1"`     // gcommsparen1
	Gcommspared1     pq.NullTime     `json:"gcommspared1"`     // gcommspared1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllGascomm(db XODB, callback func(x Gascomm) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gascomm `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		g := Gascomm{}

		// scan
		err = q.Scan(&g.Gcommdate, &g.Gcommcomment, &g.Gcommenteredby, &g.Gcommcommentcode, &g.Gcommspecific, &g.Gcommoutcome, &g.Gcommsparec1, &g.Gcommsparen1, &g.Gcommspared1, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(g) {
			return nil
		}
	}

	return nil
}

// GascommByEquinoxLrn retrieves a row from 'equinox.gascomm' as a Gascomm.
//
// Generated from index 'gascomm_pkey'.
func GascommByEquinoxLrn(db XODB, equinoxLrn int64) (*Gascomm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gascomm ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gascomm{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gcommdate, &g.Gcommcomment, &g.Gcommenteredby, &g.Gcommcommentcode, &g.Gcommspecific, &g.Gcommoutcome, &g.Gcommsparec1, &g.Gcommsparen1, &g.Gcommspared1, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
