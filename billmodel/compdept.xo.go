// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Compdept represents a row from 'equinox.compdept'.
type Compdept struct {
	Compdeptno      sql.NullInt64  `json:"compdeptno"`      // compdeptno
	Compdeptname    sql.NullString `json:"compdeptname"`    // compdeptname
	Compdeptsparec1 sql.NullString `json:"compdeptsparec1"` // compdeptsparec1
	Compdeptsparen1 sql.NullInt64  `json:"compdeptsparen1"` // compdeptsparen1
	Compdeptsparel1 sql.NullInt64  `json:"compdeptsparel1"` // compdeptsparel1
	Compdeptspared1 pq.NullTime    `json:"compdeptspared1"` // compdeptspared1
	Compdeptsparem1 sql.NullInt64  `json:"compdeptsparem1"` // compdeptsparem1
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Compdept exists in the database.
func (c *Compdept) Exists() bool {
	return c._exists
}

// Deleted provides information if the Compdept has been deleted from the database.
func (c *Compdept) Deleted() bool {
	return c._deleted
}

// Insert inserts the Compdept to the database.
func (c *Compdept) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.compdept (` +
		`compdeptno, compdeptname, compdeptsparec1, compdeptsparen1, compdeptsparel1, compdeptspared1, compdeptsparem1, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Compdeptno, c.Compdeptname, c.Compdeptsparec1, c.Compdeptsparen1, c.Compdeptsparel1, c.Compdeptspared1, c.Compdeptsparem1, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Compdeptno, c.Compdeptname, c.Compdeptsparec1, c.Compdeptsparen1, c.Compdeptsparel1, c.Compdeptspared1, c.Compdeptsparem1, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Compdept in the database.
func (c *Compdept) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.compdept SET (` +
		`compdeptno, compdeptname, compdeptsparec1, compdeptsparen1, compdeptsparel1, compdeptspared1, compdeptsparem1, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE equinox_lrn = $9`

	// run query
	XOLog(sqlstr, c.Compdeptno, c.Compdeptname, c.Compdeptsparec1, c.Compdeptsparen1, c.Compdeptsparel1, c.Compdeptspared1, c.Compdeptsparem1, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Compdeptno, c.Compdeptname, c.Compdeptsparec1, c.Compdeptsparen1, c.Compdeptsparel1, c.Compdeptspared1, c.Compdeptsparem1, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Compdept to the database.
func (c *Compdept) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Compdept.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Compdept) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.compdept (` +
		`compdeptno, compdeptname, compdeptsparec1, compdeptsparen1, compdeptsparel1, compdeptspared1, compdeptsparem1, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`compdeptno, compdeptname, compdeptsparec1, compdeptsparen1, compdeptsparel1, compdeptspared1, compdeptsparem1, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.compdeptno, EXCLUDED.compdeptname, EXCLUDED.compdeptsparec1, EXCLUDED.compdeptsparen1, EXCLUDED.compdeptsparel1, EXCLUDED.compdeptspared1, EXCLUDED.compdeptsparem1, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Compdeptno, c.Compdeptname, c.Compdeptsparec1, c.Compdeptsparen1, c.Compdeptsparel1, c.Compdeptspared1, c.Compdeptsparem1, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Compdeptno, c.Compdeptname, c.Compdeptsparec1, c.Compdeptsparen1, c.Compdeptsparel1, c.Compdeptspared1, c.Compdeptsparem1, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Compdept from the database.
func (c *Compdept) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.compdept WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CompdeptByEquinoxLrn retrieves a row from 'equinox.compdept' as a Compdept.
//
// Generated from index 'compdept_pkey'.
func CompdeptByEquinoxLrn(db XODB, equinoxLrn int64) (*Compdept, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`compdeptno, compdeptname, compdeptsparec1, compdeptsparen1, compdeptsparel1, compdeptspared1, compdeptsparem1, equinox_lrn, equinox_sec ` +
		`FROM equinox.compdept ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Compdept{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Compdeptno, &c.Compdeptname, &c.Compdeptsparec1, &c.Compdeptsparen1, &c.Compdeptsparel1, &c.Compdeptspared1, &c.Compdeptsparem1, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
