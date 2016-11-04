// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Campaign represents a row from 'equinox.campaign'.
type Campaign struct {
	CampCode        sql.NullString `json:"camp_code"`        // camp_code
	CampDescription sql.NullString `json:"camp_description"` // camp_description
	CampStart       pq.NullTime    `json:"camp_start"`       // camp_start
	CampClose       pq.NullTime    `json:"camp_close"`       // camp_close
	CampVolumeSent  sql.NullInt64  `json:"camp_volume_sent"` // camp_volume_sent
	CampResponses   sql.NullInt64  `json:"camp_responses"`   // camp_responses
	CampOrders      sql.NullInt64  `json:"camp_orders"`      // camp_orders
	CampNotes       sql.NullInt64  `json:"camp_notes"`       // camp_notes
	CampRmoption    sql.NullString `json:"camp_rmoption"`    // camp_rmoption
	CampRmduration  sql.NullInt64  `json:"camp_rmduration"`  // camp_rmduration
	CampType        sql.NullInt64  `json:"camp_type"`        // camp_type
	CampCliservlvl  sql.NullString `json:"camp_cliservlvl"`  // camp_cliservlvl
	CampSpareC2     sql.NullString `json:"camp_spare_c2"`    // camp_spare_c2
	CampSpareN1     sql.NullInt64  `json:"camp_spare_n1"`    // camp_spare_n1
	CampSpareN2     sql.NullInt64  `json:"camp_spare_n2"`    // camp_spare_n2
	CampSpareD1     pq.NullTime    `json:"camp_spare_d1"`    // camp_spare_d1
	CampSpareD2     pq.NullTime    `json:"camp_spare_d2"`    // camp_spare_d2
	EquinoxLrn      int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Campaign exists in the database.
func (c *Campaign) Exists() bool {
	return c._exists
}

// Deleted provides information if the Campaign has been deleted from the database.
func (c *Campaign) Deleted() bool {
	return c._deleted
}

// Insert inserts the Campaign to the database.
func (c *Campaign) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.campaign (` +
		`camp_code, camp_description, camp_start, camp_close, camp_volume_sent, camp_responses, camp_orders, camp_notes, camp_rmoption, camp_rmduration, camp_type, camp_cliservlvl, camp_spare_c2, camp_spare_n1, camp_spare_n2, camp_spare_d1, camp_spare_d2, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.CampCode, c.CampDescription, c.CampStart, c.CampClose, c.CampVolumeSent, c.CampResponses, c.CampOrders, c.CampNotes, c.CampRmoption, c.CampRmduration, c.CampType, c.CampCliservlvl, c.CampSpareC2, c.CampSpareN1, c.CampSpareN2, c.CampSpareD1, c.CampSpareD2, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.CampCode, c.CampDescription, c.CampStart, c.CampClose, c.CampVolumeSent, c.CampResponses, c.CampOrders, c.CampNotes, c.CampRmoption, c.CampRmduration, c.CampType, c.CampCliservlvl, c.CampSpareC2, c.CampSpareN1, c.CampSpareN2, c.CampSpareD1, c.CampSpareD2, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Campaign in the database.
func (c *Campaign) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.campaign SET (` +
		`camp_code, camp_description, camp_start, camp_close, camp_volume_sent, camp_responses, camp_orders, camp_notes, camp_rmoption, camp_rmduration, camp_type, camp_cliservlvl, camp_spare_c2, camp_spare_n1, camp_spare_n2, camp_spare_d1, camp_spare_d2, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`) WHERE equinox_lrn = $19`

	// run query
	XOLog(sqlstr, c.CampCode, c.CampDescription, c.CampStart, c.CampClose, c.CampVolumeSent, c.CampResponses, c.CampOrders, c.CampNotes, c.CampRmoption, c.CampRmduration, c.CampType, c.CampCliservlvl, c.CampSpareC2, c.CampSpareN1, c.CampSpareN2, c.CampSpareD1, c.CampSpareD2, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.CampCode, c.CampDescription, c.CampStart, c.CampClose, c.CampVolumeSent, c.CampResponses, c.CampOrders, c.CampNotes, c.CampRmoption, c.CampRmduration, c.CampType, c.CampCliservlvl, c.CampSpareC2, c.CampSpareN1, c.CampSpareN2, c.CampSpareD1, c.CampSpareD2, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Campaign to the database.
func (c *Campaign) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Campaign.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Campaign) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.campaign (` +
		`camp_code, camp_description, camp_start, camp_close, camp_volume_sent, camp_responses, camp_orders, camp_notes, camp_rmoption, camp_rmduration, camp_type, camp_cliservlvl, camp_spare_c2, camp_spare_n1, camp_spare_n2, camp_spare_d1, camp_spare_d2, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`camp_code, camp_description, camp_start, camp_close, camp_volume_sent, camp_responses, camp_orders, camp_notes, camp_rmoption, camp_rmduration, camp_type, camp_cliservlvl, camp_spare_c2, camp_spare_n1, camp_spare_n2, camp_spare_d1, camp_spare_d2, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.camp_code, EXCLUDED.camp_description, EXCLUDED.camp_start, EXCLUDED.camp_close, EXCLUDED.camp_volume_sent, EXCLUDED.camp_responses, EXCLUDED.camp_orders, EXCLUDED.camp_notes, EXCLUDED.camp_rmoption, EXCLUDED.camp_rmduration, EXCLUDED.camp_type, EXCLUDED.camp_cliservlvl, EXCLUDED.camp_spare_c2, EXCLUDED.camp_spare_n1, EXCLUDED.camp_spare_n2, EXCLUDED.camp_spare_d1, EXCLUDED.camp_spare_d2, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.CampCode, c.CampDescription, c.CampStart, c.CampClose, c.CampVolumeSent, c.CampResponses, c.CampOrders, c.CampNotes, c.CampRmoption, c.CampRmduration, c.CampType, c.CampCliservlvl, c.CampSpareC2, c.CampSpareN1, c.CampSpareN2, c.CampSpareD1, c.CampSpareD2, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.CampCode, c.CampDescription, c.CampStart, c.CampClose, c.CampVolumeSent, c.CampResponses, c.CampOrders, c.CampNotes, c.CampRmoption, c.CampRmduration, c.CampType, c.CampCliservlvl, c.CampSpareC2, c.CampSpareN1, c.CampSpareN2, c.CampSpareD1, c.CampSpareD2, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Campaign from the database.
func (c *Campaign) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.campaign WHERE equinox_lrn = $1`

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

// CampaignByEquinoxLrn retrieves a row from 'equinox.campaign' as a Campaign.
//
// Generated from index 'campaign_pkey'.
func CampaignByEquinoxLrn(db XODB, equinoxLrn int64) (*Campaign, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`camp_code, camp_description, camp_start, camp_close, camp_volume_sent, camp_responses, camp_orders, camp_notes, camp_rmoption, camp_rmduration, camp_type, camp_cliservlvl, camp_spare_c2, camp_spare_n1, camp_spare_n2, camp_spare_d1, camp_spare_d2, equinox_lrn, equinox_sec ` +
		`FROM equinox.campaign ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Campaign{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.CampCode, &c.CampDescription, &c.CampStart, &c.CampClose, &c.CampVolumeSent, &c.CampResponses, &c.CampOrders, &c.CampNotes, &c.CampRmoption, &c.CampRmduration, &c.CampType, &c.CampCliservlvl, &c.CampSpareC2, &c.CampSpareN1, &c.CampSpareN2, &c.CampSpareD1, &c.CampSpareD2, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}