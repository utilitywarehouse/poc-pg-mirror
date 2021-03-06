// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Prmopt represents a row from 'equinox.prmopt'.
type Prmopt struct {
	PromoCategory   sql.NullString `json:"promo_category"`   // promo_category
	PromoDescrip    sql.NullString `json:"promo_descrip"`    // promo_descrip
	PromoData       sql.NullString `json:"promo_data"`       // promo_data
	PromoOptsuspend sql.NullInt64  `json:"promo_optsuspend"` // promo_optsuspend
	EquinoxPrn      sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn      int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllPrmopt(db XODB, callback func(x Prmopt) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`promo_category, promo_descrip, promo_data, promo_optsuspend, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.prmopt `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Prmopt{}

		// scan
		err = q.Scan(&p.PromoCategory, &p.PromoDescrip, &p.PromoData, &p.PromoOptsuspend, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PrmoptByEquinoxLrn retrieves a row from 'equinox.prmopt' as a Prmopt.
//
// Generated from index 'prmopt_pkey'.
func PrmoptByEquinoxLrn(db XODB, equinoxLrn int64) (*Prmopt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`promo_category, promo_descrip, promo_data, promo_optsuspend, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.prmopt ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Prmopt{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.PromoCategory, &p.PromoDescrip, &p.PromoData, &p.PromoOptsuspend, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
