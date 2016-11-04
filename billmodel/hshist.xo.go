// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// HsHist represents a row from 'equinox.hs_hist'.
type HsHist struct {
	HsHistDate      pq.NullTime     `json:"hs_hist_date"`     // hs_hist_date
	HsHistcost      sql.NullFloat64 `json:"hs_histcost"`      // hs_histcost
	HsHistRPC       sql.NullString  `json:"hs_hist_rpc"`      // hs_hist_rpc
	HsHistcontribut sql.NullFloat64 `json:"hs_histcontribut"` // hs_histcontribut
	HsHisthsguide   sql.NullString  `json:"hs_histhsguide"`   // hs_histhsguide
	HsHistspared1   pq.NullTime     `json:"hs_histspared1"`   // hs_histspared1
	HsHistphonetype sql.NullInt64   `json:"hs_histphonetype"` // hs_histphonetype
	HsHistgoldcont  sql.NullFloat64 `json:"hs_histgoldcont"`  // hs_histgoldcont
	HsHist24mcharge sql.NullString  `json:"hs_hist24mcharge"` // hs_hist24mcharge
	HsHist12mcharge sql.NullString  `json:"hs_hist12mcharge"` // hs_hist12mcharge
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the HsHist exists in the database.
func (hh *HsHist) Exists() bool {
	return hh._exists
}

// Deleted provides information if the HsHist has been deleted from the database.
func (hh *HsHist) Deleted() bool {
	return hh._deleted
}

// Insert inserts the HsHist to the database.
func (hh *HsHist) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if hh._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.hs_hist (` +
		`hs_hist_date, hs_histcost, hs_hist_rpc, hs_histcontribut, hs_histhsguide, hs_histspared1, hs_histphonetype, hs_histgoldcont, hs_hist24mcharge, hs_hist12mcharge, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, hh.HsHistDate, hh.HsHistcost, hh.HsHistRPC, hh.HsHistcontribut, hh.HsHisthsguide, hh.HsHistspared1, hh.HsHistphonetype, hh.HsHistgoldcont, hh.HsHist24mcharge, hh.HsHist12mcharge, hh.EquinoxPrn, hh.EquinoxSec)
	err = db.QueryRow(sqlstr, hh.HsHistDate, hh.HsHistcost, hh.HsHistRPC, hh.HsHistcontribut, hh.HsHisthsguide, hh.HsHistspared1, hh.HsHistphonetype, hh.HsHistgoldcont, hh.HsHist24mcharge, hh.HsHist12mcharge, hh.EquinoxPrn, hh.EquinoxSec).Scan(&hh.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	hh._exists = true

	return nil
}

// Update updates the HsHist in the database.
func (hh *HsHist) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !hh._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if hh._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.hs_hist SET (` +
		`hs_hist_date, hs_histcost, hs_hist_rpc, hs_histcontribut, hs_histhsguide, hs_histspared1, hs_histphonetype, hs_histgoldcont, hs_hist24mcharge, hs_hist12mcharge, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) WHERE equinox_lrn = $13`

	// run query
	XOLog(sqlstr, hh.HsHistDate, hh.HsHistcost, hh.HsHistRPC, hh.HsHistcontribut, hh.HsHisthsguide, hh.HsHistspared1, hh.HsHistphonetype, hh.HsHistgoldcont, hh.HsHist24mcharge, hh.HsHist12mcharge, hh.EquinoxPrn, hh.EquinoxSec, hh.EquinoxLrn)
	_, err = db.Exec(sqlstr, hh.HsHistDate, hh.HsHistcost, hh.HsHistRPC, hh.HsHistcontribut, hh.HsHisthsguide, hh.HsHistspared1, hh.HsHistphonetype, hh.HsHistgoldcont, hh.HsHist24mcharge, hh.HsHist12mcharge, hh.EquinoxPrn, hh.EquinoxSec, hh.EquinoxLrn)
	return err
}

// Save saves the HsHist to the database.
func (hh *HsHist) Save(db XODB) error {
	if hh.Exists() {
		return hh.Update(db)
	}

	return hh.Insert(db)
}

// Upsert performs an upsert for HsHist.
//
// NOTE: PostgreSQL 9.5+ only
func (hh *HsHist) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if hh._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.hs_hist (` +
		`hs_hist_date, hs_histcost, hs_hist_rpc, hs_histcontribut, hs_histhsguide, hs_histspared1, hs_histphonetype, hs_histgoldcont, hs_hist24mcharge, hs_hist12mcharge, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`hs_hist_date, hs_histcost, hs_hist_rpc, hs_histcontribut, hs_histhsguide, hs_histspared1, hs_histphonetype, hs_histgoldcont, hs_hist24mcharge, hs_hist12mcharge, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.hs_hist_date, EXCLUDED.hs_histcost, EXCLUDED.hs_hist_rpc, EXCLUDED.hs_histcontribut, EXCLUDED.hs_histhsguide, EXCLUDED.hs_histspared1, EXCLUDED.hs_histphonetype, EXCLUDED.hs_histgoldcont, EXCLUDED.hs_hist24mcharge, EXCLUDED.hs_hist12mcharge, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, hh.HsHistDate, hh.HsHistcost, hh.HsHistRPC, hh.HsHistcontribut, hh.HsHisthsguide, hh.HsHistspared1, hh.HsHistphonetype, hh.HsHistgoldcont, hh.HsHist24mcharge, hh.HsHist12mcharge, hh.EquinoxPrn, hh.EquinoxLrn, hh.EquinoxSec)
	_, err = db.Exec(sqlstr, hh.HsHistDate, hh.HsHistcost, hh.HsHistRPC, hh.HsHistcontribut, hh.HsHisthsguide, hh.HsHistspared1, hh.HsHistphonetype, hh.HsHistgoldcont, hh.HsHist24mcharge, hh.HsHist12mcharge, hh.EquinoxPrn, hh.EquinoxLrn, hh.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	hh._exists = true

	return nil
}

// Delete deletes the HsHist from the database.
func (hh *HsHist) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !hh._exists {
		return nil
	}

	// if deleted, bail
	if hh._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.hs_hist WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, hh.EquinoxLrn)
	_, err = db.Exec(sqlstr, hh.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	hh._deleted = true

	return nil
}

// HsHistByEquinoxLrn retrieves a row from 'equinox.hs_hist' as a HsHist.
//
// Generated from index 'hs_hist_pkey'.
func HsHistByEquinoxLrn(db XODB, equinoxLrn int64) (*HsHist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`hs_hist_date, hs_histcost, hs_hist_rpc, hs_histcontribut, hs_histhsguide, hs_histspared1, hs_histphonetype, hs_histgoldcont, hs_hist24mcharge, hs_hist12mcharge, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.hs_hist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	hh := HsHist{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&hh.HsHistDate, &hh.HsHistcost, &hh.HsHistRPC, &hh.HsHistcontribut, &hh.HsHisthsguide, &hh.HsHistspared1, &hh.HsHistphonetype, &hh.HsHistgoldcont, &hh.HsHist24mcharge, &hh.HsHist12mcharge, &hh.EquinoxPrn, &hh.EquinoxLrn, &hh.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &hh, nil
}
