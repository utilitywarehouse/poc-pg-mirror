// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Custopt represents a row from 'equinox.custopts'.
type Custopt struct {
	Custrmoptid      sql.NullString  `json:"custrmoptid"`      // custrmoptid
	Custrmsetup      sql.NullFloat64 `json:"custrmsetup"`      // custrmsetup
	Custrmmonthly    sql.NullFloat64 `json:"custrmmonthly"`    // custrmmonthly
	Custrmfreemins   sql.NullInt64   `json:"custrmfreemins"`   // custrmfreemins
	Custrmstartdate  pq.NullTime     `json:"custrmstartdate"`  // custrmstartdate
	Custrmenddate    pq.NullTime     `json:"custrmenddate"`    // custrmenddate
	Custrmpaiduntil  pq.NullTime     `json:"custrmpaiduntil"`  // custrmpaiduntil
	Custrmpayuntil   pq.NullTime     `json:"custrmpayuntil"`   // custrmpayuntil
	Custrmprorata    sql.NullInt64   `json:"custrmprorata"`    // custrmprorata
	Custrmmonthsadv  sql.NullInt64   `json:"custrmmonthsadv"`  // custrmmonthsadv
	Custrmoneoffbill sql.NullFloat64 `json:"custrmoneoffbill"` // custrmoneoffbill
	Custrmspecdeal   sql.NullInt64   `json:"custrmspecdeal"`   // custrmspecdeal
	Custrmpduntilbr  pq.NullTime     `json:"custrmpduntilbr"`  // custrmpduntilbr
	Custrmsparec1    sql.NullString  `json:"custrmsparec1"`    // custrmsparec1
	Custrmspared1    pq.NullTime     `json:"custrmspared1"`    // custrmspared1
	Custrmdiscband   sql.NullString  `json:"custrmdiscband"`   // custrmdiscband
	Custrmdiscpercen sql.NullInt64   `json:"custrmdiscpercen"` // custrmdiscpercen
	Custrmwlrorder   sql.NullString  `json:"custrmwlrorder"`   // custrmwlrorder
	Custrmquantity   sql.NullInt64   `json:"custrmquantity"`   // custrmquantity
	Custrmsparen1    sql.NullFloat64 `json:"custrmsparen1"`    // custrmsparen1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCustopt(db XODB, callback func(x Custopt) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`custrmoptid, custrmsetup, custrmmonthly, custrmfreemins, custrmstartdate, custrmenddate, custrmpaiduntil, custrmpayuntil, custrmprorata, custrmmonthsadv, custrmoneoffbill, custrmspecdeal, custrmpduntilbr, custrmsparec1, custrmspared1, custrmdiscband, custrmdiscpercen, custrmwlrorder, custrmquantity, custrmsparen1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.custopts `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Custopt{}

		// scan
		err = q.Scan(&c.Custrmoptid, &c.Custrmsetup, &c.Custrmmonthly, &c.Custrmfreemins, &c.Custrmstartdate, &c.Custrmenddate, &c.Custrmpaiduntil, &c.Custrmpayuntil, &c.Custrmprorata, &c.Custrmmonthsadv, &c.Custrmoneoffbill, &c.Custrmspecdeal, &c.Custrmpduntilbr, &c.Custrmsparec1, &c.Custrmspared1, &c.Custrmdiscband, &c.Custrmdiscpercen, &c.Custrmwlrorder, &c.Custrmquantity, &c.Custrmsparen1, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CustoptByEquinoxLrn retrieves a row from 'equinox.custopts' as a Custopt.
//
// Generated from index 'custopts_pkey'.
func CustoptByEquinoxLrn(db XODB, equinoxLrn int64) (*Custopt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`custrmoptid, custrmsetup, custrmmonthly, custrmfreemins, custrmstartdate, custrmenddate, custrmpaiduntil, custrmpayuntil, custrmprorata, custrmmonthsadv, custrmoneoffbill, custrmspecdeal, custrmpduntilbr, custrmsparec1, custrmspared1, custrmdiscband, custrmdiscpercen, custrmwlrorder, custrmquantity, custrmsparen1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.custopts ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Custopt{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Custrmoptid, &c.Custrmsetup, &c.Custrmmonthly, &c.Custrmfreemins, &c.Custrmstartdate, &c.Custrmenddate, &c.Custrmpaiduntil, &c.Custrmpayuntil, &c.Custrmprorata, &c.Custrmmonthsadv, &c.Custrmoneoffbill, &c.Custrmspecdeal, &c.Custrmpduntilbr, &c.Custrmsparec1, &c.Custrmspared1, &c.Custrmdiscband, &c.Custrmdiscpercen, &c.Custrmwlrorder, &c.Custrmquantity, &c.Custrmsparen1, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
