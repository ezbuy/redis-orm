
CREATE TABLE `user_blogs` (
	`user_id` INT(11) NOT NULL DEFAULT '0',
	`blog_id` INT(11) NOT NULL DEFAULT '0',
	PRIMARY KEY(`user_id`,`blog_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'user_blogs';

