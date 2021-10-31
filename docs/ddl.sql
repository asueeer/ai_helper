-- 用户表
-- 存放用户的基本信息: 昵称、头像等
create table user_center
(
    id         bigserial primary key   not null,
    user_id    bigint unique           not null,
    nickname   varchar(32)  default '' not null,
    head_url   varchar(256) default '' not null,
    created_at timestamp               not null,
    updated_at timestamp,
    deleted_at timestamp    default null
);
comment
on column "user_center"."id" is '主键自增id,无业务意义';
comment
on column "user_center"."user_id" is '用户id';
comment
on column "user_center"."nickname" is '昵称';
comment
on column "user_center"."created_at" is '创建时间戳';
comment
on column "user_center"."updated_at" is '更新时间戳';
comment
on column "user_center"."deleted_at" is '删除时间戳';
comment
on table "user_center" is '用户表';

-- 用户账户信息
-- 存放用户的登陆信息: 手机号、邮箱、微信唯一标识、用户名等
create table user_account
(
    id           bigserial primary key         not null,
    user_id      bigint unique                 not null,
    username     varchar(32) default ''        not null,
    phone_number varchar(32) unique unique     not null,
    email        varchar(32) default ''        not null,
    wx_open_id   varchar(32) default '' unique not null,
    created_at   timestamp                     not null,
    updated_at   timestamp,
    deleted_at   timestamp   default null
);

comment
on column "user_account"."id" is '主键自增id,无业务意义';
comment
on column "user_account"."user_id" is '用户id';
comment
on column "user_account"."username" is '用户名';
comment
on column "user_account"."phone_number" is '手机号';
comment
on column "user_account"."email" is '电子邮箱';
comment
on column "user_account"."wx_open_id" is '微信open_id';
comment
on column "user_account"."created_at" is '创建时间戳';
comment
on column "user_account"."updated_at" is '更新时间戳';
comment
on column "user_account"."deleted_at" is '删除时间戳';
comment
on table "user_center" is '用户表';


INSERT INTO user_center (user_id, nickname, head_url, created_at, updated_at)
VALUES (3926077, '测试账号1', 'https://images.gitee.com/uploads/images/2021/0731/134131_a864d20c_7809561.jpeg', now(),
        now());

insert into user_account (user_id, phone_number, created_at, updated_at)
VALUES (3926077, '18878907302', now(), now());

INSERT INTO user_center (user_id, nickname, head_url, created_at, updated_at)
VALUES (305088049, '阿苏EEer', 'https://images.gitee.com/uploads/images/2021/0731/134131_a864d20c_7809561.jpeg', now(),
        now());

insert into user_account (user_id, username, phone_number, created_at, updated_at)
VALUES (305088049, 'asueeer', 'admin', now(), now());

-- 会话表
create table conversation
(
    id         bigserial primary key  not null,
    conv_id    bigint unique          not null,
    type       varchar(64) default '' not null,
    creator    bigint                 not null,
    status     varchar(64) default '' not null,
    timestamp  timestamp              not null,
    created_at timestamp              not null,
    updated_at timestamp,
    deleted_at timestamp   default null
);

comment
on table "conversation" is '会话表';
comment
on column "conversation"."id" is '主键自增id,无业务意义';
comment
on column "conversation"."conv_id" is '会话id';
comment
on column "conversation".type is '会话类型; 枚举值: "helper"-客服, "chat"-普通会话, "group"-群聊';
comment
on column "conversation".creator is '会话创建者';
comment
on column "conversation".status is '会话状态';
comment
on column "conversation".created_at is '创建时间戳';
comment
on column "conversation".updated_at is '更新时间戳';
comment
on column "conversation".deleted_at is '删除时间戳';
comment
on column "conversation".timestamp is '会话时间戳';


-- 用户与会话的关系表
create table user_conv_rel
(
    id         bigserial primary key  not null,
    rel_id     bigint unique          not null,
    conv_id    bigint                 not null,
    user_id    bigint                 not null,
    role       varchar(32) default '' not null,
    created_at timestamp              not null,
    updated_at timestamp,
    deleted_at timestamp   default null
);

comment
on table "user_conv_rel" is '用户会话关系表';
comment
on column "user_conv_rel"."id" is '主键自增id,无业务意义';
comment
on column "user_conv_rel"."conv_id" is '会话id';
comment
on column "user_conv_rel".user_id is '用户id';
comment
on column "user_conv_rel".role is '用户在会话中的角色';
comment
on column "user_conv_rel".created_at is '创建时间戳';
comment
on column "user_conv_rel".updated_at is '更新时间戳';
comment
on column "user_conv_rel".deleted_at is '删除时间戳';



-- 消息发件箱
create table message_from
(
    id          bigserial primary key  not null,
    message_id  bigint unique          not null,
    conv_id     bigint                 not null,
    sender_id   bigint                 not null,
    receiver_id bigint                 not null,
    ref_msg_id  bigint      default 0  not null,
    content     json                   not null,
    type        varchar(16) default '' not null,
    timestamp   timestamp              not null,
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
on column "message_from".ref_msg_id is '引用的消息id';
comment
on column "message_from".content is '消息发送内容';
comment
on column "message_from".timestamp is '消息时间戳';
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
    timestamp  timestamp             not null,
    has_read   int4      default 1   not null,
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
on column "message_to".has_read is '消息是否已读; 1: 未读; 2: 已读';
comment
on column "message_to".timestamp is '消息时间戳';
comment
on column "message_to".created_at is '创建时间戳';
comment
on column "message_to".updated_at is '更新时间戳';
comment
on column "message_to".deleted_at is '删除时间戳';