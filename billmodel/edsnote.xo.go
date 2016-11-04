// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Edsnote represents a row from 'equinox.edsnotes'.
type Edsnote struct {
	Edsnnotes        sql.NullInt64  `json:"edsnnotes"`        // edsnnotes
	Edsnentereddate  pq.NullTime    `json:"edsnentereddate"`  // edsnentereddate
	Edsnenteredtime  pq.NullTime    `json:"edsnenteredtime"`  // edsnenteredtime
	Edsnenteredby    sql.NullString `json:"edsnenteredby"`    // edsnenteredby
	Edsndatereminder pq.NullTime    `json:"edsndatereminder"` // edsndatereminder
	Edsndiarise      sql.NullInt64  `json:"edsndiarise"`      // edsndiarise
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Edsnote exists in the database.
func (e *Edsnote) Exists() bool {
	return e._exists
}

// Deleted provides information if the Edsnote has been deleted from the database.
func (e *Edsnote) Deleted() bool {
	return e._deleted
}

// Insert inserts the Edsnote to the database.
func (e *Edsnote) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.edsnotes (` +
		`edsnnotes, edsnentereddate, edsnenteredtime, edsnenteredby, edsndatereminder, edsndiarise, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.Edsnnotes, e.Edsnentereddate, e.Edsnenteredtime, e.Edsnenteredby, e.Edsndatereminder, e.Edsndiarise, e.EquinoxPrn, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.Edsnnotes, e.Edsnentereddate, e.Edsnenteredtime, e.Edsnenteredby, e.Edsndatereminder, e.Edsndiarise, e.EquinoxPrn, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Edsnote in the database.
func (e *Edsnote) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.edsnotes SET (` +
		`edsnnotes, edsnentereddate, edsnenteredtime, edsnenteredby, edsndatereminder, edsndiarise, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE equinox_lrn = $9`

	// run query
	XOLog(sqlstr, e.Edsnnotes, e.Edsnentereddate, e.Edsnenteredtime, e.Edsnenteredby, e.Edsndatereminder, e.Edsndiarise, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.Edsnnotes, e.Edsnentereddate, e.Edsnenteredtime, e.Edsnenteredby, e.Edsndatereminder, e.Edsndiarise, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Edsnote to the database.
func (e *Edsnote) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Edsnote.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Edsnote) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.edsnotes (` +
		`edsnnotes, edsnentereddate, edsnenteredtime, edsnenteredby, edsndatereminder, edsndiarise, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`edsnnotes, edsnentereddate, edsnenteredtime, edsnenteredby, edsndatereminder, edsndiarise, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.edsnnotes, EXCLUDED.edsnentereddate, EXCLUDED.edsnenteredtime, EXCLUDED.edsnenteredby, EXCLUDED.edsndatereminder, EXCLUDED.edsndiarise, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.Edsnnotes, e.Edsnentereddate, e.Edsnenteredtime, e.Edsnenteredby, e.Edsndatereminder, e.Edsndiarise, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.Edsnnotes, e.Edsnentereddate, e.Edsnenteredtime, e.Edsnenteredby, e.Edsndatereminder, e.Edsndiarise, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Edsnote from the database.
func (e *Edsnote) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.edsnotes WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// EdsnoteByEquinoxLrn retrieves a row from 'equinox.edsnotes' as a Edsnote.
//
// Generated from index 'edsnotes_pkey'.
func EdsnoteByEquinoxLrn(db XODB, equinoxLrn int64) (*Edsnote, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`edsnnotes, edsnentereddate, edsnenteredtime, edsnenteredby, edsndatereminder, edsndiarise, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.edsnotes ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Edsnote{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Edsnnotes, &e.Edsnentereddate, &e.Edsnenteredtime, &e.Edsnenteredby, &e.Edsndatereminder, &e.Edsndiarise, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}