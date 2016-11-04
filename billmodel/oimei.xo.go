// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Oimei represents a row from 'equinox.oimei'.
type Oimei struct {
	Oimeikey         sql.NullString `json:"oimeikey"`         // oimeikey
	Oimeimake        sql.NullString `json:"oimeimake"`        // oimeimake
	Oimeimodel       sql.NullString `json:"oimeimodel"`       // oimeimodel
	Oimeiother       sql.NullString `json:"oimeiother"`       // oimeiother
	Oimeidateentered pq.NullTime    `json:"oimeidateentered"` // oimeidateentered
	Oimeistocklevel  sql.NullInt64  `json:"oimeistocklevel"`  // oimeistocklevel
	Oimeicriticallev sql.NullInt64  `json:"oimeicriticallev"` // oimeicriticallev
	Oimeiorder       sql.NullString `json:"oimeiorder"`       // oimeiorder
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Oimei exists in the database.
func (o *Oimei) Exists() bool {
	return o._exists
}

// Deleted provides information if the Oimei has been deleted from the database.
func (o *Oimei) Deleted() bool {
	return o._deleted
}

// Insert inserts the Oimei to the database.
func (o *Oimei) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.oimei (` +
		`oimeikey, oimeimake, oimeimodel, oimeiother, oimeidateentered, oimeistocklevel, oimeicriticallev, oimeiorder, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, o.Oimeikey, o.Oimeimake, o.Oimeimodel, o.Oimeiother, o.Oimeidateentered, o.Oimeistocklevel, o.Oimeicriticallev, o.Oimeiorder, o.EquinoxPrn, o.EquinoxSec)
	err = db.QueryRow(sqlstr, o.Oimeikey, o.Oimeimake, o.Oimeimodel, o.Oimeiother, o.Oimeidateentered, o.Oimeistocklevel, o.Oimeicriticallev, o.Oimeiorder, o.EquinoxPrn, o.EquinoxSec).Scan(&o.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Update updates the Oimei in the database.
func (o *Oimei) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if o._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.oimei SET (` +
		`oimeikey, oimeimake, oimeimodel, oimeiother, oimeidateentered, oimeistocklevel, oimeicriticallev, oimeiorder, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) WHERE equinox_lrn = $11`

	// run query
	XOLog(sqlstr, o.Oimeikey, o.Oimeimake, o.Oimeimodel, o.Oimeiother, o.Oimeidateentered, o.Oimeistocklevel, o.Oimeicriticallev, o.Oimeiorder, o.EquinoxPrn, o.EquinoxSec, o.EquinoxLrn)
	_, err = db.Exec(sqlstr, o.Oimeikey, o.Oimeimake, o.Oimeimodel, o.Oimeiother, o.Oimeidateentered, o.Oimeistocklevel, o.Oimeicriticallev, o.Oimeiorder, o.EquinoxPrn, o.EquinoxSec, o.EquinoxLrn)
	return err
}

// Save saves the Oimei to the database.
func (o *Oimei) Save(db XODB) error {
	if o.Exists() {
		return o.Update(db)
	}

	return o.Insert(db)
}

// Upsert performs an upsert for Oimei.
//
// NOTE: PostgreSQL 9.5+ only
func (o *Oimei) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.oimei (` +
		`oimeikey, oimeimake, oimeimodel, oimeiother, oimeidateentered, oimeistocklevel, oimeicriticallev, oimeiorder, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`oimeikey, oimeimake, oimeimodel, oimeiother, oimeidateentered, oimeistocklevel, oimeicriticallev, oimeiorder, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.oimeikey, EXCLUDED.oimeimake, EXCLUDED.oimeimodel, EXCLUDED.oimeiother, EXCLUDED.oimeidateentered, EXCLUDED.oimeistocklevel, EXCLUDED.oimeicriticallev, EXCLUDED.oimeiorder, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, o.Oimeikey, o.Oimeimake, o.Oimeimodel, o.Oimeiother, o.Oimeidateentered, o.Oimeistocklevel, o.Oimeicriticallev, o.Oimeiorder, o.EquinoxPrn, o.EquinoxLrn, o.EquinoxSec)
	_, err = db.Exec(sqlstr, o.Oimeikey, o.Oimeimake, o.Oimeimodel, o.Oimeiother, o.Oimeidateentered, o.Oimeistocklevel, o.Oimeicriticallev, o.Oimeiorder, o.EquinoxPrn, o.EquinoxLrn, o.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Delete deletes the Oimei from the database.
func (o *Oimei) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return nil
	}

	// if deleted, bail
	if o._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.oimei WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, o.EquinoxLrn)
	_, err = db.Exec(sqlstr, o.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	o._deleted = true

	return nil
}

// OimeiByEquinoxLrn retrieves a row from 'equinox.oimei' as a Oimei.
//
// Generated from index 'oimei_pkey'.
func OimeiByEquinoxLrn(db XODB, equinoxLrn int64) (*Oimei, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`oimeikey, oimeimake, oimeimodel, oimeiother, oimeidateentered, oimeistocklevel, oimeicriticallev, oimeiorder, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.oimei ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	o := Oimei{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&o.Oimeikey, &o.Oimeimake, &o.Oimeimodel, &o.Oimeiother, &o.Oimeidateentered, &o.Oimeistocklevel, &o.Oimeicriticallev, &o.Oimeiorder, &o.EquinoxPrn, &o.EquinoxLrn, &o.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
