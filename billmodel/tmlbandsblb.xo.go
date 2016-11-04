// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// TmlbandsBlb represents a row from 'equinox.tmlbands_blb'.
type TmlbandsBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TmlbandsBlb exists in the database.
func (tb *TmlbandsBlb) Exists() bool {
	return tb._exists
}

// Deleted provides information if the TmlbandsBlb has been deleted from the database.
func (tb *TmlbandsBlb) Deleted() bool {
	return tb._deleted
}

// Insert inserts the TmlbandsBlb to the database.
func (tb *TmlbandsBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if tb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.tmlbands_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, tb.BlbData, tb.BlbText)
	err = db.QueryRow(sqlstr, tb.BlbData, tb.BlbText).Scan(&tb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	tb._exists = true

	return nil
}

// Update updates the TmlbandsBlb in the database.
func (tb *TmlbandsBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if tb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.tmlbands_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, tb.BlbData, tb.BlbText, tb.BlbLrn)
	_, err = db.Exec(sqlstr, tb.BlbData, tb.BlbText, tb.BlbLrn)
	return err
}

// Save saves the TmlbandsBlb to the database.
func (tb *TmlbandsBlb) Save(db XODB) error {
	if tb.Exists() {
		return tb.Update(db)
	}

	return tb.Insert(db)
}

// Upsert performs an upsert for TmlbandsBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (tb *TmlbandsBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if tb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.tmlbands_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, tb.BlbLrn, tb.BlbData, tb.BlbText)
	_, err = db.Exec(sqlstr, tb.BlbLrn, tb.BlbData, tb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	tb._exists = true

	return nil
}

// Delete deletes the TmlbandsBlb from the database.
func (tb *TmlbandsBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tb._exists {
		return nil
	}

	// if deleted, bail
	if tb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.tmlbands_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, tb.BlbLrn)
	_, err = db.Exec(sqlstr, tb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	tb._deleted = true

	return nil
}

// TmlbandsBlbByBlbLrn retrieves a row from 'equinox.tmlbands_blb' as a TmlbandsBlb.
//
// Generated from index 'tmlbands_blb_pkey'.
func TmlbandsBlbByBlbLrn(db XODB, blbLrn int64) (*TmlbandsBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.tmlbands_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	tb := TmlbandsBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&tb.BlbLrn, &tb.BlbData, &tb.BlbText)
	if err != nil {
		return nil, err
	}

	return &tb, nil
}
