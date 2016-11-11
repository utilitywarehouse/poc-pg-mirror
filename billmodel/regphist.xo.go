// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Regphist represents a row from 'equinox.regphist'.
type Regphist struct {
	RphPromonamex   sql.NullString `json:"rph_promonamex"`   // rph_promonamex
	RphQualifyingx  sql.NullInt64  `json:"rph_qualifyingx"`  // rph_qualifyingx
	RphDatex        pq.NullTime    `json:"rph_datex"`        // rph_datex
	RphNotesx       sql.NullInt64  `json:"rph_notesx"`       // rph_notesx
	RphPeriodcountx sql.NullInt64  `json:"rph_periodcountx"` // rph_periodcountx
	RphPeriodstartx pq.NullTime    `json:"rph_periodstartx"` // rph_periodstartx
	RphPeriodendx   pq.NullTime    `json:"rph_periodendx"`   // rph_periodendx
	RphStatusx      sql.NullInt64  `json:"rph_statusx"`      // rph_statusx
	EquinoxPrn      sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn      int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllRegphist(db XODB, callback func(x Regphist) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`rph_promonamex, rph_qualifyingx, rph_datex, rph_notesx, rph_periodcountx, rph_periodstartx, rph_periodendx, rph_statusx, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.regphist `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		r := Regphist{}

		// scan
		err = q.Scan(&r.RphPromonamex, &r.RphQualifyingx, &r.RphDatex, &r.RphNotesx, &r.RphPeriodcountx, &r.RphPeriodstartx, &r.RphPeriodendx, &r.RphStatusx, &r.EquinoxPrn, &r.EquinoxLrn, &r.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(r) {
			return nil
		}
	}

	return nil
}

// RegphistByEquinoxLrn retrieves a row from 'equinox.regphist' as a Regphist.
//
// Generated from index 'regphist_pkey'.
func RegphistByEquinoxLrn(db XODB, equinoxLrn int64) (*Regphist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`rph_promonamex, rph_qualifyingx, rph_datex, rph_notesx, rph_periodcountx, rph_periodstartx, rph_periodendx, rph_statusx, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.regphist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	r := Regphist{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&r.RphPromonamex, &r.RphQualifyingx, &r.RphDatex, &r.RphNotesx, &r.RphPeriodcountx, &r.RphPeriodstartx, &r.RphPeriodendx, &r.RphStatusx, &r.EquinoxPrn, &r.EquinoxLrn, &r.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
