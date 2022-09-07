create table instr_instruments
(
    instr_instref    text not null
        primary key,
    instr_instgrp    text not null,
    instr_longdesc   text not null,
    instr_denminst   text not null,
    instr_pricedps   text not null,
    instr_divisor    text not null,
    instr_multiplier text not null,
    instr_pricetype  text not null,
    instr_tick       text not null,
    instr_accrued    text not null,
    instr_minqty     text not null,
    instr_qtydps     text not null,
    instr_shrtdesc   text not null,
    instr_market     text not null,
    instr_settinst   text not null,
    instr_book       text not null,
    instr_maxmovt    text not null,
    instr_minmovt    text not null,
    instr_active     text not null,
    instr_verdat     text not null
);

alter table instr_instruments
    owner to root;

