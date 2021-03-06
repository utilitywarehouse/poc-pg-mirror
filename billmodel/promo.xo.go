// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Promo represents a row from 'equinox.promo'.
type Promo struct {
	PromoCode  sql.NullString `json:"promo_code"`  // promo_code
	EquinoxLrn int64          `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64  `json:"equinox_sec"` // equinox_sec
}

func AllPromo(db XODB, callback func(x Promo) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`promo_code, equinox_lrn, equinox_sec ` +
		`FROM equinox.promo `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Promo{}

		// scan
		err = q.Scan(&p.PromoCode, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PromoByEquinoxLrn retrieves a row from 'equinox.promo' as a Promo.
//
// Generated from index 'promo_pkey'.
func PromoByEquinoxLrn(db XODB, equinoxLrn int64) (*Promo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`promo_code, equinox_lrn, equinox_sec ` +
		`FROM equinox.promo ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Promo{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.PromoCode, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
