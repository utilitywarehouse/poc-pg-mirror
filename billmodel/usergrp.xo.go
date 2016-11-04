// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Usergrp represents a row from 'equinox.usergrp'.
type Usergrp struct {
	Usergroupid   sql.NullString `json:"usergroupid"`   // usergroupid
	Ugname        sql.NullString `json:"ugname"`        // ugname
	Ugdescription sql.NullInt64  `json:"ugdescription"` // ugdescription
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Usergrp exists in the database.
func (u *Usergrp) Exists() bool {
	return u._exists
}

// Deleted provides information if the Usergrp has been deleted from the database.
func (u *Usergrp) Deleted() bool {
	return u._deleted
}

// Insert inserts the Usergrp to the database.
func (u *Usergrp) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.usergrp (` +
		`usergroupid, ugname, ugdescription, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, u.Usergroupid, u.Ugname, u.Ugdescription, u.EquinoxSec)
	err = db.QueryRow(sqlstr, u.Usergroupid, u.Ugname, u.Ugdescription, u.EquinoxSec).Scan(&u.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Update updates the Usergrp in the database.
func (u *Usergrp) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.usergrp SET (` +
		`usergroupid, ugname, ugdescription, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE equinox_lrn = $5`

	// run query
	XOLog(sqlstr, u.Usergroupid, u.Ugname, u.Ugdescription, u.EquinoxSec, u.EquinoxLrn)
	_, err = db.Exec(sqlstr, u.Usergroupid, u.Ugname, u.Ugdescription, u.EquinoxSec, u.EquinoxLrn)
	return err
}

// Save saves the Usergrp to the database.
func (u *Usergrp) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Upsert performs an upsert for Usergrp.
//
// NOTE: PostgreSQL 9.5+ only
func (u *Usergrp) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.usergrp (` +
		`usergroupid, ugname, ugdescription, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`usergroupid, ugname, ugdescription, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.usergroupid, EXCLUDED.ugname, EXCLUDED.ugdescription, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, u.Usergroupid, u.Ugname, u.Ugdescription, u.EquinoxLrn, u.EquinoxSec)
	_, err = db.Exec(sqlstr, u.Usergroupid, u.Ugname, u.Ugdescription, u.EquinoxLrn, u.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Delete deletes the Usergrp from the database.
func (u *Usergrp) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.usergrp WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, u.EquinoxLrn)
	_, err = db.Exec(sqlstr, u.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UsergrpByEquinoxLrn retrieves a row from 'equinox.usergrp' as a Usergrp.
//
// Generated from index 'usergrp_pkey'.
func UsergrpByEquinoxLrn(db XODB, equinoxLrn int64) (*Usergrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`usergroupid, ugname, ugdescription, equinox_lrn, equinox_sec ` +
		`FROM equinox.usergrp ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	u := Usergrp{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&u.Usergroupid, &u.Ugname, &u.Ugdescription, &u.EquinoxLrn, &u.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &u, nil
}