// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// CalltypeBlb represents a row from 'equinox.calltype_blb'.
type CalltypeBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CalltypeBlb exists in the database.
func (cb *CalltypeBlb) Exists() bool {
	return cb._exists
}

// Deleted provides information if the CalltypeBlb has been deleted from the database.
func (cb *CalltypeBlb) Deleted() bool {
	return cb._deleted
}

// Insert inserts the CalltypeBlb to the database.
func (cb *CalltypeBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if cb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.calltype_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, cb.BlbData, cb.BlbText)
	err = db.QueryRow(sqlstr, cb.BlbData, cb.BlbText).Scan(&cb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	cb._exists = true

	return nil
}

// Update updates the CalltypeBlb in the database.
func (cb *CalltypeBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !cb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if cb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.calltype_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, cb.BlbData, cb.BlbText, cb.BlbLrn)
	_, err = db.Exec(sqlstr, cb.BlbData, cb.BlbText, cb.BlbLrn)
	return err
}

// Save saves the CalltypeBlb to the database.
func (cb *CalltypeBlb) Save(db XODB) error {
	if cb.Exists() {
		return cb.Update(db)
	}

	return cb.Insert(db)
}

// Upsert performs an upsert for CalltypeBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (cb *CalltypeBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if cb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.calltype_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, cb.BlbLrn, cb.BlbData, cb.BlbText)
	_, err = db.Exec(sqlstr, cb.BlbLrn, cb.BlbData, cb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	cb._exists = true

	return nil
}

// Delete deletes the CalltypeBlb from the database.
func (cb *CalltypeBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !cb._exists {
		return nil
	}

	// if deleted, bail
	if cb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.calltype_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, cb.BlbLrn)
	_, err = db.Exec(sqlstr, cb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	cb._deleted = true

	return nil
}

// CalltypeBlbByBlbLrn retrieves a row from 'equinox.calltype_blb' as a CalltypeBlb.
//
// Generated from index 'calltype_blb_pkey'.
func CalltypeBlbByBlbLrn(db XODB, blbLrn int64) (*CalltypeBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.calltype_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CalltypeBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
