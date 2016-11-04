// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Gvisit represents a row from 'equinox.gvisit'.
type Gvisit struct {
	Gvisitdate       pq.NullTime    `json:"gvisitdate"`       // gvisitdate
	Gvisitoutcome    sql.NullString `json:"gvisitoutcome"`    // gvisitoutcome
	Gvisitnotes      sql.NullInt64  `json:"gvisitnotes"`      // gvisitnotes
	Gvisitreading    sql.NullString `json:"gvisitreading"`    // gvisitreading
	Gvisitreadtype   sql.NullString `json:"gvisitreadtype"`   // gvisitreadtype
	Gvisitfilename   sql.NullString `json:"gvisitfilename"`   // gvisitfilename
	Gvisitimportdate pq.NullTime    `json:"gvisitimportdate"` // gvisitimportdate
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Gvisit exists in the database.
func (g *Gvisit) Exists() bool {
	return g._exists
}

// Deleted provides information if the Gvisit has been deleted from the database.
func (g *Gvisit) Deleted() bool {
	return g._deleted
}

// Insert inserts the Gvisit to the database.
func (g *Gvisit) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gvisit (` +
		`gvisitdate, gvisitoutcome, gvisitnotes, gvisitreading, gvisitreadtype, gvisitfilename, gvisitimportdate, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, g.Gvisitdate, g.Gvisitoutcome, g.Gvisitnotes, g.Gvisitreading, g.Gvisitreadtype, g.Gvisitfilename, g.Gvisitimportdate, g.EquinoxPrn, g.EquinoxSec)
	err = db.QueryRow(sqlstr, g.Gvisitdate, g.Gvisitoutcome, g.Gvisitnotes, g.Gvisitreading, g.Gvisitreadtype, g.Gvisitfilename, g.Gvisitimportdate, g.EquinoxPrn, g.EquinoxSec).Scan(&g.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Update updates the Gvisit in the database.
func (g *Gvisit) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !g._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if g._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.gvisit SET (` +
		`gvisitdate, gvisitoutcome, gvisitnotes, gvisitreading, gvisitreadtype, gvisitfilename, gvisitimportdate, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) WHERE equinox_lrn = $10`

	// run query
	XOLog(sqlstr, g.Gvisitdate, g.Gvisitoutcome, g.Gvisitnotes, g.Gvisitreading, g.Gvisitreadtype, g.Gvisitfilename, g.Gvisitimportdate, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	_, err = db.Exec(sqlstr, g.Gvisitdate, g.Gvisitoutcome, g.Gvisitnotes, g.Gvisitreading, g.Gvisitreadtype, g.Gvisitfilename, g.Gvisitimportdate, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	return err
}

// Save saves the Gvisit to the database.
func (g *Gvisit) Save(db XODB) error {
	if g.Exists() {
		return g.Update(db)
	}

	return g.Insert(db)
}

// Upsert performs an upsert for Gvisit.
//
// NOTE: PostgreSQL 9.5+ only
func (g *Gvisit) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gvisit (` +
		`gvisitdate, gvisitoutcome, gvisitnotes, gvisitreading, gvisitreadtype, gvisitfilename, gvisitimportdate, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`gvisitdate, gvisitoutcome, gvisitnotes, gvisitreading, gvisitreadtype, gvisitfilename, gvisitimportdate, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.gvisitdate, EXCLUDED.gvisitoutcome, EXCLUDED.gvisitnotes, EXCLUDED.gvisitreading, EXCLUDED.gvisitreadtype, EXCLUDED.gvisitfilename, EXCLUDED.gvisitimportdate, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, g.Gvisitdate, g.Gvisitoutcome, g.Gvisitnotes, g.Gvisitreading, g.Gvisitreadtype, g.Gvisitfilename, g.Gvisitimportdate, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	_, err = db.Exec(sqlstr, g.Gvisitdate, g.Gvisitoutcome, g.Gvisitnotes, g.Gvisitreading, g.Gvisitreadtype, g.Gvisitfilename, g.Gvisitimportdate, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Delete deletes the Gvisit from the database.
func (g *Gvisit) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !g._exists {
		return nil
	}

	// if deleted, bail
	if g._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.gvisit WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, g.EquinoxLrn)
	_, err = db.Exec(sqlstr, g.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	g._deleted = true

	return nil
}

// GvisitByEquinoxLrn retrieves a row from 'equinox.gvisit' as a Gvisit.
//
// Generated from index 'gvisit_pkey'.
func GvisitByEquinoxLrn(db XODB, equinoxLrn int64) (*Gvisit, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gvisitdate, gvisitoutcome, gvisitnotes, gvisitreading, gvisitreadtype, gvisitfilename, gvisitimportdate, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gvisit ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gvisit{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gvisitdate, &g.Gvisitoutcome, &g.Gvisitnotes, &g.Gvisitreading, &g.Gvisitreadtype, &g.Gvisitfilename, &g.Gvisitimportdate, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
