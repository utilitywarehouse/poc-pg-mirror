// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// BishistBlb represents a row from 'equinox.bishist_blb'.
type BishistBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BishistBlb exists in the database.
func (bb *BishistBlb) Exists() bool {
	return bb._exists
}

// Deleted provides information if the BishistBlb has been deleted from the database.
func (bb *BishistBlb) Deleted() bool {
	return bb._deleted
}

// Insert inserts the BishistBlb to the database.
func (bb *BishistBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if bb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.bishist_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, bb.BlbData, bb.BlbText)
	err = db.QueryRow(sqlstr, bb.BlbData, bb.BlbText).Scan(&bb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	bb._exists = true

	return nil
}

// Update updates the BishistBlb in the database.
func (bb *BishistBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !bb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if bb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.bishist_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, bb.BlbData, bb.BlbText, bb.BlbLrn)
	_, err = db.Exec(sqlstr, bb.BlbData, bb.BlbText, bb.BlbLrn)
	return err
}

// Save saves the BishistBlb to the database.
func (bb *BishistBlb) Save(db XODB) error {
	if bb.Exists() {
		return bb.Update(db)
	}

	return bb.Insert(db)
}

// Upsert performs an upsert for BishistBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (bb *BishistBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if bb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.bishist_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, bb.BlbLrn, bb.BlbData, bb.BlbText)
	_, err = db.Exec(sqlstr, bb.BlbLrn, bb.BlbData, bb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	bb._exists = true

	return nil
}

// Delete deletes the BishistBlb from the database.
func (bb *BishistBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !bb._exists {
		return nil
	}

	// if deleted, bail
	if bb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.bishist_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, bb.BlbLrn)
	_, err = db.Exec(sqlstr, bb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	bb._deleted = true

	return nil
}

// BishistBlbByBlbLrn retrieves a row from 'equinox.bishist_blb' as a BishistBlb.
//
// Generated from index 'bishist_blb_pkey'.
func BishistBlbByBlbLrn(db XODB, blbLrn int64) (*BishistBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.bishist_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	bb := BishistBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
	if err != nil {
		return nil, err
	}

	return &bb, nil
}
