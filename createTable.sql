create table "user"
(
    id         serial
        constraint user_pk
            primary key,
    name varchar(100) not null,
    email varchar(100) not null,
    login varchar(100) not null,
    password varchar(100) not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    deleted_at timestamp,
    is_deleted boolean   default false             not null
);

alter table "user"
    owner to max;

create unique index user_id_uindex
    on "user" (id);

create table "chat"
(
    id         serial
        constraint chat_pk
            primary key,
    name varchar(100) not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    deleted_at timestamp,
    is_deleted boolean   default false             not null
);

alter table "chat"
    owner to max;

create unique index chat_id_uindex
    on "chat" (id);

create table "chat_user_relation"
(
    id         serial
        constraint chat_user_relation_pk
            primary key,
    chat_id int not null,
    user_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    deleted_at timestamp,
    is_deleted boolean   default false             not null,
    constraint chat_user_relation_chat_id_fk
        foreign key (chat_id) references "chat"(id),
    constraint chat_user_relation_user_id_fk
        foreign key (user_id) references "user"(id)
);

alter table "chat_user_relation"
    owner to max;

create unique index chat_user_relation_id_uindex
    on "chat_user_relation" (id);


create table "message"
(
    id         serial
        constraint message_pk
            primary key,
    message varchar(255) not null,
    chat_user_relation_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    deleted_at timestamp,
    is_deleted boolean   default false             not null,

    constraint message_chat_user_relation_id_fk
        foreign key (chat_user_relation_id) references "chat_user_relation"(id)
);

alter table "message"
    owner to max;

create unique index message_id_uindex
    on "message" (id);


select * from user;
SELECT email
FROM "user";


