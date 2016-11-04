// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// LlulogBlb represents a row from 'equinox.llulog_blb'.
type LlulogBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the LlulogBlb exists in the database.
func (lb *LlulogBlb) Exists() bool {
	return lb._exists
}

// Deleted provides information if the LlulogBlb has been deleted from the database.
func (lb *LlulogBlb) Deleted() bool {
	return lb._deleted
}

// Insert inserts the LlulogBlb to the database.
func (lb *LlulogBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if lb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.llulog_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, lb.BlbData, lb.BlbText)
	err = db.QueryRow(sqlstr, lb.BlbData, lb.BlbText).Scan(&lb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	lb._exists = true

	return nil
}

// Update updates the LlulogBlb in the database.
func (lb *LlulogBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !lb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if lb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.llulog_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, lb.BlbData, lb.BlbText, lb.BlbLrn)
	_, err = db.Exec(sqlstr, lb.BlbData, lb.BlbText, lb.BlbLrn)
	return err
}

// Save saves the LlulogBlb to the database.
func (lb *LlulogBlb) Save(db XODB) error {
	if lb.Exists() {
		return lb.Update(db)
	}

	return lb.Insert(db)
}

// Upsert performs an upsert for LlulogBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (lb *LlulogBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if lb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.llulog_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, lb.BlbLrn, lb.BlbData, lb.BlbText)
	_, err = db.Exec(sqlstr, lb.BlbLrn, lb.BlbData, lb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	lb._exists = true

	return nil
}

// Delete deletes the LlulogBlb from the database.
func (lb *LlulogBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !lb._exists {
		return nil
	}

	// if deleted, bail
	if lb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.llulog_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, lb.BlbLrn)
	_, err = db.Exec(sqlstr, lb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	lb._deleted = true

	return nil
}

// LlulogBlbByBlbLrn retrieves a row from 'equinox.llulog_blb' as a LlulogBlb.
//
// Generated from index 'llulog_blb_pkey'.
func LlulogBlbByBlbLrn(db XODB, blbLrn int64) (*LlulogBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.llulog_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	lb := LlulogBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&lb.BlbLrn, &lb.BlbData, &lb.BlbText)
	if err != nil {
		return nil, err
	}

	return &lb, nil
}
