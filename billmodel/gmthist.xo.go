// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Gmthist represents a row from 'equinox.gmthist'.
type Gmthist struct {
	Gmthfiledate   pq.NullTime    `json:"gmthfiledate"`   // gmthfiledate
	Gmthfilename   sql.NullString `json:"gmthfilename"`   // gmthfilename
	Gmthgref       sql.NullString `json:"gmthgref"`       // gmthgref
	Gmthdateloaded pq.NullTime    `json:"gmthdateloaded"` // gmthdateloaded
	Gmthaction     sql.NullString `json:"gmthaction"`     // gmthaction
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Gmthist exists in the database.
func (g *Gmthist) Exists() bool {
	return g._exists
}

// Deleted provides information if the Gmthist has been deleted from the database.
func (g *Gmthist) Deleted() bool {
	return g._deleted
}

// Insert inserts the Gmthist to the database.
func (g *Gmthist) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gmthist (` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, g.Gmthfiledate, g.Gmthfilename, g.Gmthgref, g.Gmthdateloaded, g.Gmthaction, g.EquinoxPrn, g.EquinoxSec)
	err = db.QueryRow(sqlstr, g.Gmthfiledate, g.Gmthfilename, g.Gmthgref, g.Gmthdateloaded, g.Gmthaction, g.EquinoxPrn, g.EquinoxSec).Scan(&g.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Update updates the Gmthist in the database.
func (g *Gmthist) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.gmthist SET (` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) WHERE equinox_lrn = $8`

	// run query
	XOLog(sqlstr, g.Gmthfiledate, g.Gmthfilename, g.Gmthgref, g.Gmthdateloaded, g.Gmthaction, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	_, err = db.Exec(sqlstr, g.Gmthfiledate, g.Gmthfilename, g.Gmthgref, g.Gmthdateloaded, g.Gmthaction, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	return err
}

// Save saves the Gmthist to the database.
func (g *Gmthist) Save(db XODB) error {
	if g.Exists() {
		return g.Update(db)
	}

	return g.Insert(db)
}

// Upsert performs an upsert for Gmthist.
//
// NOTE: PostgreSQL 9.5+ only
func (g *Gmthist) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gmthist (` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.gmthfiledate, EXCLUDED.gmthfilename, EXCLUDED.gmthgref, EXCLUDED.gmthdateloaded, EXCLUDED.gmthaction, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, g.Gmthfiledate, g.Gmthfilename, g.Gmthgref, g.Gmthdateloaded, g.Gmthaction, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	_, err = db.Exec(sqlstr, g.Gmthfiledate, g.Gmthfilename, g.Gmthgref, g.Gmthdateloaded, g.Gmthaction, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Delete deletes the Gmthist from the database.
func (g *Gmthist) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.gmthist WHERE equinox_lrn = $1`

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

// GmthistByEquinoxLrn retrieves a row from 'equinox.gmthist' as a Gmthist.
//
// Generated from index 'gmthist_pkey'.
func GmthistByEquinoxLrn(db XODB, equinoxLrn int64) (*Gmthist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gmthfiledate, gmthfilename, gmthgref, gmthdateloaded, gmthaction, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gmthist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gmthist{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gmthfiledate, &g.Gmthfilename, &g.Gmthgref, &g.Gmthdateloaded, &g.Gmthaction, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
