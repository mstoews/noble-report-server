// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type TrdAmount struct {
	TrdRecordno           int32  `json:"trd_recordno"`
	TrdChargeLevyTypeP2k  string `json:"trd_charge_levy_type_p2k"`
	TrdChargeLevyInstrP2k string `json:"trd_charge_levy_instr_p2k"`
	TrdChargeDiscountWil  string `json:"trd_charge_discount_wil"`
	TrdChargeLevyQtyP2k   string `json:"trd_charge_levy_qty_p2k"`
	TrdChargeLevyrateP2k  string `json:"trd_charge_levyrate_p2k"`
}

type TrdDriver struct {
	TrdRecordno   int32  `json:"trd_recordno"`
	TrdDrivertype string `json:"trd_drivertype"`
	TrdDrivercode string `json:"trd_drivercode"`
}

type TrdEvent struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	TrdEventtype     string `json:"trd_eventtype"`
	TrdEventdate     string `json:"trd_eventdate"`
	TrdEventdateto   string `json:"trd_eventdateto"`
	TrdEntrydatetime string `json:"trd_entrydatetime"`
	TrdEventcode     string `json:"trd_eventcode"`
	TrdExceptiontype string `json:"trd_exceptiontype"`
	TrdExpirydate    string `json:"trd_expirydate"`
}

type TrdEventNarrative struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	TrdNarrativeCode string `json:"trd_narrative_code"`
	TrdSeqno         string `json:"trd_seqno"`
	TrdNarrative     string `json:"trd_narrative"`
}

type TrdExternalRef struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	ExtRefExtreftype string `json:"ext_ref_extreftype"`
	ExtRefExtref     string `json:"ext_ref_extref"`
}

type TrdInstExt struct {
	TrdRecordno int32  `json:"trd_recordno"`
	TrdService  string `json:"trd_service"`
	TrdExtref   string `json:"trd_extref"`
}

type TrdInstruction struct {
	TrdRecordno         int32  `json:"trd_recordno"`
	TrdProcaction       string `json:"trd_procaction"`
	TrdDestination      string `json:"trd_destination"`
	TrdProcstatus       string `json:"trd_procstatus"`
	TrdRecordidentifier string `json:"trd_recordidentifier"`
	TrdRecordversion    string `json:"trd_recordversion"`
	TrdInstformat       string `json:"trd_instformat"`
	TrdTradeparty       string `json:"trd_tradeparty"`
	TrdPartyref         string `json:"trd_partyref"`
	TrdDeliverytype     string `json:"trd_deliverytype"`
	TrdAddresscode      string `json:"trd_addresscode"`
	TrdServicestatus    string `json:"trd_servicestatus"`
	TrdNoofcopies       string `json:"trd_noofcopies"`
	TrdDuedatetime      string `json:"trd_duedatetime"`
}

type TrdInstructionEffect struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	TrdEventtype     string `json:"trd_eventtype"`
	TrdEventdate     string `json:"trd_eventdate"`
	TrdEventdateto   string `json:"trd_eventdateto"`
	TrdEntrydatetime string `json:"trd_entrydatetime"`
	TrdEventcode     string `json:"trd_eventcode"`
	TrdExceptiontype string `json:"trd_exceptiontype"`
	TrdExpirydate    string `json:"trd_expirydate"`
}

type TrdInstrument struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	TrdInstrtype     string `json:"trd_instrtype"`
	TrdP2000instrref string `json:"trd_p2000instrref"`
	TrdInstrreftype  string `json:"trd_instrreftype"`
	TrdInstrref      string `json:"trd_instrref"`
	TrdLongname      string `json:"trd_longname"`
	TrdQuantity      string `json:"trd_quantity"`
}

type TrdJournal struct {
	TrdRecordno        int32  `json:"trd_recordno"`
	TrdAccountsCompany string `json:"trd_accounts_company"`
	TrdJournalType     string `json:"trd_journal_type"`
	TrdPostingType     string `json:"trd_posting_type"`
	TrdJournalNo       string `json:"trd_journal_no"`
	TrdProcaction      string `json:"trd_procaction"`
}

type TrdLink struct {
	TrdRecordno        int32  `json:"trd_recordno"`
	TrdLinkTypeWil     string `json:"trd_link_type_wil"`
	TrdMainRecordNoWil string `json:"trd_main_record_no_wil"`
	TrdSubRecordnoWil  string `json:"trd_sub_recordno_wil"`
	TrdLinkQtyWil      string `json:"trd_link_qty_wil"`
	TrdLinkStatusWil   string `json:"trd_link_status_wil"`
}

type TrdParty struct {
	TrdRecordno         int32  `json:"trd_recordno"`
	TrdTradeParty       string `json:"trd_trade_party"`
	TrdPartyref         string `json:"trd_partyref"`
	TrdPartyrefTypeText string `json:"trd_partyref_type_text"`
	TrdExtPartyref      string `json:"trd_ext_partyref"`
	TrdLongname         string `json:"trd_longname"`
}

type TrdPartyDriver struct {
	TrdRecordno   int32  `json:"trd_recordno"`
	TrdTradeParty string `json:"trd_trade_party"`
	TrdDriverType string `json:"trd_driver_type"`
	TrdDriverCode string `json:"trd_driver_code"`
}

type TrdProcessing struct {
	TrdRecordno     int32  `json:"trd_recordno"`
	TrdProcAlias    string `json:"trd_proc_alias"`
	TrdProcAction   string `json:"trd_proc_action"`
	TrdDueDatetime  string `json:"trd_due_datetime"`
	TrdDoneDatetime string `json:"trd_done_datetime"`
}

type TrdProcessingEvent struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	TrdEventtype     string `json:"trd_eventtype"`
	TrdEventdate     string `json:"trd_eventdate"`
	TrdEventdateto   string `json:"trd_eventdateto"`
	TrdEntrydatetime string `json:"trd_entrydatetime"`
	TrdEventcode     string `json:"trd_eventcode"`
	TrdExceptiontype string `json:"trd_exceptiontype"`
	TrdExpirydate    string `json:"trd_expirydate"`
}

type TrdRate struct {
	TrdRecordno         int32  `json:"trd_recordno"`
	TrdChargeLevyType   string `json:"trd_charge_levy_type"`
	TrdActualCharge     string `json:"trd_actual_charge"`
	TrdAmountType       string `json:"trd_amount_type"`
	TrdRateInstrrefType string `json:"trd_rate_instrref_type"`
	TrdRateInstrref     string `json:"trd_rate_instrref"`
	TrdRateInstrid      string `json:"trd_rate_instrid"`
	TrdRateEntered      string `json:"trd_rate_entered"`
	TrdChargeRate       string `json:"trd_charge_rate"`
	TrdMultDivind       string `json:"trd_mult_divind"`
}

type TrdRefDate struct {
	TrdRecordno     int32     `json:"trd_recordno"`
	Datetype        string    `json:"datetype"`
	Datewil         time.Time `json:"datewil"`
	Refdatetime     string    `json:"refdatetime"`
	Dateversionuser string    `json:"dateversionuser"`
}

type TrdSettlement struct {
	TrdRecordno         int32  `json:"trd_recordno"`
	TrdDeliverytype     string `json:"trd_deliverytype"`
	TrdSettleeventinstr string `json:"trd_settleeventinstr"`
	TrdSettleterms      string `json:"trd_settleterms"`
	TrdOriginalqty      string `json:"trd_originalqty"`
	TrdOpenqty          string `json:"trd_openqty"`
	TrdSettledate       string `json:"trd_settledate"`
	TrdDelrecind        string `json:"trd_delrecind"`
	TrdSettlestatus     string `json:"trd_settlestatus"`
	TrdTradestatus      string `json:"trd_tradestatus"`
	TrdSettlenarrative1 string `json:"trd_settlenarrative1"`
	TrdSettlenarrative2 string `json:"trd_settlenarrative2"`
	TrdSettlenarrative3 string `json:"trd_settlenarrative3"`
	TrdSettlenarrative4 string `json:"trd_settlenarrative4"`
	TrdSettlenarrative5 string `json:"trd_settlenarrative5"`
	TrdSettlenarrative6 string `json:"trd_settlenarrative6"`
	TrdSettlenarrative7 string `json:"trd_settlenarrative7"`
	TrdSettlenarrative8 string `json:"trd_settlenarrative8"`
	TrdDompaliaswil     string `json:"trd_dompaliaswil"`
	TrdDompaliasdesc    string `json:"trd_dompaliasdesc"`
	TrdDompdepottypewil string `json:"trd_dompdepottypewil"`
	TrdDompdaccwil      string `json:"trd_dompdaccwil"`
	TrdDompservicewil   string `json:"trd_dompservicewil"`
	TrdSecpaliaswil     string `json:"trd_secpaliaswil"`
	TrdSecpservicewil   string `json:"trd_secpservicewil"`
}

type TrdTrade struct {
	TrdUuid             uuid.UUID `json:"trd_uuid"`
	TrdRecordno         int32     `json:"trd_recordno"`
	TrdGlosstraderef    int32     `json:"trd_glosstraderef"`
	TrdVersiono         int32     `json:"trd_versiono"`
	TrdOrigin           string    `json:"trd_origin"`
	TrdTradetype        string    `json:"trd_tradetype"`
	TrdSettlementstatus string    `json:"trd_settlementstatus"`
	TrdTradestatus      string    `json:"trd_tradestatus"`
	TrdOriginversion    int32     `json:"trd_originversion"`
}

type TrdTradeCode struct {
	TrdRecordno  int32  `json:"trd_recordno"`
	TrdExt       string `json:"trd_ext"`
	TrdRadeclass string `json:"trd_radeclass"`
	TrdRadecode  string `json:"trd_radecode"`
}

type TrdTradeNarrative struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	TrdNarrativeCode string `json:"trd_narrative_code"`
	TrdSeqno         string `json:"trd_seqno"`
	TrdNarrative     string `json:"trd_narrative"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
