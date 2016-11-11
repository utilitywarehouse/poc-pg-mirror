// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Replog represents a row from 'equinox.replog'.
type Replog struct {
	Replogdaterun pq.NullTime    `json:"replogdaterun"` // replogdaterun
	Replogtimerun pq.NullTime    `json:"replogtimerun"` // replogtimerun
	Replogbyuser  sql.NullString `json:"replogbyuser"`  // replogbyuser
	Replogonmc    sql.NullString `json:"replogonmc"`    // replogonmc
	EquinoxPrn    sql.NullInt64  `json:"equinox_prn"`   // equinox_prn
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec
}

func AllReplog(db XODB, callback func(x Replog) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`replogdaterun, replogtimerun, replogbyuser, replogonmc, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.replog `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		r := Replog{}

		// scan
		err = q.Scan(&r.Replogdaterun, &r.Replogtimerun, &r.Replogbyuser, &r.Replogonmc, &r.EquinoxPrn, &r.EquinoxLrn, &r.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(r) {
			return nil
		}
	}

	return nil
}

// ReplogByEquinoxLrn retrieves a row from 'equinox.replog' as a Replog.
//
// Generated from index 'replog_pkey'.
func ReplogByEquinoxLrn(db XODB, equinoxLrn int64) (*Replog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`replogdaterun, replogtimerun, replogbyuser, replogonmc, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.replog ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	r := Replog{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&r.Replogdaterun, &r.Replogtimerun, &r.Replogbyuser, &r.Replogonmc, &r.EquinoxPrn, &r.EquinoxLrn, &r.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
