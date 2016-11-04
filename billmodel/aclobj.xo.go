// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Aclobj represents a row from 'equinox.aclobj'.
type Aclobj struct {
	Aclobjectid    sql.NullInt64   `json:"aclobjectid"`    // aclobjectid
	Aclobjecttype  sql.NullString  `json:"aclobjecttype"`  // aclobjecttype
	Aclobjectname  sql.NullString  `json:"aclobjectname"`  // aclobjectname
	Acldescription sql.NullString  `json:"acldescription"` // acldescription
	Aclnotes       sql.NullInt64   `json:"aclnotes"`       // aclnotes
	Aclocount      sql.NullInt64   `json:"aclocount"`      // aclocount
	Aclosparec1    sql.NullString  `json:"aclosparec1"`    // aclosparec1
	Aclosparec2    sql.NullString  `json:"aclosparec2"`    // aclosparec2
	Aclosparen1    sql.NullFloat64 `json:"aclosparen1"`    // aclosparen1
	Aclosparen2    sql.NullFloat64 `json:"aclosparen2"`    // aclosparen2
	Aclolastaction pq.NullTime     `json:"aclolastaction"` // aclolastaction
	Aclolastuser   sql.NullString  `json:"aclolastuser"`   // aclolastuser
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Aclobj exists in the database.
func (a *Aclobj) Exists() bool {
	return a._exists
}

// Deleted provides information if the Aclobj has been deleted from the database.
func (a *Aclobj) Deleted() bool {
	return a._deleted
}

// Insert inserts the Aclobj to the database.
func (a *Aclobj) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.aclobj (` +
		`aclobjectid, aclobjecttype, aclobjectname, acldescription, aclnotes, aclocount, aclosparec1, aclosparec2, aclosparen1, aclosparen2, aclolastaction, aclolastuser, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, a.Aclobjectid, a.Aclobjecttype, a.Aclobjectname, a.Acldescription, a.Aclnotes, a.Aclocount, a.Aclosparec1, a.Aclosparec2, a.Aclosparen1, a.Aclosparen2, a.Aclolastaction, a.Aclolastuser, a.EquinoxSec)
	err = db.QueryRow(sqlstr, a.Aclobjectid, a.Aclobjecttype, a.Aclobjectname, a.Acldescription, a.Aclnotes, a.Aclocount, a.Aclosparec1, a.Aclosparec2, a.Aclosparen1, a.Aclosparen2, a.Aclolastaction, a.Aclolastuser, a.EquinoxSec).Scan(&a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Aclobj in the database.
func (a *Aclobj) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.aclobj SET (` +
		`aclobjectid, aclobjecttype, aclobjectname, acldescription, aclnotes, aclocount, aclosparec1, aclosparec2, aclosparen1, aclosparen2, aclolastaction, aclolastuser, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) WHERE equinox_lrn = $14`

	// run query
	XOLog(sqlstr, a.Aclobjectid, a.Aclobjecttype, a.Aclobjectname, a.Acldescription, a.Aclnotes, a.Aclocount, a.Aclosparec1, a.Aclosparec2, a.Aclosparen1, a.Aclosparen2, a.Aclolastaction, a.Aclolastuser, a.EquinoxSec, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.Aclobjectid, a.Aclobjecttype, a.Aclobjectname, a.Acldescription, a.Aclnotes, a.Aclocount, a.Aclosparec1, a.Aclosparec2, a.Aclosparen1, a.Aclosparen2, a.Aclolastaction, a.Aclolastuser, a.EquinoxSec, a.EquinoxLrn)
	return err
}

// Save saves the Aclobj to the database.
func (a *Aclobj) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Aclobj.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Aclobj) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.aclobj (` +
		`aclobjectid, aclobjecttype, aclobjectname, acldescription, aclnotes, aclocount, aclosparec1, aclosparec2, aclosparen1, aclosparen2, aclolastaction, aclolastuser, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`aclobjectid, aclobjecttype, aclobjectname, acldescription, aclnotes, aclocount, aclosparec1, aclosparec2, aclosparen1, aclosparen2, aclolastaction, aclolastuser, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.aclobjectid, EXCLUDED.aclobjecttype, EXCLUDED.aclobjectname, EXCLUDED.acldescription, EXCLUDED.aclnotes, EXCLUDED.aclocount, EXCLUDED.aclosparec1, EXCLUDED.aclosparec2, EXCLUDED.aclosparen1, EXCLUDED.aclosparen2, EXCLUDED.aclolastaction, EXCLUDED.aclolastuser, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, a.Aclobjectid, a.Aclobjecttype, a.Aclobjectname, a.Acldescription, a.Aclnotes, a.Aclocount, a.Aclosparec1, a.Aclosparec2, a.Aclosparen1, a.Aclosparen2, a.Aclolastaction, a.Aclolastuser, a.EquinoxLrn, a.EquinoxSec)
	_, err = db.Exec(sqlstr, a.Aclobjectid, a.Aclobjecttype, a.Aclobjectname, a.Acldescription, a.Aclnotes, a.Aclocount, a.Aclosparec1, a.Aclosparec2, a.Aclosparen1, a.Aclosparen2, a.Aclolastaction, a.Aclolastuser, a.EquinoxLrn, a.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Aclobj from the database.
func (a *Aclobj) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.aclobj WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AclobjByEquinoxLrn retrieves a row from 'equinox.aclobj' as a Aclobj.
//
// Generated from index 'aclobj_pkey'.
func AclobjByEquinoxLrn(db XODB, equinoxLrn int64) (*Aclobj, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`aclobjectid, aclobjecttype, aclobjectname, acldescription, aclnotes, aclocount, aclosparec1, aclosparec2, aclosparen1, aclosparen2, aclolastaction, aclolastuser, equinox_lrn, equinox_sec ` +
		`FROM equinox.aclobj ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Aclobj{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Aclobjectid, &a.Aclobjecttype, &a.Aclobjectname, &a.Acldescription, &a.Aclnotes, &a.Aclocount, &a.Aclosparec1, &a.Aclosparec2, &a.Aclosparen1, &a.Aclosparen2, &a.Aclolastaction, &a.Aclolastuser, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}