// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gpppay represents a row from 'equinox.gpppay'.
type Gpppay struct {
	Gpppoutlet    sql.NullString  `json:"gpppoutlet"`    // gpppoutlet
	Gppptransdate pq.NullTime     `json:"gppptransdate"` // gppptransdate
	Gppptranstime pq.NullTime     `json:"gppptranstime"` // gppptranstime
	Gpppamount    sql.NullFloat64 `json:"gpppamount"`    // gpppamount
	Gpppmsnmatch  sql.NullInt64   `json:"gpppmsnmatch"`  // gpppmsnmatch
	Gpppmpid      sql.NullString  `json:"gpppmpid"`      // gpppmpid
	Gppptca       sql.NullInt64   `json:"gppptca"`       // gppptca
	Gpppinfdate   pq.NullTime     `json:"gpppinfdate"`   // gpppinfdate
	Gpppbalance   sql.NullFloat64 `json:"gpppbalance"`   // gpppbalance
	Gpppstdchrg   sql.NullFloat64 `json:"gpppstdchrg"`   // gpppstdchrg
	Gpppdayrate   sql.NullFloat64 `json:"gpppdayrate"`   // gpppdayrate
	Gpppnightrate sql.NullFloat64 `json:"gpppnightrate"` // gpppnightrate
	Gppptransno   sql.NullInt64   `json:"gppptransno"`   // gppptransno
	EquinoxPrn    sql.NullInt64   `json:"equinox_prn"`   // equinox_prn
	EquinoxLrn    int64           `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64   `json:"equinox_sec"`   // equinox_sec
}

// GpppayByEquinoxLrn retrieves a row from 'equinox.gpppay' as a Gpppay.
//
// Generated from index 'gpppay_pkey'.
func GpppayByEquinoxLrn(db XODB, equinoxLrn int64) (*Gpppay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gpppoutlet, gppptransdate, gppptranstime, gpppamount, gpppmsnmatch, gpppmpid, gppptca, gpppinfdate, gpppbalance, gpppstdchrg, gpppdayrate, gpppnightrate, gppptransno, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gpppay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gpppay{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gpppoutlet, &g.Gppptransdate, &g.Gppptranstime, &g.Gpppamount, &g.Gpppmsnmatch, &g.Gpppmpid, &g.Gppptca, &g.Gpppinfdate, &g.Gpppbalance, &g.Gpppstdchrg, &g.Gpppdayrate, &g.Gpppnightrate, &g.Gppptransno, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
