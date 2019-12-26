-- create db if not exists
CREATE SCHEMA IF NOT EXISTS appdb CHARSET utf8mb4;

-- select db
USE appdb;

-- create table if not exist
CREATE TABLE IF NOT EXISTS `user` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` TIMESTAMP NULL DEFAULT NULL,
    `updated_at` TIMESTAMP NULL DEFAULT NULL,
    `email` VARCHAR(50) NOT NULL UNIQUE,
    `password` VARCHAR(50) NOT NULL,
    PRIMARY KEY(`id`)
);
