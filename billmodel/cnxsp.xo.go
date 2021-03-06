// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cnxsp represents a row from 'equinox.cnxsp'.
type Cnxsp struct {
	Cnxspcli         sql.NullString  `json:"cnxspcli"`         // cnxspcli
	Cnxspboxes       sql.NullInt64   `json:"cnxspboxes"`       // cnxspboxes
	Cnxspfeature     sql.NullInt64   `json:"cnxspfeature"`     // cnxspfeature
	Cnxspcable       sql.NullInt64   `json:"cnxspcable"`       // cnxspcable
	Cnxspni          sql.NullInt64   `json:"cnxspni"`          // cnxspni
	Cnxspservicelvl  sql.NullString  `json:"cnxspservicelvl"`  // cnxspservicelvl
	Cnxspdatetocs    pq.NullTime     `json:"cnxspdatetocs"`    // cnxspdatetocs
	Cnxspdatefromc1  pq.NullTime     `json:"cnxspdatefromc1"`  // cnxspdatefromc1
	Cnxspdatefromc2  pq.NullTime     `json:"cnxspdatefromc2"`  // cnxspdatefromc2
	Cnxspdatetoopt   pq.NullTime     `json:"cnxspdatetoopt"`   // cnxspdatetoopt
	Cnxspdatefromopt pq.NullTime     `json:"cnxspdatefromopt"` // cnxspdatefromopt
	Cnxspoptordnum   sql.NullString  `json:"cnxspoptordnum"`   // cnxspoptordnum
	Cnxspserialnum   sql.NullString  `json:"cnxspserialnum"`   // cnxspserialnum
	Cnxspdateentered pq.NullTime     `json:"cnxspdateentered"` // cnxspdateentered
	Cnxspweborder    sql.NullInt64   `json:"cnxspweborder"`    // cnxspweborder
	Cnxspdatetobill  pq.NullTime     `json:"cnxspdatetobill"`  // cnxspdatetobill
	Cnxspdateletter  pq.NullTime     `json:"cnxspdateletter"`  // cnxspdateletter
	Cnxspprepayment  sql.NullFloat64 `json:"cnxspprepayment"`  // cnxspprepayment
	Cnxspdateprepay  pq.NullTime     `json:"cnxspdateprepay"`  // cnxspdateprepay
	Cnxspenteredby   sql.NullString  `json:"cnxspenteredby"`   // cnxspenteredby
	Cnxspfreebox     sql.NullInt64   `json:"cnxspfreebox"`     // cnxspfreebox
	Cnxsppromocd     sql.NullString  `json:"cnxsppromocd"`     // cnxsppromocd
	Cnxspprice       sql.NullFloat64 `json:"cnxspprice"`       // cnxspprice
	Cnxsppackageno   sql.NullString  `json:"cnxsppackageno"`   // cnxsppackageno
	Cnxspsubtarrif   sql.NullString  `json:"cnxspsubtarrif"`   // cnxspsubtarrif
	Cnxspcpspostcode sql.NullString  `json:"cnxspcpspostcode"` // cnxspcpspostcode
	Cnxspdatecps     pq.NullTime     `json:"cnxspdatecps"`     // cnxspdatecps
	Cnxspcpsreqd     sql.NullInt64   `json:"cnxspcpsreqd"`     // cnxspcpsreqd
	Cnxspoptimalplan sql.NullInt64   `json:"cnxspoptimalplan"` // cnxspoptimalplan
	Cnxspcpsstrikes  sql.NullInt64   `json:"cnxspcpsstrikes"`  // cnxspcpsstrikes
	Cnxspcpsdstrke1  pq.NullTime     `json:"cnxspcpsdstrke1"`  // cnxspcpsdstrke1
	Cnxspcpsdstrke2  pq.NullTime     `json:"cnxspcpsdstrke2"`  // cnxspcpsdstrke2
	Cnxspcpsdstrke3  pq.NullTime     `json:"cnxspcpsdstrke3"`  // cnxspcpsdstrke3
	Cnxspcpscarrier  sql.NullString  `json:"cnxspcpscarrier"`  // cnxspcpscarrier
	Cnxspextras      sql.NullString  `json:"cnxspextras"`      // cnxspextras
	Cnxspwlrtype     sql.NullString  `json:"cnxspwlrtype"`     // cnxspwlrtype
	Cnxspcpscancel   pq.NullTime     `json:"cnxspcpscancel"`   // cnxspcpscancel
	Cnxspmaccode     sql.NullString  `json:"cnxspmaccode"`     // cnxspmaccode
	Cnxspordtype     sql.NullString  `json:"cnxspordtype"`     // cnxspordtype
	Cnxsptariff      sql.NullString  `json:"cnxsptariff"`      // cnxsptariff
	Cnxspsparen1     sql.NullFloat64 `json:"cnxspsparen1"`     // cnxspsparen1
	Cnxspsparen2     sql.NullFloat64 `json:"cnxspsparen2"`     // cnxspsparen2
	Cnxspcontterm    sql.NullInt64   `json:"cnxspcontterm"`    // cnxspcontterm
	Cnxspcsigndate   pq.NullTime     `json:"cnxspcsigndate"`   // cnxspcsigndate
	Cnxspspared3     pq.NullTime     `json:"cnxspspared3"`     // cnxspspared3
	Cnxspcnf         sql.NullString  `json:"cnxspcnf"`         // cnxspcnf
	Cnxspdonoracc    sql.NullString  `json:"cnxspdonoracc"`    // cnxspdonoracc
	Cnxspdoncliuni   sql.NullFloat64 `json:"cnxspdoncliuni"`   // cnxspdoncliuni
	Cnxspdiscband    sql.NullString  `json:"cnxspdiscband"`    // cnxspdiscband
	Cnxspholduntil   pq.NullTime     `json:"cnxspholduntil"`   // cnxspholduntil
	Cnxspbtaccountno sql.NullString  `json:"cnxspbtaccountno"` // cnxspbtaccountno
	Cnxsptps         sql.NullInt64   `json:"cnxsptps"`         // cnxsptps
	Cnxsphmdate      pq.NullTime     `json:"cnxsphmdate"`      // cnxsphmdate
	Cnxspbttermend   pq.NullTime     `json:"cnxspbttermend"`   // cnxspbttermend
	Cnxsplinewithbt  sql.NullInt64   `json:"cnxsplinewithbt"`  // cnxsplinewithbt
	Cnxspwelclett    sql.NullInt64   `json:"cnxspwelclett"`    // cnxspwelclett
	Cnxspaddserv     sql.NullInt64   `json:"cnxspaddserv"`     // cnxspaddserv
	Cnxspcallbundle  sql.NullString  `json:"cnxspcallbundle"`  // cnxspcallbundle
	Cnxspvaiddate    pq.NullTime     `json:"cnxspvaiddate"`    // cnxspvaiddate
	Cnxspvalidby     sql.NullString  `json:"cnxspvalidby"`     // cnxspvalidby
	Cnxspcampaign    sql.NullString  `json:"cnxspcampaign"`    // cnxspcampaign
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCnxsp(db XODB, callback func(x Cnxsp) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cnxspcli, cnxspboxes, cnxspfeature, cnxspcable, cnxspni, cnxspservicelvl, cnxspdatetocs, cnxspdatefromc1, cnxspdatefromc2, cnxspdatetoopt, cnxspdatefromopt, cnxspoptordnum, cnxspserialnum, cnxspdateentered, cnxspweborder, cnxspdatetobill, cnxspdateletter, cnxspprepayment, cnxspdateprepay, cnxspenteredby, cnxspfreebox, cnxsppromocd, cnxspprice, cnxsppackageno, cnxspsubtarrif, cnxspcpspostcode, cnxspdatecps, cnxspcpsreqd, cnxspoptimalplan, cnxspcpsstrikes, cnxspcpsdstrke1, cnxspcpsdstrke2, cnxspcpsdstrke3, cnxspcpscarrier, cnxspextras, cnxspwlrtype, cnxspcpscancel, cnxspmaccode, cnxspordtype, cnxsptariff, cnxspsparen1, cnxspsparen2, cnxspcontterm, cnxspcsigndate, cnxspspared3, cnxspcnf, cnxspdonoracc, cnxspdoncliuni, cnxspdiscband, cnxspholduntil, cnxspbtaccountno, cnxsptps, cnxsphmdate, cnxspbttermend, cnxsplinewithbt, cnxspwelclett, cnxspaddserv, cnxspcallbundle, cnxspvaiddate, cnxspvalidby, cnxspcampaign, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxsp `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cnxsp{}

		// scan
		err = q.Scan(&c.Cnxspcli, &c.Cnxspboxes, &c.Cnxspfeature, &c.Cnxspcable, &c.Cnxspni, &c.Cnxspservicelvl, &c.Cnxspdatetocs, &c.Cnxspdatefromc1, &c.Cnxspdatefromc2, &c.Cnxspdatetoopt, &c.Cnxspdatefromopt, &c.Cnxspoptordnum, &c.Cnxspserialnum, &c.Cnxspdateentered, &c.Cnxspweborder, &c.Cnxspdatetobill, &c.Cnxspdateletter, &c.Cnxspprepayment, &c.Cnxspdateprepay, &c.Cnxspenteredby, &c.Cnxspfreebox, &c.Cnxsppromocd, &c.Cnxspprice, &c.Cnxsppackageno, &c.Cnxspsubtarrif, &c.Cnxspcpspostcode, &c.Cnxspdatecps, &c.Cnxspcpsreqd, &c.Cnxspoptimalplan, &c.Cnxspcpsstrikes, &c.Cnxspcpsdstrke1, &c.Cnxspcpsdstrke2, &c.Cnxspcpsdstrke3, &c.Cnxspcpscarrier, &c.Cnxspextras, &c.Cnxspwlrtype, &c.Cnxspcpscancel, &c.Cnxspmaccode, &c.Cnxspordtype, &c.Cnxsptariff, &c.Cnxspsparen1, &c.Cnxspsparen2, &c.Cnxspcontterm, &c.Cnxspcsigndate, &c.Cnxspspared3, &c.Cnxspcnf, &c.Cnxspdonoracc, &c.Cnxspdoncliuni, &c.Cnxspdiscband, &c.Cnxspholduntil, &c.Cnxspbtaccountno, &c.Cnxsptps, &c.Cnxsphmdate, &c.Cnxspbttermend, &c.Cnxsplinewithbt, &c.Cnxspwelclett, &c.Cnxspaddserv, &c.Cnxspcallbundle, &c.Cnxspvaiddate, &c.Cnxspvalidby, &c.Cnxspcampaign, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CnxspByEquinoxLrn retrieves a row from 'equinox.cnxsp' as a Cnxsp.
//
// Generated from index 'cnxsp_pkey'.
func CnxspByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxsp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxspcli, cnxspboxes, cnxspfeature, cnxspcable, cnxspni, cnxspservicelvl, cnxspdatetocs, cnxspdatefromc1, cnxspdatefromc2, cnxspdatetoopt, cnxspdatefromopt, cnxspoptordnum, cnxspserialnum, cnxspdateentered, cnxspweborder, cnxspdatetobill, cnxspdateletter, cnxspprepayment, cnxspdateprepay, cnxspenteredby, cnxspfreebox, cnxsppromocd, cnxspprice, cnxsppackageno, cnxspsubtarrif, cnxspcpspostcode, cnxspdatecps, cnxspcpsreqd, cnxspoptimalplan, cnxspcpsstrikes, cnxspcpsdstrke1, cnxspcpsdstrke2, cnxspcpsdstrke3, cnxspcpscarrier, cnxspextras, cnxspwlrtype, cnxspcpscancel, cnxspmaccode, cnxspordtype, cnxsptariff, cnxspsparen1, cnxspsparen2, cnxspcontterm, cnxspcsigndate, cnxspspared3, cnxspcnf, cnxspdonoracc, cnxspdoncliuni, cnxspdiscband, cnxspholduntil, cnxspbtaccountno, cnxsptps, cnxsphmdate, cnxspbttermend, cnxsplinewithbt, cnxspwelclett, cnxspaddserv, cnxspcallbundle, cnxspvaiddate, cnxspvalidby, cnxspcampaign, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxsp ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxsp{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxspcli, &c.Cnxspboxes, &c.Cnxspfeature, &c.Cnxspcable, &c.Cnxspni, &c.Cnxspservicelvl, &c.Cnxspdatetocs, &c.Cnxspdatefromc1, &c.Cnxspdatefromc2, &c.Cnxspdatetoopt, &c.Cnxspdatefromopt, &c.Cnxspoptordnum, &c.Cnxspserialnum, &c.Cnxspdateentered, &c.Cnxspweborder, &c.Cnxspdatetobill, &c.Cnxspdateletter, &c.Cnxspprepayment, &c.Cnxspdateprepay, &c.Cnxspenteredby, &c.Cnxspfreebox, &c.Cnxsppromocd, &c.Cnxspprice, &c.Cnxsppackageno, &c.Cnxspsubtarrif, &c.Cnxspcpspostcode, &c.Cnxspdatecps, &c.Cnxspcpsreqd, &c.Cnxspoptimalplan, &c.Cnxspcpsstrikes, &c.Cnxspcpsdstrke1, &c.Cnxspcpsdstrke2, &c.Cnxspcpsdstrke3, &c.Cnxspcpscarrier, &c.Cnxspextras, &c.Cnxspwlrtype, &c.Cnxspcpscancel, &c.Cnxspmaccode, &c.Cnxspordtype, &c.Cnxsptariff, &c.Cnxspsparen1, &c.Cnxspsparen2, &c.Cnxspcontterm, &c.Cnxspcsigndate, &c.Cnxspspared3, &c.Cnxspcnf, &c.Cnxspdonoracc, &c.Cnxspdoncliuni, &c.Cnxspdiscband, &c.Cnxspholduntil, &c.Cnxspbtaccountno, &c.Cnxsptps, &c.Cnxsphmdate, &c.Cnxspbttermend, &c.Cnxsplinewithbt, &c.Cnxspwelclett, &c.Cnxspaddserv, &c.Cnxspcallbundle, &c.Cnxspvaiddate, &c.Cnxspvalidby, &c.Cnxspcampaign, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
