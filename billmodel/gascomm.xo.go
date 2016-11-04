// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Gascomm represents a row from 'equinox.gascomm'.
type Gascomm struct {
	Gcommdate        pq.NullTime     `json:"gcommdate"`        // gcommdate
	Gcommcomment     sql.NullInt64   `json:"gcommcomment"`     // gcommcomment
	Gcommenteredby   sql.NullString  `json:"gcommenteredby"`   // gcommenteredby
	Gcommcommentcode sql.NullString  `json:"gcommcommentcode"` // gcommcommentcode
	Gcommspecific    sql.NullString  `json:"gcommspecific"`    // gcommspecific
	Gcommoutcome     sql.NullString  `json:"gcommoutcome"`     // gcommoutcome
	Gcommsparec1     sql.NullString  `json:"gcommsparec1"`     // gcommsparec1
	Gcommsparen1     sql.NullFloat64 `json:"gcommsparen1"`     // gcommsparen1
	Gcommspared1     pq.NullTime     `json:"gcommspared1"`     // gcommspared1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Gascomm exists in the database.
func (g *Gascomm) Exists() bool {
	return g._exists
}

// Deleted provides information if the Gascomm has been deleted from the database.
func (g *Gascomm) Deleted() bool {
	return g._deleted
}

// Insert inserts the Gascomm to the database.
func (g *Gascomm) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gascomm (` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, g.Gcommdate, g.Gcommcomment, g.Gcommenteredby, g.Gcommcommentcode, g.Gcommspecific, g.Gcommoutcome, g.Gcommsparec1, g.Gcommsparen1, g.Gcommspared1, g.EquinoxPrn, g.EquinoxSec)
	err = db.QueryRow(sqlstr, g.Gcommdate, g.Gcommcomment, g.Gcommenteredby, g.Gcommcommentcode, g.Gcommspecific, g.Gcommoutcome, g.Gcommsparec1, g.Gcommsparen1, g.Gcommspared1, g.EquinoxPrn, g.EquinoxSec).Scan(&g.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Update updates the Gascomm in the database.
func (g *Gascomm) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.gascomm SET (` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) WHERE equinox_lrn = $12`

	// run query
	XOLog(sqlstr, g.Gcommdate, g.Gcommcomment, g.Gcommenteredby, g.Gcommcommentcode, g.Gcommspecific, g.Gcommoutcome, g.Gcommsparec1, g.Gcommsparen1, g.Gcommspared1, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	_, err = db.Exec(sqlstr, g.Gcommdate, g.Gcommcomment, g.Gcommenteredby, g.Gcommcommentcode, g.Gcommspecific, g.Gcommoutcome, g.Gcommsparec1, g.Gcommsparen1, g.Gcommspared1, g.EquinoxPrn, g.EquinoxSec, g.EquinoxLrn)
	return err
}

// Save saves the Gascomm to the database.
func (g *Gascomm) Save(db XODB) error {
	if g.Exists() {
		return g.Update(db)
	}

	return g.Insert(db)
}

// Upsert performs an upsert for Gascomm.
//
// NOTE: PostgreSQL 9.5+ only
func (g *Gascomm) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if g._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gascomm (` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.gcommdate, EXCLUDED.gcommcomment, EXCLUDED.gcommenteredby, EXCLUDED.gcommcommentcode, EXCLUDED.gcommspecific, EXCLUDED.gcommoutcome, EXCLUDED.gcommsparec1, EXCLUDED.gcommsparen1, EXCLUDED.gcommspared1, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, g.Gcommdate, g.Gcommcomment, g.Gcommenteredby, g.Gcommcommentcode, g.Gcommspecific, g.Gcommoutcome, g.Gcommsparec1, g.Gcommsparen1, g.Gcommspared1, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	_, err = db.Exec(sqlstr, g.Gcommdate, g.Gcommcomment, g.Gcommenteredby, g.Gcommcommentcode, g.Gcommspecific, g.Gcommoutcome, g.Gcommsparec1, g.Gcommsparen1, g.Gcommspared1, g.EquinoxPrn, g.EquinoxLrn, g.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	g._exists = true

	return nil
}

// Delete deletes the Gascomm from the database.
func (g *Gascomm) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.gascomm WHERE equinox_lrn = $1`

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

// GascommByEquinoxLrn retrieves a row from 'equinox.gascomm' as a Gascomm.
//
// Generated from index 'gascomm_pkey'.
func GascommByEquinoxLrn(db XODB, equinoxLrn int64) (*Gascomm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gcommdate, gcommcomment, gcommenteredby, gcommcommentcode, gcommspecific, gcommoutcome, gcommsparec1, gcommsparen1, gcommspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gascomm ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gascomm{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gcommdate, &g.Gcommcomment, &g.Gcommenteredby, &g.Gcommcommentcode, &g.Gcommspecific, &g.Gcommoutcome, &g.Gcommsparec1, &g.Gcommsparen1, &g.Gcommspared1, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
