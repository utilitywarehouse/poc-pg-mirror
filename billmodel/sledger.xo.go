// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Sledger represents a row from 'equinox.sledger'.
type Sledger struct {
	Slaccountno    sql.NullString  `json:"slaccountno"`    // slaccountno
	Sloutstanding  sql.NullFloat64 `json:"sloutstanding"`  // sloutstanding
	Slonledger     sql.NullFloat64 `json:"slonledger"`     // slonledger
	Slonprepay     sql.NullFloat64 `json:"slonprepay"`     // slonprepay
	Slarchivedate  pq.NullTime     `json:"slarchivedate"`  // slarchivedate
	Slcomments     sql.NullInt64   `json:"slcomments"`     // slcomments
	Sldebits       sql.NullFloat64 `json:"sldebits"`       // sldebits
	Sldebititems   sql.NullInt64   `json:"sldebititems"`   // sldebititems
	Slcredits      sql.NullFloat64 `json:"slcredits"`      // slcredits
	Slcredititems  sql.NullInt64   `json:"slcredititems"`  // slcredititems
	Sldeposits     sql.NullFloat64 `json:"sldeposits"`     // sldeposits
	Slrecalc       pq.NullTime     `json:"slrecalc"`       // slrecalc
	Slearliestitem pq.NullTime     `json:"slearliestitem"` // slearliestitem
	Sllatestitem   pq.NullTime     `json:"sllatestitem"`   // sllatestitem
	Slosinvoices   sql.NullInt64   `json:"slosinvoices"`   // slosinvoices
	Slnettp        sql.NullFloat64 `json:"slnettp"`        // slnettp
	Slvattp        sql.NullFloat64 `json:"slvattp"`        // slvattp
	Slnetgp        sql.NullFloat64 `json:"slnetgp"`        // slnetgp
	Slvatgp        sql.NullFloat64 `json:"slvatgp"`        // slvatgp
	Slnetep        sql.NullFloat64 `json:"slnetep"`        // slnetep
	Slvatep        sql.NullFloat64 `json:"slvatep"`        // slvatep
	Slnettotal     sql.NullFloat64 `json:"slnettotal"`     // slnettotal
	Slvattotal     sql.NullFloat64 `json:"slvattotal"`     // slvattotal
	Slgpover3      sql.NullFloat64 `json:"slgpover3"`      // slgpover3
	Slepover3      sql.NullFloat64 `json:"slepover3"`      // slepover3
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Sledger exists in the database.
func (s *Sledger) Exists() bool {
	return s._exists
}

// Deleted provides information if the Sledger has been deleted from the database.
func (s *Sledger) Deleted() bool {
	return s._deleted
}

// Insert inserts the Sledger to the database.
func (s *Sledger) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.sledger (` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, s.Slaccountno, s.Sloutstanding, s.Slonledger, s.Slonprepay, s.Slarchivedate, s.Slcomments, s.Sldebits, s.Sldebititems, s.Slcredits, s.Slcredititems, s.Sldeposits, s.Slrecalc, s.Slearliestitem, s.Sllatestitem, s.Slosinvoices, s.Slnettp, s.Slvattp, s.Slnetgp, s.Slvatgp, s.Slnetep, s.Slvatep, s.Slnettotal, s.Slvattotal, s.Slgpover3, s.Slepover3, s.EquinoxSec)
	err = db.QueryRow(sqlstr, s.Slaccountno, s.Sloutstanding, s.Slonledger, s.Slonprepay, s.Slarchivedate, s.Slcomments, s.Sldebits, s.Sldebititems, s.Slcredits, s.Slcredititems, s.Sldeposits, s.Slrecalc, s.Slearliestitem, s.Sllatestitem, s.Slosinvoices, s.Slnettp, s.Slvattp, s.Slnetgp, s.Slvatgp, s.Slnetep, s.Slvatep, s.Slnettotal, s.Slvattotal, s.Slgpover3, s.Slepover3, s.EquinoxSec).Scan(&s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the Sledger in the database.
func (s *Sledger) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.sledger SET (` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26` +
		`) WHERE equinox_lrn = $27`

	// run query
	XOLog(sqlstr, s.Slaccountno, s.Sloutstanding, s.Slonledger, s.Slonprepay, s.Slarchivedate, s.Slcomments, s.Sldebits, s.Sldebititems, s.Slcredits, s.Slcredititems, s.Sldeposits, s.Slrecalc, s.Slearliestitem, s.Sllatestitem, s.Slosinvoices, s.Slnettp, s.Slvattp, s.Slnetgp, s.Slvatgp, s.Slnetep, s.Slvatep, s.Slnettotal, s.Slvattotal, s.Slgpover3, s.Slepover3, s.EquinoxSec, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.Slaccountno, s.Sloutstanding, s.Slonledger, s.Slonprepay, s.Slarchivedate, s.Slcomments, s.Sldebits, s.Sldebititems, s.Slcredits, s.Slcredititems, s.Sldeposits, s.Slrecalc, s.Slearliestitem, s.Sllatestitem, s.Slosinvoices, s.Slnettp, s.Slvattp, s.Slnetgp, s.Slvatgp, s.Slnetep, s.Slvatep, s.Slnettotal, s.Slvattotal, s.Slgpover3, s.Slepover3, s.EquinoxSec, s.EquinoxLrn)
	return err
}

// Save saves the Sledger to the database.
func (s *Sledger) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for Sledger.
//
// NOTE: PostgreSQL 9.5+ only
func (s *Sledger) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.sledger (` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.slaccountno, EXCLUDED.sloutstanding, EXCLUDED.slonledger, EXCLUDED.slonprepay, EXCLUDED.slarchivedate, EXCLUDED.slcomments, EXCLUDED.sldebits, EXCLUDED.sldebititems, EXCLUDED.slcredits, EXCLUDED.slcredititems, EXCLUDED.sldeposits, EXCLUDED.slrecalc, EXCLUDED.slearliestitem, EXCLUDED.sllatestitem, EXCLUDED.slosinvoices, EXCLUDED.slnettp, EXCLUDED.slvattp, EXCLUDED.slnetgp, EXCLUDED.slvatgp, EXCLUDED.slnetep, EXCLUDED.slvatep, EXCLUDED.slnettotal, EXCLUDED.slvattotal, EXCLUDED.slgpover3, EXCLUDED.slepover3, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, s.Slaccountno, s.Sloutstanding, s.Slonledger, s.Slonprepay, s.Slarchivedate, s.Slcomments, s.Sldebits, s.Sldebititems, s.Slcredits, s.Slcredititems, s.Sldeposits, s.Slrecalc, s.Slearliestitem, s.Sllatestitem, s.Slosinvoices, s.Slnettp, s.Slvattp, s.Slnetgp, s.Slvatgp, s.Slnetep, s.Slvatep, s.Slnettotal, s.Slvattotal, s.Slgpover3, s.Slepover3, s.EquinoxLrn, s.EquinoxSec)
	_, err = db.Exec(sqlstr, s.Slaccountno, s.Sloutstanding, s.Slonledger, s.Slonprepay, s.Slarchivedate, s.Slcomments, s.Sldebits, s.Sldebititems, s.Slcredits, s.Slcredititems, s.Sldeposits, s.Slrecalc, s.Slearliestitem, s.Sllatestitem, s.Slosinvoices, s.Slnettp, s.Slvattp, s.Slnetgp, s.Slvatgp, s.Slnetep, s.Slvatep, s.Slnettotal, s.Slvattotal, s.Slgpover3, s.Slepover3, s.EquinoxLrn, s.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the Sledger from the database.
func (s *Sledger) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.sledger WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

// SledgerByEquinoxLrn retrieves a row from 'equinox.sledger' as a Sledger.
//
// Generated from index 'sledger_pkey'.
func SledgerByEquinoxLrn(db XODB, equinoxLrn int64) (*Sledger, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_lrn, equinox_sec ` +
		`FROM equinox.sledger ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Sledger{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Slaccountno, &s.Sloutstanding, &s.Slonledger, &s.Slonprepay, &s.Slarchivedate, &s.Slcomments, &s.Sldebits, &s.Sldebititems, &s.Slcredits, &s.Slcredititems, &s.Sldeposits, &s.Slrecalc, &s.Slearliestitem, &s.Sllatestitem, &s.Slosinvoices, &s.Slnettp, &s.Slvattp, &s.Slnetgp, &s.Slvatgp, &s.Slnetep, &s.Slvatep, &s.Slnettotal, &s.Slvattotal, &s.Slgpover3, &s.Slepover3, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
