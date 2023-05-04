create database panDemo;

use panDemo;
-- auto-generated definition
create table user
(
    uid       int auto_increment comment '用户ID'
        primary key,
    username  varchar(70)                          not null comment '用户名',
    uemail    varchar(100)                         null comment '用户邮件',
    isAdmin   tinyint(1) default 0                 not null comment '是否是管理',
    password  varchar(255)                         not null comment '密码',
    create_at timestamp  default CURRENT_TIMESTAMP not null comment '创建时间'
);


w