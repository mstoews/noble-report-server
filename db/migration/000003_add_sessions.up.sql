CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");


CREATE TABLE trade (
     trd_recordno          INTEGER  NOT NULL PRIMARY KEY
    ,trd_glosstraderef     INTEGER  NOT NULL
    ,trd_versiono          INTEGER  NOT NULL
    ,trd_origin            VARCHAR(2) NOT NULL
    ,trd_tradetype         VARCHAR(4) NOT NULL
    ,trd_settlementstatus  VARCHAR(4) NOT NULL
    ,trd_tradestatus       VARCHAR(1) NOT NULL
    ,trd_originversion     INTEGER  NOT NULL
    ,trd_date              DATE NOT NULL

);

CREATE TABLE IF NOT EXISTS external_ref (
   trd_recordno        INTEGER  NOT NULL PRIMARY KEY     
  ,ext_ref_extreftype  VARCHAR(4) NOT NULL
  ,ext_ref_extref      VARCHAR(3) NOT NULL
);


CREATE TABLE ref_date (
   trd_recordno     INTEGER  NOT NULL PRIMARY KEY 
  ,datetype     VARCHAR(4) NOT NULL
  ,datewil      DATE  NOT NULL
  ,refdatetime         VARCHAR(8) NOT NULL
  ,dateversionuser  VARCHAR(3) NOT NULL
);




