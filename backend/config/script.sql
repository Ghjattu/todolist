create table category
(
    id   bigint auto_increment comment '编号'
        primary key,
    name varchar(255) not null comment '分类名'
);

create table item
(
    id            bigint auto_increment comment '编号'
        primary key,
    author_id     bigint       not null comment '作者编号',
    body          varchar(255) not null comment '内容',
    reminder_time datetime     not null comment '提醒时间',
    priority      bigint       not null comment '优先级',
    category_id   bigint       not null comment '分类编号',
    status        varchar(255) not null comment '状态'
);

create table user
(
    id       bigint auto_increment comment '用户编号'
        primary key,
    username varchar(255) not null comment '用户名',
    password varchar(255) not null comment '密码'
);


