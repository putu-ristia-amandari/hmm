-- Create Database
DROP DATABASE IF EXISTS gym;
CREATE DATABASE IF NOT EXISTS gym;
USE gym;

CREATE TABLE IF NOT EXISTS `memberships` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(25) NOT NULL,
    `image` varchar(200) DEFAULT NULL,
    `limited_class` bigint(3) DEFAULT NULL,
    `limited_time` bigint(3) DEFAULT NULL,
    `description` longtext DEFAULT NULL,
    `details` longtext DEFAULT NULL,
    `price` double(15,2) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `membership_id` bigint(20) DEFAULT NULL,
    `name`  varchar(200) NOT NULL,
    `email`  varchar(100) NOT NULL,
    `password`  varchar(100) NOT NULL,
    `handphone`  varchar(15) DEFAULT NULL,
    `address` longtext DEFAULT NULL,
    `city`  varchar(100) DEFAULT NULL,
    `province`  varchar(100) DEFAULT NULL,
    `nationality`  varchar(100) DEFAULT NULL,
    `gender`  varchar(6) DEFAULT NULL,
    `birth_of_date` date DEFAULT NULL,
    `height` bigint(3) DEFAULT NULL,
    `weight` bigint(3) DEFAULT NULL,
    `photo` varchar(200) DEFAULT NULL,
    `status` tinyint(1) DEFAULT 0,
    `status_membership` tinyint(1) DEFAULT 0,
    `remember_token` varchar(64) DEFAULT NULL,
    `is_reset` tinyint(1) DEFAULT 0,
    `verified_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`membership_id`) REFERENCES `memberships` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
)

CREATE TABLE IF NOT EXISTS `instructors` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(100) NOT NULL,
    `email`  varchar(100) NOT NULL,
    `handphone`  varchar(15) DEFAULT NULL,
    `address` longtext DEFAULT NULL,
    `city`  varchar(100) DEFAULT NULL,
    `province`  varchar(100) DEFAULT NULL,
    `nationality`  varchar(100) DEFAULT NULL,
    `gender`  varchar(6) DEFAULT NULL,
    `birth_of_date` date DEFAULT NULL,
    `height` bigint(3) DEFAULT NULL,
    `weight` bigint(3) DEFAULT NULL,
    `photo` varchar(200) DEFAULT NULL,
    `status` tinyint(1) DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS `categories` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(100) NOT NULL,
    `photo` varchar(200) DEFAULT NULL,
    `description` longtext DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS `contents` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `category_id` bigint(20) DEFAULT NULL,
    `title`  varchar(200) NOT NULL,
    `image` varchar(200) DEFAULT NULL,
    `description` longtext DEFAULT NULL,
    `created_by` longtext DEFAULT NULL,
    `updated_by` longtext DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
)