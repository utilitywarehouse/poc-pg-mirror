// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// WhldcommBlb represents a row from 'equinox.whldcomm_blb'.
type WhldcommBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WhldcommBlb exists in the database.
func (wb *WhldcommBlb) Exists() bool {
	return wb._exists
}

// Deleted provides information if the WhldcommBlb has been deleted from the database.
func (wb *WhldcommBlb) Deleted() bool {
	return wb._deleted
}

// Insert inserts the WhldcommBlb to the database.
func (wb *WhldcommBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.whldcomm_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, wb.BlbData, wb.BlbText)
	err = db.QueryRow(sqlstr, wb.BlbData, wb.BlbText).Scan(&wb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	wb._exists = true

	return nil
}

// Update updates the WhldcommBlb in the database.
func (wb *WhldcommBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.whldcomm_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, wb.BlbData, wb.BlbText, wb.BlbLrn)
	_, err = db.Exec(sqlstr, wb.BlbData, wb.BlbText, wb.BlbLrn)
	return err
}

// Save saves the WhldcommBlb to the database.
func (wb *WhldcommBlb) Save(db XODB) error {
	if wb.Exists() {
		return wb.Update(db)
	}

	return wb.Insert(db)
}

// Upsert performs an upsert for WhldcommBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (wb *WhldcommBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if wb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.whldcomm_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, wb.BlbLrn, wb.BlbData, wb.BlbText)
	_, err = db.Exec(sqlstr, wb.BlbLrn, wb.BlbData, wb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	wb._exists = true

	return nil
}

// Delete deletes the WhldcommBlb from the database.
func (wb *WhldcommBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wb._exists {
		return nil
	}

	// if deleted, bail
	if wb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.whldcomm_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, wb.BlbLrn)
	_, err = db.Exec(sqlstr, wb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	wb._deleted = true

	return nil
}

// WhldcommBlbByBlbLrn retrieves a row from 'equinox.whldcomm_blb' as a WhldcommBlb.
//
// Generated from index 'whldcomm_blb_pkey'.
func WhldcommBlbByBlbLrn(db XODB, blbLrn int64) (*WhldcommBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.whldcomm_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	wb := WhldcommBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&wb.BlbLrn, &wb.BlbData, &wb.BlbText)
	if err != nil {
		return nil, err
	}

	return &wb, nil
}
