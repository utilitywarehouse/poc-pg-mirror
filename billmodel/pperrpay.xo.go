// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Pperrpay represents a row from 'equinox.pperrpay'.
type Pperrpay struct {
	Pperrpaydt       pq.NullTime     `json:"pperrpaydt"`       // pperrpaydt
	Pperrpaytm       pq.NullTime     `json:"pperrpaytm"`       // pperrpaytm
	Pperrpayamt      sql.NullFloat64 `json:"pperrpayamt"`      // pperrpayamt
	Pperrpayoutlet   sql.NullString  `json:"pperrpayoutlet"`   // pperrpayoutlet
	Pperrpaympid     sql.NullString  `json:"pperrpaympid"`     // pperrpaympid
	Pperrpaytca      sql.NullInt64   `json:"pperrpaytca"`      // pperrpaytca
	Pperrpayinfdt    pq.NullTime     `json:"pperrpayinfdt"`    // pperrpayinfdt
	Pperrpaybal      sql.NullFloat64 `json:"pperrpaybal"`      // pperrpaybal
	Pperrpaystdchrg  sql.NullFloat64 `json:"pperrpaystdchrg"`  // pperrpaystdchrg
	Pperrpaydyrt     sql.NullFloat64 `json:"pperrpaydyrt"`     // pperrpaydyrt
	Pperrpayngtrt    sql.NullFloat64 `json:"pperrpayngtrt"`    // pperrpayngtrt
	Pperrpaytransno  sql.NullInt64   `json:"pperrpaytransno"`  // pperrpaytransno
	Pperrpayadded    pq.NullTime     `json:"pperrpayadded"`    // pperrpayadded
	Pperrpayalloc    pq.NullTime     `json:"pperrpayalloc"`    // pperrpayalloc
	Pperrpayallocto  sql.NullString  `json:"pperrpayallocto"`  // pperrpayallocto
	Pperrpayrefunded pq.NullTime     `json:"pperrpayrefunded"` // pperrpayrefunded
	Pperrpayrefto    sql.NullString  `json:"pperrpayrefto"`    // pperrpayrefto
	Pperrpayorigref  sql.NullString  `json:"pperrpayorigref"`  // pperrpayorigref
	Pperrpaysrc      sql.NullString  `json:"pperrpaysrc"`      // pperrpaysrc
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllPperrpay(db XODB, callback func(x Pperrpay) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`pperrpaydt, pperrpaytm, pperrpayamt, pperrpayoutlet, pperrpaympid, pperrpaytca, pperrpayinfdt, pperrpaybal, pperrpaystdchrg, pperrpaydyrt, pperrpayngtrt, pperrpaytransno, pperrpayadded, pperrpayalloc, pperrpayallocto, pperrpayrefunded, pperrpayrefto, pperrpayorigref, pperrpaysrc, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pperrpay `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Pperrpay{}

		// scan
		err = q.Scan(&p.Pperrpaydt, &p.Pperrpaytm, &p.Pperrpayamt, &p.Pperrpayoutlet, &p.Pperrpaympid, &p.Pperrpaytca, &p.Pperrpayinfdt, &p.Pperrpaybal, &p.Pperrpaystdchrg, &p.Pperrpaydyrt, &p.Pperrpayngtrt, &p.Pperrpaytransno, &p.Pperrpayadded, &p.Pperrpayalloc, &p.Pperrpayallocto, &p.Pperrpayrefunded, &p.Pperrpayrefto, &p.Pperrpayorigref, &p.Pperrpaysrc, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PperrpayByEquinoxLrn retrieves a row from 'equinox.pperrpay' as a Pperrpay.
//
// Generated from index 'pperrpay_pkey'.
func PperrpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Pperrpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pperrpaydt, pperrpaytm, pperrpayamt, pperrpayoutlet, pperrpaympid, pperrpaytca, pperrpayinfdt, pperrpaybal, pperrpaystdchrg, pperrpaydyrt, pperrpayngtrt, pperrpaytransno, pperrpayadded, pperrpayalloc, pperrpayallocto, pperrpayrefunded, pperrpayrefto, pperrpayorigref, pperrpaysrc, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pperrpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pperrpay{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pperrpaydt, &p.Pperrpaytm, &p.Pperrpayamt, &p.Pperrpayoutlet, &p.Pperrpaympid, &p.Pperrpaytca, &p.Pperrpayinfdt, &p.Pperrpaybal, &p.Pperrpaystdchrg, &p.Pperrpaydyrt, &p.Pperrpayngtrt, &p.Pperrpaytransno, &p.Pperrpayadded, &p.Pperrpayalloc, &p.Pperrpayallocto, &p.Pperrpayrefunded, &p.Pperrpayrefto, &p.Pperrpayorigref, &p.Pperrpaysrc, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
