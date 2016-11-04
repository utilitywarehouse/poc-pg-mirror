// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// College represents a row from 'equinox.college'.
type College struct {
	CollID       sql.NullInt64  `json:"coll_id"`       // coll_id
	CollName     sql.NullString `json:"coll_name"`     // coll_name
	CollAddr1    sql.NullString `json:"coll_addr1"`    // coll_addr1
	CollAddr2    sql.NullString `json:"coll_addr2"`    // coll_addr2
	CollAddr3    sql.NullString `json:"coll_addr3"`    // coll_addr3
	CollAddr4    sql.NullString `json:"coll_addr4"`    // coll_addr4
	CollAddr5    sql.NullString `json:"coll_addr5"`    // coll_addr5
	CollPcode    sql.NullString `json:"coll_pcode"`    // coll_pcode
	CollTel1     sql.NullString `json:"coll_tel1"`     // coll_tel1
	CollTel2     sql.NullString `json:"coll_tel2"`     // coll_tel2
	CollFax      sql.NullString `json:"coll_fax"`      // coll_fax
	CollCapacity sql.NullInt64  `json:"coll_capacity"` // coll_capacity
	CollTrainer  sql.NullString `json:"coll_trainer"`  // coll_trainer
	CollTime1    pq.NullTime    `json:"coll_time1"`    // coll_time1
	CollTime2    pq.NullTime    `json:"coll_time2"`    // coll_time2
	EquinoxLrn   int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the College exists in the database.
func (c *College) Exists() bool {
	return c._exists
}

// Deleted provides information if the College has been deleted from the database.
func (c *College) Deleted() bool {
	return c._deleted
}

// Insert inserts the College to the database.
func (c *College) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.college (` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.CollID, c.CollName, c.CollAddr1, c.CollAddr2, c.CollAddr3, c.CollAddr4, c.CollAddr5, c.CollPcode, c.CollTel1, c.CollTel2, c.CollFax, c.CollCapacity, c.CollTrainer, c.CollTime1, c.CollTime2, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.CollID, c.CollName, c.CollAddr1, c.CollAddr2, c.CollAddr3, c.CollAddr4, c.CollAddr5, c.CollPcode, c.CollTel1, c.CollTel2, c.CollFax, c.CollCapacity, c.CollTrainer, c.CollTime1, c.CollTime2, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the College in the database.
func (c *College) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.college SET (` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) WHERE equinox_lrn = $17`

	// run query
	XOLog(sqlstr, c.CollID, c.CollName, c.CollAddr1, c.CollAddr2, c.CollAddr3, c.CollAddr4, c.CollAddr5, c.CollPcode, c.CollTel1, c.CollTel2, c.CollFax, c.CollCapacity, c.CollTrainer, c.CollTime1, c.CollTime2, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.CollID, c.CollName, c.CollAddr1, c.CollAddr2, c.CollAddr3, c.CollAddr4, c.CollAddr5, c.CollPcode, c.CollTel1, c.CollTel2, c.CollFax, c.CollCapacity, c.CollTrainer, c.CollTime1, c.CollTime2, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the College to the database.
func (c *College) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for College.
//
// NOTE: PostgreSQL 9.5+ only
func (c *College) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.college (` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.coll_id, EXCLUDED.coll_name, EXCLUDED.coll_addr1, EXCLUDED.coll_addr2, EXCLUDED.coll_addr3, EXCLUDED.coll_addr4, EXCLUDED.coll_addr5, EXCLUDED.coll_pcode, EXCLUDED.coll_tel1, EXCLUDED.coll_tel2, EXCLUDED.coll_fax, EXCLUDED.coll_capacity, EXCLUDED.coll_trainer, EXCLUDED.coll_time1, EXCLUDED.coll_time2, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.CollID, c.CollName, c.CollAddr1, c.CollAddr2, c.CollAddr3, c.CollAddr4, c.CollAddr5, c.CollPcode, c.CollTel1, c.CollTel2, c.CollFax, c.CollCapacity, c.CollTrainer, c.CollTime1, c.CollTime2, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.CollID, c.CollName, c.CollAddr1, c.CollAddr2, c.CollAddr3, c.CollAddr4, c.CollAddr5, c.CollPcode, c.CollTel1, c.CollTel2, c.CollFax, c.CollCapacity, c.CollTrainer, c.CollTime1, c.CollTime2, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the College from the database.
func (c *College) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.college WHERE equinox_lrn = $1`

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

// CollegeByEquinoxLrn retrieves a row from 'equinox.college' as a College.
//
// Generated from index 'college_pkey'.
func CollegeByEquinoxLrn(db XODB, equinoxLrn int64) (*College, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_lrn, equinox_sec ` +
		`FROM equinox.college ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := College{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.CollID, &c.CollName, &c.CollAddr1, &c.CollAddr2, &c.CollAddr3, &c.CollAddr4, &c.CollAddr5, &c.CollPcode, &c.CollTel1, &c.CollTel2, &c.CollFax, &c.CollCapacity, &c.CollTrainer, &c.CollTime1, &c.CollTime2, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}