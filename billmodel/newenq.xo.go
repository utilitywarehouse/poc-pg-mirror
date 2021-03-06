// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Newenq represents a row from 'equinox.newenq'.
type Newenq struct {
	Enqaccountno     sql.NullString `json:"enqaccountno"`     // enqaccountno
	Enqpostcode      sql.NullString `json:"enqpostcode"`      // enqpostcode
	Enqadd1          sql.NullString `json:"enqadd1"`          // enqadd1
	Enqadd2          sql.NullString `json:"enqadd2"`          // enqadd2
	Enqadd3          sql.NullString `json:"enqadd3"`          // enqadd3
	Enqadd4          sql.NullString `json:"enqadd4"`          // enqadd4
	Enqcounty        sql.NullString `json:"enqcounty"`        // enqcounty
	Enqconttitle     sql.NullString `json:"enqconttitle"`     // enqconttitle
	Enqcontsurname   sql.NullString `json:"enqcontsurname"`   // enqcontsurname
	Enqcontfirstnam  sql.NullString `json:"enqcontfirstnam"`  // enqcontfirstnam
	Enqname          sql.NullString `json:"enqname"`          // enqname
	Enqaddverified   sql.NullInt64  `json:"enqaddverified"`   // enqaddverified
	Enqoldaccount    sql.NullInt64  `json:"enqoldaccount"`    // enqoldaccount
	Enqdayphone      sql.NullString `json:"enqdayphone"`      // enqdayphone
	Enqhomephone     sql.NullString `json:"enqhomephone"`     // enqhomephone
	Enqmobilephone   sql.NullString `json:"enqmobilephone"`   // enqmobilephone
	Enqemail         sql.NullString `json:"enqemail"`         // enqemail
	Enqservsp        sql.NullInt64  `json:"enqservsp"`        // enqservsp
	Enqservfree      sql.NullInt64  `json:"enqservfree"`      // enqservfree
	Enqservmob       sql.NullInt64  `json:"enqservmob"`       // enqservmob
	Enqservelec      sql.NullInt64  `json:"enqservelec"`      // enqservelec
	Enqservgas       sql.NullInt64  `json:"enqservgas"`       // enqservgas
	Enqservgoplus    sql.NullInt64  `json:"enqservgoplus"`    // enqservgoplus
	Enqservbroadcall sql.NullInt64  `json:"enqservbroadcall"` // enqservbroadcall
	Enqinterestfrom  sql.NullInt64  `json:"enqinterestfrom"`  // enqinterestfrom
	Enqexistphone    sql.NullString `json:"enqexistphone"`    // enqexistphone
	Enqexistmobile   sql.NullString `json:"enqexistmobile"`   // enqexistmobile
	Enqexistelectric sql.NullString `json:"enqexistelectric"` // enqexistelectric
	Enqexistgas      sql.NullString `json:"enqexistgas"`      // enqexistgas
	Enqexistinternet sql.NullString `json:"enqexistinternet"` // enqexistinternet
	Enqexecid        sql.NullString `json:"enqexecid"`        // enqexecid
	Enqexname        sql.NullString `json:"enqexname"`        // enqexname
	Enqdateentered   pq.NullTime    `json:"enqdateentered"`   // enqdateentered
	Enqdatetolive    pq.NullTime    `json:"enqdatetolive"`    // enqdatetolive
	Enquirystatus    sql.NullString `json:"enquirystatus"`    // enquirystatus
	Enqnumquote      sql.NullInt64  `json:"enqnumquote"`      // enqnumquote
	Enqoldaccountnum sql.NullString `json:"enqoldaccountnum"` // enqoldaccountnum
	Enqbillaccountno sql.NullString `json:"enqbillaccountno"` // enqbillaccountno
	Enqbanksort      sql.NullString `json:"enqbanksort"`      // enqbanksort
	Enqbankaccno     sql.NullString `json:"enqbankaccno"`     // enqbankaccno
	Enqbankaccname   sql.NullString `json:"enqbankaccname"`   // enqbankaccname
	Enqdob           pq.NullTime    `json:"enqdob"`           // enqdob
	Enqcompletebank  sql.NullString `json:"enqcompletebank"`  // enqcompletebank
	Enqemailbill     sql.NullInt64  `json:"enqemailbill"`     // enqemailbill
	Enqnomarketing   sql.NullInt64  `json:"enqnomarketing"`   // enqnomarketing
	Enqenteredby     sql.NullString `json:"enqenteredby"`     // enqenteredby
	Enqpiggybondflag sql.NullString `json:"enqpiggybondflag"` // enqpiggybondflag
	Enqsparechar1    sql.NullString `json:"enqsparechar1"`    // enqsparechar1
	Enqsparechar2    sql.NullString `json:"enqsparechar2"`    // enqsparechar2
	Enqsparenum1     sql.NullInt64  `json:"enqsparenum1"`     // enqsparenum1
	Enqsparenum2     sql.NullInt64  `json:"enqsparenum2"`     // enqsparenum2
	Enqbillinggroup  sql.NullString `json:"enqbillinggroup"`  // enqbillinggroup
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllNewenq(db XODB, callback func(x Newenq) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`enqaccountno, enqpostcode, enqadd1, enqadd2, enqadd3, enqadd4, enqcounty, enqconttitle, enqcontsurname, enqcontfirstnam, enqname, enqaddverified, enqoldaccount, enqdayphone, enqhomephone, enqmobilephone, enqemail, enqservsp, enqservfree, enqservmob, enqservelec, enqservgas, enqservgoplus, enqservbroadcall, enqinterestfrom, enqexistphone, enqexistmobile, enqexistelectric, enqexistgas, enqexistinternet, enqexecid, enqexname, enqdateentered, enqdatetolive, enquirystatus, enqnumquote, enqoldaccountnum, enqbillaccountno, enqbanksort, enqbankaccno, enqbankaccname, enqdob, enqcompletebank, enqemailbill, enqnomarketing, enqenteredby, enqpiggybondflag, enqsparechar1, enqsparechar2, enqsparenum1, enqsparenum2, enqbillinggroup, equinox_lrn, equinox_sec ` +
		`FROM equinox.newenq `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		n := Newenq{}

		// scan
		err = q.Scan(&n.Enqaccountno, &n.Enqpostcode, &n.Enqadd1, &n.Enqadd2, &n.Enqadd3, &n.Enqadd4, &n.Enqcounty, &n.Enqconttitle, &n.Enqcontsurname, &n.Enqcontfirstnam, &n.Enqname, &n.Enqaddverified, &n.Enqoldaccount, &n.Enqdayphone, &n.Enqhomephone, &n.Enqmobilephone, &n.Enqemail, &n.Enqservsp, &n.Enqservfree, &n.Enqservmob, &n.Enqservelec, &n.Enqservgas, &n.Enqservgoplus, &n.Enqservbroadcall, &n.Enqinterestfrom, &n.Enqexistphone, &n.Enqexistmobile, &n.Enqexistelectric, &n.Enqexistgas, &n.Enqexistinternet, &n.Enqexecid, &n.Enqexname, &n.Enqdateentered, &n.Enqdatetolive, &n.Enquirystatus, &n.Enqnumquote, &n.Enqoldaccountnum, &n.Enqbillaccountno, &n.Enqbanksort, &n.Enqbankaccno, &n.Enqbankaccname, &n.Enqdob, &n.Enqcompletebank, &n.Enqemailbill, &n.Enqnomarketing, &n.Enqenteredby, &n.Enqpiggybondflag, &n.Enqsparechar1, &n.Enqsparechar2, &n.Enqsparenum1, &n.Enqsparenum2, &n.Enqbillinggroup, &n.EquinoxLrn, &n.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(n) {
			return nil
		}
	}

	return nil
}

// NewenqByEquinoxLrn retrieves a row from 'equinox.newenq' as a Newenq.
//
// Generated from index 'newenq_pkey'.
func NewenqByEquinoxLrn(db XODB, equinoxLrn int64) (*Newenq, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`enqaccountno, enqpostcode, enqadd1, enqadd2, enqadd3, enqadd4, enqcounty, enqconttitle, enqcontsurname, enqcontfirstnam, enqname, enqaddverified, enqoldaccount, enqdayphone, enqhomephone, enqmobilephone, enqemail, enqservsp, enqservfree, enqservmob, enqservelec, enqservgas, enqservgoplus, enqservbroadcall, enqinterestfrom, enqexistphone, enqexistmobile, enqexistelectric, enqexistgas, enqexistinternet, enqexecid, enqexname, enqdateentered, enqdatetolive, enquirystatus, enqnumquote, enqoldaccountnum, enqbillaccountno, enqbanksort, enqbankaccno, enqbankaccname, enqdob, enqcompletebank, enqemailbill, enqnomarketing, enqenteredby, enqpiggybondflag, enqsparechar1, enqsparechar2, enqsparenum1, enqsparenum2, enqbillinggroup, equinox_lrn, equinox_sec ` +
		`FROM equinox.newenq ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	n := Newenq{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&n.Enqaccountno, &n.Enqpostcode, &n.Enqadd1, &n.Enqadd2, &n.Enqadd3, &n.Enqadd4, &n.Enqcounty, &n.Enqconttitle, &n.Enqcontsurname, &n.Enqcontfirstnam, &n.Enqname, &n.Enqaddverified, &n.Enqoldaccount, &n.Enqdayphone, &n.Enqhomephone, &n.Enqmobilephone, &n.Enqemail, &n.Enqservsp, &n.Enqservfree, &n.Enqservmob, &n.Enqservelec, &n.Enqservgas, &n.Enqservgoplus, &n.Enqservbroadcall, &n.Enqinterestfrom, &n.Enqexistphone, &n.Enqexistmobile, &n.Enqexistelectric, &n.Enqexistgas, &n.Enqexistinternet, &n.Enqexecid, &n.Enqexname, &n.Enqdateentered, &n.Enqdatetolive, &n.Enquirystatus, &n.Enqnumquote, &n.Enqoldaccountnum, &n.Enqbillaccountno, &n.Enqbanksort, &n.Enqbankaccno, &n.Enqbankaccname, &n.Enqdob, &n.Enqcompletebank, &n.Enqemailbill, &n.Enqnomarketing, &n.Enqenteredby, &n.Enqpiggybondflag, &n.Enqsparechar1, &n.Enqsparechar2, &n.Enqsparenum1, &n.Enqsparenum2, &n.Enqbillinggroup, &n.EquinoxLrn, &n.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &n, nil
}
