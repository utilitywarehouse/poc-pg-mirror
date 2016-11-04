// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Cpcode represents a row from 'equinox.cpcodes'.
type Cpcode struct {
	Cpcnumber        sql.NullInt64   `json:"cpcnumber"`        // cpcnumber
	Cpccodeid        sql.NullInt64   `json:"cpccodeid"`        // cpccodeid
	Cpccode          sql.NullString  `json:"cpccode"`          // cpccode
	Cpcdescription   sql.NullString  `json:"cpcdescription"`   // cpcdescription
	Cpcdisplayorder  sql.NullInt64   `json:"cpcdisplayorder"`  // cpcdisplayorder
	Cpcpercentnett   sql.NullFloat64 `json:"cpcpercentnett"`   // cpcpercentnett
	Cpcpercentnett2  sql.NullFloat64 `json:"cpcpercentnett2"`  // cpcpercentnett2
	Cpcpercentnett3  sql.NullFloat64 `json:"cpcpercentnett3"`  // cpcpercentnett3
	Cpcpercentnett4  sql.NullFloat64 `json:"cpcpercentnett4"`  // cpcpercentnett4
	Cpcpercentnett5  sql.NullFloat64 `json:"cpcpercentnett5"`  // cpcpercentnett5
	Cpcpercentnett6  sql.NullFloat64 `json:"cpcpercentnett6"`  // cpcpercentnett6
	Cpcpercentnett7  sql.NullFloat64 `json:"cpcpercentnett7"`  // cpcpercentnett7
	Cpcpercentnett8  sql.NullFloat64 `json:"cpcpercentnett8"`  // cpcpercentnett8
	Cpcpercentnett9  sql.NullFloat64 `json:"cpcpercentnett9"`  // cpcpercentnett9
	Cpcpercentnett10 sql.NullFloat64 `json:"cpcpercentnett10"` // cpcpercentnett10
	Cpcpercentnett11 sql.NullFloat64 `json:"cpcpercentnett11"` // cpcpercentnett11
	Cpcpercentnett12 sql.NullFloat64 `json:"cpcpercentnett12"` // cpcpercentnett12
	Cpcpercentnett13 sql.NullFloat64 `json:"cpcpercentnett13"` // cpcpercentnett13
	Cpcsparen1       sql.NullFloat64 `json:"cpcsparen1"`       // cpcsparen1
	Cpcsparen2       sql.NullFloat64 `json:"cpcsparen2"`       // cpcsparen2
	Cpcsparec1       sql.NullString  `json:"cpcsparec1"`       // cpcsparec1
	Cpcsparec2       sql.NullString  `json:"cpcsparec2"`       // cpcsparec2
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cpcode exists in the database.
func (c *Cpcode) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cpcode has been deleted from the database.
func (c *Cpcode) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cpcode to the database.
func (c *Cpcode) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cpcodes (` +
		`cpcnumber, cpccodeid, cpccode, cpcdescription, cpcdisplayorder, cpcpercentnett, cpcpercentnett2, cpcpercentnett3, cpcpercentnett4, cpcpercentnett5, cpcpercentnett6, cpcpercentnett7, cpcpercentnett8, cpcpercentnett9, cpcpercentnett10, cpcpercentnett11, cpcpercentnett12, cpcpercentnett13, cpcsparen1, cpcsparen2, cpcsparec1, cpcsparec2, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cpcnumber, c.Cpccodeid, c.Cpccode, c.Cpcdescription, c.Cpcdisplayorder, c.Cpcpercentnett, c.Cpcpercentnett2, c.Cpcpercentnett3, c.Cpcpercentnett4, c.Cpcpercentnett5, c.Cpcpercentnett6, c.Cpcpercentnett7, c.Cpcpercentnett8, c.Cpcpercentnett9, c.Cpcpercentnett10, c.Cpcpercentnett11, c.Cpcpercentnett12, c.Cpcpercentnett13, c.Cpcsparen1, c.Cpcsparen2, c.Cpcsparec1, c.Cpcsparec2, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cpcnumber, c.Cpccodeid, c.Cpccode, c.Cpcdescription, c.Cpcdisplayorder, c.Cpcpercentnett, c.Cpcpercentnett2, c.Cpcpercentnett3, c.Cpcpercentnett4, c.Cpcpercentnett5, c.Cpcpercentnett6, c.Cpcpercentnett7, c.Cpcpercentnett8, c.Cpcpercentnett9, c.Cpcpercentnett10, c.Cpcpercentnett11, c.Cpcpercentnett12, c.Cpcpercentnett13, c.Cpcsparen1, c.Cpcsparen2, c.Cpcsparec1, c.Cpcsparec2, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cpcode in the database.
func (c *Cpcode) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.cpcodes SET (` +
		`cpcnumber, cpccodeid, cpccode, cpcdescription, cpcdisplayorder, cpcpercentnett, cpcpercentnett2, cpcpercentnett3, cpcpercentnett4, cpcpercentnett5, cpcpercentnett6, cpcpercentnett7, cpcpercentnett8, cpcpercentnett9, cpcpercentnett10, cpcpercentnett11, cpcpercentnett12, cpcpercentnett13, cpcsparen1, cpcsparen2, cpcsparec1, cpcsparec2, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24` +
		`) WHERE equinox_lrn = $25`

	// run query
	XOLog(sqlstr, c.Cpcnumber, c.Cpccodeid, c.Cpccode, c.Cpcdescription, c.Cpcdisplayorder, c.Cpcpercentnett, c.Cpcpercentnett2, c.Cpcpercentnett3, c.Cpcpercentnett4, c.Cpcpercentnett5, c.Cpcpercentnett6, c.Cpcpercentnett7, c.Cpcpercentnett8, c.Cpcpercentnett9, c.Cpcpercentnett10, c.Cpcpercentnett11, c.Cpcpercentnett12, c.Cpcpercentnett13, c.Cpcsparen1, c.Cpcsparen2, c.Cpcsparec1, c.Cpcsparec2, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cpcnumber, c.Cpccodeid, c.Cpccode, c.Cpcdescription, c.Cpcdisplayorder, c.Cpcpercentnett, c.Cpcpercentnett2, c.Cpcpercentnett3, c.Cpcpercentnett4, c.Cpcpercentnett5, c.Cpcpercentnett6, c.Cpcpercentnett7, c.Cpcpercentnett8, c.Cpcpercentnett9, c.Cpcpercentnett10, c.Cpcpercentnett11, c.Cpcpercentnett12, c.Cpcpercentnett13, c.Cpcsparen1, c.Cpcsparen2, c.Cpcsparec1, c.Cpcsparec2, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cpcode to the database.
func (c *Cpcode) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cpcode.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cpcode) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cpcodes (` +
		`cpcnumber, cpccodeid, cpccode, cpcdescription, cpcdisplayorder, cpcpercentnett, cpcpercentnett2, cpcpercentnett3, cpcpercentnett4, cpcpercentnett5, cpcpercentnett6, cpcpercentnett7, cpcpercentnett8, cpcpercentnett9, cpcpercentnett10, cpcpercentnett11, cpcpercentnett12, cpcpercentnett13, cpcsparen1, cpcsparen2, cpcsparec1, cpcsparec2, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cpcnumber, cpccodeid, cpccode, cpcdescription, cpcdisplayorder, cpcpercentnett, cpcpercentnett2, cpcpercentnett3, cpcpercentnett4, cpcpercentnett5, cpcpercentnett6, cpcpercentnett7, cpcpercentnett8, cpcpercentnett9, cpcpercentnett10, cpcpercentnett11, cpcpercentnett12, cpcpercentnett13, cpcsparen1, cpcsparen2, cpcsparec1, cpcsparec2, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cpcnumber, EXCLUDED.cpccodeid, EXCLUDED.cpccode, EXCLUDED.cpcdescription, EXCLUDED.cpcdisplayorder, EXCLUDED.cpcpercentnett, EXCLUDED.cpcpercentnett2, EXCLUDED.cpcpercentnett3, EXCLUDED.cpcpercentnett4, EXCLUDED.cpcpercentnett5, EXCLUDED.cpcpercentnett6, EXCLUDED.cpcpercentnett7, EXCLUDED.cpcpercentnett8, EXCLUDED.cpcpercentnett9, EXCLUDED.cpcpercentnett10, EXCLUDED.cpcpercentnett11, EXCLUDED.cpcpercentnett12, EXCLUDED.cpcpercentnett13, EXCLUDED.cpcsparen1, EXCLUDED.cpcsparen2, EXCLUDED.cpcsparec1, EXCLUDED.cpcsparec2, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cpcnumber, c.Cpccodeid, c.Cpccode, c.Cpcdescription, c.Cpcdisplayorder, c.Cpcpercentnett, c.Cpcpercentnett2, c.Cpcpercentnett3, c.Cpcpercentnett4, c.Cpcpercentnett5, c.Cpcpercentnett6, c.Cpcpercentnett7, c.Cpcpercentnett8, c.Cpcpercentnett9, c.Cpcpercentnett10, c.Cpcpercentnett11, c.Cpcpercentnett12, c.Cpcpercentnett13, c.Cpcsparen1, c.Cpcsparen2, c.Cpcsparec1, c.Cpcsparec2, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cpcnumber, c.Cpccodeid, c.Cpccode, c.Cpcdescription, c.Cpcdisplayorder, c.Cpcpercentnett, c.Cpcpercentnett2, c.Cpcpercentnett3, c.Cpcpercentnett4, c.Cpcpercentnett5, c.Cpcpercentnett6, c.Cpcpercentnett7, c.Cpcpercentnett8, c.Cpcpercentnett9, c.Cpcpercentnett10, c.Cpcpercentnett11, c.Cpcpercentnett12, c.Cpcpercentnett13, c.Cpcsparen1, c.Cpcsparen2, c.Cpcsparec1, c.Cpcsparec2, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cpcode from the database.
func (c *Cpcode) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.cpcodes WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CpcodeByEquinoxLrn retrieves a row from 'equinox.cpcodes' as a Cpcode.
//
// Generated from index 'cpcodes_pkey'.
func CpcodeByEquinoxLrn(db XODB, equinoxLrn int64) (*Cpcode, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cpcnumber, cpccodeid, cpccode, cpcdescription, cpcdisplayorder, cpcpercentnett, cpcpercentnett2, cpcpercentnett3, cpcpercentnett4, cpcpercentnett5, cpcpercentnett6, cpcpercentnett7, cpcpercentnett8, cpcpercentnett9, cpcpercentnett10, cpcpercentnett11, cpcpercentnett12, cpcpercentnett13, cpcsparen1, cpcsparen2, cpcsparec1, cpcsparec2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cpcodes ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cpcode{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cpcnumber, &c.Cpccodeid, &c.Cpccode, &c.Cpcdescription, &c.Cpcdisplayorder, &c.Cpcpercentnett, &c.Cpcpercentnett2, &c.Cpcpercentnett3, &c.Cpcpercentnett4, &c.Cpcpercentnett5, &c.Cpcpercentnett6, &c.Cpcpercentnett7, &c.Cpcpercentnett8, &c.Cpcpercentnett9, &c.Cpcpercentnett10, &c.Cpcpercentnett11, &c.Cpcpercentnett12, &c.Cpcpercentnett13, &c.Cpcsparen1, &c.Cpcsparen2, &c.Cpcsparec1, &c.Cpcsparec2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
