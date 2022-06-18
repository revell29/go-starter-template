CREATE TABLE IF NOT EXISTS users (
    user_id            char(22) not null constraint users_pk primary key,
    name               varchar,
    email              varchar,
    hash               varchar,
    password           varchar,
    country            varchar,
    e164_code          varchar,
    phone_number       varchar,
    is_email_confirmed bool,
    is_phone_confirmed bool,
    created_date       timestamptz default now(),
    created_by         bigint,
    updated_date       timestamptz,
    updated_by         bigint
);

create unique index users_user_id_uindex
    on users (user_id);

