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
    sendDate TIMESTAMP    NOT NULL,
    msg      VARCHAR(255) NOT NULL,
    unRead   BOOLEAN      NOT NULL
);

-- 测试数据
INSERT INTO user_t (uid, token, email, name, password, createAt) VALUES
    (1001, 'qvqryt36gvret34', 'beta@example.com', 'beta', '123456', '2020-01-01'),
    (1002, '1nfuin34jtbr323', 'alpha@example.com', 'alpha', '123456', '2020-01-02'),
    (1003, '34y5whaebray45w', 'stable@example.com', 'stable', '123456', '2020-01-03');

INSERT INTO friend_t (uid, fid) VALUES
    (1001, 1002), (1002, 1001),
    (1001, 1003), (1003, 1001);

INSERT INTO chat_t (sendId, recvId, sendDate, msg, unRead) VALUES
    (1001, 1002, '2021-06-06 10:10:10', 'hello', TRUE),
    (1001, 1002, '2021-06-06 10:10:12', 'world', TRUE),
    (1002, 1001, '2021-06-07 10:10:14', '你好', TRUE),
    (1002, 1001, '2021-06-07 10:10:20', '世界', TRUE);