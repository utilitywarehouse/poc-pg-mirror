// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Cnxip represents a row from 'equinox.cnxip'.
type Cnxip struct {
	Cnxipcli         sql.NullString  `json:"cnxipcli"`         // cnxipcli
	Cnxipcliline1    sql.NullString  `json:"cnxipcliline1"`    // cnxipcliline1
	Cnxipcliline2    sql.NullString  `json:"cnxipcliline2"`    // cnxipcliline2
	Cnxipsecondline  sql.NullInt64   `json:"cnxipsecondline"`  // cnxipsecondline
	Cnxipservtype    sql.NullString  `json:"cnxipservtype"`    // cnxipservtype
	Cnxipsurcharge   sql.NullInt64   `json:"cnxipsurcharge"`   // cnxipsurcharge
	Cnxipentityno    sql.NullString  `json:"cnxipentityno"`    // cnxipentityno
	Cnxipextensionno sql.NullString  `json:"cnxipextensionno"` // cnxipextensionno
	Cnxipextnnol2    sql.NullString  `json:"cnxipextnnol2"`    // cnxipextnnol2
	Cnxipdatetocs    pq.NullTime     `json:"cnxipdatetocs"`    // cnxipdatetocs
	Cnxipdatefromc1  pq.NullTime     `json:"cnxipdatefromc1"`  // cnxipdatefromc1
	Cnxipdatefromc2  pq.NullTime     `json:"cnxipdatefromc2"`  // cnxipdatefromc2
	Cnxipdatetoopt   pq.NullTime     `json:"cnxipdatetoopt"`   // cnxipdatetoopt
	Cnxipdatefromopt pq.NullTime     `json:"cnxipdatefromopt"` // cnxipdatefromopt
	Cnxipdatetobill  pq.NullTime     `json:"cnxipdatetobill"`  // cnxipdatetobill
	Cnxippackageno   sql.NullString  `json:"cnxippackageno"`   // cnxippackageno
	Cnxipmaccode     sql.NullString  `json:"cnxipmaccode"`     // cnxipmaccode
	Cnxipdateentered pq.NullTime     `json:"cnxipdateentered"` // cnxipdateentered
	Cnxipenteredby   sql.NullString  `json:"cnxipenteredby"`   // cnxipenteredby
	Cnxipdateletter  pq.NullTime     `json:"cnxipdateletter"`  // cnxipdateletter
	Cnxipbtpostcode  sql.NullString  `json:"cnxipbtpostcode"`  // cnxipbtpostcode
	Cnxipstdcode     sql.NullString  `json:"cnxipstdcode"`     // cnxipstdcode
	Cnxipordtype     sql.NullString  `json:"cnxipordtype"`     // cnxipordtype
	Cnxipcsigndate   pq.NullTime     `json:"cnxipcsigndate"`   // cnxipcsigndate
	Cnxipdeladdr     sql.NullInt64   `json:"cnxipdeladdr"`     // cnxipdeladdr
	Cnxipdeladdref   sql.NullInt64   `json:"cnxipdeladdref"`   // cnxipdeladdref
	Cnxipmobsaver    sql.NullInt64   `json:"cnxipmobsaver"`    // cnxipmobsaver
	Cnxipintsaver    sql.NullInt64   `json:"cnxipintsaver"`    // cnxipintsaver
	Cnxipequipment   sql.NullString  `json:"cnxipequipment"`   // cnxipequipment
	Cnxipweborder    sql.NullInt64   `json:"cnxipweborder"`    // cnxipweborder
	Cnxipprepayment  sql.NullFloat64 `json:"cnxipprepayment"`  // cnxipprepayment
	Cnxipdateprepay  pq.NullTime     `json:"cnxipdateprepay"`  // cnxipdateprepay
	Cnxipservicelvl  sql.NullString  `json:"cnxipservicelvl"`  // cnxipservicelvl
	Cnxipfreebox     sql.NullInt64   `json:"cnxipfreebox"`     // cnxipfreebox
	Cnxipprice       sql.NullFloat64 `json:"cnxipprice"`       // cnxipprice
	Cnxipsubtarrif   sql.NullString  `json:"cnxipsubtarrif"`   // cnxipsubtarrif
	Cnxipextras      sql.NullString  `json:"cnxipextras"`      // cnxipextras
	Cnxiptariff      sql.NullString  `json:"cnxiptariff"`      // cnxiptariff
	Cnxipcontterm    sql.NullInt64   `json:"cnxipcontterm"`    // cnxipcontterm
	Cnxipdonoracc    sql.NullString  `json:"cnxipdonoracc"`    // cnxipdonoracc
	Cnxipdoncliuni   sql.NullFloat64 `json:"cnxipdoncliuni"`   // cnxipdoncliuni
	Cnxipdiscband    sql.NullString  `json:"cnxipdiscband"`    // cnxipdiscband
	Cnxipholduntil   pq.NullTime     `json:"cnxipholduntil"`   // cnxipholduntil
	Cnxipbtaccountno sql.NullString  `json:"cnxipbtaccountno"` // cnxipbtaccountno
	Cnxiptps         sql.NullInt64   `json:"cnxiptps"`         // cnxiptps
	Cnxiphmdate      pq.NullTime     `json:"cnxiphmdate"`      // cnxiphmdate
	Cnxipbttermend   pq.NullTime     `json:"cnxipbttermend"`   // cnxipbttermend
	Cnxiplinewithbt  sql.NullInt64   `json:"cnxiplinewithbt"`  // cnxiplinewithbt
	Cnxippromocode   sql.NullString  `json:"cnxippromocode"`   // cnxippromocode
	Cnxipss07        sql.NullInt64   `json:"cnxipss07"`        // cnxipss07
	Cnxipss01        sql.NullInt64   `json:"cnxipss01"`        // cnxipss01
	Cnxipss1a        sql.NullInt64   `json:"cnxipss1a"`        // cnxipss1a
	Cnxipss08        sql.NullInt64   `json:"cnxipss08"`        // cnxipss08
	Cnxipss06        sql.NullInt64   `json:"cnxipss06"`        // cnxipss06
	Cnxipss09        sql.NullInt64   `json:"cnxipss09"`        // cnxipss09
	Cnxipss02        sql.NullInt64   `json:"cnxipss02"`        // cnxipss02
	Cnxipss05        sql.NullInt64   `json:"cnxipss05"`        // cnxipss05
	Cnxipss14        sql.NullInt64   `json:"cnxipss14"`        // cnxipss14
	Cnxipnumwithheld sql.NullInt64   `json:"cnxipnumwithheld"` // cnxipnumwithheld
	Cnxipextnpassl1  sql.NullString  `json:"cnxipextnpassl1"`  // cnxipextnpassl1
	Cnxipextnpassl2  sql.NullString  `json:"cnxipextnpassl2"`  // cnxipextnpassl2
	Cnxipvmailcli    sql.NullString  `json:"cnxipvmailcli"`    // cnxipvmailcli
	Cnxipmoretalk    sql.NullString  `json:"cnxipmoretalk"`    // cnxipmoretalk
	Cnxiprpc         sql.NullString  `json:"cnxiprpc"`         // cnxiprpc
	Cnxipspared1     pq.NullTime     `json:"cnxipspared1"`     // cnxipspared1
	Cnxipspared2     pq.NullTime     `json:"cnxipspared2"`     // cnxipspared2
	Cnxipsparen1     sql.NullFloat64 `json:"cnxipsparen1"`     // cnxipsparen1
	Cnxipsparen2     sql.NullFloat64 `json:"cnxipsparen2"`     // cnxipsparen2
	Cnxipaddserv     sql.NullInt64   `json:"cnxipaddserv"`     // cnxipaddserv
	Cnxipvaliddate   pq.NullTime     `json:"cnxipvaliddate"`   // cnxipvaliddate
	Cnxipvalidby     sql.NullString  `json:"cnxipvalidby"`     // cnxipvalidby
	Cnxipcampaign    sql.NullString  `json:"cnxipcampaign"`    // cnxipcampaign
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cnxip exists in the database.
func (c *Cnxip) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cnxip has been deleted from the database.
func (c *Cnxip) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cnxip to the database.
func (c *Cnxip) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cnxip (` +
		`cnxipcli, cnxipcliline1, cnxipcliline2, cnxipsecondline, cnxipservtype, cnxipsurcharge, cnxipentityno, cnxipextensionno, cnxipextnnol2, cnxipdatetocs, cnxipdatefromc1, cnxipdatefromc2, cnxipdatetoopt, cnxipdatefromopt, cnxipdatetobill, cnxippackageno, cnxipmaccode, cnxipdateentered, cnxipenteredby, cnxipdateletter, cnxipbtpostcode, cnxipstdcode, cnxipordtype, cnxipcsigndate, cnxipdeladdr, cnxipdeladdref, cnxipmobsaver, cnxipintsaver, cnxipequipment, cnxipweborder, cnxipprepayment, cnxipdateprepay, cnxipservicelvl, cnxipfreebox, cnxipprice, cnxipsubtarrif, cnxipextras, cnxiptariff, cnxipcontterm, cnxipdonoracc, cnxipdoncliuni, cnxipdiscband, cnxipholduntil, cnxipbtaccountno, cnxiptps, cnxiphmdate, cnxipbttermend, cnxiplinewithbt, cnxippromocode, cnxipss07, cnxipss01, cnxipss1a, cnxipss08, cnxipss06, cnxipss09, cnxipss02, cnxipss05, cnxipss14, cnxipnumwithheld, cnxipextnpassl1, cnxipextnpassl2, cnxipvmailcli, cnxipmoretalk, cnxiprpc, cnxipspared1, cnxipspared2, cnxipsparen1, cnxipsparen2, cnxipaddserv, cnxipvaliddate, cnxipvalidby, cnxipcampaign, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cnxipcli, c.Cnxipcliline1, c.Cnxipcliline2, c.Cnxipsecondline, c.Cnxipservtype, c.Cnxipsurcharge, c.Cnxipentityno, c.Cnxipextensionno, c.Cnxipextnnol2, c.Cnxipdatetocs, c.Cnxipdatefromc1, c.Cnxipdatefromc2, c.Cnxipdatetoopt, c.Cnxipdatefromopt, c.Cnxipdatetobill, c.Cnxippackageno, c.Cnxipmaccode, c.Cnxipdateentered, c.Cnxipenteredby, c.Cnxipdateletter, c.Cnxipbtpostcode, c.Cnxipstdcode, c.Cnxipordtype, c.Cnxipcsigndate, c.Cnxipdeladdr, c.Cnxipdeladdref, c.Cnxipmobsaver, c.Cnxipintsaver, c.Cnxipequipment, c.Cnxipweborder, c.Cnxipprepayment, c.Cnxipdateprepay, c.Cnxipservicelvl, c.Cnxipfreebox, c.Cnxipprice, c.Cnxipsubtarrif, c.Cnxipextras, c.Cnxiptariff, c.Cnxipcontterm, c.Cnxipdonoracc, c.Cnxipdoncliuni, c.Cnxipdiscband, c.Cnxipholduntil, c.Cnxipbtaccountno, c.Cnxiptps, c.Cnxiphmdate, c.Cnxipbttermend, c.Cnxiplinewithbt, c.Cnxippromocode, c.Cnxipss07, c.Cnxipss01, c.Cnxipss1a, c.Cnxipss08, c.Cnxipss06, c.Cnxipss09, c.Cnxipss02, c.Cnxipss05, c.Cnxipss14, c.Cnxipnumwithheld, c.Cnxipextnpassl1, c.Cnxipextnpassl2, c.Cnxipvmailcli, c.Cnxipmoretalk, c.Cnxiprpc, c.Cnxipspared1, c.Cnxipspared2, c.Cnxipsparen1, c.Cnxipsparen2, c.Cnxipaddserv, c.Cnxipvaliddate, c.Cnxipvalidby, c.Cnxipcampaign, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cnxipcli, c.Cnxipcliline1, c.Cnxipcliline2, c.Cnxipsecondline, c.Cnxipservtype, c.Cnxipsurcharge, c.Cnxipentityno, c.Cnxipextensionno, c.Cnxipextnnol2, c.Cnxipdatetocs, c.Cnxipdatefromc1, c.Cnxipdatefromc2, c.Cnxipdatetoopt, c.Cnxipdatefromopt, c.Cnxipdatetobill, c.Cnxippackageno, c.Cnxipmaccode, c.Cnxipdateentered, c.Cnxipenteredby, c.Cnxipdateletter, c.Cnxipbtpostcode, c.Cnxipstdcode, c.Cnxipordtype, c.Cnxipcsigndate, c.Cnxipdeladdr, c.Cnxipdeladdref, c.Cnxipmobsaver, c.Cnxipintsaver, c.Cnxipequipment, c.Cnxipweborder, c.Cnxipprepayment, c.Cnxipdateprepay, c.Cnxipservicelvl, c.Cnxipfreebox, c.Cnxipprice, c.Cnxipsubtarrif, c.Cnxipextras, c.Cnxiptariff, c.Cnxipcontterm, c.Cnxipdonoracc, c.Cnxipdoncliuni, c.Cnxipdiscband, c.Cnxipholduntil, c.Cnxipbtaccountno, c.Cnxiptps, c.Cnxiphmdate, c.Cnxipbttermend, c.Cnxiplinewithbt, c.Cnxippromocode, c.Cnxipss07, c.Cnxipss01, c.Cnxipss1a, c.Cnxipss08, c.Cnxipss06, c.Cnxipss09, c.Cnxipss02, c.Cnxipss05, c.Cnxipss14, c.Cnxipnumwithheld, c.Cnxipextnpassl1, c.Cnxipextnpassl2, c.Cnxipvmailcli, c.Cnxipmoretalk, c.Cnxiprpc, c.Cnxipspared1, c.Cnxipspared2, c.Cnxipsparen1, c.Cnxipsparen2, c.Cnxipaddserv, c.Cnxipvaliddate, c.Cnxipvalidby, c.Cnxipcampaign, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cnxip in the database.
func (c *Cnxip) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cnxip SET (` +
		`cnxipcli, cnxipcliline1, cnxipcliline2, cnxipsecondline, cnxipservtype, cnxipsurcharge, cnxipentityno, cnxipextensionno, cnxipextnnol2, cnxipdatetocs, cnxipdatefromc1, cnxipdatefromc2, cnxipdatetoopt, cnxipdatefromopt, cnxipdatetobill, cnxippackageno, cnxipmaccode, cnxipdateentered, cnxipenteredby, cnxipdateletter, cnxipbtpostcode, cnxipstdcode, cnxipordtype, cnxipcsigndate, cnxipdeladdr, cnxipdeladdref, cnxipmobsaver, cnxipintsaver, cnxipequipment, cnxipweborder, cnxipprepayment, cnxipdateprepay, cnxipservicelvl, cnxipfreebox, cnxipprice, cnxipsubtarrif, cnxipextras, cnxiptariff, cnxipcontterm, cnxipdonoracc, cnxipdoncliuni, cnxipdiscband, cnxipholduntil, cnxipbtaccountno, cnxiptps, cnxiphmdate, cnxipbttermend, cnxiplinewithbt, cnxippromocode, cnxipss07, cnxipss01, cnxipss1a, cnxipss08, cnxipss06, cnxipss09, cnxipss02, cnxipss05, cnxipss14, cnxipnumwithheld, cnxipextnpassl1, cnxipextnpassl2, cnxipvmailcli, cnxipmoretalk, cnxiprpc, cnxipspared1, cnxipspared2, cnxipsparen1, cnxipsparen2, cnxipaddserv, cnxipvaliddate, cnxipvalidby, cnxipcampaign, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74` +
		`) WHERE equinox_lrn = $75`

	// run query
	XOLog(sqlstr, c.Cnxipcli, c.Cnxipcliline1, c.Cnxipcliline2, c.Cnxipsecondline, c.Cnxipservtype, c.Cnxipsurcharge, c.Cnxipentityno, c.Cnxipextensionno, c.Cnxipextnnol2, c.Cnxipdatetocs, c.Cnxipdatefromc1, c.Cnxipdatefromc2, c.Cnxipdatetoopt, c.Cnxipdatefromopt, c.Cnxipdatetobill, c.Cnxippackageno, c.Cnxipmaccode, c.Cnxipdateentered, c.Cnxipenteredby, c.Cnxipdateletter, c.Cnxipbtpostcode, c.Cnxipstdcode, c.Cnxipordtype, c.Cnxipcsigndate, c.Cnxipdeladdr, c.Cnxipdeladdref, c.Cnxipmobsaver, c.Cnxipintsaver, c.Cnxipequipment, c.Cnxipweborder, c.Cnxipprepayment, c.Cnxipdateprepay, c.Cnxipservicelvl, c.Cnxipfreebox, c.Cnxipprice, c.Cnxipsubtarrif, c.Cnxipextras, c.Cnxiptariff, c.Cnxipcontterm, c.Cnxipdonoracc, c.Cnxipdoncliuni, c.Cnxipdiscband, c.Cnxipholduntil, c.Cnxipbtaccountno, c.Cnxiptps, c.Cnxiphmdate, c.Cnxipbttermend, c.Cnxiplinewithbt, c.Cnxippromocode, c.Cnxipss07, c.Cnxipss01, c.Cnxipss1a, c.Cnxipss08, c.Cnxipss06, c.Cnxipss09, c.Cnxipss02, c.Cnxipss05, c.Cnxipss14, c.Cnxipnumwithheld, c.Cnxipextnpassl1, c.Cnxipextnpassl2, c.Cnxipvmailcli, c.Cnxipmoretalk, c.Cnxiprpc, c.Cnxipspared1, c.Cnxipspared2, c.Cnxipsparen1, c.Cnxipsparen2, c.Cnxipaddserv, c.Cnxipvaliddate, c.Cnxipvalidby, c.Cnxipcampaign, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cnxipcli, c.Cnxipcliline1, c.Cnxipcliline2, c.Cnxipsecondline, c.Cnxipservtype, c.Cnxipsurcharge, c.Cnxipentityno, c.Cnxipextensionno, c.Cnxipextnnol2, c.Cnxipdatetocs, c.Cnxipdatefromc1, c.Cnxipdatefromc2, c.Cnxipdatetoopt, c.Cnxipdatefromopt, c.Cnxipdatetobill, c.Cnxippackageno, c.Cnxipmaccode, c.Cnxipdateentered, c.Cnxipenteredby, c.Cnxipdateletter, c.Cnxipbtpostcode, c.Cnxipstdcode, c.Cnxipordtype, c.Cnxipcsigndate, c.Cnxipdeladdr, c.Cnxipdeladdref, c.Cnxipmobsaver, c.Cnxipintsaver, c.Cnxipequipment, c.Cnxipweborder, c.Cnxipprepayment, c.Cnxipdateprepay, c.Cnxipservicelvl, c.Cnxipfreebox, c.Cnxipprice, c.Cnxipsubtarrif, c.Cnxipextras, c.Cnxiptariff, c.Cnxipcontterm, c.Cnxipdonoracc, c.Cnxipdoncliuni, c.Cnxipdiscband, c.Cnxipholduntil, c.Cnxipbtaccountno, c.Cnxiptps, c.Cnxiphmdate, c.Cnxipbttermend, c.Cnxiplinewithbt, c.Cnxippromocode, c.Cnxipss07, c.Cnxipss01, c.Cnxipss1a, c.Cnxipss08, c.Cnxipss06, c.Cnxipss09, c.Cnxipss02, c.Cnxipss05, c.Cnxipss14, c.Cnxipnumwithheld, c.Cnxipextnpassl1, c.Cnxipextnpassl2, c.Cnxipvmailcli, c.Cnxipmoretalk, c.Cnxiprpc, c.Cnxipspared1, c.Cnxipspared2, c.Cnxipsparen1, c.Cnxipsparen2, c.Cnxipaddserv, c.Cnxipvaliddate, c.Cnxipvalidby, c.Cnxipcampaign, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cnxip to the database.
func (c *Cnxip) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cnxip.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cnxip) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cnxip (` +
		`cnxipcli, cnxipcliline1, cnxipcliline2, cnxipsecondline, cnxipservtype, cnxipsurcharge, cnxipentityno, cnxipextensionno, cnxipextnnol2, cnxipdatetocs, cnxipdatefromc1, cnxipdatefromc2, cnxipdatetoopt, cnxipdatefromopt, cnxipdatetobill, cnxippackageno, cnxipmaccode, cnxipdateentered, cnxipenteredby, cnxipdateletter, cnxipbtpostcode, cnxipstdcode, cnxipordtype, cnxipcsigndate, cnxipdeladdr, cnxipdeladdref, cnxipmobsaver, cnxipintsaver, cnxipequipment, cnxipweborder, cnxipprepayment, cnxipdateprepay, cnxipservicelvl, cnxipfreebox, cnxipprice, cnxipsubtarrif, cnxipextras, cnxiptariff, cnxipcontterm, cnxipdonoracc, cnxipdoncliuni, cnxipdiscband, cnxipholduntil, cnxipbtaccountno, cnxiptps, cnxiphmdate, cnxipbttermend, cnxiplinewithbt, cnxippromocode, cnxipss07, cnxipss01, cnxipss1a, cnxipss08, cnxipss06, cnxipss09, cnxipss02, cnxipss05, cnxipss14, cnxipnumwithheld, cnxipextnpassl1, cnxipextnpassl2, cnxipvmailcli, cnxipmoretalk, cnxiprpc, cnxipspared1, cnxipspared2, cnxipsparen1, cnxipsparen2, cnxipaddserv, cnxipvaliddate, cnxipvalidby, cnxipcampaign, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cnxipcli, cnxipcliline1, cnxipcliline2, cnxipsecondline, cnxipservtype, cnxipsurcharge, cnxipentityno, cnxipextensionno, cnxipextnnol2, cnxipdatetocs, cnxipdatefromc1, cnxipdatefromc2, cnxipdatetoopt, cnxipdatefromopt, cnxipdatetobill, cnxippackageno, cnxipmaccode, cnxipdateentered, cnxipenteredby, cnxipdateletter, cnxipbtpostcode, cnxipstdcode, cnxipordtype, cnxipcsigndate, cnxipdeladdr, cnxipdeladdref, cnxipmobsaver, cnxipintsaver, cnxipequipment, cnxipweborder, cnxipprepayment, cnxipdateprepay, cnxipservicelvl, cnxipfreebox, cnxipprice, cnxipsubtarrif, cnxipextras, cnxiptariff, cnxipcontterm, cnxipdonoracc, cnxipdoncliuni, cnxipdiscband, cnxipholduntil, cnxipbtaccountno, cnxiptps, cnxiphmdate, cnxipbttermend, cnxiplinewithbt, cnxippromocode, cnxipss07, cnxipss01, cnxipss1a, cnxipss08, cnxipss06, cnxipss09, cnxipss02, cnxipss05, cnxipss14, cnxipnumwithheld, cnxipextnpassl1, cnxipextnpassl2, cnxipvmailcli, cnxipmoretalk, cnxiprpc, cnxipspared1, cnxipspared2, cnxipsparen1, cnxipsparen2, cnxipaddserv, cnxipvaliddate, cnxipvalidby, cnxipcampaign, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cnxipcli, EXCLUDED.cnxipcliline1, EXCLUDED.cnxipcliline2, EXCLUDED.cnxipsecondline, EXCLUDED.cnxipservtype, EXCLUDED.cnxipsurcharge, EXCLUDED.cnxipentityno, EXCLUDED.cnxipextensionno, EXCLUDED.cnxipextnnol2, EXCLUDED.cnxipdatetocs, EXCLUDED.cnxipdatefromc1, EXCLUDED.cnxipdatefromc2, EXCLUDED.cnxipdatetoopt, EXCLUDED.cnxipdatefromopt, EXCLUDED.cnxipdatetobill, EXCLUDED.cnxippackageno, EXCLUDED.cnxipmaccode, EXCLUDED.cnxipdateentered, EXCLUDED.cnxipenteredby, EXCLUDED.cnxipdateletter, EXCLUDED.cnxipbtpostcode, EXCLUDED.cnxipstdcode, EXCLUDED.cnxipordtype, EXCLUDED.cnxipcsigndate, EXCLUDED.cnxipdeladdr, EXCLUDED.cnxipdeladdref, EXCLUDED.cnxipmobsaver, EXCLUDED.cnxipintsaver, EXCLUDED.cnxipequipment, EXCLUDED.cnxipweborder, EXCLUDED.cnxipprepayment, EXCLUDED.cnxipdateprepay, EXCLUDED.cnxipservicelvl, EXCLUDED.cnxipfreebox, EXCLUDED.cnxipprice, EXCLUDED.cnxipsubtarrif, EXCLUDED.cnxipextras, EXCLUDED.cnxiptariff, EXCLUDED.cnxipcontterm, EXCLUDED.cnxipdonoracc, EXCLUDED.cnxipdoncliuni, EXCLUDED.cnxipdiscband, EXCLUDED.cnxipholduntil, EXCLUDED.cnxipbtaccountno, EXCLUDED.cnxiptps, EXCLUDED.cnxiphmdate, EXCLUDED.cnxipbttermend, EXCLUDED.cnxiplinewithbt, EXCLUDED.cnxippromocode, EXCLUDED.cnxipss07, EXCLUDED.cnxipss01, EXCLUDED.cnxipss1a, EXCLUDED.cnxipss08, EXCLUDED.cnxipss06, EXCLUDED.cnxipss09, EXCLUDED.cnxipss02, EXCLUDED.cnxipss05, EXCLUDED.cnxipss14, EXCLUDED.cnxipnumwithheld, EXCLUDED.cnxipextnpassl1, EXCLUDED.cnxipextnpassl2, EXCLUDED.cnxipvmailcli, EXCLUDED.cnxipmoretalk, EXCLUDED.cnxiprpc, EXCLUDED.cnxipspared1, EXCLUDED.cnxipspared2, EXCLUDED.cnxipsparen1, EXCLUDED.cnxipsparen2, EXCLUDED.cnxipaddserv, EXCLUDED.cnxipvaliddate, EXCLUDED.cnxipvalidby, EXCLUDED.cnxipcampaign, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cnxipcli, c.Cnxipcliline1, c.Cnxipcliline2, c.Cnxipsecondline, c.Cnxipservtype, c.Cnxipsurcharge, c.Cnxipentityno, c.Cnxipextensionno, c.Cnxipextnnol2, c.Cnxipdatetocs, c.Cnxipdatefromc1, c.Cnxipdatefromc2, c.Cnxipdatetoopt, c.Cnxipdatefromopt, c.Cnxipdatetobill, c.Cnxippackageno, c.Cnxipmaccode, c.Cnxipdateentered, c.Cnxipenteredby, c.Cnxipdateletter, c.Cnxipbtpostcode, c.Cnxipstdcode, c.Cnxipordtype, c.Cnxipcsigndate, c.Cnxipdeladdr, c.Cnxipdeladdref, c.Cnxipmobsaver, c.Cnxipintsaver, c.Cnxipequipment, c.Cnxipweborder, c.Cnxipprepayment, c.Cnxipdateprepay, c.Cnxipservicelvl, c.Cnxipfreebox, c.Cnxipprice, c.Cnxipsubtarrif, c.Cnxipextras, c.Cnxiptariff, c.Cnxipcontterm, c.Cnxipdonoracc, c.Cnxipdoncliuni, c.Cnxipdiscband, c.Cnxipholduntil, c.Cnxipbtaccountno, c.Cnxiptps, c.Cnxiphmdate, c.Cnxipbttermend, c.Cnxiplinewithbt, c.Cnxippromocode, c.Cnxipss07, c.Cnxipss01, c.Cnxipss1a, c.Cnxipss08, c.Cnxipss06, c.Cnxipss09, c.Cnxipss02, c.Cnxipss05, c.Cnxipss14, c.Cnxipnumwithheld, c.Cnxipextnpassl1, c.Cnxipextnpassl2, c.Cnxipvmailcli, c.Cnxipmoretalk, c.Cnxiprpc, c.Cnxipspared1, c.Cnxipspared2, c.Cnxipsparen1, c.Cnxipsparen2, c.Cnxipaddserv, c.Cnxipvaliddate, c.Cnxipvalidby, c.Cnxipcampaign, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cnxipcli, c.Cnxipcliline1, c.Cnxipcliline2, c.Cnxipsecondline, c.Cnxipservtype, c.Cnxipsurcharge, c.Cnxipentityno, c.Cnxipextensionno, c.Cnxipextnnol2, c.Cnxipdatetocs, c.Cnxipdatefromc1, c.Cnxipdatefromc2, c.Cnxipdatetoopt, c.Cnxipdatefromopt, c.Cnxipdatetobill, c.Cnxippackageno, c.Cnxipmaccode, c.Cnxipdateentered, c.Cnxipenteredby, c.Cnxipdateletter, c.Cnxipbtpostcode, c.Cnxipstdcode, c.Cnxipordtype, c.Cnxipcsigndate, c.Cnxipdeladdr, c.Cnxipdeladdref, c.Cnxipmobsaver, c.Cnxipintsaver, c.Cnxipequipment, c.Cnxipweborder, c.Cnxipprepayment, c.Cnxipdateprepay, c.Cnxipservicelvl, c.Cnxipfreebox, c.Cnxipprice, c.Cnxipsubtarrif, c.Cnxipextras, c.Cnxiptariff, c.Cnxipcontterm, c.Cnxipdonoracc, c.Cnxipdoncliuni, c.Cnxipdiscband, c.Cnxipholduntil, c.Cnxipbtaccountno, c.Cnxiptps, c.Cnxiphmdate, c.Cnxipbttermend, c.Cnxiplinewithbt, c.Cnxippromocode, c.Cnxipss07, c.Cnxipss01, c.Cnxipss1a, c.Cnxipss08, c.Cnxipss06, c.Cnxipss09, c.Cnxipss02, c.Cnxipss05, c.Cnxipss14, c.Cnxipnumwithheld, c.Cnxipextnpassl1, c.Cnxipextnpassl2, c.Cnxipvmailcli, c.Cnxipmoretalk, c.Cnxiprpc, c.Cnxipspared1, c.Cnxipspared2, c.Cnxipsparen1, c.Cnxipsparen2, c.Cnxipaddserv, c.Cnxipvaliddate, c.Cnxipvalidby, c.Cnxipcampaign, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cnxip from the database.
func (c *Cnxip) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cnxip WHERE equinox_lrn = $1`

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

// CnxipByEquinoxLrn retrieves a row from 'equinox.cnxip' as a Cnxip.
//
// Generated from index 'cnxip_pkey'.
func CnxipByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxip, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxipcli, cnxipcliline1, cnxipcliline2, cnxipsecondline, cnxipservtype, cnxipsurcharge, cnxipentityno, cnxipextensionno, cnxipextnnol2, cnxipdatetocs, cnxipdatefromc1, cnxipdatefromc2, cnxipdatetoopt, cnxipdatefromopt, cnxipdatetobill, cnxippackageno, cnxipmaccode, cnxipdateentered, cnxipenteredby, cnxipdateletter, cnxipbtpostcode, cnxipstdcode, cnxipordtype, cnxipcsigndate, cnxipdeladdr, cnxipdeladdref, cnxipmobsaver, cnxipintsaver, cnxipequipment, cnxipweborder, cnxipprepayment, cnxipdateprepay, cnxipservicelvl, cnxipfreebox, cnxipprice, cnxipsubtarrif, cnxipextras, cnxiptariff, cnxipcontterm, cnxipdonoracc, cnxipdoncliuni, cnxipdiscband, cnxipholduntil, cnxipbtaccountno, cnxiptps, cnxiphmdate, cnxipbttermend, cnxiplinewithbt, cnxippromocode, cnxipss07, cnxipss01, cnxipss1a, cnxipss08, cnxipss06, cnxipss09, cnxipss02, cnxipss05, cnxipss14, cnxipnumwithheld, cnxipextnpassl1, cnxipextnpassl2, cnxipvmailcli, cnxipmoretalk, cnxiprpc, cnxipspared1, cnxipspared2, cnxipsparen1, cnxipsparen2, cnxipaddserv, cnxipvaliddate, cnxipvalidby, cnxipcampaign, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxip ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxip{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxipcli, &c.Cnxipcliline1, &c.Cnxipcliline2, &c.Cnxipsecondline, &c.Cnxipservtype, &c.Cnxipsurcharge, &c.Cnxipentityno, &c.Cnxipextensionno, &c.Cnxipextnnol2, &c.Cnxipdatetocs, &c.Cnxipdatefromc1, &c.Cnxipdatefromc2, &c.Cnxipdatetoopt, &c.Cnxipdatefromopt, &c.Cnxipdatetobill, &c.Cnxippackageno, &c.Cnxipmaccode, &c.Cnxipdateentered, &c.Cnxipenteredby, &c.Cnxipdateletter, &c.Cnxipbtpostcode, &c.Cnxipstdcode, &c.Cnxipordtype, &c.Cnxipcsigndate, &c.Cnxipdeladdr, &c.Cnxipdeladdref, &c.Cnxipmobsaver, &c.Cnxipintsaver, &c.Cnxipequipment, &c.Cnxipweborder, &c.Cnxipprepayment, &c.Cnxipdateprepay, &c.Cnxipservicelvl, &c.Cnxipfreebox, &c.Cnxipprice, &c.Cnxipsubtarrif, &c.Cnxipextras, &c.Cnxiptariff, &c.Cnxipcontterm, &c.Cnxipdonoracc, &c.Cnxipdoncliuni, &c.Cnxipdiscband, &c.Cnxipholduntil, &c.Cnxipbtaccountno, &c.Cnxiptps, &c.Cnxiphmdate, &c.Cnxipbttermend, &c.Cnxiplinewithbt, &c.Cnxippromocode, &c.Cnxipss07, &c.Cnxipss01, &c.Cnxipss1a, &c.Cnxipss08, &c.Cnxipss06, &c.Cnxipss09, &c.Cnxipss02, &c.Cnxipss05, &c.Cnxipss14, &c.Cnxipnumwithheld, &c.Cnxipextnpassl1, &c.Cnxipextnpassl2, &c.Cnxipvmailcli, &c.Cnxipmoretalk, &c.Cnxiprpc, &c.Cnxipspared1, &c.Cnxipspared2, &c.Cnxipsparen1, &c.Cnxipsparen2, &c.Cnxipaddserv, &c.Cnxipvaliddate, &c.Cnxipvalidby, &c.Cnxipcampaign, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
