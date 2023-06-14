CREATE TABLE IF NOT EXISTS history
(
    id               uuid                     default uuid_generate_v4()    not null primary key,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP     not null,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP     not null,
    deleted_at       timestamp with time zone,
    endpoint         varchar(255)             default ''::character varying not null
);