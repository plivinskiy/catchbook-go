create table user
(
    id         binary(36)         not null primary key,
    status     smallint default 1 not null,
    email      varchar(255)       not null,
    username   varchar(255)       not null,
    password   varchar(255)       not null,
    firstname  varchar(255)       not null,
    lastname   varchar(255)       null,
    created_at datetime           not null,
    constraint user_uniq unique (email)
);

