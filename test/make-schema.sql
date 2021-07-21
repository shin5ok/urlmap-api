create table IF not exists `redirects`
(
 `user`             VARCHAR(20),
 `redirect_path`	VARCHAR(64),
 `org`			VARCHAR(10240),
 `host`			VARCHAR(128),
 `comment`		TEXT,
 `active`		int(1) DEFAULT 0,
 `begin_at`       Datetime DEFAULT NULL,
 `end_at`       Datetime DEFAULT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`redirect_path`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;