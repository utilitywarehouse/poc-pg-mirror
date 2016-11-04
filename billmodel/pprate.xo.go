// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Pprate represents a row from 'equinox.pprates'.
type Pprate struct {
	Pprregion        sql.NullString  `json:"pprregion"`        // pprregion
	Pprmeter         sql.NullString  `json:"pprmeter"`         // pprmeter
	Pprvalidfrom     pq.NullTime     `json:"pprvalidfrom"`     // pprvalidfrom
	Pprvalidto       pq.NullTime     `json:"pprvalidto"`       // pprvalidto
	Pprstandingcharg sql.NullFloat64 `json:"pprstandingcharg"` // pprstandingcharg
	Pprppkwh1        sql.NullFloat64 `json:"pprppkwh1"`        // pprppkwh1
	Pprppkwh2        sql.NullFloat64 `json:"pprppkwh2"`        // pprppkwh2
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Pprate exists in the database.
func (p *Pprate) Exists() bool {
	return p._exists
}

// Deleted provides information if the Pprate has been deleted from the database.
func (p *Pprate) Deleted() bool {
	return p._deleted
}

// Insert inserts the Pprate to the database.
func (p *Pprate) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pprates (` +
		`pprregion, pprmeter, pprvalidfrom, pprvalidto, pprstandingcharg, pprppkwh1, pprppkwh2, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Pprregion, p.Pprmeter, p.Pprvalidfrom, p.Pprvalidto, p.Pprstandingcharg, p.Pprppkwh1, p.Pprppkwh2, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Pprregion, p.Pprmeter, p.Pprvalidfrom, p.Pprvalidto, p.Pprstandingcharg, p.Pprppkwh1, p.Pprppkwh2, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Pprate in the database.
func (p *Pprate) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.pprates SET (` +
		`pprregion, pprmeter, pprvalidfrom, pprvalidto, pprstandingcharg, pprppkwh1, pprppkwh2, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE equinox_lrn = $9`

	// run query
	XOLog(sqlstr, p.Pprregion, p.Pprmeter, p.Pprvalidfrom, p.Pprvalidto, p.Pprstandingcharg, p.Pprppkwh1, p.Pprppkwh2, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Pprregion, p.Pprmeter, p.Pprvalidfrom, p.Pprvalidto, p.Pprstandingcharg, p.Pprppkwh1, p.Pprppkwh2, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Pprate to the database.
func (p *Pprate) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Pprate.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Pprate) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pprates (` +
		`pprregion, pprmeter, pprvalidfrom, pprvalidto, pprstandingcharg, pprppkwh1, pprppkwh2, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`pprregion, pprmeter, pprvalidfrom, pprvalidto, pprstandingcharg, pprppkwh1, pprppkwh2, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.pprregion, EXCLUDED.pprmeter, EXCLUDED.pprvalidfrom, EXCLUDED.pprvalidto, EXCLUDED.pprstandingcharg, EXCLUDED.pprppkwh1, EXCLUDED.pprppkwh2, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Pprregion, p.Pprmeter, p.Pprvalidfrom, p.Pprvalidto, p.Pprstandingcharg, p.Pprppkwh1, p.Pprppkwh2, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Pprregion, p.Pprmeter, p.Pprvalidfrom, p.Pprvalidto, p.Pprstandingcharg, p.Pprppkwh1, p.Pprppkwh2, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Pprate from the database.
func (p *Pprate) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.pprates WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// PprateByEquinoxLrn retrieves a row from 'equinox.pprates' as a Pprate.
//
// Generated from index 'pprates_pkey'.
func PprateByEquinoxLrn(db XODB, equinoxLrn int64) (*Pprate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pprregion, pprmeter, pprvalidfrom, pprvalidto, pprstandingcharg, pprppkwh1, pprppkwh2, equinox_lrn, equinox_sec ` +
		`FROM equinox.pprates ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pprate{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pprregion, &p.Pprmeter, &p.Pprvalidfrom, &p.Pprvalidto, &p.Pprstandingcharg, &p.Pprppkwh1, &p.Pprppkwh2, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}