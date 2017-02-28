
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
	`id` INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL DEFAULT '',
	`mailbox` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
	`sex` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0',
	`age` INT(11) NULL ,
	`longitude` FLOAT NOT NULL DEFAULT '0',
	`latitude` FLOAT NOT NULL DEFAULT '0',
	`description` VARCHAR(100) NULL ,
	`password` VARCHAR(100) NOT NULL DEFAULT '',
	`head_url` VARCHAR(100) NULL ,
	`status` INT(11) NOT NULL DEFAULT '0',
	`created_at` BIGINT(20) NOT NULL DEFAULT '0',
	`updated_at` BIGINT(20) NOT NULL DEFAULT '0',
	`deleted_at` BIGINT(20) NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE UNIQUE INDEX `id_of_user_uk` ON `users`(`id`);
CREATE UNIQUE INDEX `mailbox_password_of_user_uk` ON `users`(`mailbox`,`password`);
CREATE INDEX `sex_of_user_idx` ON `users`(`sex`);
CREATE INDEX `id_of_user_rng` ON `users`(`id`);
CREATE INDEX `age_of_user_rng` ON `users`(`age`);

