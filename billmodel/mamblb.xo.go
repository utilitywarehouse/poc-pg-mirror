// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// MamBlb represents a row from 'equinox.mam_blb'.
type MamBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MamBlb exists in the database.
func (mb *MamBlb) Exists() bool {
	return mb._exists
}

// Deleted provides information if the MamBlb has been deleted from the database.
func (mb *MamBlb) Deleted() bool {
	return mb._deleted
}

// Insert inserts the MamBlb to the database.
func (mb *MamBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if mb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mam_blb (` +
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

// Update updates the MamBlb in the database.
func (mb *MamBlb) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.mam_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, mb.BlbData, mb.BlbText, mb.BlbLrn)
	_, err = db.Exec(sqlstr, mb.BlbData, mb.BlbText, mb.BlbLrn)
	return err
}

// Save saves the MamBlb to the database.
func (mb *MamBlb) Save(db XODB) error {
	if mb.Exists() {
		return mb.Update(db)
	}

	return mb.Insert(db)
}

// Upsert performs an upsert for MamBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (mb *MamBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if mb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mam_blb (` +
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

// Delete deletes the MamBlb from the database.
func (mb *MamBlb) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.mam_blb WHERE blb_lrn = $1`

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

// MamBlbByBlbLrn retrieves a row from 'equinox.mam_blb' as a MamBlb.
//
// Generated from index 'mam_blb_pkey'.
func MamBlbByBlbLrn(db XODB, blbLrn int64) (*MamBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mam_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MamBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
