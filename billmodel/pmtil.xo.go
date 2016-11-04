// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Pmtil represents a row from 'equinox.pmtil'.
type Pmtil struct {
	Pmtid        sql.NullString `json:"pmtid"`        // pmtid
	Pmttariff    sql.NullString `json:"pmttariff"`    // pmttariff
	Pmtregion    sql.NullString `json:"pmtregion"`    // pmtregion
	Pmtfuel      sql.NullString `json:"pmtfuel"`      // pmtfuel
	Pmtpaymethod sql.NullString `json:"pmtpaymethod"` // pmtpaymethod
	Pmtentry1    sql.NullInt64  `json:"pmtentry1"`    // pmtentry1
	Pmtentry2    sql.NullInt64  `json:"pmtentry2"`    // pmtentry2
	Pmtentry3    sql.NullInt64  `json:"pmtentry3"`    // pmtentry3
	Pmtentry4    sql.NullInt64  `json:"pmtentry4"`    // pmtentry4
	Pmtentry5    sql.NullInt64  `json:"pmtentry5"`    // pmtentry5
	Pmtentry6    sql.NullInt64  `json:"pmtentry6"`    // pmtentry6
	Pmtentry7    sql.NullInt64  `json:"pmtentry7"`    // pmtentry7
	Pmtentry8    sql.NullInt64  `json:"pmtentry8"`    // pmtentry8
	Pmtentry9    sql.NullInt64  `json:"pmtentry9"`    // pmtentry9
	Pmtentry10   sql.NullInt64  `json:"pmtentry10"`   // pmtentry10
	Pmtentry11   sql.NullInt64  `json:"pmtentry11"`   // pmtentry11
	Pmtentry12   sql.NullInt64  `json:"pmtentry12"`   // pmtentry12
	Pmtentry13   sql.NullInt64  `json:"pmtentry13"`   // pmtentry13
	Pmtentry14   sql.NullInt64  `json:"pmtentry14"`   // pmtentry14
	Pmtentry15   sql.NullInt64  `json:"pmtentry15"`   // pmtentry15
	Pmtentry16   sql.NullInt64  `json:"pmtentry16"`   // pmtentry16
	Pmtentry17   sql.NullInt64  `json:"pmtentry17"`   // pmtentry17
	Pmtentry18   sql.NullInt64  `json:"pmtentry18"`   // pmtentry18
	Pmtentry19   sql.NullInt64  `json:"pmtentry19"`   // pmtentry19
	Pmtentry20   sql.NullInt64  `json:"pmtentry20"`   // pmtentry20
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Pmtil exists in the database.
func (p *Pmtil) Exists() bool {
	return p._exists
}

// Deleted provides information if the Pmtil has been deleted from the database.
func (p *Pmtil) Deleted() bool {
	return p._deleted
}

// Insert inserts the Pmtil to the database.
func (p *Pmtil) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pmtil (` +
		`pmtid, pmttariff, pmtregion, pmtfuel, pmtpaymethod, pmtentry1, pmtentry2, pmtentry3, pmtentry4, pmtentry5, pmtentry6, pmtentry7, pmtentry8, pmtentry9, pmtentry10, pmtentry11, pmtentry12, pmtentry13, pmtentry14, pmtentry15, pmtentry16, pmtentry17, pmtentry18, pmtentry19, pmtentry20, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Pmtid, p.Pmttariff, p.Pmtregion, p.Pmtfuel, p.Pmtpaymethod, p.Pmtentry1, p.Pmtentry2, p.Pmtentry3, p.Pmtentry4, p.Pmtentry5, p.Pmtentry6, p.Pmtentry7, p.Pmtentry8, p.Pmtentry9, p.Pmtentry10, p.Pmtentry11, p.Pmtentry12, p.Pmtentry13, p.Pmtentry14, p.Pmtentry15, p.Pmtentry16, p.Pmtentry17, p.Pmtentry18, p.Pmtentry19, p.Pmtentry20, p.EquinoxPrn, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Pmtid, p.Pmttariff, p.Pmtregion, p.Pmtfuel, p.Pmtpaymethod, p.Pmtentry1, p.Pmtentry2, p.Pmtentry3, p.Pmtentry4, p.Pmtentry5, p.Pmtentry6, p.Pmtentry7, p.Pmtentry8, p.Pmtentry9, p.Pmtentry10, p.Pmtentry11, p.Pmtentry12, p.Pmtentry13, p.Pmtentry14, p.Pmtentry15, p.Pmtentry16, p.Pmtentry17, p.Pmtentry18, p.Pmtentry19, p.Pmtentry20, p.EquinoxPrn, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Pmtil in the database.
func (p *Pmtil) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.pmtil SET (` +
		`pmtid, pmttariff, pmtregion, pmtfuel, pmtpaymethod, pmtentry1, pmtentry2, pmtentry3, pmtentry4, pmtentry5, pmtentry6, pmtentry7, pmtentry8, pmtentry9, pmtentry10, pmtentry11, pmtentry12, pmtentry13, pmtentry14, pmtentry15, pmtentry16, pmtentry17, pmtentry18, pmtentry19, pmtentry20, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27` +
		`) WHERE equinox_lrn = $28`

	// run query
	XOLog(sqlstr, p.Pmtid, p.Pmttariff, p.Pmtregion, p.Pmtfuel, p.Pmtpaymethod, p.Pmtentry1, p.Pmtentry2, p.Pmtentry3, p.Pmtentry4, p.Pmtentry5, p.Pmtentry6, p.Pmtentry7, p.Pmtentry8, p.Pmtentry9, p.Pmtentry10, p.Pmtentry11, p.Pmtentry12, p.Pmtentry13, p.Pmtentry14, p.Pmtentry15, p.Pmtentry16, p.Pmtentry17, p.Pmtentry18, p.Pmtentry19, p.Pmtentry20, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Pmtid, p.Pmttariff, p.Pmtregion, p.Pmtfuel, p.Pmtpaymethod, p.Pmtentry1, p.Pmtentry2, p.Pmtentry3, p.Pmtentry4, p.Pmtentry5, p.Pmtentry6, p.Pmtentry7, p.Pmtentry8, p.Pmtentry9, p.Pmtentry10, p.Pmtentry11, p.Pmtentry12, p.Pmtentry13, p.Pmtentry14, p.Pmtentry15, p.Pmtentry16, p.Pmtentry17, p.Pmtentry18, p.Pmtentry19, p.Pmtentry20, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Pmtil to the database.
func (p *Pmtil) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Pmtil.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Pmtil) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.pmtil (` +
		`pmtid, pmttariff, pmtregion, pmtfuel, pmtpaymethod, pmtentry1, pmtentry2, pmtentry3, pmtentry4, pmtentry5, pmtentry6, pmtentry7, pmtentry8, pmtentry9, pmtentry10, pmtentry11, pmtentry12, pmtentry13, pmtentry14, pmtentry15, pmtentry16, pmtentry17, pmtentry18, pmtentry19, pmtentry20, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`pmtid, pmttariff, pmtregion, pmtfuel, pmtpaymethod, pmtentry1, pmtentry2, pmtentry3, pmtentry4, pmtentry5, pmtentry6, pmtentry7, pmtentry8, pmtentry9, pmtentry10, pmtentry11, pmtentry12, pmtentry13, pmtentry14, pmtentry15, pmtentry16, pmtentry17, pmtentry18, pmtentry19, pmtentry20, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.pmtid, EXCLUDED.pmttariff, EXCLUDED.pmtregion, EXCLUDED.pmtfuel, EXCLUDED.pmtpaymethod, EXCLUDED.pmtentry1, EXCLUDED.pmtentry2, EXCLUDED.pmtentry3, EXCLUDED.pmtentry4, EXCLUDED.pmtentry5, EXCLUDED.pmtentry6, EXCLUDED.pmtentry7, EXCLUDED.pmtentry8, EXCLUDED.pmtentry9, EXCLUDED.pmtentry10, EXCLUDED.pmtentry11, EXCLUDED.pmtentry12, EXCLUDED.pmtentry13, EXCLUDED.pmtentry14, EXCLUDED.pmtentry15, EXCLUDED.pmtentry16, EXCLUDED.pmtentry17, EXCLUDED.pmtentry18, EXCLUDED.pmtentry19, EXCLUDED.pmtentry20, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Pmtid, p.Pmttariff, p.Pmtregion, p.Pmtfuel, p.Pmtpaymethod, p.Pmtentry1, p.Pmtentry2, p.Pmtentry3, p.Pmtentry4, p.Pmtentry5, p.Pmtentry6, p.Pmtentry7, p.Pmtentry8, p.Pmtentry9, p.Pmtentry10, p.Pmtentry11, p.Pmtentry12, p.Pmtentry13, p.Pmtentry14, p.Pmtentry15, p.Pmtentry16, p.Pmtentry17, p.Pmtentry18, p.Pmtentry19, p.Pmtentry20, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Pmtid, p.Pmttariff, p.Pmtregion, p.Pmtfuel, p.Pmtpaymethod, p.Pmtentry1, p.Pmtentry2, p.Pmtentry3, p.Pmtentry4, p.Pmtentry5, p.Pmtentry6, p.Pmtentry7, p.Pmtentry8, p.Pmtentry9, p.Pmtentry10, p.Pmtentry11, p.Pmtentry12, p.Pmtentry13, p.Pmtentry14, p.Pmtentry15, p.Pmtentry16, p.Pmtentry17, p.Pmtentry18, p.Pmtentry19, p.Pmtentry20, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Pmtil from the database.
func (p *Pmtil) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.pmtil WHERE equinox_lrn = $1`

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

// PmtilByEquinoxLrn retrieves a row from 'equinox.pmtil' as a Pmtil.
//
// Generated from index 'pmtil_pkey'.
func PmtilByEquinoxLrn(db XODB, equinoxLrn int64) (*Pmtil, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pmtid, pmttariff, pmtregion, pmtfuel, pmtpaymethod, pmtentry1, pmtentry2, pmtentry3, pmtentry4, pmtentry5, pmtentry6, pmtentry7, pmtentry8, pmtentry9, pmtentry10, pmtentry11, pmtentry12, pmtentry13, pmtentry14, pmtentry15, pmtentry16, pmtentry17, pmtentry18, pmtentry19, pmtentry20, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.pmtil ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Pmtil{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pmtid, &p.Pmttariff, &p.Pmtregion, &p.Pmtfuel, &p.Pmtpaymethod, &p.Pmtentry1, &p.Pmtentry2, &p.Pmtentry3, &p.Pmtentry4, &p.Pmtentry5, &p.Pmtentry6, &p.Pmtentry7, &p.Pmtentry8, &p.Pmtentry9, &p.Pmtentry10, &p.Pmtentry11, &p.Pmtentry12, &p.Pmtentry13, &p.Pmtentry14, &p.Pmtentry15, &p.Pmtentry16, &p.Pmtentry17, &p.Pmtentry18, &p.Pmtentry19, &p.Pmtentry20, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}