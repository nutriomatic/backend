CREATE TABLE `activity_levels` (
  `al_id` varchar(36),
  `al_type` bigint DEFAULT NULL,
  `al_desc` longtext DEFAULT NULL,
  `al_value` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`al_id`)
);

CREATE TABLE `health_goals` (
  `hg_id` varchar(36),
  `hg_type` bigint,
  `hg_desc` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`hg_id`)
);

CREATE TABLE `nutrition_info` (
  `ni_id` varchar(36),
  `ni_type` varchar(255) DEFAULT NULL,
  `ni_text` varchar(500) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`ni_id`)
);

CREATE TABLE `product_types` (
  `pt_id` varchar(36) NOT NULL,
  `pt_name` longtext DEFAULT NULL,
  `pt_type` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`pt_id`)
);

CREATE TABLE `products` (
  `product_id` varchar(36) NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `product_price` decimal(10,2) DEFAULT NULL,
  `product_desc` text,
  `product_isshow` tinyint(1) DEFAULT NULL,
  `product_lemaktotal` decimal(10,2) DEFAULT NULL,
  `product_protein` decimal(10,2) DEFAULT NULL,
  `product_karbohidrat` decimal(10,2) DEFAULT NULL,
  `product_garam` decimal(10,2) DEFAULT NULL,
  `product_servingsize` decimal(10,2) DEFAULT NULL,
  `product_picture` varchar(255) DEFAULT NULL,
  `product_expshow` timestamp NULL DEFAULT NULL,
  `createdat` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedat` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `store_id` varchar(36) DEFAULT NULL,
  `pt_id` varchar(36) DEFAULT NULL,
  `product_grade` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`product_id`),
  KEY `store_id` (`store_id`),
  KEY `pt_id` (`pt_id`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`store_id`) REFERENCES `stores` (`store_id`),
  CONSTRAINT `products_ibfk_2` FOREIGN KEY (`pt_id`) REFERENCES `product_types` (`pt_id`)
);

CREATE TABLE `stores` (
  `store_id` varchar(36),
  `store_name` longtext,
  `store_username` longtext,
  `store_address` longtext,
  `store_contact` longtext,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `user_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`store_id`),
  KEY `fk_stores_user` (`user_id`),
  CONSTRAINT `fk_stores_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
); 

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `username` varchar(191) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role` longtext DEFAULT NULL,
  `gender` bigint DEFAULT NULL,
  `telp` longtext DEFAULT NULL,
  `profpic` longtext DEFAULT NULL,
  `birthdate` longtext DEFAULT NULL,
  `place` longtext DEFAULT NULL,
  `height` double DEFAULT NULL,
  `weight` double DEFAULT NULL,
  `weight_goal` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `hg_id` varchar(36) DEFAULT NULL,
  `al_id` varchar(36) DEFAULT NULL,
  `calories` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `FK_users_health_goals` (`hg_id`),
  KEY `FK_users_activity_levels` (`al_id`),
  CONSTRAINT `FK_users_activity_levels` FOREIGN KEY (`al_id`) REFERENCES `activity_levels` (`al_id`),
  CONSTRAINT `FK_users_health_goals` FOREIGN KEY (`hg_id`) REFERENCES `health_goals` (`hg_id`)
);

CREATE TABLE `tokens` (
  `user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `token` varchar(255) NOT NULL,
  `expires_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  UNIQUE KEY `token` (`token`),
  UNIQUE KEY `idx_tokens_deleted_at` (`deleted_at`) USING BTREE,
  KEY `unique_user_id` (`user_id`) USING BTREE
)