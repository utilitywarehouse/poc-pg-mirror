// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cnxgp represents a row from 'equinox.cnxgp'.
type Cnxgp struct {
	Cnxgpcli         sql.NullString  `json:"cnxgpcli"`         // cnxgpcli
	Cnxgpboxes       sql.NullInt64   `json:"cnxgpboxes"`       // cnxgpboxes
	Cnxgpfeature     sql.NullInt64   `json:"cnxgpfeature"`     // cnxgpfeature
	Cnxgpcable       sql.NullInt64   `json:"cnxgpcable"`       // cnxgpcable
	Cnxgpni          sql.NullInt64   `json:"cnxgpni"`          // cnxgpni
	Cnxgpservicelvl  sql.NullString  `json:"cnxgpservicelvl"`  // cnxgpservicelvl
	Cnxgpdatetocs    pq.NullTime     `json:"cnxgpdatetocs"`    // cnxgpdatetocs
	Cnxgpdatefromc1  pq.NullTime     `json:"cnxgpdatefromc1"`  // cnxgpdatefromc1
	Cnxgpdatefromc2  pq.NullTime     `json:"cnxgpdatefromc2"`  // cnxgpdatefromc2
	Cnxgpdatetoopt   pq.NullTime     `json:"cnxgpdatetoopt"`   // cnxgpdatetoopt
	Cnxgpdatefromopt pq.NullTime     `json:"cnxgpdatefromopt"` // cnxgpdatefromopt
	Cnxgpstate       sql.NullString  `json:"cnxgpstate"`       // cnxgpstate
	Cnxgptermnum     sql.NullString  `json:"cnxgptermnum"`     // cnxgptermnum
	Cnxgpdateentered pq.NullTime     `json:"cnxgpdateentered"` // cnxgpdateentered
	Cnxgpweborder    sql.NullInt64   `json:"cnxgpweborder"`    // cnxgpweborder
	Cnxgpdatetobill  pq.NullTime     `json:"cnxgpdatetobill"`  // cnxgpdatetobill
	Cnxgpdateletter  pq.NullTime     `json:"cnxgpdateletter"`  // cnxgpdateletter
	Cnxgpdateprepay  pq.NullTime     `json:"cnxgpdateprepay"`  // cnxgpdateprepay
	Cnxgpenteredby   sql.NullString  `json:"cnxgpenteredby"`   // cnxgpenteredby
	Cnxgppromocd     sql.NullString  `json:"cnxgppromocd"`     // cnxgppromocd
	Cnxgppackageno   sql.NullString  `json:"cnxgppackageno"`   // cnxgppackageno
	Cnxgpsubtariff   sql.NullString  `json:"cnxgpsubtariff"`   // cnxgpsubtariff
	Cnxgpcpspostcode sql.NullString  `json:"cnxgpcpspostcode"` // cnxgpcpspostcode
	Cnxgpextras      sql.NullString  `json:"cnxgpextras"`      // cnxgpextras
	Cnxgpmaccode     sql.NullString  `json:"cnxgpmaccode"`     // cnxgpmaccode
	Cnxgpordtype     sql.NullString  `json:"cnxgpordtype"`     // cnxgpordtype
	Cnxgptariff      sql.NullString  `json:"cnxgptariff"`      // cnxgptariff
	Cnxgpfilters     sql.NullFloat64 `json:"cnxgpfilters"`     // cnxgpfilters
	Cnxgpsparen2     sql.NullFloat64 `json:"cnxgpsparen2"`     // cnxgpsparen2
	Cnxgpsparen3     sql.NullFloat64 `json:"cnxgpsparen3"`     // cnxgpsparen3
	Cnxgpcsigndate   pq.NullTime     `json:"cnxgpcsigndate"`   // cnxgpcsigndate
	Cnxgpagreeddate  pq.NullTime     `json:"cnxgpagreeddate"`  // cnxgpagreeddate
	Cnxgpequipment   sql.NullString  `json:"cnxgpequipment"`   // cnxgpequipment
	Cnxgpnetwork     sql.NullString  `json:"cnxgpnetwork"`     // cnxgpnetwork
	Cnxgpdonoracc    sql.NullString  `json:"cnxgpdonoracc"`    // cnxgpdonoracc
	Cnxgpdeladdr     sql.NullInt64   `json:"cnxgpdeladdr"`     // cnxgpdeladdr
	Cnxgpdeladdref   sql.NullInt64   `json:"cnxgpdeladdref"`   // cnxgpdeladdref
	Cnxgpdoncliuni   sql.NullFloat64 `json:"cnxgpdoncliuni"`   // cnxgpdoncliuni
	Cnxgpholduntil   pq.NullTime     `json:"cnxgpholduntil"`   // cnxgpholduntil
	Cnxgphmdate      pq.NullTime     `json:"cnxgphmdate"`      // cnxgphmdate
	Cnxgpaddserv     sql.NullInt64   `json:"cnxgpaddserv"`     // cnxgpaddserv
	Cnxgpsparec1     sql.NullString  `json:"cnxgpsparec1"`     // cnxgpsparec1
	Cnxgpnousbadap   sql.NullInt64   `json:"cnxgpnousbadap"`   // cnxgpnousbadap
	Cnxgpvaliddate   pq.NullTime     `json:"cnxgpvaliddate"`   // cnxgpvaliddate
	Cnxgpvalidby     sql.NullString  `json:"cnxgpvalidby"`     // cnxgpvalidby
	Cnxgpcampaign    sql.NullString  `json:"cnxgpcampaign"`    // cnxgpcampaign
	Cnxgptype        sql.NullString  `json:"cnxgptype"`        // cnxgptype
	Cnxgpcommandid   sql.NullString  `json:"cnxgpcommandid"`   // cnxgpcommandid
	Cnxgpreqdate     pq.NullTime     `json:"cnxgpreqdate"`     // cnxgpreqdate
	Cnxgpstatedate   pq.NullTime     `json:"cnxgpstatedate"`   // cnxgpstatedate
	Cnxgppassword    sql.NullString  `json:"cnxgppassword"`    // cnxgppassword
	Cnxgpcancelrid   sql.NullString  `json:"cnxgpcancelrid"`   // cnxgpcancelrid
	Cnxgpgainingrid  sql.NullString  `json:"cnxgpgainingrid"`  // cnxgpgainingrid
	Cnxgpcpwnref     sql.NullString  `json:"cnxgpcpwnref"`     // cnxgpcpwnref
	Cnxgplorn        sql.NullString  `json:"cnxgplorn"`        // cnxgplorn
	Cnxgpsimprovord  sql.NullString  `json:"cnxgpsimprovord"`  // cnxgpsimprovord
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// CnxgpByEquinoxLrn retrieves a row from 'equinox.cnxgp' as a Cnxgp.
//
// Generated from index 'cnxgp_pkey'.
func CnxgpByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxgp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxgpcli, cnxgpboxes, cnxgpfeature, cnxgpcable, cnxgpni, cnxgpservicelvl, cnxgpdatetocs, cnxgpdatefromc1, cnxgpdatefromc2, cnxgpdatetoopt, cnxgpdatefromopt, cnxgpstate, cnxgptermnum, cnxgpdateentered, cnxgpweborder, cnxgpdatetobill, cnxgpdateletter, cnxgpdateprepay, cnxgpenteredby, cnxgppromocd, cnxgppackageno, cnxgpsubtariff, cnxgpcpspostcode, cnxgpextras, cnxgpmaccode, cnxgpordtype, cnxgptariff, cnxgpfilters, cnxgpsparen2, cnxgpsparen3, cnxgpcsigndate, cnxgpagreeddate, cnxgpequipment, cnxgpnetwork, cnxgpdonoracc, cnxgpdeladdr, cnxgpdeladdref, cnxgpdoncliuni, cnxgpholduntil, cnxgphmdate, cnxgpaddserv, cnxgpsparec1, cnxgpnousbadap, cnxgpvaliddate, cnxgpvalidby, cnxgpcampaign, cnxgptype, cnxgpcommandid, cnxgpreqdate, cnxgpstatedate, cnxgppassword, cnxgpcancelrid, cnxgpgainingrid, cnxgpcpwnref, cnxgplorn, cnxgpsimprovord, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxgp ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxgp{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxgpcli, &c.Cnxgpboxes, &c.Cnxgpfeature, &c.Cnxgpcable, &c.Cnxgpni, &c.Cnxgpservicelvl, &c.Cnxgpdatetocs, &c.Cnxgpdatefromc1, &c.Cnxgpdatefromc2, &c.Cnxgpdatetoopt, &c.Cnxgpdatefromopt, &c.Cnxgpstate, &c.Cnxgptermnum, &c.Cnxgpdateentered, &c.Cnxgpweborder, &c.Cnxgpdatetobill, &c.Cnxgpdateletter, &c.Cnxgpdateprepay, &c.Cnxgpenteredby, &c.Cnxgppromocd, &c.Cnxgppackageno, &c.Cnxgpsubtariff, &c.Cnxgpcpspostcode, &c.Cnxgpextras, &c.Cnxgpmaccode, &c.Cnxgpordtype, &c.Cnxgptariff, &c.Cnxgpfilters, &c.Cnxgpsparen2, &c.Cnxgpsparen3, &c.Cnxgpcsigndate, &c.Cnxgpagreeddate, &c.Cnxgpequipment, &c.Cnxgpnetwork, &c.Cnxgpdonoracc, &c.Cnxgpdeladdr, &c.Cnxgpdeladdref, &c.Cnxgpdoncliuni, &c.Cnxgpholduntil, &c.Cnxgphmdate, &c.Cnxgpaddserv, &c.Cnxgpsparec1, &c.Cnxgpnousbadap, &c.Cnxgpvaliddate, &c.Cnxgpvalidby, &c.Cnxgpcampaign, &c.Cnxgptype, &c.Cnxgpcommandid, &c.Cnxgpreqdate, &c.Cnxgpstatedate, &c.Cnxgppassword, &c.Cnxgpcancelrid, &c.Cnxgpgainingrid, &c.Cnxgpcpwnref, &c.Cnxgplorn, &c.Cnxgpsimprovord, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
