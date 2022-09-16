create table trd_ref_date
(
    trd_recordno    integer    not null
        constraint ref_date_pkey
            primary key,
    datetype        varchar(4) not null,
    datewil         date       not null,
    refdatetime     varchar(8) not null,
    dateversionuser varchar(3) not null
);

alter table trd_ref_date
    owner to root;

create table trd_trade
(
    trd_uuid             uuid default gen_random_uuid() not null
        constraint pkey_tbl
            primary key,
    trd_recordno         integer                        not null,
    trd_glosstraderef    integer                        not null,
    trd_versiono         integer                        not null,
    trd_origin           text                           not null,
    trd_tradetype        text                           not null,
    trd_settlementstatus text                           not null,
    trd_tradestatus      varchar(1)                     not null,
    trd_originversion    integer                        not null
);

alter table trd_trade
    owner to root;

create table trd_external_ref
(
    trd_recordno       integer not null,
    ext_ref_extreftype text    not null,
    ext_ref_extref     text    not null,
    constraint external_ref_pkey
        primary key (trd_recordno, ext_ref_extreftype)
);

alter table trd_external_ref
    owner to root;

create table trd_driver
(
    trd_recordno   integer not null,
    trd_drivertype text    not null,
    trd_drivercode text    not null,
    primary key (trd_recordno, trd_drivertype)
);

alter table trd_driver
    owner to root;

create table trd_instrument
(
    trd_recordno      integer not null,
    trd_instrtype     text    not null,
    trd_p2000instrref text    not null,
    trd_instrreftype  text    not null,
    trd_instrref      text    not null,
    trd_longname      text    not null,
    trd_quantity      text    not null,
    primary key (trd_recordno, trd_instrtype)
);

alter table trd_instrument
    owner to root;

create table trd_journal
(
    trd_recordno         integer not null,
    trd_accounts_company text    not null,
    trd_journal_type     text    not null,
    trd_posting_type     text    not null,
    trd_journal_no       text    not null,
    trd_procaction       text    not null,
    primary key (trd_recordno, trd_accounts_company)
);

alter table trd_journal
    owner to root;

create table trd_link
(
    trd_recordno           integer not null,
    trd_link_type_wil      text    not null,
    trd_main_record_no_wil text    not null,
    trd_sub_recordno_wil   text    not null,
    trd_link_qty_wil       text    not null,
    trd_link_status_wil    text    not null,
    primary key (trd_recordno, trd_link_type_wil)
);

alter table trd_link
    owner to root;

create table trd_inst_ext
(
    trd_recordno integer not null,
    trd_service  text    not null,
    trd_extref   text    not null,
    primary key (trd_recordno, trd_service)
);

alter table trd_inst_ext
    owner to root;

create table trd_trade_narrative
(
    trd_recordno       integer not null,
    trd_narrative_code text    not null,
    trd_seqno          text    not null,
    trd_narrative      text    not null,
    primary key (trd_recordno, trd_narrative_code)
);

alter table trd_trade_narrative
    owner to root;

create table trd_party
(
    trd_recordno           integer not null,
    trd_trade_party        text    not null,
    trd_partyref           text    not null,
    trd_partyref_type_text text    not null,
    trd_ext_partyref       text    not null,
    trd_longname           text    not null,
    primary key (trd_recordno, trd_trade_party)
);

alter table trd_party
    owner to root;

create table trd_party_driver
(
    trd_recordno    integer not null,
    trd_trade_party text    not null,
    trd_driver_type text    not null,
    trd_driver_code text    not null,
    primary key (trd_recordno, trd_trade_party)
);

alter table trd_party_driver
    owner to root;

create table trd_rate
(
    trd_recordno           integer not null,
    trd_charge_levy_type   text    not null,
    trd_actual_charge      text    not null,
    trd_amount_type        text    not null,
    trd_rate_instrref_type text    not null,
    trd_rate_instrref      text    not null,
    trd_rate_instrid       text    not null,
    trd_rate_entered       text    not null,
    trd_charge_rate        text    not null,
    trd_mult_divind        text    not null,
    primary key (trd_recordno, trd_charge_levy_type)
);

alter table trd_rate
    owner to root;

create table trd_amount
(
    trd_recordno              integer not null,
    trd_charge_levy_type_p2k  text    not null,
    trd_charge_levy_instr_p2k text    not null,
    trd_charge_discount_wil   text    not null,
    trd_charge_levy_qty_p2k   text    not null,
    trd_charge_levyrate_p2k   text    not null,
    primary key (trd_recordno, trd_charge_levy_type_p2k, trd_charge_levy_instr_p2k)
);

alter table trd_amount
    owner to root;

create table trd_settlement
(
    trd_recordno         integer not null,
    trd_deliverytype     text    not null,
    trd_settleeventinstr text    not null,
    trd_settleterms      text    not null,
    trd_originalqty      text    not null,
    trd_openqty          text    not null,
    trd_settledate       text    not null,
    trd_delrecind        text    not null,
    trd_settlestatus     text    not null,
    trd_tradestatus      text    not null,
    trd_settlenarrative1 text    not null,
    trd_settlenarrative2 text    not null,
    trd_settlenarrative3 text    not null,
    trd_settlenarrative4 text    not null,
    trd_settlenarrative5 text    not null,
    trd_settlenarrative6 text    not null,
    trd_settlenarrative7 text    not null,
    trd_settlenarrative8 text    not null,
    trd_dompaliaswil     text    not null,
    trd_dompaliasdesc    text    not null,
    trd_dompdepottypewil text    not null,
    trd_dompdaccwil      text    not null,
    trd_dompservicewil   text    not null,
    trd_secpaliaswil     text    not null,
    trd_secpservicewil   text    not null,
    primary key (trd_recordno, trd_deliverytype)
);

alter table trd_settlement
    owner to root;

create table trd_processing
(
    trd_recordno      integer not null,
    trd_proc_alias    text    not null,
    trd_proc_action   text    not null,
    trd_due_datetime  text    not null,
    trd_done_datetime text    not null,
    primary key (trd_recordno, trd_proc_alias)
);

alter table trd_processing
    owner to root;

create table trd_processing_event
(
    trd_recordno      integer not null,
    trd_eventtype     text    not null,
    trd_eventdate     text    not null,
    trd_eventdateto   text    not null,
    trd_entrydatetime text    not null,
    trd_eventcode     text    not null,
    trd_exceptiontype text    not null,
    trd_expirydate    text    not null,
    primary key (trd_recordno, trd_eventtype)
);

alter table trd_processing_event
    owner to root;

create table trd_event_narrative
(
    trd_recordno       integer not null,
    trd_narrative_code text    not null,
    trd_seqno          text    not null,
    trd_narrative      text    not null,
    primary key (trd_recordno, trd_narrative_code)
);

alter table trd_event_narrative
    owner to root;

create table trd_instruction
(
    trd_recordno         integer not null,
    trd_procaction       text    not null,
    trd_destination      text    not null,
    trd_procstatus       text    not null,
    trd_recordidentifier text    not null,
    trd_recordversion    text    not null,
    trd_instformat       text    not null,
    trd_tradeparty       text    not null,
    trd_partyref         text    not null,
    trd_deliverytype     text    not null,
    trd_addresscode      text    not null,
    trd_servicestatus    text    not null,
    trd_noofcopies       text    not null,
    trd_duedatetime      text    not null,
    primary key (trd_recordno, trd_deliverytype)
);

alter table trd_instruction
    owner to root;

create table trd_instruction_effect
(
    trd_recordno      integer not null,
    trd_eventtype     text    not null,
    trd_eventdate     text    not null,
    trd_eventdateto   text    not null,
    trd_entrydatetime text    not null,
    trd_eventcode     text    not null,
    trd_exceptiontype text    not null,
    trd_expirydate    text    not null,
    primary key (trd_recordno, trd_eventtype)
);

alter table trd_instruction_effect
    owner to root;

create table trd_event
(
    trd_recordno      integer not null,
    trd_eventtype     text    not null,
    trd_eventdate     text    not null,
    trd_eventdateto   text    not null,
    trd_entrydatetime text    not null,
    trd_eventcode     text    not null,
    trd_exceptiontype text    not null,
    trd_expirydate    text    not null,
    primary key (trd_recordno, trd_eventtype)
);

alter table trd_event
    owner to root;

create table trd_trade_code
(
    trd_recordno  integer not null,
    trd_ext       text    not null,
    trd_radeclass text    not null,
    trd_radecode  text    not null,
    primary key (trd_recordno, trd_ext)
);

alter table trd_trade_code
    owner to root;


create or replace view trades
            (trd_uuid, trd_recordno, trd_glosstraderef, trd_versiono, trd_origin, trd_tradetype, trd_settlementstatus,
             trd_tradestatus, trd_originversion, trd_accounts_company, trd_journal_no)
as
SELECT t.trd_uuid,
       t.trd_recordno,
       t.trd_glosstraderef,
       t.trd_versiono,
       t.trd_origin,
       t.trd_tradetype,
       t.trd_settlementstatus,
       t.trd_tradestatus,
       t.trd_originversion,
       j.trd_accounts_company,
       j.trd_journal_no
FROM trd_trade t,
     trd_journal j
WHERE t.trd_recordno = j.trd_recordno;

alter table trades
    owner to root;

