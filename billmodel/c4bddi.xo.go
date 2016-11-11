// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// C4bddi represents a row from 'equinox.c4bddi'.
type C4bddi struct {
	C4bddiaccountno  sql.NullString `json:"c4bddiaccountno"`  // c4bddiaccountno
	C4bddiclifrom    sql.NullString `json:"c4bddiclifrom"`    // c4bddiclifrom
	C4bddiclito      sql.NullString `json:"c4bddiclito"`      // c4bddiclito
	C4bddiclibase    sql.NullString `json:"c4bddiclibase"`    // c4bddiclibase
	C4bnddistartdate pq.NullTime    `json:"c4bnddistartdate"` // c4bnddistartdate
	C4bddienddate    pq.NullTime    `json:"c4bddienddate"`    // c4bddienddate
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllC4bddi(db XODB, callback func(x C4bddi) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`c4bddiaccountno, c4bddiclifrom, c4bddiclito, c4bddiclibase, c4bnddistartdate, c4bddienddate, equinox_lrn, equinox_sec ` +
		`FROM equinox.c4bddi `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := C4bddi{}

		// scan
		err = q.Scan(&c.C4bddiaccountno, &c.C4bddiclifrom, &c.C4bddiclito, &c.C4bddiclibase, &c.C4bnddistartdate, &c.C4bddienddate, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// C4bddiByEquinoxLrn retrieves a row from 'equinox.c4bddi' as a C4bddi.
//
// Generated from index 'c4bddi_pkey'.
func C4bddiByEquinoxLrn(db XODB, equinoxLrn int64) (*C4bddi, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`c4bddiaccountno, c4bddiclifrom, c4bddiclito, c4bddiclibase, c4bnddistartdate, c4bddienddate, equinox_lrn, equinox_sec ` +
		`FROM equinox.c4bddi ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := C4bddi{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.C4bddiaccountno, &c.C4bddiclifrom, &c.C4bddiclito, &c.C4bddiclibase, &c.C4bnddistartdate, &c.C4bddienddate, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
