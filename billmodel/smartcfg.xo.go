// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Smartcfg represents a row from 'equinox.smartcfg'.
type Smartcfg struct {
	Scfgid         sql.NullInt64  `json:"scfgid"`         // scfgid
	Scfgname       sql.NullString `json:"scfgname"`       // scfgname
	Scfgprov       sql.NullString `json:"scfgprov"`       // scfgprov
	Scfgfileprefix sql.NullString `json:"scfgfileprefix"` // scfgfileprefix
	Scfgmaindir    sql.NullString `json:"scfgmaindir"`    // scfgmaindir
	Scfgspooldir   sql.NullString `json:"scfgspooldir"`   // scfgspooldir
	Scfgsenddir    sql.NullString `json:"scfgsenddir"`    // scfgsenddir
	Scfglogdir     sql.NullString `json:"scfglogdir"`     // scfglogdir
	Scfglive       sql.NullInt64  `json:"scfglive"`       // scfglive
	Scfgmon        sql.NullInt64  `json:"scfgmon"`        // scfgmon
	Scfgtue        sql.NullInt64  `json:"scfgtue"`        // scfgtue
	Scfgwed        sql.NullInt64  `json:"scfgwed"`        // scfgwed
	Scfgthu        sql.NullInt64  `json:"scfgthu"`        // scfgthu
	Scfgfri        sql.NullInt64  `json:"scfgfri"`        // scfgfri
	Scfgsat        sql.NullInt64  `json:"scfgsat"`        // scfgsat
	Scfgsun        sql.NullInt64  `json:"scfgsun"`        // scfgsun
	Scfglastdate   pq.NullTime    `json:"scfglastdate"`   // scfglastdate
	Scfglasttime   pq.NullTime    `json:"scfglasttime"`   // scfglasttime
	Scfgrunonce    sql.NullInt64  `json:"scfgrunonce"`    // scfgrunonce
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Smartcfg exists in the database.
func (s *Smartcfg) Exists() bool {
	return s._exists
}

// Deleted provides information if the Smartcfg has been deleted from the database.
func (s *Smartcfg) Deleted() bool {
	return s._deleted
}

// Insert inserts the Smartcfg to the database.
func (s *Smartcfg) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.smartcfg (` +
		`scfgid, scfgname, scfgprov, scfgfileprefix, scfgmaindir, scfgspooldir, scfgsenddir, scfglogdir, scfglive, scfgmon, scfgtue, scfgwed, scfgthu, scfgfri, scfgsat, scfgsun, scfglastdate, scfglasttime, scfgrunonce, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, s.Scfgid, s.Scfgname, s.Scfgprov, s.Scfgfileprefix, s.Scfgmaindir, s.Scfgspooldir, s.Scfgsenddir, s.Scfglogdir, s.Scfglive, s.Scfgmon, s.Scfgtue, s.Scfgwed, s.Scfgthu, s.Scfgfri, s.Scfgsat, s.Scfgsun, s.Scfglastdate, s.Scfglasttime, s.Scfgrunonce, s.EquinoxSec)
	err = db.QueryRow(sqlstr, s.Scfgid, s.Scfgname, s.Scfgprov, s.Scfgfileprefix, s.Scfgmaindir, s.Scfgspooldir, s.Scfgsenddir, s.Scfglogdir, s.Scfglive, s.Scfgmon, s.Scfgtue, s.Scfgwed, s.Scfgthu, s.Scfgfri, s.Scfgsat, s.Scfgsun, s.Scfglastdate, s.Scfglasttime, s.Scfgrunonce, s.EquinoxSec).Scan(&s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the Smartcfg in the database.
func (s *Smartcfg) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.smartcfg SET (` +
		`scfgid, scfgname, scfgprov, scfgfileprefix, scfgmaindir, scfgspooldir, scfgsenddir, scfglogdir, scfglive, scfgmon, scfgtue, scfgwed, scfgthu, scfgfri, scfgsat, scfgsun, scfglastdate, scfglasttime, scfgrunonce, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20` +
		`) WHERE equinox_lrn = $21`

	// run query
	XOLog(sqlstr, s.Scfgid, s.Scfgname, s.Scfgprov, s.Scfgfileprefix, s.Scfgmaindir, s.Scfgspooldir, s.Scfgsenddir, s.Scfglogdir, s.Scfglive, s.Scfgmon, s.Scfgtue, s.Scfgwed, s.Scfgthu, s.Scfgfri, s.Scfgsat, s.Scfgsun, s.Scfglastdate, s.Scfglasttime, s.Scfgrunonce, s.EquinoxSec, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.Scfgid, s.Scfgname, s.Scfgprov, s.Scfgfileprefix, s.Scfgmaindir, s.Scfgspooldir, s.Scfgsenddir, s.Scfglogdir, s.Scfglive, s.Scfgmon, s.Scfgtue, s.Scfgwed, s.Scfgthu, s.Scfgfri, s.Scfgsat, s.Scfgsun, s.Scfglastdate, s.Scfglasttime, s.Scfgrunonce, s.EquinoxSec, s.EquinoxLrn)
	return err
}

// Save saves the Smartcfg to the database.
func (s *Smartcfg) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for Smartcfg.
//
// NOTE: PostgreSQL 9.5+ only
func (s *Smartcfg) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.smartcfg (` +
		`scfgid, scfgname, scfgprov, scfgfileprefix, scfgmaindir, scfgspooldir, scfgsenddir, scfglogdir, scfglive, scfgmon, scfgtue, scfgwed, scfgthu, scfgfri, scfgsat, scfgsun, scfglastdate, scfglasttime, scfgrunonce, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`scfgid, scfgname, scfgprov, scfgfileprefix, scfgmaindir, scfgspooldir, scfgsenddir, scfglogdir, scfglive, scfgmon, scfgtue, scfgwed, scfgthu, scfgfri, scfgsat, scfgsun, scfglastdate, scfglasttime, scfgrunonce, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.scfgid, EXCLUDED.scfgname, EXCLUDED.scfgprov, EXCLUDED.scfgfileprefix, EXCLUDED.scfgmaindir, EXCLUDED.scfgspooldir, EXCLUDED.scfgsenddir, EXCLUDED.scfglogdir, EXCLUDED.scfglive, EXCLUDED.scfgmon, EXCLUDED.scfgtue, EXCLUDED.scfgwed, EXCLUDED.scfgthu, EXCLUDED.scfgfri, EXCLUDED.scfgsat, EXCLUDED.scfgsun, EXCLUDED.scfglastdate, EXCLUDED.scfglasttime, EXCLUDED.scfgrunonce, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, s.Scfgid, s.Scfgname, s.Scfgprov, s.Scfgfileprefix, s.Scfgmaindir, s.Scfgspooldir, s.Scfgsenddir, s.Scfglogdir, s.Scfglive, s.Scfgmon, s.Scfgtue, s.Scfgwed, s.Scfgthu, s.Scfgfri, s.Scfgsat, s.Scfgsun, s.Scfglastdate, s.Scfglasttime, s.Scfgrunonce, s.EquinoxLrn, s.EquinoxSec)
	_, err = db.Exec(sqlstr, s.Scfgid, s.Scfgname, s.Scfgprov, s.Scfgfileprefix, s.Scfgmaindir, s.Scfgspooldir, s.Scfgsenddir, s.Scfglogdir, s.Scfglive, s.Scfgmon, s.Scfgtue, s.Scfgwed, s.Scfgthu, s.Scfgfri, s.Scfgsat, s.Scfgsun, s.Scfglastdate, s.Scfglasttime, s.Scfgrunonce, s.EquinoxLrn, s.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the Smartcfg from the database.
func (s *Smartcfg) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.smartcfg WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

// SmartcfgByEquinoxLrn retrieves a row from 'equinox.smartcfg' as a Smartcfg.
//
// Generated from index 'smartcfg_pkey'.
func SmartcfgByEquinoxLrn(db XODB, equinoxLrn int64) (*Smartcfg, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`scfgid, scfgname, scfgprov, scfgfileprefix, scfgmaindir, scfgspooldir, scfgsenddir, scfglogdir, scfglive, scfgmon, scfgtue, scfgwed, scfgthu, scfgfri, scfgsat, scfgsun, scfglastdate, scfglasttime, scfgrunonce, equinox_lrn, equinox_sec ` +
		`FROM equinox.smartcfg ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Smartcfg{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Scfgid, &s.Scfgname, &s.Scfgprov, &s.Scfgfileprefix, &s.Scfgmaindir, &s.Scfgspooldir, &s.Scfgsenddir, &s.Scfglogdir, &s.Scfglive, &s.Scfgmon, &s.Scfgtue, &s.Scfgwed, &s.Scfgthu, &s.Scfgfri, &s.Scfgsat, &s.Scfgsun, &s.Scfglastdate, &s.Scfglasttime, &s.Scfgrunonce, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
