-- iot_api_db.access_keys definição
CREATE TABLE `access_keys` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `type_key` enum('rfid','qrcode','pin','other') NOT NULL,
  `status_key` enum('open','expired','used') NOT NULL,
  `value` varchar(255) NOT NULL,
  `assigned_to` varchar(36) DEFAULT NULL,
  `reservation_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_access_keys_value` (`value`),
  KEY `idx_access_keys_deleted_at` (`deleted_at`),
  KEY `idx_access_keys_assigned_to` (`assigned_to`),
  KEY `idx_access_keys_reservation_id` (`reservation_id`),
  CONSTRAINT `fk_access_keys_reservation` FOREIGN KEY (`reservation_id`) REFERENCES `reservations` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_access_keys_user` FOREIGN KEY (`assigned_to`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- iot_api_db.event_logs definição
CREATE TABLE `event_logs` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `access_key_id` varchar(36) DEFAULT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `reservation_id` varchar(36) DEFAULT NULL,
  `equipment_id` varchar(36) DEFAULT NULL,
  `action` enum('created','updated','deleted') NOT NULL,
  `message` text,
  PRIMARY KEY (`id`),
  KEY `idx_event_logs_deleted_at` (`deleted_at`),
  KEY `idx_event_logs_access_key_id` (`access_key_id`),
  KEY `idx_event_logs_user_id` (`user_id`),
  KEY `idx_event_logs_reservation_id` (`reservation_id`),
  KEY `idx_event_logs_equipment_id` (`equipment_id`),
  CONSTRAINT `fk_event_logs_access_key` FOREIGN KEY (`access_key_id`) REFERENCES `access_keys` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_event_logs_reservation` FOREIGN KEY (`reservation_id`) REFERENCES `reservations` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_event_logs_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- iot_api_db.equipment definição
CREATE TABLE `equipment` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `description` text,
  `status` enum('available','in_use','maintenance') DEFAULT 'available',
  PRIMARY KEY (`id`),
  KEY `idx_equipment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- iot_api_db.reservations definição
CREATE TABLE `reservations` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` varchar(36) NOT NULL,
  `responsible_id` varchar(36) DEFAULT NULL,
  `reservation_start` datetime(3) NOT NULL,
  `reservation_end` datetime(3) NOT NULL,
  `equipment_id` varchar(36) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_reservations_deleted_at` (`deleted_at`),
  KEY `idx_reservations_user_id` (`user_id`),
  KEY `idx_reservations_responsible_id` (`responsible_id`),
  KEY `idx_reservations_equipment_id` (`equipment_id`),
  CONSTRAINT `fk_reservations_responsible` FOREIGN KEY (`responsible_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_reservations_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- iot_api_db.users definição
CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `email` varchar(150) NOT NULL,
  `discord_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`),
  UNIQUE KEY `idx_users_discord_id` (`discord_id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;