CREATE TABLE IF NOT EXISTS `users` (
  `user_id` int PRIMARY KEY AUTO_INCREMENT,
  `image_url` varchar(100),
  `username` varchar(100) UNIQUE NOT NULL,
  `email` varchar(100) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `identification` varchar(20) UNIQUE NOT NULL,
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

CREATE TABLE IF NOT EXISTS `admins` (
  `admin_id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `admin_level` int,
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

CREATE TABLE IF NOT EXISTS `bookings` (
  `booking_id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `court_id` int,
  `sports_equipment_id` int,
  `date_booked` date NOT NULL,
  `start_time` time NOT NULL,
  `end_time` time NOT NULL,
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

CREATE TABLE IF NOT EXISTS `courts` (
  `courts_id` int PRIMARY KEY AUTO_INCREMENT,
  `image_url` varchar(100),
  `courts_name` varchar(255) NOT NULL,
  `courts_description` varchar(255),
  `price` int,
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

CREATE TABLE IF NOT EXISTS `sports_equipment` (
  `sports_equipment_id` int PRIMARY KEY AUTO_INCREMENT,
  `image_url` varchar(100),
  `sports_equipment_name` varchar(255) NOT NULL,
  `sports_equipment_description` varchar(255),
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

CREATE TABLE IF NOT EXISTS `payments` (
  `payment_id` int PRIMARY KEY AUTO_INCREMENT,
  `booking_id` int,
  `user_id` int,
  `amount` decimal(10, 2),
  `payment_method` varchar(255),
  `status` varchar(100),
  `transaction_date` timestamp DEFAULT (current_timestamp()),
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

CREATE TABLE IF NOT EXISTS `history` (
  `history_id` int PRIMARY KEY AUTO_INCREMENT,
  `booking_id` int,
  `user_id` int,
  `courts_id` int,
  `sports_equipment_id` int,
  `date_booked` date,
  `start_time` time,
  `end_time` time,
  `status` varchar(100),
  `price` decimal(10,2),
  `payment_status` varchar(100),
  `created_at` timestamp DEFAULT (current_timestamp()),
  `updated_at` timestamp DEFAULT (current_timestamp())
);

ALTER TABLE `admins` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `bookings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `bookings` ADD FOREIGN KEY (`court_id`) REFERENCES `courts` (`courts_id`);

ALTER TABLE `bookings` ADD FOREIGN KEY (`sports_equipment_id`) REFERENCES `sports_equipment` (`sports_equipment_id`);

ALTER TABLE `payments` ADD FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`booking_id`);

ALTER TABLE `payments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `history` ADD FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`booking_id`);

ALTER TABLE `history` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `history` ADD FOREIGN KEY (`courts_id`) REFERENCES `courts` (`courts_id`);

ALTER TABLE `history` ADD FOREIGN KEY (`sports_equipment_id`) REFERENCES `sports_equipment` (`sports_equipment_id`);
