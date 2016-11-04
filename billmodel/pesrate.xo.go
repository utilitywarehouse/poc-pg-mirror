// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Pesrate represents a row from 'equinox.pesrates'.
type Pesrate struct {
	Pesregion      sql.NullString  `json:"pesregion"`      // pesregion
	Pesmeter       sql.NullString  `json:"pesmeter"`       // pesmeter
	Pestariffmeter sql.NullString  `json:"pestariffmeter"` // pestariffmeter
	Pesppkwh       sql.NullFloat64 `json:"pesppkwh"`       // pesppkwh
	Pesmonthly     sql.NullFloat64 `json:"pesmonthly"`     // pesmonthly
	Pesdaily       sql.NullFloat64 `json:"pesdaily"`       // pesdaily
	Pesmonthorday  sql.NullString  `json:"pesmonthorday"`  // pesmonthorday
	Pesvalidfrom   pq.NullTime     `json:"pesvalidfrom"`   // pesvalidfrom
	Pesvalidto     pq.NullTime     `json:"pesvalidto"`     // pesvalidto
	Pestcr         sql.NullFloat64 `json:"pestcr"`         // pestcr
	Pestcrdual     sql.NullFloat64 `json:"pestcrdual"`     // pestcrdual
	EquinoxPrn     sql.NullInt64   `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Pesrate exists in the database.
func (p *Pesrate) Exists() bool {
	return p._exists
}

// Deleted provides information if the Pesrate has been deleted from the database.
func (p *Pesrate) Deleted() bool {
	return p._deleted
}

// Insert inserts the Pesrate to the database.
func (p *Pesrate) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pesrates (` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Pesregion, p.Pesmeter, p.Pestariffmeter, p.Pesppkwh, p.Pesmonthly, p.Pesdaily, p.Pesmonthorday, p.Pesvalidfrom, p.Pesvalidto, p.Pestcr, p.Pestcrdual, p.EquinoxPrn, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Pesregion, p.Pesmeter, p.Pestariffmeter, p.Pesppkwh, p.Pesmonthly, p.Pesdaily, p.Pesmonthorday, p.Pesvalidfrom, p.Pesvalidto, p.Pestcr, p.Pestcrdual, p.EquinoxPrn, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Pesrate in the database.
func (p *Pesrate) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.pesrates SET (` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) WHERE equinox_lrn = $14`

	// run query
	XOLog(sqlstr, p.Pesregion, p.Pesmeter, p.Pestariffmeter, p.Pesppkwh, p.Pesmonthly, p.Pesdaily, p.Pesmonthorday, p.Pesvalidfrom, p.Pesvalidto, p.Pestcr, p.Pestcrdual, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Pesregion, p.Pesmeter, p.Pestariffmeter, p.Pesppkwh, p.Pesmonthly, p.Pesdaily, p.Pesmonthorday, p.Pesvalidfrom, p.Pesvalidto, p.Pestcr, p.Pestcrdual, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Pesrate to the database.
func (p *Pesrate) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Pesrate.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Pesrate) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pesrates (` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.pesregion, EXCLUDED.pesmeter, EXCLUDED.pestariffmeter, EXCLUDED.pesppkwh, EXCLUDED.pesmonthly, EXCLUDED.pesdaily, EXCLUDED.pesmonthorday, EXCLUDED.pesvalidfrom, EXCLUDED.pesvalidto, EXCLUDED.pestcr, EXCLUDED.pestcrdual, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Pesregion, p.Pesmeter, p.Pestariffmeter, p.Pesppkwh, p.Pesmonthly, p.Pesdaily, p.Pesmonthorday, p.Pesvalidfrom, p.Pesvalidto, p.Pestcr, p.Pestcrdual, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Pesregion, p.Pesmeter, p.Pestariffmeter, p.Pesppkwh, p.Pesmonthly, p.Pesdaily, p.Pesmonthorday, p.Pesvalidfrom, p.Pesvalidto, p.Pestcr, p.Pestcrdual, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Pesrate from the database.
func (p *Pesrate) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.pesrates WHERE equinox_lrn = $1`

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

// PesrateByEquinoxLrn retrieves a row from 'equinox.pesrates' as a Pesrate.
//
// Generated from index 'pesrates_pkey'.
func PesrateByEquinoxLrn(db XODB, equinoxLrn int64) (*Pesrate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pesregion, pesmeter, pestariffmeter, pesppkwh, pesmonthly, pesdaily, pesmonthorday, pesvalidfrom, pesvalidto, pestcr, pestcrdual, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pesrates ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pesrate{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pesregion, &p.Pesmeter, &p.Pestariffmeter, &p.Pesppkwh, &p.Pesmonthly, &p.Pesdaily, &p.Pesmonthorday, &p.Pesvalidfrom, &p.Pesvalidto, &p.Pestcr, &p.Pestcrdual, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
