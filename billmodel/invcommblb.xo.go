// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// InvcommBlb represents a row from 'equinox.invcomm_blb'.
type InvcommBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the InvcommBlb exists in the database.
func (ib *InvcommBlb) Exists() bool {
	return ib._exists
}

// Deleted provides information if the InvcommBlb has been deleted from the database.
func (ib *InvcommBlb) Deleted() bool {
	return ib._deleted
}

// Insert inserts the InvcommBlb to the database.
func (ib *InvcommBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ib._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.invcomm_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, ib.BlbData, ib.BlbText)
	err = db.QueryRow(sqlstr, ib.BlbData, ib.BlbText).Scan(&ib.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	ib._exists = true

	return nil
}

// Update updates the InvcommBlb in the database.
func (ib *InvcommBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ib._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ib._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.invcomm_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, ib.BlbData, ib.BlbText, ib.BlbLrn)
	_, err = db.Exec(sqlstr, ib.BlbData, ib.BlbText, ib.BlbLrn)
	return err
}

// Save saves the InvcommBlb to the database.
func (ib *InvcommBlb) Save(db XODB) error {
	if ib.Exists() {
		return ib.Update(db)
	}

	return ib.Insert(db)
}

// Upsert performs an upsert for InvcommBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (ib *InvcommBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ib._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.invcomm_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, ib.BlbLrn, ib.BlbData, ib.BlbText)
	_, err = db.Exec(sqlstr, ib.BlbLrn, ib.BlbData, ib.BlbText)
	if err != nil {
		return err
	}

	// set existence
	ib._exists = true

	return nil
}

// Delete deletes the InvcommBlb from the database.
func (ib *InvcommBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ib._exists {
		return nil
	}

	// if deleted, bail
	if ib._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.invcomm_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, ib.BlbLrn)
	_, err = db.Exec(sqlstr, ib.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	ib._deleted = true

	return nil
}

// InvcommBlbByBlbLrn retrieves a row from 'equinox.invcomm_blb' as a InvcommBlb.
//
// Generated from index 'invcomm_blb_pkey'.
func InvcommBlbByBlbLrn(db XODB, blbLrn int64) (*InvcommBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.invcomm_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ib := InvcommBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ib.BlbLrn, &ib.BlbData, &ib.BlbText)
	if err != nil {
		return nil, err
	}

	return &ib, nil
}