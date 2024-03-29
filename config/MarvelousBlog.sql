﻿create database MarvelousBlog;

use MarvelousBlog;

CREATE TABLE author (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '作者（管理员）id' ,
    nickname VARCHAR(20) NOT NULL COMMENT '作者（管理员）登录名',
    password VARCHAR(50) NOT NULL COMMENT '作者（管理员）登录密码',
    avatar VARCHAR(100) DEFAULT NULL COMMENT '作者（管理员）头像',
    register_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间',
    last_login_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账户上次登录时间',
	role TINYINT NOT NULL DEFAULT 0  COMMENT '作者身份，1为站主，0为普通作者'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE label (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '标签id',
    label_name VARCHAR(20) NOT NULL COMMENT '标签名字',
    description VARCHAR(150) DEFAULT NULL COMMENT '标签描述'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE article(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '文章id',
    title VARCHAR(50) NOT NULL COMMENT '文章标题' ,
    abstract VARCHAR(75) DEFAULT NULL COMMENT '文章摘要' ,
    content LONGTEXT DEFAULT NULL COMMENT '文章内容',
    label_id BIGINT DEFAULT NULL COMMENT '文章标签' ,
    post_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间' ,
    author_id BIGINT NOT NULL COMMENT '作者ID' ,
    visit_count BIGINT DEFAULT 0 COMMENT '浏览量',
    last_modify_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次修改时间',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '评论状态，1为可用；0为禁用',
    INDEX label_id(label_id) USING BTREE ,
    INDEX post_time(post_time) USING BTREE ,
    CONSTRAINT fk_label_id_1 FOREIGN KEY (label_id) REFERENCES label(id) ON DELETE CASCADE ON UPDATE CASCADE ,
    CONSTRAINT fk_author_id_1 FOREIGN KEY (author_id) REFERENCES author(id) ON DELETE CASCADE ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE comment(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '评论id',
    mobile VARCHAR(20) DEFAULT NULL COMMENT '评论者手机号',
    email VARCHAR(50) DEFAULT NULL COMMENT '评论者邮箱' ,
    context TEXT DEFAULT NULL COMMENT '评论内容',
    comment_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '评论时间' ,
    article_id BIGINT NOT NULL COMMENT '评论文章id',
    INDEX comment_time(comment_time) USING BTREE ,
    CONSTRAINT fk_article_id_1 FOREIGN KEY (article_id) REFERENCES article(id) ON DELETE CASCADE ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

show tables ;

#DROP DATABASE MarvelousBlog;
