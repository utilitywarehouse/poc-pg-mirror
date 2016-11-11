// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Pesrate represents a row from 'equinox.pesrates'.
type Pesrate struct {
	Pesregion      sql.NullString  `json:"pesregion"`      // pesregion
	Pesmeter       sql.NullString  `json:"pesmeter"`       // pesmeter
	Pestariffmeter sql.NullString  `json:"pestariffmeter"` // pestariffmeter
	Pesppkwh       sql.NullFloat64 `json:"pesppkwh"`       // pesppkwh
	Pesmonthly     sql.NullFloat64 `json:"pesmonthly"`     // pesmonthly
	Pesdaily       sql.NullFloat64 `json:"pesdaily"`       // pesdaily
	Pesmonthorday  sql.NullString  `json:"pesmonthorday"`  // pesmonthorday
	Pesvalidfrom   pq.NullTime     `json:"pesvalidfrom"`   // pesvalidfrom
	Pesvalidto     pq.NullTime     `json:"pesvalidto"`     // pesvalidto
	Pestcr         sql.NullFloat64 `json:"pestcr"`         // pestcr
	Pestcrdual     sql.NullFloat64 `json:"pestcrdual"`     // pestcrdual
	EquinoxPrn     sql.NullInt64   `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec
}

func AllPesrate(db XODB, callback func(x Pesrate) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pesrates `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Pesrate{}

		// scan
		err = q.Scan(&p.Pesregion, &p.Pesmeter, &p.Pestariffmeter, &p.Pesppkwh, &p.Pesmonthly, &p.Pesdaily, &p.Pesmonthorday, &p.Pesvalidfrom, &p.Pesvalidto, &p.Pestcr, &p.Pestcrdual, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PesrateByEquinoxLrn retrieves a row from 'equinox.pesrates' as a Pesrate.
//
// Generated from index 'pesrates_pkey'.
func PesrateByEquinoxLrn(db XODB, equinoxLrn int64) (*Pesrate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pesrates ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pesrate{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pesregion, &p.Pesmeter, &p.Pestariffmeter, &p.Pesppkwh, &p.Pesmonthly, &p.Pesdaily, &p.Pesmonthorday, &p.Pesvalidfrom, &p.Pesvalidto, &p.Pestcr, &p.Pestcrdual, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
