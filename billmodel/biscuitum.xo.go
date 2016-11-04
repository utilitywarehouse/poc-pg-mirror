// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Biscuitum represents a row from 'equinox.biscuita'.
type Biscuitum struct {
	Bisref           sql.NullString `json:"bisref"`           // bisref
	Biscli           sql.NullString `json:"biscli"`           // biscli
	Bisinitid        sql.NullString `json:"bisinitid"`        // bisinitid
	Bisinitdept      sql.NullString `json:"bisinitdept"`      // bisinitdept
	Bisrecipientid   sql.NullString `json:"bisrecipientid"`   // bisrecipientid
	Bisrecdeptid     sql.NullString `json:"bisrecdeptid"`     // bisrecdeptid
	Bisothercontact  sql.NullString `json:"bisothercontact"`  // bisothercontact
	Bisflow          sql.NullString `json:"bisflow"`          // bisflow
	Bisname          sql.NullString `json:"bisname"`          // bisname
	Bishousenumber   sql.NullString `json:"bishousenumber"`   // bishousenumber
	Bisstreetname    sql.NullString `json:"bisstreetname"`    // bisstreetname
	Bisaddress       sql.NullString `json:"bisaddress"`       // bisaddress
	Bispostcode      sql.NullString `json:"bispostcode"`      // bispostcode
	Bistransferdate  pq.NullTime    `json:"bistransferdate"`  // bistransferdate
	Bistranscoread   sql.NullInt64  `json:"bistranscoread"`   // bistranscoread
	Bistranscoae     sql.NullString `json:"bistranscoae"`     // bistranscoae
	Biscriteria      sql.NullString `json:"biscriteria"`      // biscriteria
	Bisproposedread  sql.NullInt64  `json:"bisproposedread"`  // bisproposedread
	Bisproposedae    sql.NullString `json:"bisproposedae"`    // bisproposedae
	Bisroot          sql.NullString `json:"bisroot"`          // bisroot
	Bisstatus        sql.NullString `json:"bisstatus"`        // bisstatus
	Bisfilesend      sql.NullInt64  `json:"bisfilesend"`      // bisfilesend
	Bisanum1         sql.NullInt64  `json:"bisanum1"`         // bisanum1
	Bisanum2         sql.NullInt64  `json:"bisanum2"`         // bisanum2
	Bissendaread     pq.NullTime    `json:"bissendaread"`     // bissendaread
	Bissendmstatus   pq.NullTime    `json:"bissendmstatus"`   // bissendmstatus
	Bisdatetotransco pq.NullTime    `json:"bisdatetotransco"` // bisdatetotransco
	Bismetastatus    sql.NullString `json:"bismetastatus"`    // bismetastatus
	Bismpr           sql.NullString `json:"bismpr"`           // bismpr
	Bissendtotransco sql.NullString `json:"bissendtotransco"` // bissendtotransco
	Bisreadagreed    pq.NullTime    `json:"bisreadagreed"`    // bisreadagreed
	Biscomplete      pq.NullTime    `json:"biscomplete"`      // biscomplete
	Bistempmemo      sql.NullInt64  `json:"bistempmemo"`      // bistempmemo
	Bisac3           sql.NullString `json:"bisac3"`           // bisac3
	Bisac4           sql.NullString `json:"bisac4"`           // bisac4
	Bissendrejection pq.NullTime    `json:"bissendrejection"` // bissendrejection
	Bistoescalate    pq.NullTime    `json:"bistoescalate"`    // bistoescalate
	Bisrefid         sql.NullString `json:"bisrefid"`         // bisrefid
	Bisrejcode       sql.NullString `json:"bisrejcode"`       // bisrejcode
	Bisescalations   pq.NullTime    `json:"bisescalations"`   // bisescalations
	Bissentproposal  pq.NullTime    `json:"bissentproposal"`  // bissentproposal
	Bisad8           pq.NullTime    `json:"bisad8"`           // bisad8
	Bisad9           pq.NullTime    `json:"bisad9"`           // bisad9
	Bisdatecreated   pq.NullTime    `json:"bisdatecreated"`   // bisdatecreated
	Bisourstatus     sql.NullInt64  `json:"bisourstatus"`     // bisourstatus
	Bislastchangeby  sql.NullString `json:"bislastchangeby"`  // bislastchangeby
	Bislastchange    sql.NullString `json:"bislastchange"`    // bislastchange
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Biscuitum exists in the database.
func (b *Biscuitum) Exists() bool {
	return b._exists
}

// Deleted provides information if the Biscuitum has been deleted from the database.
func (b *Biscuitum) Deleted() bool {
	return b._deleted
}

// Insert inserts the Biscuitum to the database.
func (b *Biscuitum) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.biscuita (` +
		`bisref, biscli, bisinitid, bisinitdept, bisrecipientid, bisrecdeptid, bisothercontact, bisflow, bisname, bishousenumber, bisstreetname, bisaddress, bispostcode, bistransferdate, bistranscoread, bistranscoae, biscriteria, bisproposedread, bisproposedae, bisroot, bisstatus, bisfilesend, bisanum1, bisanum2, bissendaread, bissendmstatus, bisdatetotransco, bismetastatus, bismpr, bissendtotransco, bisreadagreed, biscomplete, bistempmemo, bisac3, bisac4, bissendrejection, bistoescalate, bisrefid, bisrejcode, bisescalations, bissentproposal, bisad8, bisad9, bisdatecreated, bisourstatus, bislastchangeby, bislastchange, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, b.Bisref, b.Biscli, b.Bisinitid, b.Bisinitdept, b.Bisrecipientid, b.Bisrecdeptid, b.Bisothercontact, b.Bisflow, b.Bisname, b.Bishousenumber, b.Bisstreetname, b.Bisaddress, b.Bispostcode, b.Bistransferdate, b.Bistranscoread, b.Bistranscoae, b.Biscriteria, b.Bisproposedread, b.Bisproposedae, b.Bisroot, b.Bisstatus, b.Bisfilesend, b.Bisanum1, b.Bisanum2, b.Bissendaread, b.Bissendmstatus, b.Bisdatetotransco, b.Bismetastatus, b.Bismpr, b.Bissendtotransco, b.Bisreadagreed, b.Biscomplete, b.Bistempmemo, b.Bisac3, b.Bisac4, b.Bissendrejection, b.Bistoescalate, b.Bisrefid, b.Bisrejcode, b.Bisescalations, b.Bissentproposal, b.Bisad8, b.Bisad9, b.Bisdatecreated, b.Bisourstatus, b.Bislastchangeby, b.Bislastchange, b.EquinoxSec)
	err = db.QueryRow(sqlstr, b.Bisref, b.Biscli, b.Bisinitid, b.Bisinitdept, b.Bisrecipientid, b.Bisrecdeptid, b.Bisothercontact, b.Bisflow, b.Bisname, b.Bishousenumber, b.Bisstreetname, b.Bisaddress, b.Bispostcode, b.Bistransferdate, b.Bistranscoread, b.Bistranscoae, b.Biscriteria, b.Bisproposedread, b.Bisproposedae, b.Bisroot, b.Bisstatus, b.Bisfilesend, b.Bisanum1, b.Bisanum2, b.Bissendaread, b.Bissendmstatus, b.Bisdatetotransco, b.Bismetastatus, b.Bismpr, b.Bissendtotransco, b.Bisreadagreed, b.Biscomplete, b.Bistempmemo, b.Bisac3, b.Bisac4, b.Bissendrejection, b.Bistoescalate, b.Bisrefid, b.Bisrejcode, b.Bisescalations, b.Bissentproposal, b.Bisad8, b.Bisad9, b.Bisdatecreated, b.Bisourstatus, b.Bislastchangeby, b.Bislastchange, b.EquinoxSec).Scan(&b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Update updates the Biscuitum in the database.
func (b *Biscuitum) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if b._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.biscuita SET (` +
		`bisref, biscli, bisinitid, bisinitdept, bisrecipientid, bisrecdeptid, bisothercontact, bisflow, bisname, bishousenumber, bisstreetname, bisaddress, bispostcode, bistransferdate, bistranscoread, bistranscoae, biscriteria, bisproposedread, bisproposedae, bisroot, bisstatus, bisfilesend, bisanum1, bisanum2, bissendaread, bissendmstatus, bisdatetotransco, bismetastatus, bismpr, bissendtotransco, bisreadagreed, biscomplete, bistempmemo, bisac3, bisac4, bissendrejection, bistoescalate, bisrefid, bisrejcode, bisescalations, bissentproposal, bisad8, bisad9, bisdatecreated, bisourstatus, bislastchangeby, bislastchange, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48` +
		`) WHERE equinox_lrn = $49`

	// run query
	XOLog(sqlstr, b.Bisref, b.Biscli, b.Bisinitid, b.Bisinitdept, b.Bisrecipientid, b.Bisrecdeptid, b.Bisothercontact, b.Bisflow, b.Bisname, b.Bishousenumber, b.Bisstreetname, b.Bisaddress, b.Bispostcode, b.Bistransferdate, b.Bistranscoread, b.Bistranscoae, b.Biscriteria, b.Bisproposedread, b.Bisproposedae, b.Bisroot, b.Bisstatus, b.Bisfilesend, b.Bisanum1, b.Bisanum2, b.Bissendaread, b.Bissendmstatus, b.Bisdatetotransco, b.Bismetastatus, b.Bismpr, b.Bissendtotransco, b.Bisreadagreed, b.Biscomplete, b.Bistempmemo, b.Bisac3, b.Bisac4, b.Bissendrejection, b.Bistoescalate, b.Bisrefid, b.Bisrejcode, b.Bisescalations, b.Bissentproposal, b.Bisad8, b.Bisad9, b.Bisdatecreated, b.Bisourstatus, b.Bislastchangeby, b.Bislastchange, b.EquinoxSec, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.Bisref, b.Biscli, b.Bisinitid, b.Bisinitdept, b.Bisrecipientid, b.Bisrecdeptid, b.Bisothercontact, b.Bisflow, b.Bisname, b.Bishousenumber, b.Bisstreetname, b.Bisaddress, b.Bispostcode, b.Bistransferdate, b.Bistranscoread, b.Bistranscoae, b.Biscriteria, b.Bisproposedread, b.Bisproposedae, b.Bisroot, b.Bisstatus, b.Bisfilesend, b.Bisanum1, b.Bisanum2, b.Bissendaread, b.Bissendmstatus, b.Bisdatetotransco, b.Bismetastatus, b.Bismpr, b.Bissendtotransco, b.Bisreadagreed, b.Biscomplete, b.Bistempmemo, b.Bisac3, b.Bisac4, b.Bissendrejection, b.Bistoescalate, b.Bisrefid, b.Bisrejcode, b.Bisescalations, b.Bissentproposal, b.Bisad8, b.Bisad9, b.Bisdatecreated, b.Bisourstatus, b.Bislastchangeby, b.Bislastchange, b.EquinoxSec, b.EquinoxLrn)
	return err
}

// Save saves the Biscuitum to the database.
func (b *Biscuitum) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Upsert performs an upsert for Biscuitum.
//
// NOTE: PostgreSQL 9.5+ only
func (b *Biscuitum) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.biscuita (` +
		`bisref, biscli, bisinitid, bisinitdept, bisrecipientid, bisrecdeptid, bisothercontact, bisflow, bisname, bishousenumber, bisstreetname, bisaddress, bispostcode, bistransferdate, bistranscoread, bistranscoae, biscriteria, bisproposedread, bisproposedae, bisroot, bisstatus, bisfilesend, bisanum1, bisanum2, bissendaread, bissendmstatus, bisdatetotransco, bismetastatus, bismpr, bissendtotransco, bisreadagreed, biscomplete, bistempmemo, bisac3, bisac4, bissendrejection, bistoescalate, bisrefid, bisrejcode, bisescalations, bissentproposal, bisad8, bisad9, bisdatecreated, bisourstatus, bislastchangeby, bislastchange, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`bisref, biscli, bisinitid, bisinitdept, bisrecipientid, bisrecdeptid, bisothercontact, bisflow, bisname, bishousenumber, bisstreetname, bisaddress, bispostcode, bistransferdate, bistranscoread, bistranscoae, biscriteria, bisproposedread, bisproposedae, bisroot, bisstatus, bisfilesend, bisanum1, bisanum2, bissendaread, bissendmstatus, bisdatetotransco, bismetastatus, bismpr, bissendtotransco, bisreadagreed, biscomplete, bistempmemo, bisac3, bisac4, bissendrejection, bistoescalate, bisrefid, bisrejcode, bisescalations, bissentproposal, bisad8, bisad9, bisdatecreated, bisourstatus, bislastchangeby, bislastchange, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.bisref, EXCLUDED.biscli, EXCLUDED.bisinitid, EXCLUDED.bisinitdept, EXCLUDED.bisrecipientid, EXCLUDED.bisrecdeptid, EXCLUDED.bisothercontact, EXCLUDED.bisflow, EXCLUDED.bisname, EXCLUDED.bishousenumber, EXCLUDED.bisstreetname, EXCLUDED.bisaddress, EXCLUDED.bispostcode, EXCLUDED.bistransferdate, EXCLUDED.bistranscoread, EXCLUDED.bistranscoae, EXCLUDED.biscriteria, EXCLUDED.bisproposedread, EXCLUDED.bisproposedae, EXCLUDED.bisroot, EXCLUDED.bisstatus, EXCLUDED.bisfilesend, EXCLUDED.bisanum1, EXCLUDED.bisanum2, EXCLUDED.bissendaread, EXCLUDED.bissendmstatus, EXCLUDED.bisdatetotransco, EXCLUDED.bismetastatus, EXCLUDED.bismpr, EXCLUDED.bissendtotransco, EXCLUDED.bisreadagreed, EXCLUDED.biscomplete, EXCLUDED.bistempmemo, EXCLUDED.bisac3, EXCLUDED.bisac4, EXCLUDED.bissendrejection, EXCLUDED.bistoescalate, EXCLUDED.bisrefid, EXCLUDED.bisrejcode, EXCLUDED.bisescalations, EXCLUDED.bissentproposal, EXCLUDED.bisad8, EXCLUDED.bisad9, EXCLUDED.bisdatecreated, EXCLUDED.bisourstatus, EXCLUDED.bislastchangeby, EXCLUDED.bislastchange, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, b.Bisref, b.Biscli, b.Bisinitid, b.Bisinitdept, b.Bisrecipientid, b.Bisrecdeptid, b.Bisothercontact, b.Bisflow, b.Bisname, b.Bishousenumber, b.Bisstreetname, b.Bisaddress, b.Bispostcode, b.Bistransferdate, b.Bistranscoread, b.Bistranscoae, b.Biscriteria, b.Bisproposedread, b.Bisproposedae, b.Bisroot, b.Bisstatus, b.Bisfilesend, b.Bisanum1, b.Bisanum2, b.Bissendaread, b.Bissendmstatus, b.Bisdatetotransco, b.Bismetastatus, b.Bismpr, b.Bissendtotransco, b.Bisreadagreed, b.Biscomplete, b.Bistempmemo, b.Bisac3, b.Bisac4, b.Bissendrejection, b.Bistoescalate, b.Bisrefid, b.Bisrejcode, b.Bisescalations, b.Bissentproposal, b.Bisad8, b.Bisad9, b.Bisdatecreated, b.Bisourstatus, b.Bislastchangeby, b.Bislastchange, b.EquinoxLrn, b.EquinoxSec)
	_, err = db.Exec(sqlstr, b.Bisref, b.Biscli, b.Bisinitid, b.Bisinitdept, b.Bisrecipientid, b.Bisrecdeptid, b.Bisothercontact, b.Bisflow, b.Bisname, b.Bishousenumber, b.Bisstreetname, b.Bisaddress, b.Bispostcode, b.Bistransferdate, b.Bistranscoread, b.Bistranscoae, b.Biscriteria, b.Bisproposedread, b.Bisproposedae, b.Bisroot, b.Bisstatus, b.Bisfilesend, b.Bisanum1, b.Bisanum2, b.Bissendaread, b.Bissendmstatus, b.Bisdatetotransco, b.Bismetastatus, b.Bismpr, b.Bissendtotransco, b.Bisreadagreed, b.Biscomplete, b.Bistempmemo, b.Bisac3, b.Bisac4, b.Bissendrejection, b.Bistoescalate, b.Bisrefid, b.Bisrejcode, b.Bisescalations, b.Bissentproposal, b.Bisad8, b.Bisad9, b.Bisdatecreated, b.Bisourstatus, b.Bislastchangeby, b.Bislastchange, b.EquinoxLrn, b.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Delete deletes the Biscuitum from the database.
func (b *Biscuitum) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.biscuita WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	b._deleted = true

	return nil
}

// BiscuitumByEquinoxLrn retrieves a row from 'equinox.biscuita' as a Biscuitum.
//
// Generated from index 'biscuita_pkey'.
func BiscuitumByEquinoxLrn(db XODB, equinoxLrn int64) (*Biscuitum, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`bisref, biscli, bisinitid, bisinitdept, bisrecipientid, bisrecdeptid, bisothercontact, bisflow, bisname, bishousenumber, bisstreetname, bisaddress, bispostcode, bistransferdate, bistranscoread, bistranscoae, biscriteria, bisproposedread, bisproposedae, bisroot, bisstatus, bisfilesend, bisanum1, bisanum2, bissendaread, bissendmstatus, bisdatetotransco, bismetastatus, bismpr, bissendtotransco, bisreadagreed, biscomplete, bistempmemo, bisac3, bisac4, bissendrejection, bistoescalate, bisrefid, bisrejcode, bisescalations, bissentproposal, bisad8, bisad9, bisdatecreated, bisourstatus, bislastchangeby, bislastchange, equinox_lrn, equinox_sec ` +
		`FROM equinox.biscuita ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Biscuitum{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Bisref, &b.Biscli, &b.Bisinitid, &b.Bisinitdept, &b.Bisrecipientid, &b.Bisrecdeptid, &b.Bisothercontact, &b.Bisflow, &b.Bisname, &b.Bishousenumber, &b.Bisstreetname, &b.Bisaddress, &b.Bispostcode, &b.Bistransferdate, &b.Bistranscoread, &b.Bistranscoae, &b.Biscriteria, &b.Bisproposedread, &b.Bisproposedae, &b.Bisroot, &b.Bisstatus, &b.Bisfilesend, &b.Bisanum1, &b.Bisanum2, &b.Bissendaread, &b.Bissendmstatus, &b.Bisdatetotransco, &b.Bismetastatus, &b.Bismpr, &b.Bissendtotransco, &b.Bisreadagreed, &b.Biscomplete, &b.Bistempmemo, &b.Bisac3, &b.Bisac4, &b.Bissendrejection, &b.Bistoescalate, &b.Bisrefid, &b.Bisrejcode, &b.Bisescalations, &b.Bissentproposal, &b.Bisad8, &b.Bisad9, &b.Bisdatecreated, &b.Bisourstatus, &b.Bislastchangeby, &b.Bislastchange, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}