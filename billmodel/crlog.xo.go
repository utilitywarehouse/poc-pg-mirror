// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Crlog represents a row from 'equinox.crlog'.
type Crlog struct {
	Crldate      pq.NullTime    `json:"crldate"`      // crldate
	Crltime      pq.NullTime    `json:"crltime"`      // crltime
	Crlparamfile sql.NullString `json:"crlparamfile"` // crlparamfile
	Crltitle     sql.NullString `json:"crltitle"`     // crltitle
	Crlentry     sql.NullInt64  `json:"crlentry"`     // crlentry
	Crlorder     sql.NullInt64  `json:"crlorder"`     // crlorder
	Crlaction    sql.NullString `json:"crlaction"`    // crlaction
	Crlstep      sql.NullString `json:"crlstep"`      // crlstep
	Crlcompleted sql.NullInt64  `json:"crlcompleted"` // crlcompleted
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec
}

func AllCrlog(db XODB, callback func(x Crlog) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`crldate, crltime, crlparamfile, crltitle, crlentry, crlorder, crlaction, crlstep, crlcompleted, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crlog `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Crlog{}

		// scan
		err = q.Scan(&c.Crldate, &c.Crltime, &c.Crlparamfile, &c.Crltitle, &c.Crlentry, &c.Crlorder, &c.Crlaction, &c.Crlstep, &c.Crlcompleted, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CrlogByEquinoxLrn retrieves a row from 'equinox.crlog' as a Crlog.
//
// Generated from index 'crlog_pkey'.
func CrlogByEquinoxLrn(db XODB, equinoxLrn int64) (*Crlog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`crldate, crltime, crlparamfile, crltitle, crlentry, crlorder, crlaction, crlstep, crlcompleted, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crlog ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Crlog{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Crldate, &c.Crltime, &c.Crlparamfile, &c.Crltitle, &c.Crlentry, &c.Crlorder, &c.Crlaction, &c.Crlstep, &c.Crlcompleted, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
