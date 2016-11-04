// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Gasdatum represents a row from 'equinox.gasdata'.
type Gasdatum struct {
	Gwldz      sql.NullString  `json:"gwldz"`       // gwldz
	Gwdate     pq.NullTime     `json:"gwdate"`      // gwdate
	Gwewcf     sql.NullFloat64 `json:"gwewcf"`      // gwewcf
	Gwalp01    sql.NullFloat64 `json:"gwalp01"`     // gwalp01
	Gwalp02    sql.NullFloat64 `json:"gwalp02"`     // gwalp02
	Gwdaf01    sql.NullFloat64 `json:"gwdaf01"`     // gwdaf01
	Gwdaf02    sql.NullFloat64 `json:"gwdaf02"`     // gwdaf02
	EquinoxLrn int64           `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64   `json:"equinox_sec"` // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Gasdatum exists in the database.
func (g *Gasdatum) Exists() bool {
	return g._exists
}

// Deleted provides information if the Gasdatum has been deleted from the database.
func (g *Gasdatum) Deleted() bool {
	return g._deleted
}

// Insert inserts the Gasdatum to the database.
func (g *Gasdatum) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gasdata (` +
		`gwldz, gwdate, gwewcf, gwalp01, gwalp02, gwdaf01, gwdaf02, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, g.Gwldz, g.Gwdate, g.Gwewcf, g.Gwalp01, g.Gwalp02, g.Gwdaf01, g.Gwdaf02, g.EquinoxSec)
	err = db.QueryRow(sqlstr, g.Gwldz, g.Gwdate, g.Gwewcf, g.Gwalp01, g.Gwalp02, g.Gwdaf01, g.Gwdaf02, g.EquinoxSec).Scan(&g.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Update updates the Gasdatum in the database.
func (g *Gasdatum) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.gasdata SET (` +
		`gwldz, gwdate, gwewcf, gwalp01, gwalp02, gwdaf01, gwdaf02, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE equinox_lrn = $9`

	// run query
	XOLog(sqlstr, g.Gwldz, g.Gwdate, g.Gwewcf, g.Gwalp01, g.Gwalp02, g.Gwdaf01, g.Gwdaf02, g.EquinoxSec, g.EquinoxLrn)
	_, err = db.Exec(sqlstr, g.Gwldz, g.Gwdate, g.Gwewcf, g.Gwalp01, g.Gwalp02, g.Gwdaf01, g.Gwdaf02, g.EquinoxSec, g.EquinoxLrn)
	return err
}

// Save saves the Gasdatum to the database.
func (g *Gasdatum) Save(db XODB) error {
	if g.Exists() {
		return g.Update(db)
	}

	return g.Insert(db)
}

// Upsert performs an upsert for Gasdatum.
//
// NOTE: PostgreSQL 9.5+ only
func (g *Gasdatum) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gasdata (` +
		`gwldz, gwdate, gwewcf, gwalp01, gwalp02, gwdaf01, gwdaf02, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`gwldz, gwdate, gwewcf, gwalp01, gwalp02, gwdaf01, gwdaf02, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.gwldz, EXCLUDED.gwdate, EXCLUDED.gwewcf, EXCLUDED.gwalp01, EXCLUDED.gwalp02, EXCLUDED.gwdaf01, EXCLUDED.gwdaf02, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, g.Gwldz, g.Gwdate, g.Gwewcf, g.Gwalp01, g.Gwalp02, g.Gwdaf01, g.Gwdaf02, g.EquinoxLrn, g.EquinoxSec)
	_, err = db.Exec(sqlstr, g.Gwldz, g.Gwdate, g.Gwewcf, g.Gwalp01, g.Gwalp02, g.Gwdaf01, g.Gwdaf02, g.EquinoxLrn, g.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Delete deletes the Gasdatum from the database.
func (g *Gasdatum) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.gasdata WHERE equinox_lrn = $1`

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

// GasdatumByEquinoxLrn retrieves a row from 'equinox.gasdata' as a Gasdatum.
//
// Generated from index 'gasdata_pkey'.
func GasdatumByEquinoxLrn(db XODB, equinoxLrn int64) (*Gasdatum, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gwldz, gwdate, gwewcf, gwalp01, gwalp02, gwdaf01, gwdaf02, equinox_lrn, equinox_sec ` +
		`FROM equinox.gasdata ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gasdatum{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gwldz, &g.Gwdate, &g.Gwewcf, &g.Gwalp01, &g.Gwalp02, &g.Gwdaf01, &g.Gwdaf02, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}