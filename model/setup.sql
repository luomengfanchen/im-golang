-- 数据库
CREATE DATABASE im_db;

-- 用户表
CREATE TABLE user_t (
    uid      SERIAL       PRIMARY KEY,
    token    VARCHAR(64)  NOT NULL,
    email    VARCHAR(64)  NOT NULL,
    name     VARCHAR(10)  NOT NULL,
    password VARCHAR(255) NOT NULL,
    createAt DATE         NOT NULL
);

-- 好友表
CREATE TABLE friend_t (
    uid INT NOT NULL,
    fid INT NOT NULL
);

-- 聊天数据表
CREATE TABLE chat_t (
    cid      SERIAL       PRIMARY KEY,
    sendId   INT          NOT NULL,
    recvId   INT          NOT NULL,
    sendDate DATE         NOT NULL,
    msg      VARCHAR(255) NOT NULL,
    unRead   BOOLEAN      NOT NULL
);

-- 测试数据
INSERT INTO user_t (uid, token, email, name, password, createAt) VALUES
    (1001, 'qvqryt36gvret34', 'beta@example.com', 'beta', '123456', '2020-01-01'),
    (1002, '1nfuin34jtbr323', 'alpha@example.com', 'alpha', '123456', '2020-01-02');

INSERT INTO friend_t (uid, fid) VALUES
    (1001, 1002), (1002, 1001);

INSERT INTO chat_t (sendId, recvId, sendDate, msg, unRead) VALUES
    (1001, 1002, '2021-06-06', 'hello', TRUE),
    (1001, 1002, '2021-06-06', 'world', TRUE),
    (1002, 1001, '2021-06-07', '你好', TRUE),
    (1002, 1001, '2021-06-07', '世界', TRUE);