// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Landreg represents a row from 'equinox.landreg'.
type Landreg struct {
	LrCustref       sql.NullString  `json:"lr_custref"`       // lr_custref
	LrSurname       sql.NullString  `json:"lr_surname"`       // lr_surname
	LrPostcode      sql.NullString  `json:"lr_postcode"`      // lr_postcode
	LrAMatchresult  sql.NullString  `json:"lr_a_matchresult"` // lr_a_matchresult
	LrTitleno       sql.NullString  `json:"lr_titleno"`       // lr_titleno
	LrNMatchresult  sql.NullString  `json:"lr_n_matchresult"` // lr_n_matchresult
	LrDateOut       pq.NullTime     `json:"lr_date_out"`      // lr_date_out
	LrDateIn        pq.NullTime     `json:"lr_date_in"`       // lr_date_in
	LrBillreference sql.NullString  `json:"lr_billreference"` // lr_billreference
	LrCommisvalue   sql.NullFloat64 `json:"lr_commisvalue"`   // lr_commisvalue
	LrQty           sql.NullInt64   `json:"lr_qty"`           // lr_qty
	LrMemo          sql.NullInt64   `json:"lr_memo"`          // lr_memo
	EquinoxLrn      int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllLandreg(db XODB, callback func(x Landreg) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`lr_custref, lr_surname, lr_postcode, lr_a_matchresult, lr_titleno, lr_n_matchresult, lr_date_out, lr_date_in, lr_billreference, lr_commisvalue, lr_qty, lr_memo, equinox_lrn, equinox_sec ` +
		`FROM equinox.landreg `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		l := Landreg{}

		// scan
		err = q.Scan(&l.LrCustref, &l.LrSurname, &l.LrPostcode, &l.LrAMatchresult, &l.LrTitleno, &l.LrNMatchresult, &l.LrDateOut, &l.LrDateIn, &l.LrBillreference, &l.LrCommisvalue, &l.LrQty, &l.LrMemo, &l.EquinoxLrn, &l.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(l) {
			return nil
		}
	}

	return nil
}

// LandregByEquinoxLrn retrieves a row from 'equinox.landreg' as a Landreg.
//
// Generated from index 'landreg_pkey'.
func LandregByEquinoxLrn(db XODB, equinoxLrn int64) (*Landreg, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`lr_custref, lr_surname, lr_postcode, lr_a_matchresult, lr_titleno, lr_n_matchresult, lr_date_out, lr_date_in, lr_billreference, lr_commisvalue, lr_qty, lr_memo, equinox_lrn, equinox_sec ` +
		`FROM equinox.landreg ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	l := Landreg{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&l.LrCustref, &l.LrSurname, &l.LrPostcode, &l.LrAMatchresult, &l.LrTitleno, &l.LrNMatchresult, &l.LrDateOut, &l.LrDateIn, &l.LrBillreference, &l.LrCommisvalue, &l.LrQty, &l.LrMemo, &l.EquinoxLrn, &l.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
