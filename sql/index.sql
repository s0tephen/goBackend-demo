/*
 Date: 30/05/2023 11:15:59
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for feedback
-- ----------------------------
DROP TABLE IF EXISTS `feedback`;
CREATE TABLE `feedback`
(
    `fid`   bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '反馈ID',
    `fUser` varchar(70) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL COMMENT '反馈者',
    `fMsg`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '反馈内容',
    `fTime` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '反馈时间',
    PRIMARY KEY (`fid`),
    KEY     `idx_fUser` (`fUser`),
    KEY     `idx_fTime` (`fTime`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of feedback
-- ----------------------------
BEGIN;
INSERT INTO `feedback` (`fid`, `fUser`, `fMsg`, `fTime`)
VALUES (1, 'test', '垃圾网站', '2023-05-02 14:22:31');
INSERT INTO `feedback` (`fid`, `fUser`, `fMsg`, `fTime`)
VALUES (2, 'test', '垃圾网站', '2023-05-02 14:24:01');
INSERT INTO `feedback` (`fid`, `fUser`, `fMsg`, `fTime`)
VALUES (3, 'stephen', '垃圾网站', '2023-05-13 10:21:13');
INSERT INTO `feedback` (`fid`, `fUser`, `fMsg`, `fTime`)
VALUES (4, 'stephen', '垃圾网站1111', '2023-05-13 10:21:41');
COMMIT;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`
(
    `id`       int                                                           NOT NULL AUTO_INCREMENT COMMENT 'zhujian',
    `filename` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `md5`      varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL,
    `uploaded` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL,
    `size`     int                                                           NOT NULL,
    `path`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `up_time`  datetime                                                      NOT NULL,
    PRIMARY KEY (`id`),
    KEY        `file_path` (`path`) USING BTREE,
    KEY        `md5` (`md5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of file
-- ----------------------------
BEGIN;
INSERT INTO `file` (`id`, `filename`, `md5`, `uploaded`, `size`, `path`, `up_time`)
VALUES (6, '5349163920ddd64848bb2c7e8b6062b6.gif', '109d5374bce8aad50fa95e60c6681a23', 'stephen', 21209,
        'static/data/109d5374bce8aad50fa95e60c6681a23', '2023-05-04 07:41:11');
INSERT INTO `file` (`id`, `filename`, `md5`, `uploaded`, `size`, `path`, `up_time`)
VALUES (7, '5f3bb0343073f56868dcd68936857cb6.gif', 'fe6a498315cf8a7b89947d86692cc771', 'stephen', 320012,
        'static/data/fe6a498315cf8a7b89947d86692cc771', '2023-05-04 07:41:11');
INSERT INTO `file` (`id`, `filename`, `md5`, `uploaded`, `size`, `path`, `up_time`)
VALUES (8, '5349163920ddd64848bb2c7e8b6062b6.gif', '109d5374bce8aad50fa95e60c6681a23', 'stephen', 21209,
        'static/data/109d5374bce8aad50fa95e60c6681a23', '2023-05-04 07:42:19');
INSERT INTO `file` (`id`, `filename`, `md5`, `uploaded`, `size`, `path`, `up_time`)
VALUES (9, '5f3bb0343073f56868dcd68936857cb6.gif', 'fe6a498315cf8a7b89947d86692cc771', 'stephen', 320012,
        'static/data/fe6a498315cf8a7b89947d86692cc771', '2023-05-04 07:42:19');
COMMIT;

-- ----------------------------
-- Table structure for label
-- ----------------------------
DROP TABLE IF EXISTS `label`;
CREATE TABLE `label`
(
    `Cid`   int                                                          NOT NULL AUTO_INCREMENT COMMENT '分类id',
    `Cname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名',
    PRIMARY KEY (`Cid`),
    KEY     `idx_Cname` (`Cname`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of label
-- ----------------------------
BEGIN;
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (3, '书籍');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (12, '公告');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (11, '其他');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (2, '动漫');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (8, '图片');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (4, '学习');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (1, '影视');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (6, '教程');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (10, '求助');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (7, '游戏');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (5, '软件');
INSERT INTO `label` (`Cid`, `Cname`)
VALUES (9, '音频');
COMMIT;

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`
(
    `ID`     int NOT NULL AUTO_INCREMENT,
    `UserID` int NOT NULL COMMENT '用户ID',
    `PostID` int NOT NULL COMMENT '帖子ID',
    PRIMARY KEY (`ID`, `UserID`, `PostID`),
    KEY      `idx_UserID` (`UserID`),
    KEY      `idx_PostID` (`PostID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of likes
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for login_session
-- ----------------------------
DROP TABLE IF EXISTS `login_session`;
CREATE TABLE `login_session`
(
    `id`         int NOT NULL AUTO_INCREMENT COMMENT 'token_Id',
    `token`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '令牌',
    `uid`        int NOT NULL COMMENT '用户ID',
    `login_time` timestamp NULL DEFAULT NULL COMMENT '登陆时间',
    `login_ip`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '登陆IP',
    PRIMARY KEY (`id`),
    KEY          `idx_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of login_session
-- ----------------------------
BEGIN;
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (26, '102f79db8a7844b7a09d66930b86e6a3', 3, '2023-05-13 10:20:32', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (27, 'a6efa4373bab44739116a3b98aaa2a10', 3, '2023-05-13 11:24:52', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (28, '993bfa184dd94d3d920f310303ed387f', 3, '2023-05-13 12:48:03', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (29, 'dec7861f62f24a51b1ba3cd7f589fa44', 3, '2023-05-17 08:51:50', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (30, 'de887e88d5e6465ba1bc361e41117481', 3, '2023-05-17 15:52:58', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (31, '6210dd3f529745a5b188dae60f62fb13', 3, '2023-05-17 15:54:14', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (32, 'e18f7f960e9c47248fae37fcf7877c07', 3, '2023-05-17 15:55:24', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (33, '35a3f882ff8740d289d89a2b9664447b', 3, '2023-05-17 19:14:57', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (34, '273a78019702450794ccba434e5480f9', 3, '2023-05-17 19:15:19', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (35, '11b74e97bdfc41f7b3cc21a6728e43bd', 3, '2023-05-17 19:15:34', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (36, '07be213c51274a63909c6142e997d72b', 3, '2023-05-17 19:15:59', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (37, '210f69e87acb48f6bcad4e0fd124a885', 3, '2023-05-17 19:17:45', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (38, 'fc06ad5ee06046c0b032c97d63b06012', 3, '2023-05-17 19:18:35', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (39, 'bc2b52a760a34b048a0759b844b1d4fc', 3, '2023-05-17 19:18:53', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (40, '3e8be28cb15547978e26607a8aaf5d35', 3, '2023-05-17 19:19:01', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (41, 'ef0cc678b8404f8f8274f6f31f76aca8', 3, '2023-05-17 19:20:15', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (42, '6e9eee1e1bd444d9b68081a825175017', 3, '2023-05-17 19:25:57', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (43, 'e29f866f23c14df3b6fabdda37769956', 3, '2023-05-17 19:26:21', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (44, '2ee5d3b20b8b48c3a43d4676114c71ed', 3, '2023-05-17 21:43:02', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (45, '910316886f54477daf3ff900ada29c6d', 3, '2023-05-18 07:46:58', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (46, 'c62d025bbf74477b827533094884c507', 9, '2023-05-18 11:11:37', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (47, '654ba507798c41b09220526fd38838c3', 3, '2023-05-18 11:12:18', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (48, '1dde35bad0ce4896aa0cfa8b9d2c92cb', 3, '2023-05-18 11:13:03', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (49, '1004c6238898410e859e93353968e042', 3, '2023-05-22 08:49:54', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (50, '0152139524b9465591164beebf1adb26', 3, '2023-05-22 08:50:37', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (51, '547f40e820454847ac6204d041956804', 3, '2023-05-22 08:51:11', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (52, '86ec92fa072c44c190fb5d0a7ce9c896', 3, '2023-05-23 09:05:12', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (53, '29cd49d5022c48db87bb4eb865e20a7d', 3, '2023-05-23 09:07:52', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (54, '6a3bed0a95844cf8a40ab051463bd689', 3, '2023-05-23 09:08:02', '127.0.0.1');
INSERT INTO `login_session` (`id`, `token`, `uid`, `login_time`, `login_ip`)
VALUES (55, 'b68411eeb90b4ca6a7b8aa3910991f52', 3, '2023-05-30 08:53:00', '127.0.0.1');
COMMIT;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`
(
    `mid`       int       NOT NULL AUTO_INCREMENT COMMENT '消息ID',
    `uname`     varchar(70) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名字',
    `content`   longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '消息内容',
    `create_at` timestamp NOT NULL                                           DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`mid`),
    KEY         `idx_uname` (`uname`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of message
-- ----------------------------
BEGIN;
INSERT INTO `message` (`mid`, `uname`, `content`, `create_at`)
VALUES (1, 'stephen', '111111111111', '2023-05-04 21:56:16');
INSERT INTO `message` (`mid`, `uname`, `content`, `create_at`)
VALUES (2, 'stephen', '1', '2023-05-04 21:56:22');
INSERT INTO `message` (`mid`, `uname`, `content`, `create_at`)
VALUES (3, 'stephen', '1', '2023-05-04 21:57:30');
INSERT INTO `message` (`mid`, `uname`, `content`, `create_at`)
VALUES (4, 'stephen', '112', '2023-05-04 21:57:54');
COMMIT;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`
(
    `pId`      int                                                          NOT NULL AUTO_INCREMENT COMMENT '帖子id',
    `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '发布者',
    `pTitle`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '帖子标题',
    `pCenter`  longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '帖子内容',
    `pImg`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '帖子图片',
    `pLabel`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '帖子标签',
    `pTime`    datetime                                                     NOT NULL COMMENT '发布时间',
    PRIMARY KEY (`pId`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of post
-- ----------------------------
BEGIN;
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (1, '张三', '国漫推荐', '推荐几部好看的国产动画：《哪吒之魔童降世》、《大鱼海棠》、《猫和老鼠》。',
        'https://unsplash.com/photos/7J4Tr4hOThg', '动漫', '2023-05-29 10:30:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (2, '李四', '我的阅读清单', '列出我最近阅读过的好书：《万历十五年》、《活着》、《百年孤独》。',
        'https://unsplash.com/photos/SbPO8c6L4u8', '书籍', '2023-05-28 14:45:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (3, '王五', 'Python 学习笔记', '记录一下自己学习 Python 的心得体会和笔记。',
        'https://unsplash.com/photos/O-RWye1b9JA', '学习', '2023-05-27 18:20:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (4, '赵六', '开发工具推荐', '分享一些常用的开发工具：VS Code、Sublime Text、PyCharm 等等。',
        'https://unsplash.com/photos/Iy59i0M7oP4', '软件', '2023-05-26 09:00:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (5, '张三', 'Java Web 教程', '分享一些 Java Web 开发的教程，包括 Spring Boot、MyBatis 等框架。',
        'https://unsplash.com/photos/Z6YxSbcIXT0', '教程', '2023-05-25 12:15:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (6, '李四', '游戏推荐', '推荐几款好玩的游戏：《英雄联盟》、《守望先锋》、《绝地求生》。',
        'https://unsplash.com/photos/qkA80ZzX2CM', '游戏', '2023-05-24 20:00:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (7, '王五', '美图欣赏', '分享一些美丽的图片，包括自然风光、人像摄影以及插画设计。',
        'https://unsplash.com/photos/lsdA8QpWN_A', '图片', '2023-05-23 16:30:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (8, '赵六', '音频剪辑软件推荐', '推荐一些用于音频剪辑和处理的软件，包括 Audacity、GarageBand 等等。',
        'https://unsplash.com/photos/36fWfW5UTXM', '音频', '2023-05-22 11:00:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (9, '张三', '求助：如何学习英语？', '想要提高自己的英语水平，不知道从何开始？请各位大佬给出建议！',
        'https://unsplash.com/photos/PiqZfESKt3k', '求助', '2023-05-21 08:00:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (10, '李四', '社区公告', '社区将于本月底进行维护，届时论坛将无法访问，请大家做好相应准备，谢谢您的配合！',
        'https://unsplash.com/photos/DQKyPCEo_TQ', '公告', '2023-05-20 15:30:00');
INSERT INTO `post` (`pId`, `username`, `pTitle`, `pCenter`, `pImg`, `pLabel`, `pTime`)
VALUES (11, '王五', '摄影技巧分享', '分享一些拍照和后期处理的技巧和心得体会。',
        'https://unsplash.com/photos/Dmg0cltK0VA', '其他', '2023-05-19 10:45:00');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `uid`       int                                                           NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username`  varchar(70) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL COMMENT '用户名',
    `avatar`    longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '用户头像',
    `uemail`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci          DEFAULT NULL COMMENT '用户邮件',
    `isAdmin`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是管理',
    `password`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
    `create_at` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`uid`),
    KEY         `idx_username` (`username`),
    KEY         `idx_uemail` (`uemail`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`uid`, `username`, `avatar`, `uemail`, `isAdmin`, `password`, `create_at`)
VALUES (2, 'test11', './static/images/test/avatar/test.png', 's0tephen@qq.com', 1,
        '$2a$10$5GZpeTPlV68C2MNMsxGjH.BR49K0S4XZcyoGK/6JMHYV77sDR0yda', '2023-05-03 17:28:41');
INSERT INTO `user` (`uid`, `username`, `avatar`, `uemail`, `isAdmin`, `password`, `create_at`)
VALUES (3, 'stephen', './static/images/stephen/avatar/stephen.png', 'stephen@qq.com', 1,
        '$2a$10$9UcvsMzLyZzaGYcNhHMAV.peUQzcFwI1uJEJcHfr36ZC0OwBimj7q', '2023-05-03 20:41:56');
INSERT INTO `user` (`uid`, `username`, `avatar`, `uemail`, `isAdmin`, `password`, `create_at`)
VALUES (4, 'stephen1', './static/images/stephen1/avatar/stephen1.png', 's1tephen@qq.com', 0,
        '$2a$10$9UcvsMzLyZzaGYcNhHMAV.peUQzcFwI1uJEJcHfr36ZC0OwBimj7q', '2023-05-04 21:29:14');
INSERT INTO `user` (`uid`, `username`, `avatar`, `uemail`, `isAdmin`, `password`, `create_at`)
VALUES (9, 'admin', './static/images/admin/avatar/admin.png', 's0tephen@outlook.com', 0,
        '$2a$10$sU7BqE8/qp/dRiP0MVMsguJ/Ma9sY5nju.9ZxeqSJpmyGeMn7Z7Mu', '2023-05-18 11:11:25');
COMMIT;

SET
FOREIGN_KEY_CHECKS = 1;
