// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Bcdetail represents a row from 'equinox.bcdetail'.
type Bcdetail struct {
	Bcdcodeid     sql.NullInt64   `json:"bcdcodeid"`     // bcdcodeid
	Bcdcode       sql.NullString  `json:"bcdcode"`       // bcdcode
	Bcdtariff     sql.NullString  `json:"bcdtariff"`     // bcdtariff
	Bcdsubtariff  sql.NullString  `json:"bcdsubtariff"`  // bcdsubtariff
	Bcdrouting    sql.NullString  `json:"bcdrouting"`    // bcdrouting
	Bcdminspend   sql.NullString  `json:"bcdminspend"`   // bcdminspend
	Bcdoptioncode sql.NullString  `json:"bcdoptioncode"` // bcdoptioncode
	Bcdcount      sql.NullInt64   `json:"bcdcount"`      // bcdcount
	Bcdgross      sql.NullFloat64 `json:"bcdgross"`      // bcdgross
	Bcdnett       sql.NullFloat64 `json:"bcdnett"`       // bcdnett
	Bcdsparen1    sql.NullFloat64 `json:"bcdsparen1"`    // bcdsparen1
	Bcdsparec1    sql.NullString  `json:"bcdsparec1"`    // bcdsparec1
	EquinoxPrn    sql.NullInt64   `json:"equinox_prn"`   // equinox_prn
	EquinoxLrn    int64           `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64   `json:"equinox_sec"`   // equinox_sec
}

func AllBcdetail(db XODB, callback func(x Bcdetail) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`bcdcodeid, bcdcode, bcdtariff, bcdsubtariff, bcdrouting, bcdminspend, bcdoptioncode, bcdcount, bcdgross, bcdnett, bcdsparen1, bcdsparec1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.bcdetail `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		b := Bcdetail{}

		// scan
		err = q.Scan(&b.Bcdcodeid, &b.Bcdcode, &b.Bcdtariff, &b.Bcdsubtariff, &b.Bcdrouting, &b.Bcdminspend, &b.Bcdoptioncode, &b.Bcdcount, &b.Bcdgross, &b.Bcdnett, &b.Bcdsparen1, &b.Bcdsparec1, &b.EquinoxPrn, &b.EquinoxLrn, &b.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(b) {
			return nil
		}
	}

	return nil
}

// BcdetailByEquinoxLrn retrieves a row from 'equinox.bcdetail' as a Bcdetail.
//
// Generated from index 'bcdetail_pkey'.
func BcdetailByEquinoxLrn(db XODB, equinoxLrn int64) (*Bcdetail, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`bcdcodeid, bcdcode, bcdtariff, bcdsubtariff, bcdrouting, bcdminspend, bcdoptioncode, bcdcount, bcdgross, bcdnett, bcdsparen1, bcdsparec1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.bcdetail ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Bcdetail{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Bcdcodeid, &b.Bcdcode, &b.Bcdtariff, &b.Bcdsubtariff, &b.Bcdrouting, &b.Bcdminspend, &b.Bcdoptioncode, &b.Bcdcount, &b.Bcdgross, &b.Bcdnett, &b.Bcdsparen1, &b.Bcdsparec1, &b.EquinoxPrn, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
