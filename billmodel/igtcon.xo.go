// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Igtcon represents a row from 'equinox.igtcon'.
type Igtcon struct {
	Igtconsuffix   sql.NullString `json:"igtconsuffix"`   // igtconsuffix
	Igtconemail    sql.NullString `json:"igtconemail"`    // igtconemail
	Igtconpassword sql.NullString `json:"igtconpassword"` // igtconpassword
	Igtcontypeid   sql.NullString `json:"igtcontypeid"`   // igtcontypeid
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Igtcon exists in the database.
func (i *Igtcon) Exists() bool {
	return i._exists
}

// Deleted provides information if the Igtcon has been deleted from the database.
func (i *Igtcon) Deleted() bool {
	return i._deleted
}

// Insert inserts the Igtcon to the database.
func (i *Igtcon) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if i._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.igtcon (` +
		`igtconsuffix, igtconemail, igtconpassword, igtcontypeid, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, i.Igtconsuffix, i.Igtconemail, i.Igtconpassword, i.Igtcontypeid, i.EquinoxPrn, i.EquinoxSec)
	err = db.QueryRow(sqlstr, i.Igtconsuffix, i.Igtconemail, i.Igtconpassword, i.Igtcontypeid, i.EquinoxPrn, i.EquinoxSec).Scan(&i.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	i._exists = true

	return nil
}

// Update updates the Igtcon in the database.
func (i *Igtcon) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !i._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if i._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.igtcon SET (` +
		`igtconsuffix, igtconemail, igtconpassword, igtcontypeid, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE equinox_lrn = $7`

	// run query
	XOLog(sqlstr, i.Igtconsuffix, i.Igtconemail, i.Igtconpassword, i.Igtcontypeid, i.EquinoxPrn, i.EquinoxSec, i.EquinoxLrn)
	_, err = db.Exec(sqlstr, i.Igtconsuffix, i.Igtconemail, i.Igtconpassword, i.Igtcontypeid, i.EquinoxPrn, i.EquinoxSec, i.EquinoxLrn)
	return err
}

// Save saves the Igtcon to the database.
func (i *Igtcon) Save(db XODB) error {
	if i.Exists() {
		return i.Update(db)
	}

	return i.Insert(db)
}

// Upsert performs an upsert for Igtcon.
//
// NOTE: PostgreSQL 9.5+ only
func (i *Igtcon) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if i._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.igtcon (` +
		`igtconsuffix, igtconemail, igtconpassword, igtcontypeid, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`igtconsuffix, igtconemail, igtconpassword, igtcontypeid, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.igtconsuffix, EXCLUDED.igtconemail, EXCLUDED.igtconpassword, EXCLUDED.igtcontypeid, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, i.Igtconsuffix, i.Igtconemail, i.Igtconpassword, i.Igtcontypeid, i.EquinoxPrn, i.EquinoxLrn, i.EquinoxSec)
	_, err = db.Exec(sqlstr, i.Igtconsuffix, i.Igtconemail, i.Igtconpassword, i.Igtcontypeid, i.EquinoxPrn, i.EquinoxLrn, i.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	i._exists = true

	return nil
}

// Delete deletes the Igtcon from the database.
func (i *Igtcon) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !i._exists {
		return nil
	}

	// if deleted, bail
	if i._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.igtcon WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, i.EquinoxLrn)
	_, err = db.Exec(sqlstr, i.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	i._deleted = true

	return nil
}

// IgtconByEquinoxLrn retrieves a row from 'equinox.igtcon' as a Igtcon.
//
// Generated from index 'igtcon_pkey'.
func IgtconByEquinoxLrn(db XODB, equinoxLrn int64) (*Igtcon, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`igtconsuffix, igtconemail, igtconpassword, igtcontypeid, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.igtcon ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	i := Igtcon{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&i.Igtconsuffix, &i.Igtconemail, &i.Igtconpassword, &i.Igtcontypeid, &i.EquinoxPrn, &i.EquinoxLrn, &i.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
