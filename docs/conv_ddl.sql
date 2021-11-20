-- 会话表
create table conversation
(
    id          bigserial primary key  not null,
    conv_id     bigint unique          not null,
    type        varchar(64) default '' not null,
    creator     bigint                 not null,
    status      varchar(64) default '' not null,
    timestamp   timestamp              not null,
    seq_id      bigint      default 0  not null,
    last_msg_id bigint      default 0  not null,
    created_at  timestamp              not null,
    updated_at  timestamp,
    deleted_at  timestamp   default null
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
on column "conversation".seq_id is '会话序列号(用于保序)';
comment
on column "conversation".last_msg_id is '最近的一条消息';
comment
on column "conversation".timestamp is '会话时间戳';

-- 用户与会话的关系表
create table user_conv_rel
(
    id           bigserial primary key  not null,
    rel_id       bigint unique          not null,
    conv_id      bigint                 not null,
    user_id      bigint                 not null,
    role         varchar(32) default '' not null,
    unread_cnt   int         default 0  not null,
    participants bigint[] not null,
    created_at   timestamp              not null,
    updated_at   timestamp,
    deleted_at   timestamp   default null
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

comment
on column "user_conv_rel".unread_cnt is '未读消息数';

comment
on column "user_conv_rel".participants is '会话参与者';
