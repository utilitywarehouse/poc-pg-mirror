// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// EdspayBlb represents a row from 'equinox.edspay_blb'.
type EdspayBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the EdspayBlb exists in the database.
func (eb *EdspayBlb) Exists() bool {
	return eb._exists
}

// Deleted provides information if the EdspayBlb has been deleted from the database.
func (eb *EdspayBlb) Deleted() bool {
	return eb._deleted
}

// Insert inserts the EdspayBlb to the database.
func (eb *EdspayBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if eb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.edspay_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, eb.BlbData, eb.BlbText)
	err = db.QueryRow(sqlstr, eb.BlbData, eb.BlbText).Scan(&eb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	eb._exists = true

	return nil
}

// Update updates the EdspayBlb in the database.
func (eb *EdspayBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !eb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if eb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.edspay_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, eb.BlbData, eb.BlbText, eb.BlbLrn)
	_, err = db.Exec(sqlstr, eb.BlbData, eb.BlbText, eb.BlbLrn)
	return err
}

// Save saves the EdspayBlb to the database.
func (eb *EdspayBlb) Save(db XODB) error {
	if eb.Exists() {
		return eb.Update(db)
	}

	return eb.Insert(db)
}

// Upsert performs an upsert for EdspayBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (eb *EdspayBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if eb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.edspay_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, eb.BlbLrn, eb.BlbData, eb.BlbText)
	_, err = db.Exec(sqlstr, eb.BlbLrn, eb.BlbData, eb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	eb._exists = true

	return nil
}

// Delete deletes the EdspayBlb from the database.
func (eb *EdspayBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !eb._exists {
		return nil
	}

	// if deleted, bail
	if eb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.edspay_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, eb.BlbLrn)
	_, err = db.Exec(sqlstr, eb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	eb._deleted = true

	return nil
}

// EdspayBlbByBlbLrn retrieves a row from 'equinox.edspay_blb' as a EdspayBlb.
//
// Generated from index 'edspay_blb_pkey'.
func EdspayBlbByBlbLrn(db XODB, blbLrn int64) (*EdspayBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.edspay_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	eb := EdspayBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
	if err != nil {
		return nil, err
	}

	return &eb, nil
}
