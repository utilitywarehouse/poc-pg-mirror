// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Gdcharge represents a row from 'equinox.gdcharge'.
type Gdcharge struct {
	Gdcstartdate pq.NullTime     `json:"gdcstartdate"` // gdcstartdate
	Gdcenddate   pq.NullTime     `json:"gdcenddate"`   // gdcenddate
	Gdcprovid    sql.NullString  `json:"gdcprovid"`    // gdcprovid
	Gdcrate      sql.NullFloat64 `json:"gdcrate"`      // gdcrate
	EquinoxPrn   sql.NullInt64   `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Gdcharge exists in the database.
func (g *Gdcharge) Exists() bool {
	return g._exists
}

// Deleted provides information if the Gdcharge has been deleted from the database.
func (g *Gdcharge) Deleted() bool {
	return g._deleted
}

// Insert inserts the Gdcharge to the database.
func (g *Gdcharge) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gdcharge (` +
		`gdcstartdate, gdcenddate, gdcprovid, gdcrate, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, g.Gdcstartdate, g.Gdcenddate, g.Gdcprovid, g.Gdcrate, g.EquinoxPrn, g.EquinoxSec)
	err = db.QueryRow(sqlstr, g.Gdcstartdate, g.Gdcenddate, g.Gdcprovid, g.Gdcrate, g.EquinoxPrn, g.EquinoxSec).Scan(&g.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Update updates the Gdcharge in the database.
func (g *Gdcharge) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.gdcharge SET (` +
		`gdcstartdate, gdcenddate, gdcprovid, gdcrate, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE equinox_lrn = $7`

	// run query
	XOLog(sqlstr, g.Gdcstartdate, g.Gdcenddate, g.Gdcprovid, g.Gdcrate, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	_, err = db.Exec(sqlstr, g.Gdcstartdate, g.Gdcenddate, g.Gdcprovid, g.Gdcrate, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	return err
}

// Save saves the Gdcharge to the database.
func (g *Gdcharge) Save(db XODB) error {
	if g.Exists() {
		return g.Update(db)
	}

	return g.Insert(db)
}

// Upsert performs an upsert for Gdcharge.
//
// NOTE: PostgreSQL 9.5+ only
func (g *Gdcharge) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gdcharge (` +
		`gdcstartdate, gdcenddate, gdcprovid, gdcrate, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`gdcstartdate, gdcenddate, gdcprovid, gdcrate, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.gdcstartdate, EXCLUDED.gdcenddate, EXCLUDED.gdcprovid, EXCLUDED.gdcrate, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, g.Gdcstartdate, g.Gdcenddate, g.Gdcprovid, g.Gdcrate, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	_, err = db.Exec(sqlstr, g.Gdcstartdate, g.Gdcenddate, g.Gdcprovid, g.Gdcrate, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Delete deletes the Gdcharge from the database.
func (g *Gdcharge) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.gdcharge WHERE equinox_lrn = $1`

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

// GdchargeByEquinoxLrn retrieves a row from 'equinox.gdcharge' as a Gdcharge.
//
// Generated from index 'gdcharge_pkey'.
func GdchargeByEquinoxLrn(db XODB, equinoxLrn int64) (*Gdcharge, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gdcstartdate, gdcenddate, gdcprovid, gdcrate, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gdcharge ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gdcharge{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gdcstartdate, &g.Gdcenddate, &g.Gdcprovid, &g.Gdcrate, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
