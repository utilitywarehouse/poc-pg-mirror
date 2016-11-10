// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Exgthdat represents a row from 'equinox.exgthdat'.
type Exgthdat struct {
	Exgthdatyr       sql.NullString  `json:"exgthdatyr"`       // exgthdatyr
	Exgthdatert      sql.NullFloat64 `json:"exgthdatert"`      // exgthdatert
	Exgthdatnert     sql.NullFloat64 `json:"exgthdatnert"`     // exgthdatnert
	Exgthdatdebt     sql.NullFloat64 `json:"exgthdatdebt"`     // exgthdatdebt
	Exgthdatpfact    sql.NullFloat64 `json:"exgthdatpfact"`    // exgthdatpfact
	Exgthlastbilldat pq.NullTime     `json:"exgthlastbilldat"` // exgthlastbilldat
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// ExgthdatByEquinoxLrn retrieves a row from 'equinox.exgthdat' as a Exgthdat.
//
// Generated from index 'exgthdat_pkey'.
func ExgthdatByEquinoxLrn(db XODB, equinoxLrn int64) (*Exgthdat, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`exgthdatyr, exgthdatert, exgthdatnert, exgthdatdebt, exgthdatpfact, exgthlastbilldat, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.exgthdat ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Exgthdat{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Exgthdatyr, &e.Exgthdatert, &e.Exgthdatnert, &e.Exgthdatdebt, &e.Exgthdatpfact, &e.Exgthlastbilldat, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
