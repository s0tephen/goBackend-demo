create database panDemo;

use panDemo;
-- auto-generated definition
create table user
(
    uid       int auto_increment comment '用户ID'
        primary key,
    username  varchar(70)                          not null comment '用户名',
    uemail    varchar(100)  null comment '用户邮件',
    isAdmin  tinyint(1) default 0                 not null comment '是否是管理',
    password  varchar(255)                         not null comment '密码',
    create_at timestamp  default CURRENT_TIMESTAMP not null comment '创建时间'
);


-- auto-generated definition
create table login_session
(
    id         int auto_increment comment 'token_Id',
    token      text         not null comment '令牌',
    uid        int          not null comment '用户ID',
    login_time timestamp    null comment '登陆时间',
    login_ip   varchar(255) null comment '登陆IP',
    primary key (id)
);


-- auto-generated definition
create table message
(
    mid       int auto_increment comment '消息ID',
    uname     varchar(70)                        null comment '用户名字',
    content   longtext                            null comment '消息内容',
    create_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '创建时间',
    primary key (mid)
);


create table feedback
(
    fid   int auto_increment comment '反馈ID',
    fUser varchar(70) not null comment '反馈者',
    fMsg  varchar(255) not null comment '反馈内容',
    fTime timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '反馈时间',
    primary key (fid)
)
