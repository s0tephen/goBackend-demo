create database demo;

use demo;

create table user
(
    uid      int auto_increment,
    username varchar(255) not null,
    password varchar(255) not null,
    creat_at timestamp    not null default NOW(),
    primary key (uid)
);

create table loginStatus
(
    id      int auto_increment,
    token   text not null,
    user_id int  not null,
    primary key (id)
);