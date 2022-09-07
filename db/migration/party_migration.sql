create table pty_party
(
    pty_partyref  text not null
        primary key,
    pty_category  text not null,
    pty_longdesc  text not null,
    pty_holiday   text not null,
    pty_country   text not null,
    pty_location  text not null,
    pty_shrtdesc  text not null,
    pty_partynam1 text not null,
    pty_partynam2 text not null,
    pty_partynam3 text not null,
    pty_active    text not null,
    pty_verdat    text not null
);

alter table pty_party
    owner to root;

