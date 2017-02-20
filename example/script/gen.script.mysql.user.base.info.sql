
DROP VIEW IF EXISTS `user_base_info`;
CREATE VIEW `user_base_info` AS SELECT `id`,`name`,`mailbox`,`sex` FROM users;

