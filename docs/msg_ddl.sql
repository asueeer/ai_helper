-- 消息发件箱
create table message_from
(
    id          bigserial primary key  not null,
    message_id  bigint unique          not null,
    conv_id     bigint                 not null,
    sender_id   bigint                 not null,
    receiver_id bigint                 not null,
    content     json                   not null,
    type        varchar(16) default '' not null,
    seq_id      bigint                 not null,
    created_at  timestamp              not null,
    updated_at  timestamp,
    deleted_at  timestamp   default null
);

comment
on table "message_from" is '消息发件箱';
comment
on column "message_from"."id" is '主键自增id,无业务意义';
comment
on column "message_from"."type" is '消息类型';
comment
on column "message_from"."conv_id" is '会话id';
comment
on column "message_from".sender_id is '发送者id';
comment
on column "message_from".receiver_id is '接收者id';
comment
on column "message_from".content is '消息发送内容';
comment
on column "message_from".seq_id is '序列号(用于保序)';
comment
on column "message_from".created_at is '创建时间戳';
comment
on column "message_from".updated_at is '更新时间戳';
comment
on column "message_from".deleted_at is '删除时间戳';


-- 消息收件箱
create table message_to
(
    id         bigserial primary key not null,
    message_id bigint                not null,
    conv_id    bigint                not null,
    owner_id   bigint                not null,
    seq_id     bigint                not null,
    created_at timestamp             not null,
    updated_at timestamp,
    deleted_at timestamp default null
);

comment
on table "message_to" is '消息发件箱';
comment
on column "message_to"."id" is '主键自增id,无业务意义';
comment
on column "message_to"."message_id" is '消息id';
comment
on column "message_to"."owner_id" is '收件箱所有者的用户id';
comment
on column "message_to"."conv_id" is '会话id';
comment
on column "message_to".seq_id is '序列号, 用于保序';
comment
on column "message_to".created_at is '创建时间戳';
comment
on column "message_to".updated_at is '更新时间戳';
comment
on column "message_to".deleted_at is '删除时间戳';
