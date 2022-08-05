CREATE TABLE sismor.users (
	id INT(20) NOT NULL,
	name VARCHAR(100) NULL,
	password VARCHAR(100) NOT NULL,
	email varchar(200) NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
CREATE UNIQUE INDEX users_id_IDX USING BTREE ON sismor.users (id);

ALTER TABLE sismor.users MODIFY COLUMN id int auto_increment NOT NULL;
