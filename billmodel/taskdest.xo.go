// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

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
}

func AllTaskdest(db XODB, callback func(x Taskdest) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`taskdest_name, taskdest_manager, taskdest_howlate, taskdest_sparec1, taskdest_email, taskdest_sparen1, taskdest_noemail, taskdest_sparem1, taskdest_suspend, equinox_lrn, equinox_sec ` +
		`FROM equinox.taskdest `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		t := Taskdest{}

		// scan
		err = q.Scan(&t.TaskdestName, &t.TaskdestManager, &t.TaskdestHowlate, &t.TaskdestSparec1, &t.TaskdestEmail, &t.TaskdestSparen1, &t.TaskdestNoemail, &t.TaskdestSparem1, &t.TaskdestSuspend, &t.EquinoxLrn, &t.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(t) {
			return nil
		}
	}

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
	t := Taskdest{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.TaskdestName, &t.TaskdestManager, &t.TaskdestHowlate, &t.TaskdestSparec1, &t.TaskdestEmail, &t.TaskdestSparen1, &t.TaskdestNoemail, &t.TaskdestSparem1, &t.TaskdestSuspend, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
