// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Propval represents a row from 'equinox.propval'.
type Propval struct {
	Pvuniquesys   sql.NullFloat64 `json:"pvuniquesys"`   // pvuniquesys
	Pvpropertyid  sql.NullFloat64 `json:"pvpropertyid"`  // pvpropertyid
	Pvuniqueid    sql.NullString  `json:"pvuniqueid"`    // pvuniqueid
	Pvvalue       sql.NullString  `json:"pvvalue"`       // pvvalue
	Pvdescription sql.NullInt64   `json:"pvdescription"` // pvdescription
	Pvbeforesave  sql.NullInt64   `json:"pvbeforesave"`  // pvbeforesave
	Pvaftersave   sql.NullInt64   `json:"pvaftersave"`   // pvaftersave
	Pvbeforefield sql.NullInt64   `json:"pvbeforefield"` // pvbeforefield
	Pvafterfield  sql.NullInt64   `json:"pvafterfield"`  // pvafterfield
	Pvadditional  sql.NullInt64   `json:"pvadditional"`  // pvadditional
	EquinoxLrn    int64           `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64   `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Propval exists in the database.
func (p *Propval) Exists() bool {
	return p._exists
}

// Deleted provides information if the Propval has been deleted from the database.
func (p *Propval) Deleted() bool {
	return p._deleted
}

// Insert inserts the Propval to the database.
func (p *Propval) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.propval (` +
		`pvuniquesys, pvpropertyid, pvuniqueid, pvvalue, pvdescription, pvbeforesave, pvaftersave, pvbeforefield, pvafterfield, pvadditional, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Pvuniquesys, p.Pvpropertyid, p.Pvuniqueid, p.Pvvalue, p.Pvdescription, p.Pvbeforesave, p.Pvaftersave, p.Pvbeforefield, p.Pvafterfield, p.Pvadditional, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Pvuniquesys, p.Pvpropertyid, p.Pvuniqueid, p.Pvvalue, p.Pvdescription, p.Pvbeforesave, p.Pvaftersave, p.Pvbeforefield, p.Pvafterfield, p.Pvadditional, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Propval in the database.
func (p *Propval) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.propval SET (` +
		`pvuniquesys, pvpropertyid, pvuniqueid, pvvalue, pvdescription, pvbeforesave, pvaftersave, pvbeforefield, pvafterfield, pvadditional, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) WHERE equinox_lrn = $12`

	// run query
	XOLog(sqlstr, p.Pvuniquesys, p.Pvpropertyid, p.Pvuniqueid, p.Pvvalue, p.Pvdescription, p.Pvbeforesave, p.Pvaftersave, p.Pvbeforefield, p.Pvafterfield, p.Pvadditional, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Pvuniquesys, p.Pvpropertyid, p.Pvuniqueid, p.Pvvalue, p.Pvdescription, p.Pvbeforesave, p.Pvaftersave, p.Pvbeforefield, p.Pvafterfield, p.Pvadditional, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Propval to the database.
func (p *Propval) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Propval.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Propval) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.propval (` +
		`pvuniquesys, pvpropertyid, pvuniqueid, pvvalue, pvdescription, pvbeforesave, pvaftersave, pvbeforefield, pvafterfield, pvadditional, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`pvuniquesys, pvpropertyid, pvuniqueid, pvvalue, pvdescription, pvbeforesave, pvaftersave, pvbeforefield, pvafterfield, pvadditional, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.pvuniquesys, EXCLUDED.pvpropertyid, EXCLUDED.pvuniqueid, EXCLUDED.pvvalue, EXCLUDED.pvdescription, EXCLUDED.pvbeforesave, EXCLUDED.pvaftersave, EXCLUDED.pvbeforefield, EXCLUDED.pvafterfield, EXCLUDED.pvadditional, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Pvuniquesys, p.Pvpropertyid, p.Pvuniqueid, p.Pvvalue, p.Pvdescription, p.Pvbeforesave, p.Pvaftersave, p.Pvbeforefield, p.Pvafterfield, p.Pvadditional, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Pvuniquesys, p.Pvpropertyid, p.Pvuniqueid, p.Pvvalue, p.Pvdescription, p.Pvbeforesave, p.Pvaftersave, p.Pvbeforefield, p.Pvafterfield, p.Pvadditional, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Propval from the database.
func (p *Propval) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.propval WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// PropvalByEquinoxLrn retrieves a row from 'equinox.propval' as a Propval.
//
// Generated from index 'propval_pkey'.
func PropvalByEquinoxLrn(db XODB, equinoxLrn int64) (*Propval, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pvuniquesys, pvpropertyid, pvuniqueid, pvvalue, pvdescription, pvbeforesave, pvaftersave, pvbeforefield, pvafterfield, pvadditional, equinox_lrn, equinox_sec ` +
		`FROM equinox.propval ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Propval{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pvuniquesys, &p.Pvpropertyid, &p.Pvuniqueid, &p.Pvvalue, &p.Pvdescription, &p.Pvbeforesave, &p.Pvaftersave, &p.Pvbeforefield, &p.Pvafterfield, &p.Pvadditional, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}