// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

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

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Replog exists in the database.
func (r *Replog) Exists() bool {
	return r._exists
}

// Deleted provides information if the Replog has been deleted from the database.
func (r *Replog) Deleted() bool {
	return r._deleted
}

// Insert inserts the Replog to the database.
func (r *Replog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.replog (` +
		`replogdaterun, replogtimerun, replogbyuser, replogonmc, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, r.Replogdaterun, r.Replogtimerun, r.Replogbyuser, r.Replogonmc, r.EquinoxPrn, r.EquinoxSec)
	err = db.QueryRow(sqlstr, r.Replogdaterun, r.Replogtimerun, r.Replogbyuser, r.Replogonmc, r.EquinoxPrn, r.EquinoxSec).Scan(&r.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Update updates the Replog in the database.
func (r *Replog) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.replog SET (` +
		`replogdaterun, replogtimerun, replogbyuser, replogonmc, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE equinox_lrn = $7`

	// run query
	XOLog(sqlstr, r.Replogdaterun, r.Replogtimerun, r.Replogbyuser, r.Replogonmc, r.EquinoxPrn, r.EquinoxSec, r.EquinoxLrn)
	_, err = db.Exec(sqlstr, r.Replogdaterun, r.Replogtimerun, r.Replogbyuser, r.Replogonmc, r.EquinoxPrn, r.EquinoxSec, r.EquinoxLrn)
	return err
}

// Save saves the Replog to the database.
func (r *Replog) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Upsert performs an upsert for Replog.
//
// NOTE: PostgreSQL 9.5+ only
func (r *Replog) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.replog (` +
		`replogdaterun, replogtimerun, replogbyuser, replogonmc, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`replogdaterun, replogtimerun, replogbyuser, replogonmc, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.replogdaterun, EXCLUDED.replogtimerun, EXCLUDED.replogbyuser, EXCLUDED.replogonmc, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, r.Replogdaterun, r.Replogtimerun, r.Replogbyuser, r.Replogonmc, r.EquinoxPrn, r.EquinoxLrn, r.EquinoxSec)
	_, err = db.Exec(sqlstr, r.Replogdaterun, r.Replogtimerun, r.Replogbyuser, r.Replogonmc, r.EquinoxPrn, r.EquinoxLrn, r.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Delete deletes the Replog from the database.
func (r *Replog) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.replog WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, r.EquinoxLrn)
	_, err = db.Exec(sqlstr, r.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

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
	r := Replog{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&r.Replogdaterun, &r.Replogtimerun, &r.Replogbyuser, &r.Replogonmc, &r.EquinoxPrn, &r.EquinoxLrn, &r.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
