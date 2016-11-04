// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Taskdest represents a row from 'equinox.taskdest'.
type Taskdest struct {
	TaskdestName    sql.NullString `json:"taskdest_name"`    // taskdest_name
	TaskdestManager sql.NullString `json:"taskdest_manager"` // taskdest_manager
	TaskdestHowlate sql.NullInt64  `json:"taskdest_howlate"` // taskdest_howlate
	TaskdestSparec1 sql.NullString `json:"taskdest_sparec1"` // taskdest_sparec1
	TaskdestEmail   sql.NullString `json:"taskdest_email"`   // taskdest_email
	TaskdestSparen1 sql.NullInt64  `json:"taskdest_sparen1"` // taskdest_sparen1
	TaskdestNoemail sql.NullInt64  `json:"taskdest_noemail"` // taskdest_noemail
	TaskdestSparem1 sql.NullInt64  `json:"taskdest_sparem1"` // taskdest_sparem1
	TaskdestSuspend sql.NullInt64  `json:"taskdest_suspend"` // taskdest_suspend
	EquinoxLrn      int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Taskdest exists in the database.
func (t *Taskdest) Exists() bool {
	return t._exists
}

// Deleted provides information if the Taskdest has been deleted from the database.
func (t *Taskdest) Deleted() bool {
	return t._deleted
}

// Insert inserts the Taskdest to the database.
func (t *Taskdest) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.taskdest (` +
		`taskdest_name, taskdest_manager, taskdest_howlate, taskdest_sparec1, taskdest_email, taskdest_sparen1, taskdest_noemail, taskdest_sparem1, taskdest_suspend, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, t.TaskdestName, t.TaskdestManager, t.TaskdestHowlate, t.TaskdestSparec1, t.TaskdestEmail, t.TaskdestSparen1, t.TaskdestNoemail, t.TaskdestSparem1, t.TaskdestSuspend, t.EquinoxSec)
	err = db.QueryRow(sqlstr, t.TaskdestName, t.TaskdestManager, t.TaskdestHowlate, t.TaskdestSparec1, t.TaskdestEmail, t.TaskdestSparen1, t.TaskdestNoemail, t.TaskdestSparem1, t.TaskdestSuspend, t.EquinoxSec).Scan(&t.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Taskdest in the database.
func (t *Taskdest) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.taskdest SET (` +
		`taskdest_name, taskdest_manager, taskdest_howlate, taskdest_sparec1, taskdest_email, taskdest_sparen1, taskdest_noemail, taskdest_sparem1, taskdest_suspend, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) WHERE equinox_lrn = $11`

	// run query
	XOLog(sqlstr, t.TaskdestName, t.TaskdestManager, t.TaskdestHowlate, t.TaskdestSparec1, t.TaskdestEmail, t.TaskdestSparen1, t.TaskdestNoemail, t.TaskdestSparem1, t.TaskdestSuspend, t.EquinoxSec, t.EquinoxLrn)
	_, err = db.Exec(sqlstr, t.TaskdestName, t.TaskdestManager, t.TaskdestHowlate, t.TaskdestSparec1, t.TaskdestEmail, t.TaskdestSparen1, t.TaskdestNoemail, t.TaskdestSparem1, t.TaskdestSuspend, t.EquinoxSec, t.EquinoxLrn)
	return err
}

// Save saves the Taskdest to the database.
func (t *Taskdest) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Upsert performs an upsert for Taskdest.
//
// NOTE: PostgreSQL 9.5+ only
func (t *Taskdest) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.taskdest (` +
		`taskdest_name, taskdest_manager, taskdest_howlate, taskdest_sparec1, taskdest_email, taskdest_sparen1, taskdest_noemail, taskdest_sparem1, taskdest_suspend, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`taskdest_name, taskdest_manager, taskdest_howlate, taskdest_sparec1, taskdest_email, taskdest_sparen1, taskdest_noemail, taskdest_sparem1, taskdest_suspend, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.taskdest_name, EXCLUDED.taskdest_manager, EXCLUDED.taskdest_howlate, EXCLUDED.taskdest_sparec1, EXCLUDED.taskdest_email, EXCLUDED.taskdest_sparen1, EXCLUDED.taskdest_noemail, EXCLUDED.taskdest_sparem1, EXCLUDED.taskdest_suspend, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, t.TaskdestName, t.TaskdestManager, t.TaskdestHowlate, t.TaskdestSparec1, t.TaskdestEmail, t.TaskdestSparen1, t.TaskdestNoemail, t.TaskdestSparem1, t.TaskdestSuspend, t.EquinoxLrn, t.EquinoxSec)
	_, err = db.Exec(sqlstr, t.TaskdestName, t.TaskdestManager, t.TaskdestHowlate, t.TaskdestSparec1, t.TaskdestEmail, t.TaskdestSparen1, t.TaskdestNoemail, t.TaskdestSparem1, t.TaskdestSuspend, t.EquinoxLrn, t.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Delete deletes the Taskdest from the database.
func (t *Taskdest) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.taskdest WHERE equinox_lrn = $1`

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

// TaskdestByEquinoxLrn retrieves a row from 'equinox.taskdest' as a Taskdest.
//
// Generated from index 'taskdest_pkey'.
func TaskdestByEquinoxLrn(db XODB, equinoxLrn int64) (*Taskdest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`taskdest_name, taskdest_manager, taskdest_howlate, taskdest_sparec1, taskdest_email, taskdest_sparen1, taskdest_noemail, taskdest_sparem1, taskdest_suspend, equinox_lrn, equinox_sec ` +
		`FROM equinox.taskdest ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	t := Taskdest{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.TaskdestName, &t.TaskdestManager, &t.TaskdestHowlate, &t.TaskdestSparec1, &t.TaskdestEmail, &t.TaskdestSparen1, &t.TaskdestNoemail, &t.TaskdestSparem1, &t.TaskdestSuspend, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
