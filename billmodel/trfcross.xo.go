// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Trfcross represents a row from 'equinox.trfcross'.
type Trfcross struct {
	Trfcregion   sql.NullInt64   `json:"trfcregion"`    // trfcregion
	Trfcdatefrom pq.NullTime     `json:"trfcdatefrom"`  // trfcdatefrom
	TrfclSGas    sql.NullFloat64 `json:"trfcl_s_gas"`   // trfcl_s_gas
	TrfcsHGas    sql.NullFloat64 `json:"trfcs_h_gas"`   // trfcs_h_gas
	TrfclSElec   sql.NullFloat64 `json:"trfcl_s_elec"`  // trfcl_s_elec
	TrfcsHElec   sql.NullFloat64 `json:"trfcs_h_elec"`  // trfcs_h_elec
	TrfcLSEco7   sql.NullFloat64 `json:"trfc_l_s_eco7"` // trfc_l_s_eco7
	TrfcSHEco7   sql.NullFloat64 `json:"trfc_s_h_eco7"` // trfc_s_h_eco7
	Trfsparechar sql.NullString  `json:"trfsparechar"`  // trfsparechar
	Trfsparenum  sql.NullFloat64 `json:"trfsparenum"`   // trfsparenum
	Trfcenddate  pq.NullTime     `json:"trfcenddate"`   // trfcenddate
	EquinoxLrn   int64           `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Trfcross exists in the database.
func (t *Trfcross) Exists() bool {
	return t._exists
}

// Deleted provides information if the Trfcross has been deleted from the database.
func (t *Trfcross) Deleted() bool {
	return t._deleted
}

// Insert inserts the Trfcross to the database.
func (t *Trfcross) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.trfcross (` +
		`trfcregion, trfcdatefrom, trfcl_s_gas, trfcs_h_gas, trfcl_s_elec, trfcs_h_elec, trfc_l_s_eco7, trfc_s_h_eco7, trfsparechar, trfsparenum, trfcenddate, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, t.Trfcregion, t.Trfcdatefrom, t.TrfclSGas, t.TrfcsHGas, t.TrfclSElec, t.TrfcsHElec, t.TrfcLSEco7, t.TrfcSHEco7, t.Trfsparechar, t.Trfsparenum, t.Trfcenddate, t.EquinoxSec)
	err = db.QueryRow(sqlstr, t.Trfcregion, t.Trfcdatefrom, t.TrfclSGas, t.TrfcsHGas, t.TrfclSElec, t.TrfcsHElec, t.TrfcLSEco7, t.TrfcSHEco7, t.Trfsparechar, t.Trfsparenum, t.Trfcenddate, t.EquinoxSec).Scan(&t.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Trfcross in the database.
func (t *Trfcross) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.trfcross SET (` +
		`trfcregion, trfcdatefrom, trfcl_s_gas, trfcs_h_gas, trfcl_s_elec, trfcs_h_elec, trfc_l_s_eco7, trfc_s_h_eco7, trfsparechar, trfsparenum, trfcenddate, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) WHERE equinox_lrn = $13`

	// run query
	XOLog(sqlstr, t.Trfcregion, t.Trfcdatefrom, t.TrfclSGas, t.TrfcsHGas, t.TrfclSElec, t.TrfcsHElec, t.TrfcLSEco7, t.TrfcSHEco7, t.Trfsparechar, t.Trfsparenum, t.Trfcenddate, t.EquinoxSec, t.EquinoxLrn)
	_, err = db.Exec(sqlstr, t.Trfcregion, t.Trfcdatefrom, t.TrfclSGas, t.TrfcsHGas, t.TrfclSElec, t.TrfcsHElec, t.TrfcLSEco7, t.TrfcSHEco7, t.Trfsparechar, t.Trfsparenum, t.Trfcenddate, t.EquinoxSec, t.EquinoxLrn)
	return err
}

// Save saves the Trfcross to the database.
func (t *Trfcross) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Upsert performs an upsert for Trfcross.
//
// NOTE: PostgreSQL 9.5+ only
func (t *Trfcross) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.trfcross (` +
		`trfcregion, trfcdatefrom, trfcl_s_gas, trfcs_h_gas, trfcl_s_elec, trfcs_h_elec, trfc_l_s_eco7, trfc_s_h_eco7, trfsparechar, trfsparenum, trfcenddate, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`trfcregion, trfcdatefrom, trfcl_s_gas, trfcs_h_gas, trfcl_s_elec, trfcs_h_elec, trfc_l_s_eco7, trfc_s_h_eco7, trfsparechar, trfsparenum, trfcenddate, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.trfcregion, EXCLUDED.trfcdatefrom, EXCLUDED.trfcl_s_gas, EXCLUDED.trfcs_h_gas, EXCLUDED.trfcl_s_elec, EXCLUDED.trfcs_h_elec, EXCLUDED.trfc_l_s_eco7, EXCLUDED.trfc_s_h_eco7, EXCLUDED.trfsparechar, EXCLUDED.trfsparenum, EXCLUDED.trfcenddate, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, t.Trfcregion, t.Trfcdatefrom, t.TrfclSGas, t.TrfcsHGas, t.TrfclSElec, t.TrfcsHElec, t.TrfcLSEco7, t.TrfcSHEco7, t.Trfsparechar, t.Trfsparenum, t.Trfcenddate, t.EquinoxLrn, t.EquinoxSec)
	_, err = db.Exec(sqlstr, t.Trfcregion, t.Trfcdatefrom, t.TrfclSGas, t.TrfcsHGas, t.TrfclSElec, t.TrfcsHElec, t.TrfcLSEco7, t.TrfcSHEco7, t.Trfsparechar, t.Trfsparenum, t.Trfcenddate, t.EquinoxLrn, t.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Delete deletes the Trfcross from the database.
func (t *Trfcross) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.trfcross WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, t.EquinoxLrn)
	_, err = db.Exec(sqlstr, t.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TrfcrossByEquinoxLrn retrieves a row from 'equinox.trfcross' as a Trfcross.
//
// Generated from index 'trfcross_pkey'.
func TrfcrossByEquinoxLrn(db XODB, equinoxLrn int64) (*Trfcross, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`trfcregion, trfcdatefrom, trfcl_s_gas, trfcs_h_gas, trfcl_s_elec, trfcs_h_elec, trfc_l_s_eco7, trfc_s_h_eco7, trfsparechar, trfsparenum, trfcenddate, equinox_lrn, equinox_sec ` +
		`FROM equinox.trfcross ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	t := Trfcross{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.Trfcregion, &t.Trfcdatefrom, &t.TrfclSGas, &t.TrfcsHGas, &t.TrfclSElec, &t.TrfcsHElec, &t.TrfcLSEco7, &t.TrfcSHEco7, &t.Trfsparechar, &t.Trfsparenum, &t.Trfcenddate, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}