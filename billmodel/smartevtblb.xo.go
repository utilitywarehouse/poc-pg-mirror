// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// SmartevtBlb represents a row from 'equinox.smartevt_blb'.
type SmartevtBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the SmartevtBlb exists in the database.
func (sb *SmartevtBlb) Exists() bool {
	return sb._exists
}

// Deleted provides information if the SmartevtBlb has been deleted from the database.
func (sb *SmartevtBlb) Deleted() bool {
	return sb._deleted
}

// Insert inserts the SmartevtBlb to the database.
func (sb *SmartevtBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if sb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.smartevt_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, sb.BlbData, sb.BlbText)
	err = db.QueryRow(sqlstr, sb.BlbData, sb.BlbText).Scan(&sb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	sb._exists = true

	return nil
}

// Update updates the SmartevtBlb in the database.
func (sb *SmartevtBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !sb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if sb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.smartevt_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, sb.BlbData, sb.BlbText, sb.BlbLrn)
	_, err = db.Exec(sqlstr, sb.BlbData, sb.BlbText, sb.BlbLrn)
	return err
}

// Save saves the SmartevtBlb to the database.
func (sb *SmartevtBlb) Save(db XODB) error {
	if sb.Exists() {
		return sb.Update(db)
	}

	return sb.Insert(db)
}

// Upsert performs an upsert for SmartevtBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (sb *SmartevtBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if sb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.smartevt_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, sb.BlbLrn, sb.BlbData, sb.BlbText)
	_, err = db.Exec(sqlstr, sb.BlbLrn, sb.BlbData, sb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	sb._exists = true

	return nil
}

// Delete deletes the SmartevtBlb from the database.
func (sb *SmartevtBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !sb._exists {
		return nil
	}

	// if deleted, bail
	if sb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.smartevt_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, sb.BlbLrn)
	_, err = db.Exec(sqlstr, sb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	sb._deleted = true

	return nil
}

// SmartevtBlbByBlbLrn retrieves a row from 'equinox.smartevt_blb' as a SmartevtBlb.
//
// Generated from index 'smartevt_blb_pkey'.
func SmartevtBlbByBlbLrn(db XODB, blbLrn int64) (*SmartevtBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.smartevt_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	sb := SmartevtBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
	if err != nil {
		return nil, err
	}

	return &sb, nil
}