
CREATE TABLE `users` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL DEFAULT '',
	`mailbox` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
	`sex` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0',
	`age` INT(11) NOT NULL DEFAULT '0',
	`longitude` FLOAT NOT NULL DEFAULT '0',
	`latitude` FLOAT NOT NULL DEFAULT '0',
	`description` VARCHAR(100) NULL ,
	`password` VARCHAR(100) NOT NULL DEFAULT '',
	`head_url` VARCHAR(100) NULL ,
	`status` INT(11) NOT NULL DEFAULT '0',
	`created_at` BIGINT(20) NOT NULL DEFAULT '0',
	`updated_at` BIGINT(20) NOT NULL DEFAULT '0',
	`deleted_at` BIGINT(20) NULL ,
	PRIMARY KEY(`id`),
	UNIQUE KEY `uniq_mailbox_password_of_user_uk` (`mailbox`,`password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';
CREATE INDEX `sex_of_user_idx` ON `users`(`sex`);
CREATE INDEX `age_of_user_rng` ON `users`(`age`);

