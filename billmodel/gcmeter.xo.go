// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gcmeter represents a row from 'equinox.gcmeter'.
type Gcmeter struct {
	Gcintmeterregid sql.NullFloat64 `json:"gcintmeterregid"` // gcintmeterregid
	Gcmetertpr      sql.NullInt64   `json:"gcmetertpr"`      // gcmetertpr
	Gcmeterserialid sql.NullString  `json:"gcmeterserialid"` // gcmeterserialid
	Gcmeterregid    sql.NullString  `json:"gcmeterregid"`    // gcmeterregid
	Gcmeterdials    sql.NullInt64   `json:"gcmeterdials"`    // gcmeterdials
	Gcregtype       sql.NullString  `json:"gcregtype"`       // gcregtype
	Gcregeac        sql.NullInt64   `json:"gcregeac"`        // gcregeac
	Gcinstalldate   pq.NullTime     `json:"gcinstalldate"`   // gcinstalldate
	Gcstartdate     pq.NullTime     `json:"gcstartdate"`     // gcstartdate
	Gcsparen1       sql.NullFloat64 `json:"gcsparen1"`       // gcsparen1
	Gcmpan          sql.NullString  `json:"gcmpan"`          // gcmpan
	Gcefffrom       pq.NullTime     `json:"gcefffrom"`       // gcefffrom
	Gcmetertype     sql.NullString  `json:"gcmetertype"`     // gcmetertype
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllGcmeter(db XODB, callback func(x Gcmeter) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gcintmeterregid, gcmetertpr, gcmeterserialid, gcmeterregid, gcmeterdials, gcregtype, gcregeac, gcinstalldate, gcstartdate, gcsparen1, gcmpan, gcefffrom, gcmetertype, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gcmeter `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		g := Gcmeter{}

		// scan
		err = q.Scan(&g.Gcintmeterregid, &g.Gcmetertpr, &g.Gcmeterserialid, &g.Gcmeterregid, &g.Gcmeterdials, &g.Gcregtype, &g.Gcregeac, &g.Gcinstalldate, &g.Gcstartdate, &g.Gcsparen1, &g.Gcmpan, &g.Gcefffrom, &g.Gcmetertype, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(g) {
			return nil
		}
	}

	return nil
}

// GcmeterByEquinoxLrn retrieves a row from 'equinox.gcmeter' as a Gcmeter.
//
// Generated from index 'gcmeter_pkey'.
func GcmeterByEquinoxLrn(db XODB, equinoxLrn int64) (*Gcmeter, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gcintmeterregid, gcmetertpr, gcmeterserialid, gcmeterregid, gcmeterdials, gcregtype, gcregeac, gcinstalldate, gcstartdate, gcsparen1, gcmpan, gcefffrom, gcmetertype, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gcmeter ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gcmeter{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gcintmeterregid, &g.Gcmetertpr, &g.Gcmeterserialid, &g.Gcmeterregid, &g.Gcmeterdials, &g.Gcregtype, &g.Gcregeac, &g.Gcinstalldate, &g.Gcstartdate, &g.Gcsparen1, &g.Gcmpan, &g.Gcefffrom, &g.Gcmetertype, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
