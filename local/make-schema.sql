create table if not exists `users`
(
    `username` varchar(64) PRIMARY KEY,
    `notify_to` varchar(128),
    `slack_url` varchar(255),
    `email`   varchar(128)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `redirects`
(
 `user`             VARCHAR(64),
 `redirect_path`	VARCHAR(128) primary key,
 `org`			VARCHAR(10240),
 `host`			VARCHAR(128),
 `comment`		TEXT,
 `active`		int(1) DEFAULT 0,
 `begin_at`       Datetime DEFAULT NULL,
 `end_at`       Datetime DEFAULT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
  constraint `user_const` FOREIGN KEY (`user`) REFERENCES users(`username`) ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
