// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Avtran represents a row from 'equinox.avtrans'.
type Avtran struct {
	Avtuniquesys   sql.NullFloat64 `json:"avtuniquesys"`   // avtuniquesys
	Avtdate        pq.NullTime     `json:"avtdate"`        // avtdate
	Avtcode        sql.NullString  `json:"avtcode"`        // avtcode
	Avtdescription sql.NullString  `json:"avtdescription"` // avtdescription
	Avtawarded     sql.NullFloat64 `json:"avtawarded"`     // avtawarded
	Avtbillno      sql.NullFloat64 `json:"avtbillno"`      // avtbillno
	Avtsenttoavios pq.NullTime     `json:"avtsenttoavios"` // avtsenttoavios
	Avtwriteoff    pq.NullTime     `json:"avtwriteoff"`    // avtwriteoff
	Avtaviosno     sql.NullString  `json:"avtaviosno"`     // avtaviosno
	Avttransactno  sql.NullFloat64 `json:"avttransactno"`  // avttransactno
	Avtrejectcode  sql.NullString  `json:"avtrejectcode"`  // avtrejectcode
	EquinoxPrn     sql.NullInt64   `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec
}

func AllAvtran(db XODB, callback func(x Avtran) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`avtuniquesys, avtdate, avtcode, avtdescription, avtawarded, avtbillno, avtsenttoavios, avtwriteoff, avtaviosno, avttransactno, avtrejectcode, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.avtrans `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		a := Avtran{}

		// scan
		err = q.Scan(&a.Avtuniquesys, &a.Avtdate, &a.Avtcode, &a.Avtdescription, &a.Avtawarded, &a.Avtbillno, &a.Avtsenttoavios, &a.Avtwriteoff, &a.Avtaviosno, &a.Avttransactno, &a.Avtrejectcode, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(a) {
			return nil
		}
	}

	return nil
}

// AvtranByEquinoxLrn retrieves a row from 'equinox.avtrans' as a Avtran.
//
// Generated from index 'avtrans_pkey'.
func AvtranByEquinoxLrn(db XODB, equinoxLrn int64) (*Avtran, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`avtuniquesys, avtdate, avtcode, avtdescription, avtawarded, avtbillno, avtsenttoavios, avtwriteoff, avtaviosno, avttransactno, avtrejectcode, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.avtrans ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Avtran{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Avtuniquesys, &a.Avtdate, &a.Avtcode, &a.Avtdescription, &a.Avtawarded, &a.Avtbillno, &a.Avtsenttoavios, &a.Avtwriteoff, &a.Avtaviosno, &a.Avttransactno, &a.Avtrejectcode, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
