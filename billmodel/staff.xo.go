// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Staff represents a row from 'equinox.staff'.
type Staff struct {
	Staffno          sql.NullString  `json:"staffno"`          // staffno
	Staffname        sql.NullString  `json:"staffname"`        // staffname
	Staffdeptid      sql.NullString  `json:"staffdeptid"`      // staffdeptid
	Staffteamid      sql.NullFloat64 `json:"staffteamid"`      // staffteamid
	Staffphoneext    sql.NullInt64   `json:"staffphoneext"`    // staffphoneext
	Staffemail       sql.NullString  `json:"staffemail"`       // staffemail
	Stafflogin       sql.NullString  `json:"stafflogin"`       // stafflogin
	Staffgroup       sql.NullString  `json:"staffgroup"`       // staffgroup
	Staffeditstaff   sql.NullInt64   `json:"staffeditstaff"`   // staffeditstaff
	Staffrestrict    sql.NullInt64   `json:"staffrestrict"`    // staffrestrict
	Staffmanager     sql.NullInt64   `json:"staffmanager"`     // staffmanager
	Staffbilling     sql.NullInt64   `json:"staffbilling"`     // staffbilling
	Staffmobadmin    sql.NullInt64   `json:"staffmobadmin"`    // staffmobadmin
	Staffdeleterecs  sql.NullInt64   `json:"staffdeleterecs"`  // staffdeleterecs
	Staffeditrecs    sql.NullInt64   `json:"staffeditrecs"`    // staffeditrecs
	Staffeditcomment sql.NullInt64   `json:"staffeditcomment"` // staffeditcomment
	Staffoutputdir   sql.NullString  `json:"staffoutputdir"`   // staffoutputdir
	Staffuserid      sql.NullString  `json:"staffuserid"`      // staffuserid
	Staffavailable   sql.NullString  `json:"staffavailable"`   // staffavailable
	Staffmobileno    sql.NullString  `json:"staffmobileno"`    // staffmobileno
	Staffposition    sql.NullInt64   `json:"staffposition"`    // staffposition
	Staffaccessflags sql.NullString  `json:"staffaccessflags"` // staffaccessflags
	Staffsparen1     sql.NullFloat64 `json:"staffsparen1"`     // staffsparen1
	Staffspared1     pq.NullTime     `json:"staffspared1"`     // staffspared1
	Staffphonelogin  sql.NullString  `json:"staffphonelogin"`  // staffphonelogin
	Staffcompdeptid  sql.NullInt64   `json:"staffcompdeptid"`  // staffcompdeptid
	Staffleftcompany sql.NullInt64   `json:"staffleftcompany"` // staffleftcompany
	Staffhrdept      sql.NullString  `json:"staffhrdept"`      // staffhrdept
	Stafftopinbox    sql.NullString  `json:"stafftopinbox"`    // stafftopinbox
	Staffsparec1     sql.NullString  `json:"staffsparec1"`     // staffsparec1
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Staff exists in the database.
func (s *Staff) Exists() bool {
	return s._exists
}

// Deleted provides information if the Staff has been deleted from the database.
func (s *Staff) Deleted() bool {
	return s._deleted
}

// Insert inserts the Staff to the database.
func (s *Staff) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.staff (` +
		`staffno, staffname, staffdeptid, staffteamid, staffphoneext, staffemail, stafflogin, staffgroup, staffeditstaff, staffrestrict, staffmanager, staffbilling, staffmobadmin, staffdeleterecs, staffeditrecs, staffeditcomment, staffoutputdir, staffuserid, staffavailable, staffmobileno, staffposition, staffaccessflags, staffsparen1, staffspared1, staffphonelogin, staffcompdeptid, staffleftcompany, staffhrdept, stafftopinbox, staffsparec1, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, s.Staffno, s.Staffname, s.Staffdeptid, s.Staffteamid, s.Staffphoneext, s.Staffemail, s.Stafflogin, s.Staffgroup, s.Staffeditstaff, s.Staffrestrict, s.Staffmanager, s.Staffbilling, s.Staffmobadmin, s.Staffdeleterecs, s.Staffeditrecs, s.Staffeditcomment, s.Staffoutputdir, s.Staffuserid, s.Staffavailable, s.Staffmobileno, s.Staffposition, s.Staffaccessflags, s.Staffsparen1, s.Staffspared1, s.Staffphonelogin, s.Staffcompdeptid, s.Staffleftcompany, s.Staffhrdept, s.Stafftopinbox, s.Staffsparec1, s.EquinoxSec)
	err = db.QueryRow(sqlstr, s.Staffno, s.Staffname, s.Staffdeptid, s.Staffteamid, s.Staffphoneext, s.Staffemail, s.Stafflogin, s.Staffgroup, s.Staffeditstaff, s.Staffrestrict, s.Staffmanager, s.Staffbilling, s.Staffmobadmin, s.Staffdeleterecs, s.Staffeditrecs, s.Staffeditcomment, s.Staffoutputdir, s.Staffuserid, s.Staffavailable, s.Staffmobileno, s.Staffposition, s.Staffaccessflags, s.Staffsparen1, s.Staffspared1, s.Staffphonelogin, s.Staffcompdeptid, s.Staffleftcompany, s.Staffhrdept, s.Stafftopinbox, s.Staffsparec1, s.EquinoxSec).Scan(&s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the Staff in the database.
func (s *Staff) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.staff SET (` +
		`staffno, staffname, staffdeptid, staffteamid, staffphoneext, staffemail, stafflogin, staffgroup, staffeditstaff, staffrestrict, staffmanager, staffbilling, staffmobadmin, staffdeleterecs, staffeditrecs, staffeditcomment, staffoutputdir, staffuserid, staffavailable, staffmobileno, staffposition, staffaccessflags, staffsparen1, staffspared1, staffphonelogin, staffcompdeptid, staffleftcompany, staffhrdept, stafftopinbox, staffsparec1, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31` +
		`) WHERE equinox_lrn = $32`

	// run query
	XOLog(sqlstr, s.Staffno, s.Staffname, s.Staffdeptid, s.Staffteamid, s.Staffphoneext, s.Staffemail, s.Stafflogin, s.Staffgroup, s.Staffeditstaff, s.Staffrestrict, s.Staffmanager, s.Staffbilling, s.Staffmobadmin, s.Staffdeleterecs, s.Staffeditrecs, s.Staffeditcomment, s.Staffoutputdir, s.Staffuserid, s.Staffavailable, s.Staffmobileno, s.Staffposition, s.Staffaccessflags, s.Staffsparen1, s.Staffspared1, s.Staffphonelogin, s.Staffcompdeptid, s.Staffleftcompany, s.Staffhrdept, s.Stafftopinbox, s.Staffsparec1, s.EquinoxSec, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.Staffno, s.Staffname, s.Staffdeptid, s.Staffteamid, s.Staffphoneext, s.Staffemail, s.Stafflogin, s.Staffgroup, s.Staffeditstaff, s.Staffrestrict, s.Staffmanager, s.Staffbilling, s.Staffmobadmin, s.Staffdeleterecs, s.Staffeditrecs, s.Staffeditcomment, s.Staffoutputdir, s.Staffuserid, s.Staffavailable, s.Staffmobileno, s.Staffposition, s.Staffaccessflags, s.Staffsparen1, s.Staffspared1, s.Staffphonelogin, s.Staffcompdeptid, s.Staffleftcompany, s.Staffhrdept, s.Stafftopinbox, s.Staffsparec1, s.EquinoxSec, s.EquinoxLrn)
	return err
}

// Save saves the Staff to the database.
func (s *Staff) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for Staff.
//
// NOTE: PostgreSQL 9.5+ only
func (s *Staff) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.staff (` +
		`staffno, staffname, staffdeptid, staffteamid, staffphoneext, staffemail, stafflogin, staffgroup, staffeditstaff, staffrestrict, staffmanager, staffbilling, staffmobadmin, staffdeleterecs, staffeditrecs, staffeditcomment, staffoutputdir, staffuserid, staffavailable, staffmobileno, staffposition, staffaccessflags, staffsparen1, staffspared1, staffphonelogin, staffcompdeptid, staffleftcompany, staffhrdept, stafftopinbox, staffsparec1, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`staffno, staffname, staffdeptid, staffteamid, staffphoneext, staffemail, stafflogin, staffgroup, staffeditstaff, staffrestrict, staffmanager, staffbilling, staffmobadmin, staffdeleterecs, staffeditrecs, staffeditcomment, staffoutputdir, staffuserid, staffavailable, staffmobileno, staffposition, staffaccessflags, staffsparen1, staffspared1, staffphonelogin, staffcompdeptid, staffleftcompany, staffhrdept, stafftopinbox, staffsparec1, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.staffno, EXCLUDED.staffname, EXCLUDED.staffdeptid, EXCLUDED.staffteamid, EXCLUDED.staffphoneext, EXCLUDED.staffemail, EXCLUDED.stafflogin, EXCLUDED.staffgroup, EXCLUDED.staffeditstaff, EXCLUDED.staffrestrict, EXCLUDED.staffmanager, EXCLUDED.staffbilling, EXCLUDED.staffmobadmin, EXCLUDED.staffdeleterecs, EXCLUDED.staffeditrecs, EXCLUDED.staffeditcomment, EXCLUDED.staffoutputdir, EXCLUDED.staffuserid, EXCLUDED.staffavailable, EXCLUDED.staffmobileno, EXCLUDED.staffposition, EXCLUDED.staffaccessflags, EXCLUDED.staffsparen1, EXCLUDED.staffspared1, EXCLUDED.staffphonelogin, EXCLUDED.staffcompdeptid, EXCLUDED.staffleftcompany, EXCLUDED.staffhrdept, EXCLUDED.stafftopinbox, EXCLUDED.staffsparec1, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, s.Staffno, s.Staffname, s.Staffdeptid, s.Staffteamid, s.Staffphoneext, s.Staffemail, s.Stafflogin, s.Staffgroup, s.Staffeditstaff, s.Staffrestrict, s.Staffmanager, s.Staffbilling, s.Staffmobadmin, s.Staffdeleterecs, s.Staffeditrecs, s.Staffeditcomment, s.Staffoutputdir, s.Staffuserid, s.Staffavailable, s.Staffmobileno, s.Staffposition, s.Staffaccessflags, s.Staffsparen1, s.Staffspared1, s.Staffphonelogin, s.Staffcompdeptid, s.Staffleftcompany, s.Staffhrdept, s.Stafftopinbox, s.Staffsparec1, s.EquinoxLrn, s.EquinoxSec)
	_, err = db.Exec(sqlstr, s.Staffno, s.Staffname, s.Staffdeptid, s.Staffteamid, s.Staffphoneext, s.Staffemail, s.Stafflogin, s.Staffgroup, s.Staffeditstaff, s.Staffrestrict, s.Staffmanager, s.Staffbilling, s.Staffmobadmin, s.Staffdeleterecs, s.Staffeditrecs, s.Staffeditcomment, s.Staffoutputdir, s.Staffuserid, s.Staffavailable, s.Staffmobileno, s.Staffposition, s.Staffaccessflags, s.Staffsparen1, s.Staffspared1, s.Staffphonelogin, s.Staffcompdeptid, s.Staffleftcompany, s.Staffhrdept, s.Stafftopinbox, s.Staffsparec1, s.EquinoxLrn, s.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the Staff from the database.
func (s *Staff) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.staff WHERE equinox_lrn = $1`

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

// StaffByEquinoxLrn retrieves a row from 'equinox.staff' as a Staff.
//
// Generated from index 'staff_pkey'.
func StaffByEquinoxLrn(db XODB, equinoxLrn int64) (*Staff, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`staffno, staffname, staffdeptid, staffteamid, staffphoneext, staffemail, stafflogin, staffgroup, staffeditstaff, staffrestrict, staffmanager, staffbilling, staffmobadmin, staffdeleterecs, staffeditrecs, staffeditcomment, staffoutputdir, staffuserid, staffavailable, staffmobileno, staffposition, staffaccessflags, staffsparen1, staffspared1, staffphonelogin, staffcompdeptid, staffleftcompany, staffhrdept, stafftopinbox, staffsparec1, equinox_lrn, equinox_sec ` +
		`FROM equinox.staff ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Staff{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Staffno, &s.Staffname, &s.Staffdeptid, &s.Staffteamid, &s.Staffphoneext, &s.Staffemail, &s.Stafflogin, &s.Staffgroup, &s.Staffeditstaff, &s.Staffrestrict, &s.Staffmanager, &s.Staffbilling, &s.Staffmobadmin, &s.Staffdeleterecs, &s.Staffeditrecs, &s.Staffeditcomment, &s.Staffoutputdir, &s.Staffuserid, &s.Staffavailable, &s.Staffmobileno, &s.Staffposition, &s.Staffaccessflags, &s.Staffsparen1, &s.Staffspared1, &s.Staffphonelogin, &s.Staffcompdeptid, &s.Staffleftcompany, &s.Staffhrdept, &s.Stafftopinbox, &s.Staffsparec1, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
