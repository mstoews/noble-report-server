create table public.gl_account_type
(
    type        varchar(20) not null
        constraint gl_accounts_type_pk
            primary key,
    description varchar(50),
    create_date date,
    create_user varchar(200),
    update_date date,
    update_user varchar(200)
);

alter table public.gl_account_type
    owner to mstoews;

create table public.gl_period
(
    period_id   integer not null,
    period_year integer not null,
    start_date  date,
    end_date    date,
    description varchar(40),
    create_date date,
    create_user varchar(200),
    update_date date,
    update_user varchar(200)
);

alter table public.gl_period
    owner to mstoews;


alter table public.gl_period
    owner to mstoews;

create table public.kb_priority
(
    priority    text not null
        primary key,
    description text,
    updatedte   date,
    updateusr   text
);

alter table public.kb_priority
    owner to mstoews;

create table public.kb_status
(
    status      text not null
        primary key,
    description text,
    updatedte   timestamp(3),
    updateusr   text
);

alter table public.kb_status
    owner to mstoews;

create table public.kb_subtask
(
    task_id  text not null,
    subid    text not null,
    "desc"   text,
    status   text,
    summary  text,
    type     text,
    estimate integer,
    primary key (task_id, subid)
);

alter table public.kb_subtask
    owner to mstoews;


create table public.kb_task
(
    task_id      text not null primary key,
    title        text,
    status       text,
    summary      text,
    type         text,
    priority     text,
    tags         text,
    estimate     integer,
    assignee     text,
    rankid       integer,
    color        text,
    classname    text,
    dependencies text default ''::text,
    description  text,
    due_date     date default CURRENT_TIMESTAMP,
    start_date   date default CURRENT_TIMESTAMP
);


alter table public.kb_task
    owner to mstoews;

create table public.kb_team
(
    team_member text not null
        primary key,
    first_name  text,
    last_name   text,
    location    text,
    title       text,
    updatedte   timestamp(3),
    updateusr   text
);

alter table public.kb_team
    owner to mstoews;

create table public.kb_type
(
    type        text not null
        primary key,
    description text,
    updatedte   timestamp(3),
    updateusr   text
);

alter table public.kb_type
    owner to mstoews;


create table public.gl_distribution
(
    dist_parent_id integer not null,
    dist_child_id  integer not null,
    journal_id     integer,
    journal_subid  integer,
    created_date   date,
    update_date    date,
    created_user   varchar(200),
    constraint gl_distribution_pk
        primary key (dist_parent_id, dist_child_id)
);

create table public.gl_funds (
    fund varchar PRIMARY KEY,
    description varchar,
    create_date date,
    create_user varchar

);

alter table public.gl_distribution
    owner to mstoews;

create table public.gl_trial_balance
(
    account              varchar(10) not null,
    child                varchar(10) not null,
    start_date           date,
    end_date             date,
    accounting_period_id integer,
    description          varchar(40),
    type                 varchar(20),
    debit                numeric(16, 2),
    credit               numeric(16, 2),
    create_date          date,
    create_user          varchar(200),
    update_date          date,
    update_user          varchar(200),
    constraint gl_trial_balance_pk
        primary key (account, child)
);

alter table public.gl_trial_balance
    owner to mstoews;

create table public.gl_accounts
(
    account        integer not null,
    child          integer not null,
    parent_account boolean,
    type           varchar(20),
    sub_type       varchar(20),
    description    varchar(40),
    balance        numeric(16, 2),
    comments       varchar(200),
    create_date    date,
    create_user    varchar(20),
    update_date    date,
    update_user    varchar(20),
    constraint gl_accounts_pk
        primary key (account, child)
);

alter table public.gl_accounts
    owner to mstoews;

create table public.gl_journal_header
(
    journal_id       integer not null
        constraint gl_journal_header_pk
            primary key,
    description      varchar(100),
    booked           boolean,
    booked_date      date,
    booked_user      varchar(200),
    create_date      date,
    create_user      varchar(200),
    period           integer,
    transaction_date date,
    status           varchar(10),
    type             varchar(20),
    amount           numeric(16, 2)
);

alter table public.gl_journal_header
    owner to mstoews;

create table public.gl_journal_detail
(
    journal_id    integer not null,
    journal_subid integer not null,
    account       integer,
    child         integer,
    sub_type      varchar(30),
    description   varchar(100),
    debit         numeric(16, 2),
    credit        numeric(16, 2),
    create_date   date,
    create_user   varchar(20),
    constraint gl_journal_detail_pk
        primary key (journal_id, journal_subid)
);

alter table public.gl_journal_detail
    owner to mstoews;

create table public.gl_distribution_ledger
(
    account         integer        not null,
    child           integer        not null,
    period          integer        not null,
    period_year     integer        not null,
    description     varchar(100),
    opening_balance numeric(16, 2) not null,
    debit_balance   numeric(16, 2) not null,
    credit_balance  numeric(16, 2) not null,
    closing_balance numeric(16, 2) not null,
    update_date     date,
    created_user    varchar(200),
    constraint gl_distribution_ledger_pk
        primary key (account, child, period)
);

alter table public.gl_distribution_ledger
    owner to mstoews;


