// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Cgb represents a row from 'equinox.cgb'.
type Cgb struct {
	Cgbid            sql.NullInt64   `json:"cgbid"`            // cgbid
	Cgbcustaccountno sql.NullString  `json:"cgbcustaccountno"` // cgbcustaccountno
	Cgbbillinggroup  sql.NullInt64   `json:"cgbbillinggroup"`  // cgbbillinggroup
	Cgbcustexecid    sql.NullString  `json:"cgbcustexecid"`    // cgbcustexecid
	Cgbcli           sql.NullString  `json:"cgbcli"`           // cgbcli
	Cgbperiod        sql.NullString  `json:"cgbperiod"`        // cgbperiod
	Cgbspamount      sql.NullFloat64 `json:"cgbspamount"`      // cgbspamount
	Cgbfpamount      sql.NullFloat64 `json:"cgbfpamount"`      // cgbfpamount
	Cgbmpamount      sql.NullFloat64 `json:"cgbmpamount"`      // cgbmpamount
	Cgbgpamount      sql.NullFloat64 `json:"cgbgpamount"`      // cgbgpamount
	Cgbelecamount    sql.NullFloat64 `json:"cgbelecamount"`    // cgbelecamount
	Cgbgasamount     sql.NullFloat64 `json:"cgbgasamount"`     // cgbgasamount
	Cgbdoubleamount  sql.NullFloat64 `json:"cgbdoubleamount"`  // cgbdoubleamount
	Cgbrating        sql.NullInt64   `json:"cgbrating"`        // cgbrating
	Cgbcustservices  sql.NullString  `json:"cgbcustservices"`  // cgbcustservices
	Cgbcustlivedate  pq.NullTime     `json:"cgbcustlivedate"`  // cgbcustlivedate
	Cgbconnectiondat pq.NullTime     `json:"cgbconnectiondat"` // cgbconnectiondat
	Cgbcustamount    sql.NullFloat64 `json:"cgbcustamount"`    // cgbcustamount
	Cgbcustclaw      sql.NullFloat64 `json:"cgbcustclaw"`      // cgbcustclaw
	Cgbdoubleclaw    sql.NullFloat64 `json:"cgbdoubleclaw"`    // cgbdoubleclaw
	Cgbcustclawperio sql.NullString  `json:"cgbcustclawperio"` // cgbcustclawperio
	Cgboldnew        sql.NullString  `json:"cgboldnew"`        // cgboldnew
	Cgbftcgb         sql.NullFloat64 `json:"cgbftcgb"`         // cgbftcgb
	Cgbsparec1       sql.NullString  `json:"cgbsparec1"`       // cgbsparec1
	Cgbsparec2       sql.NullString  `json:"cgbsparec2"`       // cgbsparec2
	Cgbbcamount      sql.NullFloat64 `json:"cgbbcamount"`      // cgbbcamount
	Cgbspared1       pq.NullTime     `json:"cgbspared1"`       // cgbspared1
	Cgbtbbpartner    sql.NullString  `json:"cgbtbbpartner"`    // cgbtbbpartner
	Cgbtbbamount     sql.NullFloat64 `json:"cgbtbbamount"`     // cgbtbbamount
	Cgbordownload    sql.NullFloat64 `json:"cgbordownload"`    // cgbordownload
	Cgborinstantsim  sql.NullFloat64 `json:"cgborinstantsim"`  // cgborinstantsim
	Cgbobngn         sql.NullFloat64 `json:"cgbobngn"`         // cgbobngn
	Cgbtotal         sql.NullFloat64 `json:"cgbtotal"`         // cgbtotal
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cgb exists in the database.
func (c *Cgb) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cgb has been deleted from the database.
func (c *Cgb) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cgb to the database.
func (c *Cgb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cgb (` +
		`cgbid, cgbcustaccountno, cgbbillinggroup, cgbcustexecid, cgbcli, cgbperiod, cgbspamount, cgbfpamount, cgbmpamount, cgbgpamount, cgbelecamount, cgbgasamount, cgbdoubleamount, cgbrating, cgbcustservices, cgbcustlivedate, cgbconnectiondat, cgbcustamount, cgbcustclaw, cgbdoubleclaw, cgbcustclawperio, cgboldnew, cgbftcgb, cgbsparec1, cgbsparec2, cgbbcamount, cgbspared1, cgbtbbpartner, cgbtbbamount, cgbordownload, cgborinstantsim, cgbobngn, cgbtotal, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cgbid, c.Cgbcustaccountno, c.Cgbbillinggroup, c.Cgbcustexecid, c.Cgbcli, c.Cgbperiod, c.Cgbspamount, c.Cgbfpamount, c.Cgbmpamount, c.Cgbgpamount, c.Cgbelecamount, c.Cgbgasamount, c.Cgbdoubleamount, c.Cgbrating, c.Cgbcustservices, c.Cgbcustlivedate, c.Cgbconnectiondat, c.Cgbcustamount, c.Cgbcustclaw, c.Cgbdoubleclaw, c.Cgbcustclawperio, c.Cgboldnew, c.Cgbftcgb, c.Cgbsparec1, c.Cgbsparec2, c.Cgbbcamount, c.Cgbspared1, c.Cgbtbbpartner, c.Cgbtbbamount, c.Cgbordownload, c.Cgborinstantsim, c.Cgbobngn, c.Cgbtotal, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cgbid, c.Cgbcustaccountno, c.Cgbbillinggroup, c.Cgbcustexecid, c.Cgbcli, c.Cgbperiod, c.Cgbspamount, c.Cgbfpamount, c.Cgbmpamount, c.Cgbgpamount, c.Cgbelecamount, c.Cgbgasamount, c.Cgbdoubleamount, c.Cgbrating, c.Cgbcustservices, c.Cgbcustlivedate, c.Cgbconnectiondat, c.Cgbcustamount, c.Cgbcustclaw, c.Cgbdoubleclaw, c.Cgbcustclawperio, c.Cgboldnew, c.Cgbftcgb, c.Cgbsparec1, c.Cgbsparec2, c.Cgbbcamount, c.Cgbspared1, c.Cgbtbbpartner, c.Cgbtbbamount, c.Cgbordownload, c.Cgborinstantsim, c.Cgbobngn, c.Cgbtotal, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cgb in the database.
func (c *Cgb) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cgb SET (` +
		`cgbid, cgbcustaccountno, cgbbillinggroup, cgbcustexecid, cgbcli, cgbperiod, cgbspamount, cgbfpamount, cgbmpamount, cgbgpamount, cgbelecamount, cgbgasamount, cgbdoubleamount, cgbrating, cgbcustservices, cgbcustlivedate, cgbconnectiondat, cgbcustamount, cgbcustclaw, cgbdoubleclaw, cgbcustclawperio, cgboldnew, cgbftcgb, cgbsparec1, cgbsparec2, cgbbcamount, cgbspared1, cgbtbbpartner, cgbtbbamount, cgbordownload, cgborinstantsim, cgbobngn, cgbtotal, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34` +
		`) WHERE equinox_lrn = $35`

	// run query
	XOLog(sqlstr, c.Cgbid, c.Cgbcustaccountno, c.Cgbbillinggroup, c.Cgbcustexecid, c.Cgbcli, c.Cgbperiod, c.Cgbspamount, c.Cgbfpamount, c.Cgbmpamount, c.Cgbgpamount, c.Cgbelecamount, c.Cgbgasamount, c.Cgbdoubleamount, c.Cgbrating, c.Cgbcustservices, c.Cgbcustlivedate, c.Cgbconnectiondat, c.Cgbcustamount, c.Cgbcustclaw, c.Cgbdoubleclaw, c.Cgbcustclawperio, c.Cgboldnew, c.Cgbftcgb, c.Cgbsparec1, c.Cgbsparec2, c.Cgbbcamount, c.Cgbspared1, c.Cgbtbbpartner, c.Cgbtbbamount, c.Cgbordownload, c.Cgborinstantsim, c.Cgbobngn, c.Cgbtotal, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cgbid, c.Cgbcustaccountno, c.Cgbbillinggroup, c.Cgbcustexecid, c.Cgbcli, c.Cgbperiod, c.Cgbspamount, c.Cgbfpamount, c.Cgbmpamount, c.Cgbgpamount, c.Cgbelecamount, c.Cgbgasamount, c.Cgbdoubleamount, c.Cgbrating, c.Cgbcustservices, c.Cgbcustlivedate, c.Cgbconnectiondat, c.Cgbcustamount, c.Cgbcustclaw, c.Cgbdoubleclaw, c.Cgbcustclawperio, c.Cgboldnew, c.Cgbftcgb, c.Cgbsparec1, c.Cgbsparec2, c.Cgbbcamount, c.Cgbspared1, c.Cgbtbbpartner, c.Cgbtbbamount, c.Cgbordownload, c.Cgborinstantsim, c.Cgbobngn, c.Cgbtotal, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cgb to the database.
func (c *Cgb) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cgb.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cgb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cgb (` +
		`cgbid, cgbcustaccountno, cgbbillinggroup, cgbcustexecid, cgbcli, cgbperiod, cgbspamount, cgbfpamount, cgbmpamount, cgbgpamount, cgbelecamount, cgbgasamount, cgbdoubleamount, cgbrating, cgbcustservices, cgbcustlivedate, cgbconnectiondat, cgbcustamount, cgbcustclaw, cgbdoubleclaw, cgbcustclawperio, cgboldnew, cgbftcgb, cgbsparec1, cgbsparec2, cgbbcamount, cgbspared1, cgbtbbpartner, cgbtbbamount, cgbordownload, cgborinstantsim, cgbobngn, cgbtotal, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cgbid, cgbcustaccountno, cgbbillinggroup, cgbcustexecid, cgbcli, cgbperiod, cgbspamount, cgbfpamount, cgbmpamount, cgbgpamount, cgbelecamount, cgbgasamount, cgbdoubleamount, cgbrating, cgbcustservices, cgbcustlivedate, cgbconnectiondat, cgbcustamount, cgbcustclaw, cgbdoubleclaw, cgbcustclawperio, cgboldnew, cgbftcgb, cgbsparec1, cgbsparec2, cgbbcamount, cgbspared1, cgbtbbpartner, cgbtbbamount, cgbordownload, cgborinstantsim, cgbobngn, cgbtotal, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cgbid, EXCLUDED.cgbcustaccountno, EXCLUDED.cgbbillinggroup, EXCLUDED.cgbcustexecid, EXCLUDED.cgbcli, EXCLUDED.cgbperiod, EXCLUDED.cgbspamount, EXCLUDED.cgbfpamount, EXCLUDED.cgbmpamount, EXCLUDED.cgbgpamount, EXCLUDED.cgbelecamount, EXCLUDED.cgbgasamount, EXCLUDED.cgbdoubleamount, EXCLUDED.cgbrating, EXCLUDED.cgbcustservices, EXCLUDED.cgbcustlivedate, EXCLUDED.cgbconnectiondat, EXCLUDED.cgbcustamount, EXCLUDED.cgbcustclaw, EXCLUDED.cgbdoubleclaw, EXCLUDED.cgbcustclawperio, EXCLUDED.cgboldnew, EXCLUDED.cgbftcgb, EXCLUDED.cgbsparec1, EXCLUDED.cgbsparec2, EXCLUDED.cgbbcamount, EXCLUDED.cgbspared1, EXCLUDED.cgbtbbpartner, EXCLUDED.cgbtbbamount, EXCLUDED.cgbordownload, EXCLUDED.cgborinstantsim, EXCLUDED.cgbobngn, EXCLUDED.cgbtotal, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cgbid, c.Cgbcustaccountno, c.Cgbbillinggroup, c.Cgbcustexecid, c.Cgbcli, c.Cgbperiod, c.Cgbspamount, c.Cgbfpamount, c.Cgbmpamount, c.Cgbgpamount, c.Cgbelecamount, c.Cgbgasamount, c.Cgbdoubleamount, c.Cgbrating, c.Cgbcustservices, c.Cgbcustlivedate, c.Cgbconnectiondat, c.Cgbcustamount, c.Cgbcustclaw, c.Cgbdoubleclaw, c.Cgbcustclawperio, c.Cgboldnew, c.Cgbftcgb, c.Cgbsparec1, c.Cgbsparec2, c.Cgbbcamount, c.Cgbspared1, c.Cgbtbbpartner, c.Cgbtbbamount, c.Cgbordownload, c.Cgborinstantsim, c.Cgbobngn, c.Cgbtotal, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cgbid, c.Cgbcustaccountno, c.Cgbbillinggroup, c.Cgbcustexecid, c.Cgbcli, c.Cgbperiod, c.Cgbspamount, c.Cgbfpamount, c.Cgbmpamount, c.Cgbgpamount, c.Cgbelecamount, c.Cgbgasamount, c.Cgbdoubleamount, c.Cgbrating, c.Cgbcustservices, c.Cgbcustlivedate, c.Cgbconnectiondat, c.Cgbcustamount, c.Cgbcustclaw, c.Cgbdoubleclaw, c.Cgbcustclawperio, c.Cgboldnew, c.Cgbftcgb, c.Cgbsparec1, c.Cgbsparec2, c.Cgbbcamount, c.Cgbspared1, c.Cgbtbbpartner, c.Cgbtbbamount, c.Cgbordownload, c.Cgborinstantsim, c.Cgbobngn, c.Cgbtotal, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cgb from the database.
func (c *Cgb) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cgb WHERE equinox_lrn = $1`

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

// CgbByEquinoxLrn retrieves a row from 'equinox.cgb' as a Cgb.
//
// Generated from index 'cgb_pkey'.
func CgbByEquinoxLrn(db XODB, equinoxLrn int64) (*Cgb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cgbid, cgbcustaccountno, cgbbillinggroup, cgbcustexecid, cgbcli, cgbperiod, cgbspamount, cgbfpamount, cgbmpamount, cgbgpamount, cgbelecamount, cgbgasamount, cgbdoubleamount, cgbrating, cgbcustservices, cgbcustlivedate, cgbconnectiondat, cgbcustamount, cgbcustclaw, cgbdoubleclaw, cgbcustclawperio, cgboldnew, cgbftcgb, cgbsparec1, cgbsparec2, cgbbcamount, cgbspared1, cgbtbbpartner, cgbtbbamount, cgbordownload, cgborinstantsim, cgbobngn, cgbtotal, equinox_lrn, equinox_sec ` +
		`FROM equinox.cgb ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cgb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cgbid, &c.Cgbcustaccountno, &c.Cgbbillinggroup, &c.Cgbcustexecid, &c.Cgbcli, &c.Cgbperiod, &c.Cgbspamount, &c.Cgbfpamount, &c.Cgbmpamount, &c.Cgbgpamount, &c.Cgbelecamount, &c.Cgbgasamount, &c.Cgbdoubleamount, &c.Cgbrating, &c.Cgbcustservices, &c.Cgbcustlivedate, &c.Cgbconnectiondat, &c.Cgbcustamount, &c.Cgbcustclaw, &c.Cgbdoubleclaw, &c.Cgbcustclawperio, &c.Cgboldnew, &c.Cgbftcgb, &c.Cgbsparec1, &c.Cgbsparec2, &c.Cgbbcamount, &c.Cgbspared1, &c.Cgbtbbpartner, &c.Cgbtbbamount, &c.Cgbordownload, &c.Cgborinstantsim, &c.Cgbobngn, &c.Cgbtotal, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}