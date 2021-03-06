// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cnxpay represents a row from 'equinox.cnxpay'.
type Cnxpay struct {
	Cnxpaydate       pq.NullTime     `json:"cnxpaydate"`       // cnxpaydate
	Cnxpaycctype     sql.NullInt64   `json:"cnxpaycctype"`     // cnxpaycctype
	Cnxpayccno       sql.NullString  `json:"cnxpayccno"`       // cnxpayccno
	Cnxpayccexp      sql.NullString  `json:"cnxpayccexp"`      // cnxpayccexp
	Cnxpayccstart    sql.NullString  `json:"cnxpayccstart"`    // cnxpayccstart
	Cnxpayccissue    sql.NullString  `json:"cnxpayccissue"`    // cnxpayccissue
	Cnxpayamount     sql.NullFloat64 `json:"cnxpayamount"`     // cnxpayamount
	Cnxpaychequeno   sql.NullString  `json:"cnxpaychequeno"`   // cnxpaychequeno
	Cnxpaytransaccod sql.NullString  `json:"cnxpaytransaccod"` // cnxpaytransaccod
	Cnxpayreffield   sql.NullString  `json:"cnxpayreffield"`   // cnxpayreffield
	Cnxpayexportflag sql.NullInt64   `json:"cnxpayexportflag"` // cnxpayexportflag
	Cnxpayauthorised sql.NullString  `json:"cnxpayauthorised"` // cnxpayauthorised
	Cnxpayauthamount sql.NullFloat64 `json:"cnxpayauthamount"` // cnxpayauthamount
	Cnxpayauthdate   pq.NullTime     `json:"cnxpayauthdate"`   // cnxpayauthdate
	Cnxpayauthtext   sql.NullString  `json:"cnxpayauthtext"`   // cnxpayauthtext
	Cnxpayenteredby  sql.NullString  `json:"cnxpayenteredby"`  // cnxpayenteredby
	Cnxpaycommentcod sql.NullString  `json:"cnxpaycommentcod"` // cnxpaycommentcod
	Cnxpayadvance    sql.NullFloat64 `json:"cnxpayadvance"`    // cnxpayadvance
	Cnxpayconncharge sql.NullFloat64 `json:"cnxpayconncharge"` // cnxpayconncharge
	Cnxpaypprice     sql.NullFloat64 `json:"cnxpaypprice"`     // cnxpaypprice
	Cnxpayprepay2bil pq.NullTime     `json:"cnxpayprepay2bil"` // cnxpayprepay2bil
	Cnxpayweborder   sql.NullInt64   `json:"cnxpayweborder"`   // cnxpayweborder
	Cnxmppuniquesys  sql.NullInt64   `json:"cnxmppuniquesys"`  // cnxmppuniquesys
	Cnxprosaccountno sql.NullInt64   `json:"cnxprosaccountno"` // cnxprosaccountno
	Cnxpaycv2        sql.NullString  `json:"cnxpaycv2"`        // cnxpaycv2
	Cnxpaypostcode   sql.NullString  `json:"cnxpaypostcode"`   // cnxpaypostcode
	Cnxpaycbcappfee  sql.NullInt64   `json:"cnxpaycbcappfee"`  // cnxpaycbcappfee
	Cnxpayspared1    pq.NullTime     `json:"cnxpayspared1"`    // cnxpayspared1
	Cnxpaybanked     pq.NullTime     `json:"cnxpaybanked"`     // cnxpaybanked
	Cnxpaypmm        pq.NullTime     `json:"cnxpaypmm"`        // cnxpaypmm
	Cnxpayrefund     pq.NullTime     `json:"cnxpayrefund"`     // cnxpayrefund
	Cnxpaytime       pq.NullTime     `json:"cnxpaytime"`       // cnxpaytime
	Cnxpaysparen1    sql.NullFloat64 `json:"cnxpaysparen1"`    // cnxpaysparen1
	Cnxpayvaliddate  pq.NullTime     `json:"cnxpayvaliddate"`  // cnxpayvaliddate
	Cnxpayvalidby    sql.NullString  `json:"cnxpayvalidby"`    // cnxpayvalidby
	Cnxpaymarked     sql.NullInt64   `json:"cnxpaymarked"`     // cnxpaymarked
	Cnxpaylastchange sql.NullString  `json:"cnxpaylastchange"` // cnxpaylastchange
	Cnxpaypsp        sql.NullString  `json:"cnxpaypsp"`        // cnxpaypsp
	Cnxpayref        sql.NullString  `json:"cnxpayref"`        // cnxpayref
	Cnxpayshopref    sql.NullString  `json:"cnxpayshopref"`    // cnxpayshopref
	Cnxpaycontract   sql.NullInt64   `json:"cnxpaycontract"`   // cnxpaycontract
	Cnxpayrecuniqsys sql.NullString  `json:"cnxpayrecuniqsys"` // cnxpayrecuniqsys
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCnxpay(db XODB, callback func(x Cnxpay) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxpay `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cnxpay{}

		// scan
		err = q.Scan(&c.Cnxpaydate, &c.Cnxpaycctype, &c.Cnxpayccno, &c.Cnxpayccexp, &c.Cnxpayccstart, &c.Cnxpayccissue, &c.Cnxpayamount, &c.Cnxpaychequeno, &c.Cnxpaytransaccod, &c.Cnxpayreffield, &c.Cnxpayexportflag, &c.Cnxpayauthorised, &c.Cnxpayauthamount, &c.Cnxpayauthdate, &c.Cnxpayauthtext, &c.Cnxpayenteredby, &c.Cnxpaycommentcod, &c.Cnxpayadvance, &c.Cnxpayconncharge, &c.Cnxpaypprice, &c.Cnxpayprepay2bil, &c.Cnxpayweborder, &c.Cnxmppuniquesys, &c.Cnxprosaccountno, &c.Cnxpaycv2, &c.Cnxpaypostcode, &c.Cnxpaycbcappfee, &c.Cnxpayspared1, &c.Cnxpaybanked, &c.Cnxpaypmm, &c.Cnxpayrefund, &c.Cnxpaytime, &c.Cnxpaysparen1, &c.Cnxpayvaliddate, &c.Cnxpayvalidby, &c.Cnxpaymarked, &c.Cnxpaylastchange, &c.Cnxpaypsp, &c.Cnxpayref, &c.Cnxpayshopref, &c.Cnxpaycontract, &c.Cnxpayrecuniqsys, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CnxpayByEquinoxLrn retrieves a row from 'equinox.cnxpay' as a Cnxpay.
//
// Generated from index 'cnxpay_pkey'.
func CnxpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxpay{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxpaydate, &c.Cnxpaycctype, &c.Cnxpayccno, &c.Cnxpayccexp, &c.Cnxpayccstart, &c.Cnxpayccissue, &c.Cnxpayamount, &c.Cnxpaychequeno, &c.Cnxpaytransaccod, &c.Cnxpayreffield, &c.Cnxpayexportflag, &c.Cnxpayauthorised, &c.Cnxpayauthamount, &c.Cnxpayauthdate, &c.Cnxpayauthtext, &c.Cnxpayenteredby, &c.Cnxpaycommentcod, &c.Cnxpayadvance, &c.Cnxpayconncharge, &c.Cnxpaypprice, &c.Cnxpayprepay2bil, &c.Cnxpayweborder, &c.Cnxmppuniquesys, &c.Cnxprosaccountno, &c.Cnxpaycv2, &c.Cnxpaypostcode, &c.Cnxpaycbcappfee, &c.Cnxpayspared1, &c.Cnxpaybanked, &c.Cnxpaypmm, &c.Cnxpayrefund, &c.Cnxpaytime, &c.Cnxpaysparen1, &c.Cnxpayvaliddate, &c.Cnxpayvalidby, &c.Cnxpaymarked, &c.Cnxpaylastchange, &c.Cnxpaypsp, &c.Cnxpayref, &c.Cnxpayshopref, &c.Cnxpaycontract, &c.Cnxpayrecuniqsys, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
