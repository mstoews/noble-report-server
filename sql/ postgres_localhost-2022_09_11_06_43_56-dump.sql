--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-1.pgdg20.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: root
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO root;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: root
--
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.accounts (
    id bigint NOT NULL,
    owner character varying NOT NULL,
    balance bigint NOT NULL,
    currency character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.accounts OWNER TO root;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO root;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- Name: entries; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.entries (
    id bigint NOT NULL,
    account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.entries OWNER TO root;

--
-- Name: COLUMN entries.amount; Type: COMMENT; Schema: public; Owner: root
--




--
-- Name: entries_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.entries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.entries_id_seq OWNER TO root;

--
-- Name: entries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.entries_id_seq OWNED BY public.entries.id;


--
-- Name: instr_instruments; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.instr_instruments (
    instr_instref text NOT NULL,
    instr_instgrp text NOT NULL,
    instr_longdesc text NOT NULL,
    instr_denminst text NOT NULL,
    instr_pricedps text NOT NULL,
    instr_divisor text NOT NULL,
    instr_multiplier text NOT NULL,
    instr_pricetype text NOT NULL,
    instr_tick text NOT NULL,
    instr_accrued text NOT NULL,
    instr_minqty text NOT NULL,
    instr_qtydps text NOT NULL,
    instr_shrtdesc text NOT NULL,
    instr_market text NOT NULL,
    instr_settinst text NOT NULL,
    instr_book text NOT NULL,
    instr_maxmovt text NOT NULL,
    instr_minmovt text NOT NULL,
    instr_active text NOT NULL,
    instr_verdat text NOT NULL
);


ALTER TABLE public.instr_instruments OWNER TO root;

--
-- Name: pty_party; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.pty_party (
    pty_partyref text NOT NULL,
    pty_category text NOT NULL,
    pty_longdesc text NOT NULL,
    pty_holiday text NOT NULL,
    pty_country text NOT NULL,
    pty_location text NOT NULL,
    pty_shrtdesc text NOT NULL,
    pty_partynam1 text NOT NULL,
    pty_partynam2 text NOT NULL,
    pty_partynam3 text NOT NULL,
    pty_active text NOT NULL,
    pty_verdat text NOT NULL
);


ALTER TABLE public.pty_party OWNER TO root;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO root;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    username character varying NOT NULL,
    refresh_token character varying NOT NULL,
    user_agent character varying NOT NULL,
    client_ip character varying NOT NULL,
    is_blocked boolean DEFAULT false NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.sessions OWNER TO root;

--
-- Name: transfers; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.transfers (
    id bigint NOT NULL,
    from_account_id bigint NOT NULL,
    to_account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.transfers OWNER TO root;

--
-- Name: COLUMN transfers.amount; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.transfers.amount IS 'must be positive';


--
-- Name: transfers_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.transfers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transfers_id_seq OWNER TO root;

--
-- Name: transfers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.transfers_id_seq OWNED BY public.transfers.id;


--
-- Name: trd_amount; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_amount (
    trd_recordno integer NOT NULL,
    trd_charge_levy_type_p2k text NOT NULL,
    trd_charge_levy_instr_p2k text NOT NULL,
    trd_charge_discount_wil text NOT NULL,
    trd_charge_levy_qty_p2k text NOT NULL,
    trd_charge_levyrate_p2k text NOT NULL
);


ALTER TABLE public.trd_amount OWNER TO root;

--
-- Name: trd_driver; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_driver (
    trd_recordno integer NOT NULL,
    trd_drivertype text NOT NULL,
    trd_drivercode text NOT NULL
);


ALTER TABLE public.trd_driver OWNER TO root;

--
-- Name: trd_event; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_event (
    trd_recordno integer NOT NULL,
    trd_eventtype text NOT NULL,
    trd_eventdate text NOT NULL,
    trd_eventdateto text NOT NULL,
    trd_entrydatetime text NOT NULL,
    trd_eventcode text NOT NULL,
    trd_exceptiontype text NOT NULL,
    trd_expirydate text NOT NULL
);


ALTER TABLE public.trd_event OWNER TO root;

--
-- Name: trd_event_narrative; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_event_narrative (
    trd_recordno integer NOT NULL,
    trd_narrative_code text NOT NULL,
    trd_seqno text NOT NULL,
    trd_narrative text NOT NULL
);


ALTER TABLE public.trd_event_narrative OWNER TO root;

--
-- Name: trd_external_ref; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_external_ref (
    trd_recordno integer NOT NULL,
    ext_ref_extreftype text NOT NULL,
    ext_ref_extref text NOT NULL
);


ALTER TABLE public.trd_external_ref OWNER TO root;

--
-- Name: trd_inst_ext; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_inst_ext (
    trd_recordno integer NOT NULL,
    trd_service text NOT NULL,
    trd_extref text NOT NULL
);


ALTER TABLE public.trd_inst_ext OWNER TO root;

--
-- Name: trd_instruction; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_instruction (
    trd_recordno integer NOT NULL,
    trd_procaction text NOT NULL,
    trd_destination text NOT NULL,
    trd_procstatus text NOT NULL,
    trd_recordidentifier text NOT NULL,
    trd_recordversion text NOT NULL,
    trd_instformat text NOT NULL,
    trd_tradeparty text NOT NULL,
    trd_partyref text NOT NULL,
    trd_deliverytype text NOT NULL,
    trd_addresscode text NOT NULL,
    trd_servicestatus text NOT NULL,
    trd_noofcopies text NOT NULL,
    trd_duedatetime text NOT NULL
);


ALTER TABLE public.trd_instruction OWNER TO root;

--
-- Name: trd_instruction_effect; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_instruction_effect (
    trd_recordno integer NOT NULL,
    trd_eventtype text NOT NULL,
    trd_eventdate text NOT NULL,
    trd_eventdateto text NOT NULL,
    trd_entrydatetime text NOT NULL,
    trd_eventcode text NOT NULL,
    trd_exceptiontype text NOT NULL,
    trd_expirydate text NOT NULL
);


ALTER TABLE public.trd_instruction_effect OWNER TO root;

--
-- Name: trd_instrument; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_instrument (
    trd_recordno integer NOT NULL,
    trd_instrtype text NOT NULL,
    trd_p2000instrref text NOT NULL,
    trd_instrreftype text NOT NULL,
    trd_instrref text NOT NULL,
    trd_longname text NOT NULL,
    trd_quantity text NOT NULL
);


ALTER TABLE public.trd_instrument OWNER TO root;

--
-- Name: trd_journal; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_journal (
    trd_recordno integer NOT NULL,
    trd_accounts_company text NOT NULL,
    trd_journal_type text NOT NULL,
    trd_posting_type text NOT NULL,
    trd_journal_no text NOT NULL,
    trd_procaction text NOT NULL
);


ALTER TABLE public.trd_journal OWNER TO root;

--
-- Name: trd_link; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_link (
    trd_recordno integer NOT NULL,
    trd_link_type_wil text NOT NULL,
    trd_main_record_no_wil text NOT NULL,
    trd_sub_recordno_wil text NOT NULL,
    trd_link_qty_wil text NOT NULL,
    trd_link_status_wil text NOT NULL
);


ALTER TABLE public.trd_link OWNER TO root;

--
-- Name: trd_party; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_party (
    trd_recordno integer NOT NULL,
    trd_trade_party text NOT NULL,
    trd_partyref text NOT NULL,
    trd_partyref_type_text text NOT NULL,
    trd_ext_partyref text NOT NULL,
    trd_longname text NOT NULL
);


ALTER TABLE public.trd_party OWNER TO root;

--
-- Name: trd_party_driver; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_party_driver (
    trd_recordno integer NOT NULL,
    trd_trade_party text NOT NULL,
    trd_driver_type text NOT NULL,
    trd_driver_code text NOT NULL
);


ALTER TABLE public.trd_party_driver OWNER TO root;

--
-- Name: trd_processing; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_processing (
    trd_recordno integer NOT NULL,
    trd_proc_alias text NOT NULL,
    trd_proc_action text NOT NULL,
    trd_due_datetime text NOT NULL,
    trd_done_datetime text NOT NULL
);


ALTER TABLE public.trd_processing OWNER TO root;

--
-- Name: trd_processing_event; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_processing_event (
    trd_recordno integer NOT NULL,
    trd_eventtype text NOT NULL,
    trd_eventdate text NOT NULL,
    trd_eventdateto text NOT NULL,
    trd_entrydatetime text NOT NULL,
    trd_eventcode text NOT NULL,
    trd_exceptiontype text NOT NULL,
    trd_expirydate text NOT NULL
);


ALTER TABLE public.trd_processing_event OWNER TO root;

--
-- Name: trd_rate; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_rate (
    trd_recordno integer NOT NULL,
    trd_charge_levy_type text NOT NULL,
    trd_actual_charge text NOT NULL,
    trd_amount_type text NOT NULL,
    trd_rate_instrref_type text NOT NULL,
    trd_rate_instrref text NOT NULL,
    trd_rate_instrid text NOT NULL,
    trd_rate_entered text NOT NULL,
    trd_charge_rate text NOT NULL,
    trd_mult_divind text NOT NULL
);


ALTER TABLE public.trd_rate OWNER TO root;

--
-- Name: trd_ref_date; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_ref_date (
    trd_recordno integer NOT NULL,
    datetype character varying(4) NOT NULL,
    datewil date NOT NULL,
    refdatetime character varying(8) NOT NULL,
    dateversionuser character varying(3) NOT NULL
);


ALTER TABLE public.trd_ref_date OWNER TO root;

--
-- Name: trd_settlement; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_settlement (
    trd_recordno integer NOT NULL,
    trd_deliverytype text NOT NULL,
    trd_settleeventinstr text NOT NULL,
    trd_settleterms text NOT NULL,
    trd_originalqty text NOT NULL,
    trd_openqty text NOT NULL,
    trd_settledate text NOT NULL,
    trd_delrecind text NOT NULL,
    trd_settlestatus text NOT NULL,
    trd_tradestatus text NOT NULL,
    trd_settlenarrative1 text NOT NULL,
    trd_settlenarrative2 text NOT NULL,
    trd_settlenarrative3 text NOT NULL,
    trd_settlenarrative4 text NOT NULL,
    trd_settlenarrative5 text NOT NULL,
    trd_settlenarrative6 text NOT NULL,
    trd_settlenarrative7 text NOT NULL,
    trd_settlenarrative8 text NOT NULL,
    trd_dompaliaswil text NOT NULL,
    trd_dompaliasdesc text NOT NULL,
    trd_dompdepottypewil text NOT NULL,
    trd_dompdaccwil text NOT NULL,
    trd_dompservicewil text NOT NULL,
    trd_secpaliaswil text NOT NULL,
    trd_secpservicewil text NOT NULL
);


ALTER TABLE public.trd_settlement OWNER TO root;

--
-- Name: trd_trade; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_trade (
    trd_uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    trd_recordno integer NOT NULL,
    trd_glosstraderef integer NOT NULL,
    trd_versiono integer NOT NULL,
    trd_origin text NOT NULL,
    trd_tradetype text NOT NULL,
    trd_settlementstatus text NOT NULL,
    trd_tradestatus character varying(1) NOT NULL,
    trd_originversion integer NOT NULL
);


ALTER TABLE public.trd_trade OWNER TO root;

--
-- Name: trd_trade_code; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_trade_code (
    trd_recordno integer NOT NULL,
    trd_ext text NOT NULL,
    trd_radeclass text NOT NULL,
    trd_radecode text NOT NULL
);


ALTER TABLE public.trd_trade_code OWNER TO root;

--
-- Name: trd_trade_narrative; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.trd_trade_narrative (
    trd_recordno integer NOT NULL,
    trd_narrative_code text NOT NULL,
    trd_seqno text NOT NULL,
    trd_narrative text NOT NULL
);


ALTER TABLE public.trd_trade_narrative OWNER TO root;

--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    username character varying NOT NULL,
    hashed_password character varying NOT NULL,
    full_name character varying NOT NULL,
    email character varying NOT NULL,
    password_changed_at timestamp with time zone DEFAULT now() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO root;

--
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- Name: entries id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.entries ALTER COLUMN id SET DEFAULT nextval('public.entries_id_seq'::regclass);


--
-- Name: transfers id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers ALTER COLUMN id SET DEFAULT nextval('public.transfers_id_seq'::regclass);

