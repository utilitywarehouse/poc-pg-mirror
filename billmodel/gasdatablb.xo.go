// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// GasdataBlb represents a row from 'equinox.gasdata_blb'.
type GasdataBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GasdataBlb exists in the database.
func (gb *GasdataBlb) Exists() bool {
	return gb._exists
}

// Deleted provides information if the GasdataBlb has been deleted from the database.
func (gb *GasdataBlb) Deleted() bool {
	return gb._deleted
}

// Insert inserts the GasdataBlb to the database.
func (gb *GasdataBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if gb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gasdata_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, gb.BlbData, gb.BlbText)
	err = db.QueryRow(sqlstr, gb.BlbData, gb.BlbText).Scan(&gb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	gb._exists = true

	return nil
}

// Update updates the GasdataBlb in the database.
func (gb *GasdataBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !gb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if gb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.gasdata_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, gb.BlbData, gb.BlbText, gb.BlbLrn)
	_, err = db.Exec(sqlstr, gb.BlbData, gb.BlbText, gb.BlbLrn)
	return err
}

// Save saves the GasdataBlb to the database.
func (gb *GasdataBlb) Save(db XODB) error {
	if gb.Exists() {
		return gb.Update(db)
	}

	return gb.Insert(db)
}

// Upsert performs an upsert for GasdataBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (gb *GasdataBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if gb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.gasdata_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, gb.BlbLrn, gb.BlbData, gb.BlbText)
	_, err = db.Exec(sqlstr, gb.BlbLrn, gb.BlbData, gb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	gb._exists = true

	return nil
}

// Delete deletes the GasdataBlb from the database.
func (gb *GasdataBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !gb._exists {
		return nil
	}

	// if deleted, bail
	if gb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.gasdata_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, gb.BlbLrn)
	_, err = db.Exec(sqlstr, gb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	gb._deleted = true

	return nil
}

// GasdataBlbByBlbLrn retrieves a row from 'equinox.gasdata_blb' as a GasdataBlb.
//
// Generated from index 'gasdata_blb_pkey'.
func GasdataBlbByBlbLrn(db XODB, blbLrn int64) (*GasdataBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gasdata_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GasdataBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}