
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
	`id` INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL DEFAULT '',
	`mailbox` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
	`sex` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0',
	`age` INT(11) NOT NULL DEFAULT '0',
	`longitude` FLOAT NOT NULL DEFAULT '0',
	`latitude` FLOAT NOT NULL DEFAULT '0',
	`description` VARCHAR(100) NOT NULL DEFAULT '',
	`password` VARCHAR(100) NOT NULL DEFAULT '',
	`head_url` VARCHAR(100) NOT NULL DEFAULT '',
	`status` INT(11) NOT NULL DEFAULT '0',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP INDEX `mailbox_password_of_user_u_k` ON `users`;
CREATE UNIQUE INDEX `mailbox_password_of_user_u_k` ON `users`(`mailbox`,`password`);


DROP INDEX `sex_of_user_i_d_x` ON `users`;
CREATE INDEX `sex_of_user_i_d_x` ON `users`(`sex`);


DROP INDEX `id_of_user_r_n_g` ON `users`;
CREATE INDEX `id_of_user_r_n_g` ON `users`(`id`);
DROP INDEX `age_of_user_r_n_g` ON `users`;
CREATE INDEX `age_of_user_r_n_g` ON `users`(`age`);

