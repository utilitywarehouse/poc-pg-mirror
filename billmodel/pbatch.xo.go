// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Pbatch represents a row from 'equinox.pbatches'.
type Pbatch struct {
	Pbbatchnumber  sql.NullString  `json:"pbbatchnumber"`  // pbbatchnumber
	Pbbatchdate    pq.NullTime     `json:"pbbatchdate"`    // pbbatchdate
	Pbbatchcode    sql.NullString  `json:"pbbatchcode"`    // pbbatchcode
	Pbcommitdate   pq.NullTime     `json:"pbcommitdate"`   // pbcommitdate
	Pbcommittime   pq.NullTime     `json:"pbcommittime"`   // pbcommittime
	Pbbatchtotal   sql.NullFloat64 `json:"pbbatchtotal"`   // pbbatchtotal
	Pbinputtotal   sql.NullFloat64 `json:"pbinputtotal"`   // pbinputtotal
	Pbinputrecords sql.NullInt64   `json:"pbinputrecords"` // pbinputrecords
	Pbnotes        sql.NullString  `json:"pbnotes"`        // pbnotes
	Pbpaymentdate  pq.NullTime     `json:"pbpaymentdate"`  // pbpaymentdate
	Pbbatchgroup   sql.NullString  `json:"pbbatchgroup"`   // pbbatchgroup
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Pbatch exists in the database.
func (p *Pbatch) Exists() bool {
	return p._exists
}

// Deleted provides information if the Pbatch has been deleted from the database.
func (p *Pbatch) Deleted() bool {
	return p._deleted
}

// Insert inserts the Pbatch to the database.
func (p *Pbatch) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pbatches (` +
		`pbbatchnumber, pbbatchdate, pbbatchcode, pbcommitdate, pbcommittime, pbbatchtotal, pbinputtotal, pbinputrecords, pbnotes, pbpaymentdate, pbbatchgroup, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Pbbatchnumber, p.Pbbatchdate, p.Pbbatchcode, p.Pbcommitdate, p.Pbcommittime, p.Pbbatchtotal, p.Pbinputtotal, p.Pbinputrecords, p.Pbnotes, p.Pbpaymentdate, p.Pbbatchgroup, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Pbbatchnumber, p.Pbbatchdate, p.Pbbatchcode, p.Pbcommitdate, p.Pbcommittime, p.Pbbatchtotal, p.Pbinputtotal, p.Pbinputrecords, p.Pbnotes, p.Pbpaymentdate, p.Pbbatchgroup, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Pbatch in the database.
func (p *Pbatch) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.pbatches SET (` +
		`pbbatchnumber, pbbatchdate, pbbatchcode, pbcommitdate, pbcommittime, pbbatchtotal, pbinputtotal, pbinputrecords, pbnotes, pbpaymentdate, pbbatchgroup, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) WHERE equinox_lrn = $13`

	// run query
	XOLog(sqlstr, p.Pbbatchnumber, p.Pbbatchdate, p.Pbbatchcode, p.Pbcommitdate, p.Pbcommittime, p.Pbbatchtotal, p.Pbinputtotal, p.Pbinputrecords, p.Pbnotes, p.Pbpaymentdate, p.Pbbatchgroup, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Pbbatchnumber, p.Pbbatchdate, p.Pbbatchcode, p.Pbcommitdate, p.Pbcommittime, p.Pbbatchtotal, p.Pbinputtotal, p.Pbinputrecords, p.Pbnotes, p.Pbpaymentdate, p.Pbbatchgroup, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Pbatch to the database.
func (p *Pbatch) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Pbatch.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Pbatch) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pbatches (` +
		`pbbatchnumber, pbbatchdate, pbbatchcode, pbcommitdate, pbcommittime, pbbatchtotal, pbinputtotal, pbinputrecords, pbnotes, pbpaymentdate, pbbatchgroup, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`pbbatchnumber, pbbatchdate, pbbatchcode, pbcommitdate, pbcommittime, pbbatchtotal, pbinputtotal, pbinputrecords, pbnotes, pbpaymentdate, pbbatchgroup, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.pbbatchnumber, EXCLUDED.pbbatchdate, EXCLUDED.pbbatchcode, EXCLUDED.pbcommitdate, EXCLUDED.pbcommittime, EXCLUDED.pbbatchtotal, EXCLUDED.pbinputtotal, EXCLUDED.pbinputrecords, EXCLUDED.pbnotes, EXCLUDED.pbpaymentdate, EXCLUDED.pbbatchgroup, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Pbbatchnumber, p.Pbbatchdate, p.Pbbatchcode, p.Pbcommitdate, p.Pbcommittime, p.Pbbatchtotal, p.Pbinputtotal, p.Pbinputrecords, p.Pbnotes, p.Pbpaymentdate, p.Pbbatchgroup, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Pbbatchnumber, p.Pbbatchdate, p.Pbbatchcode, p.Pbcommitdate, p.Pbcommittime, p.Pbbatchtotal, p.Pbinputtotal, p.Pbinputrecords, p.Pbnotes, p.Pbpaymentdate, p.Pbbatchgroup, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Pbatch from the database.
func (p *Pbatch) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.pbatches WHERE equinox_lrn = $1`

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

// PbatchByEquinoxLrn retrieves a row from 'equinox.pbatches' as a Pbatch.
//
// Generated from index 'pbatches_pkey'.
func PbatchByEquinoxLrn(db XODB, equinoxLrn int64) (*Pbatch, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pbbatchnumber, pbbatchdate, pbbatchcode, pbcommitdate, pbcommittime, pbbatchtotal, pbinputtotal, pbinputrecords, pbnotes, pbpaymentdate, pbbatchgroup, equinox_lrn, equinox_sec ` +
		`FROM equinox.pbatches ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pbatch{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pbbatchnumber, &p.Pbbatchdate, &p.Pbbatchcode, &p.Pbcommitdate, &p.Pbcommittime, &p.Pbbatchtotal, &p.Pbinputtotal, &p.Pbinputrecords, &p.Pbnotes, &p.Pbpaymentdate, &p.Pbbatchgroup, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
