// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// BillrunsBlb represents a row from 'equinox.billruns_blb'.
type BillrunsBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BillrunsBlb exists in the database.
func (bb *BillrunsBlb) Exists() bool {
	return bb._exists
}

// Deleted provides information if the BillrunsBlb has been deleted from the database.
func (bb *BillrunsBlb) Deleted() bool {
	return bb._deleted
}

// Insert inserts the BillrunsBlb to the database.
func (bb *BillrunsBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if bb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.billruns_blb (` +
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

// Update updates the BillrunsBlb in the database.
func (bb *BillrunsBlb) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.billruns_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, bb.BlbData, bb.BlbText, bb.BlbLrn)
	_, err = db.Exec(sqlstr, bb.BlbData, bb.BlbText, bb.BlbLrn)
	return err
}

// Save saves the BillrunsBlb to the database.
func (bb *BillrunsBlb) Save(db XODB) error {
	if bb.Exists() {
		return bb.Update(db)
	}

	return bb.Insert(db)
}

// Upsert performs an upsert for BillrunsBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (bb *BillrunsBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if bb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.billruns_blb (` +
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

// Delete deletes the BillrunsBlb from the database.
func (bb *BillrunsBlb) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.billruns_blb WHERE blb_lrn = $1`

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

// BillrunsBlbByBlbLrn retrieves a row from 'equinox.billruns_blb' as a BillrunsBlb.
//
// Generated from index 'billruns_blb_pkey'.
func BillrunsBlbByBlbLrn(db XODB, blbLrn int64) (*BillrunsBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.billruns_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	bb := BillrunsBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
	if err != nil {
		return nil, err
	}

	return &bb, nil
}
