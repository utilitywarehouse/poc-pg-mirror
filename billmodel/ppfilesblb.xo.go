// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// PpfilesBlb represents a row from 'equinox.ppfiles_blb'.
type PpfilesBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PpfilesBlb exists in the database.
func (pb *PpfilesBlb) Exists() bool {
	return pb._exists
}

// Deleted provides information if the PpfilesBlb has been deleted from the database.
func (pb *PpfilesBlb) Deleted() bool {
	return pb._deleted
}

// Insert inserts the PpfilesBlb to the database.
func (pb *PpfilesBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ppfiles_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, pb.BlbData, pb.BlbText)
	err = db.QueryRow(sqlstr, pb.BlbData, pb.BlbText).Scan(&pb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	pb._exists = true

	return nil
}

// Update updates the PpfilesBlb in the database.
func (pb *PpfilesBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.ppfiles_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, pb.BlbData, pb.BlbText, pb.BlbLrn)
	_, err = db.Exec(sqlstr, pb.BlbData, pb.BlbText, pb.BlbLrn)
	return err
}

// Save saves the PpfilesBlb to the database.
func (pb *PpfilesBlb) Save(db XODB) error {
	if pb.Exists() {
		return pb.Update(db)
	}

	return pb.Insert(db)
}

// Upsert performs an upsert for PpfilesBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (pb *PpfilesBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if pb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ppfiles_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, pb.BlbLrn, pb.BlbData, pb.BlbText)
	_, err = db.Exec(sqlstr, pb.BlbLrn, pb.BlbData, pb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	pb._exists = true

	return nil
}

// Delete deletes the PpfilesBlb from the database.
func (pb *PpfilesBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pb._exists {
		return nil
	}

	// if deleted, bail
	if pb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.ppfiles_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, pb.BlbLrn)
	_, err = db.Exec(sqlstr, pb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	pb._deleted = true

	return nil
}

// PpfilesBlbByBlbLrn retrieves a row from 'equinox.ppfiles_blb' as a PpfilesBlb.
//
// Generated from index 'ppfiles_blb_pkey'.
func PpfilesBlbByBlbLrn(db XODB, blbLrn int64) (*PpfilesBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ppfiles_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PpfilesBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}