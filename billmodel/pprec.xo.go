// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Pprec represents a row from 'equinox.pprec'.
type Pprec struct {
	Pprref           sql.NullString  `json:"pprref"`           // pprref
	Pprperiod        sql.NullString  `json:"pprperiod"`        // pprperiod
	Pprstartreadreg1 sql.NullInt64   `json:"pprstartreadreg1"` // pprstartreadreg1
	Pprstrdreg1tp    sql.NullString  `json:"pprstrdreg1tp"`    // pprstrdreg1tp
	Pprendreadreg1   sql.NullInt64   `json:"pprendreadreg1"`   // pprendreadreg1
	Ppredrdreg1tp    sql.NullString  `json:"ppredrdreg1tp"`    // ppredrdreg1tp
	Pprstartreadreg2 sql.NullInt64   `json:"pprstartreadreg2"` // pprstartreadreg2
	Pprstrdreg2tp    sql.NullString  `json:"pprstrdreg2tp"`    // pprstrdreg2tp
	Pprendreadreg2   sql.NullInt64   `json:"pprendreadreg2"`   // pprendreadreg2
	Ppredrdreg2tp    sql.NullString  `json:"ppredrdreg2tp"`    // ppredrdreg2tp
	Pprusage         sql.NullFloat64 `json:"pprusage"`         // pprusage
	Pprstdcharge     sql.NullFloat64 `json:"pprstdcharge"`     // pprstdcharge
	Pprtariff        sql.NullString  `json:"pprtariff"`        // pprtariff
	Pprpayments      sql.NullFloat64 `json:"pprpayments"`      // pprpayments
	Pprmeterstartbal sql.NullFloat64 `json:"pprmeterstartbal"` // pprmeterstartbal
	Pprmeterendbal   sql.NullFloat64 `json:"pprmeterendbal"`   // pprmeterendbal
	Pprexcode        sql.NullString  `json:"pprexcode"`        // pprexcode
	Pprcv            sql.NullFloat64 `json:"pprcv"`            // pprcv
	Pprbilltype      sql.NullString  `json:"pprbilltype"`      // pprbilltype
	Pprkwh           sql.NullInt64   `json:"pprkwh"`           // pprkwh
	Pprspared1       pq.NullTime     `json:"pprspared1"`       // pprspared1
	Pprmx            sql.NullInt64   `json:"pprmx"`            // pprmx
	Ppradjust        sql.NullFloat64 `json:"ppradjust"`        // ppradjust
	Pprdebtlanded    sql.NullFloat64 `json:"pprdebtlanded"`    // pprdebtlanded
	Pprdebtbalance   sql.NullFloat64 `json:"pprdebtbalance"`   // pprdebtbalance
	Pprimpormetric   sql.NullString  `json:"pprimpormetric"`   // pprimpormetric
	Pprinitcredit    sql.NullFloat64 `json:"pprinitcredit"`    // pprinitcredit
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllPprec(db XODB, callback func(x Pprec) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`pprref, pprperiod, pprstartreadreg1, pprstrdreg1tp, pprendreadreg1, ppredrdreg1tp, pprstartreadreg2, pprstrdreg2tp, pprendreadreg2, ppredrdreg2tp, pprusage, pprstdcharge, pprtariff, pprpayments, pprmeterstartbal, pprmeterendbal, pprexcode, pprcv, pprbilltype, pprkwh, pprspared1, pprmx, ppradjust, pprdebtlanded, pprdebtbalance, pprimpormetric, pprinitcredit, equinox_lrn, equinox_sec ` +
		`FROM equinox.pprec `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Pprec{}

		// scan
		err = q.Scan(&p.Pprref, &p.Pprperiod, &p.Pprstartreadreg1, &p.Pprstrdreg1tp, &p.Pprendreadreg1, &p.Ppredrdreg1tp, &p.Pprstartreadreg2, &p.Pprstrdreg2tp, &p.Pprendreadreg2, &p.Ppredrdreg2tp, &p.Pprusage, &p.Pprstdcharge, &p.Pprtariff, &p.Pprpayments, &p.Pprmeterstartbal, &p.Pprmeterendbal, &p.Pprexcode, &p.Pprcv, &p.Pprbilltype, &p.Pprkwh, &p.Pprspared1, &p.Pprmx, &p.Ppradjust, &p.Pprdebtlanded, &p.Pprdebtbalance, &p.Pprimpormetric, &p.Pprinitcredit, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PprecByEquinoxLrn retrieves a row from 'equinox.pprec' as a Pprec.
//
// Generated from index 'pprec_pkey'.
func PprecByEquinoxLrn(db XODB, equinoxLrn int64) (*Pprec, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pprref, pprperiod, pprstartreadreg1, pprstrdreg1tp, pprendreadreg1, ppredrdreg1tp, pprstartreadreg2, pprstrdreg2tp, pprendreadreg2, ppredrdreg2tp, pprusage, pprstdcharge, pprtariff, pprpayments, pprmeterstartbal, pprmeterendbal, pprexcode, pprcv, pprbilltype, pprkwh, pprspared1, pprmx, ppradjust, pprdebtlanded, pprdebtbalance, pprimpormetric, pprinitcredit, equinox_lrn, equinox_sec ` +
		`FROM equinox.pprec ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pprec{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pprref, &p.Pprperiod, &p.Pprstartreadreg1, &p.Pprstrdreg1tp, &p.Pprendreadreg1, &p.Ppredrdreg1tp, &p.Pprstartreadreg2, &p.Pprstrdreg2tp, &p.Pprendreadreg2, &p.Ppredrdreg2tp, &p.Pprusage, &p.Pprstdcharge, &p.Pprtariff, &p.Pprpayments, &p.Pprmeterstartbal, &p.Pprmeterendbal, &p.Pprexcode, &p.Pprcv, &p.Pprbilltype, &p.Pprkwh, &p.Pprspared1, &p.Pprmx, &p.Ppradjust, &p.Pprdebtlanded, &p.Pprdebtbalance, &p.Pprimpormetric, &p.Pprinitcredit, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
