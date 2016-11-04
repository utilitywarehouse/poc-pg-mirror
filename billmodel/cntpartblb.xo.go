// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// CntpartBlb represents a row from 'equinox.cntpart_blb'.
type CntpartBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CntpartBlb exists in the database.
func (cb *CntpartBlb) Exists() bool {
	return cb._exists
}

// Deleted provides information if the CntpartBlb has been deleted from the database.
func (cb *CntpartBlb) Deleted() bool {
	return cb._deleted
}

// Insert inserts the CntpartBlb to the database.
func (cb *CntpartBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if cb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cntpart_blb (` +
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

// Update updates the CntpartBlb in the database.
func (cb *CntpartBlb) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cntpart_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, cb.BlbData, cb.BlbText, cb.BlbLrn)
	_, err = db.Exec(sqlstr, cb.BlbData, cb.BlbText, cb.BlbLrn)
	return err
}

// Save saves the CntpartBlb to the database.
func (cb *CntpartBlb) Save(db XODB) error {
	if cb.Exists() {
		return cb.Update(db)
	}

	return cb.Insert(db)
}

// Upsert performs an upsert for CntpartBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (cb *CntpartBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if cb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cntpart_blb (` +
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

// Delete deletes the CntpartBlb from the database.
func (cb *CntpartBlb) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cntpart_blb WHERE blb_lrn = $1`

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

// CntpartBlbByBlbLrn retrieves a row from 'equinox.cntpart_blb' as a CntpartBlb.
//
// Generated from index 'cntpart_blb_pkey'.
func CntpartBlbByBlbLrn(db XODB, blbLrn int64) (*CntpartBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.cntpart_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CntpartBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
