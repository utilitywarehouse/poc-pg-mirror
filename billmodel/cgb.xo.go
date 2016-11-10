// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

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
	c := Cgb{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cgbid, &c.Cgbcustaccountno, &c.Cgbbillinggroup, &c.Cgbcustexecid, &c.Cgbcli, &c.Cgbperiod, &c.Cgbspamount, &c.Cgbfpamount, &c.Cgbmpamount, &c.Cgbgpamount, &c.Cgbelecamount, &c.Cgbgasamount, &c.Cgbdoubleamount, &c.Cgbrating, &c.Cgbcustservices, &c.Cgbcustlivedate, &c.Cgbconnectiondat, &c.Cgbcustamount, &c.Cgbcustclaw, &c.Cgbdoubleclaw, &c.Cgbcustclawperio, &c.Cgboldnew, &c.Cgbftcgb, &c.Cgbsparec1, &c.Cgbsparec2, &c.Cgbbcamount, &c.Cgbspared1, &c.Cgbtbbpartner, &c.Cgbtbbamount, &c.Cgbordownload, &c.Cgborinstantsim, &c.Cgbobngn, &c.Cgbtotal, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
