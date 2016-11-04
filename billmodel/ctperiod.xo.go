// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Ctperiod represents a row from 'equinox.ctperiod'.
type Ctperiod struct {
	Ctdate        pq.NullTime   `json:"ctdate"`        // ctdate
	Ctlgasofgem   sql.NullInt64 `json:"ctlgasofgem"`   // ctlgasofgem
	Ctsgasofgem   sql.NullInt64 `json:"ctsgasofgem"`   // ctsgasofgem
	Cthgasofgem   sql.NullInt64 `json:"cthgasofgem"`   // cthgasofgem
	Ctlelecofgem  sql.NullInt64 `json:"ctlelecofgem"`  // ctlelecofgem
	Ctselecofgem  sql.NullInt64 `json:"ctselecofgem"`  // ctselecofgem
	Cthelecofgem  sql.NullInt64 `json:"cthelecofgem"`  // cthelecofgem
	Ctle7ofgem    sql.NullInt64 `json:"ctle7ofgem"`    // ctle7ofgem
	Ctse7ofgem    sql.NullInt64 `json:"ctse7ofgem"`    // ctse7ofgem
	Cthe7ofgem    sql.NullInt64 `json:"cthe7ofgem"`    // cthe7ofgem
	Ctuwgofftake  sql.NullInt64 `json:"ctuwgofftake"`  // ctuwgofftake
	Ctuweofftake  sql.NullInt64 `json:"ctuweofftake"`  // ctuweofftake
	Ctuwe7offtake sql.NullInt64 `json:"ctuwe7offtake"` // ctuwe7offtake
	EquinoxLrn    int64         `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64 `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ctperiod exists in the database.
func (c *Ctperiod) Exists() bool {
	return c._exists
}

// Deleted provides information if the Ctperiod has been deleted from the database.
func (c *Ctperiod) Deleted() bool {
	return c._deleted
}

// Insert inserts the Ctperiod to the database.
func (c *Ctperiod) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ctperiod (` +
		`ctdate, ctlgasofgem, ctsgasofgem, cthgasofgem, ctlelecofgem, ctselecofgem, cthelecofgem, ctle7ofgem, ctse7ofgem, cthe7ofgem, ctuwgofftake, ctuweofftake, ctuwe7offtake, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Ctdate, c.Ctlgasofgem, c.Ctsgasofgem, c.Cthgasofgem, c.Ctlelecofgem, c.Ctselecofgem, c.Cthelecofgem, c.Ctle7ofgem, c.Ctse7ofgem, c.Cthe7ofgem, c.Ctuwgofftake, c.Ctuweofftake, c.Ctuwe7offtake, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Ctdate, c.Ctlgasofgem, c.Ctsgasofgem, c.Cthgasofgem, c.Ctlelecofgem, c.Ctselecofgem, c.Cthelecofgem, c.Ctle7ofgem, c.Ctse7ofgem, c.Cthe7ofgem, c.Ctuwgofftake, c.Ctuweofftake, c.Ctuwe7offtake, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Ctperiod in the database.
func (c *Ctperiod) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.ctperiod SET (` +
		`ctdate, ctlgasofgem, ctsgasofgem, cthgasofgem, ctlelecofgem, ctselecofgem, cthelecofgem, ctle7ofgem, ctse7ofgem, cthe7ofgem, ctuwgofftake, ctuweofftake, ctuwe7offtake, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`) WHERE equinox_lrn = $15`

	// run query
	XOLog(sqlstr, c.Ctdate, c.Ctlgasofgem, c.Ctsgasofgem, c.Cthgasofgem, c.Ctlelecofgem, c.Ctselecofgem, c.Cthelecofgem, c.Ctle7ofgem, c.Ctse7ofgem, c.Cthe7ofgem, c.Ctuwgofftake, c.Ctuweofftake, c.Ctuwe7offtake, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Ctdate, c.Ctlgasofgem, c.Ctsgasofgem, c.Cthgasofgem, c.Ctlelecofgem, c.Ctselecofgem, c.Cthelecofgem, c.Ctle7ofgem, c.Ctse7ofgem, c.Cthe7ofgem, c.Ctuwgofftake, c.Ctuweofftake, c.Ctuwe7offtake, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Ctperiod to the database.
func (c *Ctperiod) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Ctperiod.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Ctperiod) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ctperiod (` +
		`ctdate, ctlgasofgem, ctsgasofgem, cthgasofgem, ctlelecofgem, ctselecofgem, cthelecofgem, ctle7ofgem, ctse7ofgem, cthe7ofgem, ctuwgofftake, ctuweofftake, ctuwe7offtake, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`ctdate, ctlgasofgem, ctsgasofgem, cthgasofgem, ctlelecofgem, ctselecofgem, cthelecofgem, ctle7ofgem, ctse7ofgem, cthe7ofgem, ctuwgofftake, ctuweofftake, ctuwe7offtake, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.ctdate, EXCLUDED.ctlgasofgem, EXCLUDED.ctsgasofgem, EXCLUDED.cthgasofgem, EXCLUDED.ctlelecofgem, EXCLUDED.ctselecofgem, EXCLUDED.cthelecofgem, EXCLUDED.ctle7ofgem, EXCLUDED.ctse7ofgem, EXCLUDED.cthe7ofgem, EXCLUDED.ctuwgofftake, EXCLUDED.ctuweofftake, EXCLUDED.ctuwe7offtake, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Ctdate, c.Ctlgasofgem, c.Ctsgasofgem, c.Cthgasofgem, c.Ctlelecofgem, c.Ctselecofgem, c.Cthelecofgem, c.Ctle7ofgem, c.Ctse7ofgem, c.Cthe7ofgem, c.Ctuwgofftake, c.Ctuweofftake, c.Ctuwe7offtake, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Ctdate, c.Ctlgasofgem, c.Ctsgasofgem, c.Cthgasofgem, c.Ctlelecofgem, c.Ctselecofgem, c.Cthelecofgem, c.Ctle7ofgem, c.Ctse7ofgem, c.Cthe7ofgem, c.Ctuwgofftake, c.Ctuweofftake, c.Ctuwe7offtake, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Ctperiod from the database.
func (c *Ctperiod) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.ctperiod WHERE equinox_lrn = $1`

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

// CtperiodByEquinoxLrn retrieves a row from 'equinox.ctperiod' as a Ctperiod.
//
// Generated from index 'ctperiod_pkey'.
func CtperiodByEquinoxLrn(db XODB, equinoxLrn int64) (*Ctperiod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ctdate, ctlgasofgem, ctsgasofgem, cthgasofgem, ctlelecofgem, ctselecofgem, cthelecofgem, ctle7ofgem, ctse7ofgem, cthe7ofgem, ctuwgofftake, ctuweofftake, ctuwe7offtake, equinox_lrn, equinox_sec ` +
		`FROM equinox.ctperiod ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Ctperiod{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Ctdate, &c.Ctlgasofgem, &c.Ctsgasofgem, &c.Cthgasofgem, &c.Ctlelecofgem, &c.Ctselecofgem, &c.Cthelecofgem, &c.Ctle7ofgem, &c.Ctse7ofgem, &c.Cthe7ofgem, &c.Ctuwgofftake, &c.Ctuweofftake, &c.Ctuwe7offtake, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
