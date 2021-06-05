-- 数据库
CREATE DATABASE im_db;

-- 用户表
CREATE TABLE user_t (
    uid      INT         PRIMARY KEY,
    uuid     VARCHAR(64),
    email    VARCHAR(64),
    name     VARCHAR(10),
    password VARCHAR(255),
    createAt DATE
)

-- 聊天数据表
CREATE TABLE chat_t (
    sendId INT,
    recvId INT,
    sendDate   DATE,
    msg    VARCHAR(255)
)