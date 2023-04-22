create database panDemo;

use panDemo;
-- auto-generated definition
create table user
(
    uid       int auto_increment comment '用户ID'
        primary key,
    username  varchar(255)                        not null comment '用户名',
    uemail    varchar(100)                        null comment '用户邮件',
    password  varchar(255)                        not null comment '密码',
    create_at timestamp default CURRENT_TIMESTAMP not null comment '创建时间'
);


-- auto-generated definition
create table login_session
(
    id         int auto_increment comment 'token_Id'
        primary key,
    token      text         not null comment '令牌',
    uid        int          not null comment '用户ID',
    login_time timestamp    null comment '登陆时间',
    login_ip   varchar(255) null comment '登陆IP'
);



-- auto-generated definition
create table message
(
    mid       int auto_increment comment '消息ID'
        primary key,
    uname     varchar(155)                        null comment '用户名字',
    create_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '创建时间',
    content   longtext                            null comment '消息内容'
);





