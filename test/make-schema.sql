create table IF not exists `redirect`
(
 `user`             VARCHAR(20),
 `redirect_path`	VARCHAR(64),
 `org`			VARCHAR(10240),
 `host`			VARCHAR(128),
 `comment`		TEXT,
 `active`		int(1) DEFAULT 0,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`redirect_path`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;