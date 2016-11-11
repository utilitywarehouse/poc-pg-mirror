// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Ctfixed represents a row from 'equinox.ctfixed'.
type Ctfixed struct {
	Ctfixeduser      sql.NullString  `json:"ctfixeduser"`      // ctfixeduser
	Ctfixedscgd      sql.NullFloat64 `json:"ctfixedscgd"`      // ctfixedscgd
	Ctfixedgdrate    sql.NullFloat64 `json:"ctfixedgdrate"`    // ctfixedgdrate
	Ctfixedsce7      sql.NullFloat64 `json:"ctfixedsce7"`      // ctfixedsce7
	Ctfixede7r1      sql.NullFloat64 `json:"ctfixede7r1"`      // ctfixede7r1
	Ctfixede7r2      sql.NullFloat64 `json:"ctfixede7r2"`      // ctfixede7r2
	Ctfixedscgas     sql.NullFloat64 `json:"ctfixedscgas"`     // ctfixedscgas
	Ctfixedgasrate   sql.NullFloat64 `json:"ctfixedgasrate"`   // ctfixedgasrate
	Ctfixedannualg   sql.NullFloat64 `json:"ctfixedannualg"`   // ctfixedannualg
	Ctfixedannuale   sql.NullFloat64 `json:"ctfixedannuale"`   // ctfixedannuale
	Ctfixedannuald   sql.NullFloat64 `json:"ctfixedannuald"`   // ctfixedannuald
	Ctfixedannualde7 sql.NullFloat64 `json:"ctfixedannualde7"` // ctfixedannualde7
	Ctfixedn1        sql.NullFloat64 `json:"ctfixedn1"`        // ctfixedn1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCtfixed(db XODB, callback func(x Ctfixed) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`ctfixeduser, ctfixedscgd, ctfixedgdrate, ctfixedsce7, ctfixede7r1, ctfixede7r2, ctfixedscgas, ctfixedgasrate, ctfixedannualg, ctfixedannuale, ctfixedannuald, ctfixedannualde7, ctfixedn1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ctfixed `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Ctfixed{}

		// scan
		err = q.Scan(&c.Ctfixeduser, &c.Ctfixedscgd, &c.Ctfixedgdrate, &c.Ctfixedsce7, &c.Ctfixede7r1, &c.Ctfixede7r2, &c.Ctfixedscgas, &c.Ctfixedgasrate, &c.Ctfixedannualg, &c.Ctfixedannuale, &c.Ctfixedannuald, &c.Ctfixedannualde7, &c.Ctfixedn1, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CtfixedByEquinoxLrn retrieves a row from 'equinox.ctfixed' as a Ctfixed.
//
// Generated from index 'ctfixed_pkey'.
func CtfixedByEquinoxLrn(db XODB, equinoxLrn int64) (*Ctfixed, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ctfixeduser, ctfixedscgd, ctfixedgdrate, ctfixedsce7, ctfixede7r1, ctfixede7r2, ctfixedscgas, ctfixedgasrate, ctfixedannualg, ctfixedannuale, ctfixedannuald, ctfixedannualde7, ctfixedn1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ctfixed ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Ctfixed{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Ctfixeduser, &c.Ctfixedscgd, &c.Ctfixedgdrate, &c.Ctfixedsce7, &c.Ctfixede7r1, &c.Ctfixede7r2, &c.Ctfixedscgas, &c.Ctfixedgasrate, &c.Ctfixedannualg, &c.Ctfixedannuale, &c.Ctfixedannuald, &c.Ctfixedannualde7, &c.Ctfixedn1, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
