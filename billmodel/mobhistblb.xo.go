// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// MobhistBlb represents a row from 'equinox.mobhist_blb'.
type MobhistBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MobhistBlb exists in the database.
func (mb *MobhistBlb) Exists() bool {
	return mb._exists
}

// Deleted provides information if the MobhistBlb has been deleted from the database.
func (mb *MobhistBlb) Deleted() bool {
	return mb._deleted
}

// Insert inserts the MobhistBlb to the database.
func (mb *MobhistBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if mb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mobhist_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, mb.BlbData, mb.BlbText)
	err = db.QueryRow(sqlstr, mb.BlbData, mb.BlbText).Scan(&mb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	mb._exists = true

	return nil
}

// Update updates the MobhistBlb in the database.
func (mb *MobhistBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if mb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.mobhist_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, mb.BlbData, mb.BlbText, mb.BlbLrn)
	_, err = db.Exec(sqlstr, mb.BlbData, mb.BlbText, mb.BlbLrn)
	return err
}

// Save saves the MobhistBlb to the database.
func (mb *MobhistBlb) Save(db XODB) error {
	if mb.Exists() {
		return mb.Update(db)
	}

	return mb.Insert(db)
}

// Upsert performs an upsert for MobhistBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (mb *MobhistBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if mb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mobhist_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, mb.BlbLrn, mb.BlbData, mb.BlbText)
	_, err = db.Exec(sqlstr, mb.BlbLrn, mb.BlbData, mb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	mb._exists = true

	return nil
}

// Delete deletes the MobhistBlb from the database.
func (mb *MobhistBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mb._exists {
		return nil
	}

	// if deleted, bail
	if mb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.mobhist_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, mb.BlbLrn)
	_, err = db.Exec(sqlstr, mb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	mb._deleted = true

	return nil
}

// MobhistBlbByBlbLrn retrieves a row from 'equinox.mobhist_blb' as a MobhistBlb.
//
// Generated from index 'mobhist_blb_pkey'.
func MobhistBlbByBlbLrn(db XODB, blbLrn int64) (*MobhistBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mobhist_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MobhistBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
