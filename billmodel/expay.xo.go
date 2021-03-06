// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Expay represents a row from 'equinox.expay'.
type Expay struct {
	Expaypayment    sql.NullFloat64 `json:"expaypayment"`    // expaypayment
	Expaylevel      sql.NullFloat64 `json:"expaylevel"`      // expaylevel
	Expaytype       sql.NullInt64   `json:"expaytype"`       // expaytype
	Expaydealcode   sql.NullString  `json:"expaydealcode"`   // expaydealcode
	Expaysparen     sql.NullInt64   `json:"expaysparen"`     // expaysparen
	Expaysparec     sql.NullString  `json:"expaysparec"`     // expaysparec
	Expaynumbdb     sql.NullInt64   `json:"expaynumbdb"`     // expaynumbdb
	Expaypaymentraw sql.NullFloat64 `json:"expaypaymentraw"` // expaypaymentraw
	Expaycvcrate    sql.NullFloat64 `json:"expaycvcrate"`    // expaycvcrate
	Expaysparec1    sql.NullString  `json:"expaysparec1"`    // expaysparec1
	Expaysparen1    sql.NullFloat64 `json:"expaysparen1"`    // expaysparen1
	Expayspared1    pq.NullTime     `json:"expayspared1"`    // expayspared1
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllExpay(db XODB, callback func(x Expay) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`expaypayment, expaylevel, expaytype, expaydealcode, expaysparen, expaysparec, expaynumbdb, expaypaymentraw, expaycvcrate, expaysparec1, expaysparen1, expayspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.expay `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Expay{}

		// scan
		err = q.Scan(&e.Expaypayment, &e.Expaylevel, &e.Expaytype, &e.Expaydealcode, &e.Expaysparen, &e.Expaysparec, &e.Expaynumbdb, &e.Expaypaymentraw, &e.Expaycvcrate, &e.Expaysparec1, &e.Expaysparen1, &e.Expayspared1, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// ExpayByEquinoxLrn retrieves a row from 'equinox.expay' as a Expay.
//
// Generated from index 'expay_pkey'.
func ExpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Expay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`expaypayment, expaylevel, expaytype, expaydealcode, expaysparen, expaysparec, expaynumbdb, expaypaymentraw, expaycvcrate, expaysparec1, expaysparen1, expayspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.expay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Expay{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Expaypayment, &e.Expaylevel, &e.Expaytype, &e.Expaydealcode, &e.Expaysparen, &e.Expaysparec, &e.Expaynumbdb, &e.Expaypaymentraw, &e.Expaycvcrate, &e.Expaysparec1, &e.Expaysparen1, &e.Expayspared1, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
