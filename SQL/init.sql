drop table if exists user;
create table user
(
    id          integer      not null AUTO_INCREMENT,
    name        varchar(125) not null comment '用户名',
    password    varchar(32)  not null comment '用户密码',
    create_time timestamp    not null default current_timestamp comment '创建时间',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    PRIMARY KEY (id)
)