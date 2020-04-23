
DROP VIEW IF EXISTS user_view_base_info;
CREATE VIEW user_view_base_info AS SELECT `id`,`name`,`mailbox`,`sex` FROM users;

