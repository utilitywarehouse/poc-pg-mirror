// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Exempl represents a row from 'equinox.exempl'.
type Exempl struct {
	Exempid      sql.NullString `json:"exempid"`      // exempid
	Exempname    sql.NullString `json:"exempname"`    // exempname
	Exempnrgdecl sql.NullInt64  `json:"exempnrgdecl"` // exempnrgdecl
	Exempidcard  sql.NullInt64  `json:"exempidcard"`  // exempidcard
	Exemptrained sql.NullInt64  `json:"exemptrained"` // exemptrained
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Exempl exists in the database.
func (e *Exempl) Exists() bool {
	return e._exists
}

// Deleted provides information if the Exempl has been deleted from the database.
func (e *Exempl) Deleted() bool {
	return e._deleted
}

// Insert inserts the Exempl to the database.
func (e *Exempl) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.exempl (` +
		`exempid, exempname, exempnrgdecl, exempidcard, exemptrained, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.Exempid, e.Exempname, e.Exempnrgdecl, e.Exempidcard, e.Exemptrained, e.EquinoxPrn, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.Exempid, e.Exempname, e.Exempnrgdecl, e.Exempidcard, e.Exemptrained, e.EquinoxPrn, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Exempl in the database.
func (e *Exempl) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.exempl SET (` +
		`exempid, exempname, exempnrgdecl, exempidcard, exemptrained, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) WHERE equinox_lrn = $8`

	// run query
	XOLog(sqlstr, e.Exempid, e.Exempname, e.Exempnrgdecl, e.Exempidcard, e.Exemptrained, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.Exempid, e.Exempname, e.Exempnrgdecl, e.Exempidcard, e.Exemptrained, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Exempl to the database.
func (e *Exempl) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Exempl.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Exempl) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.exempl (` +
		`exempid, exempname, exempnrgdecl, exempidcard, exemptrained, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`exempid, exempname, exempnrgdecl, exempidcard, exemptrained, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.exempid, EXCLUDED.exempname, EXCLUDED.exempnrgdecl, EXCLUDED.exempidcard, EXCLUDED.exemptrained, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.Exempid, e.Exempname, e.Exempnrgdecl, e.Exempidcard, e.Exemptrained, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.Exempid, e.Exempname, e.Exempnrgdecl, e.Exempidcard, e.Exemptrained, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Exempl from the database.
func (e *Exempl) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.exempl WHERE equinox_lrn = $1`

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

// ExemplByEquinoxLrn retrieves a row from 'equinox.exempl' as a Exempl.
//
// Generated from index 'exempl_pkey'.
func ExemplByEquinoxLrn(db XODB, equinoxLrn int64) (*Exempl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`exempid, exempname, exempnrgdecl, exempidcard, exemptrained, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.exempl ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Exempl{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Exempid, &e.Exempname, &e.Exempnrgdecl, &e.Exempidcard, &e.Exemptrained, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
