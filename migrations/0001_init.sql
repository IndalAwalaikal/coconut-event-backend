-- Migration Update: Add event_type and price fields
-- Dialect: MySQL (compatible with XAMPP)
CREATE DATABASE IF NOT EXISTS coconut_event_hub
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;

USE coconut_event_hub;


SET FOREIGN_KEY_CHECKS = 0;

-- Drop if exists (safe to re-run during development)
DROP TABLE IF EXISTS `registrations`;
DROP TABLE IF EXISTS `documentations`;
DROP TABLE IF EXISTS `events`;
DROP TABLE IF EXISTS `admins`;

-- Admins table (for backend authentication)
CREATE TABLE `admins` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(100) NOT NULL UNIQUE,
	`password_hash` VARCHAR(255) NOT NULL,
	`name` VARCHAR(255) DEFAULT NULL,
	`role` VARCHAR(50) DEFAULT 'admin',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Events table (UPDATED with event_type and price)
CREATE TABLE `events` (
	`id` VARCHAR(50) NOT NULL,
	`category` ENUM('open-class','webinar','seminar','bootcamp') NOT NULL,
	`category_label` VARCHAR(100) NOT NULL,
	`title` VARCHAR(255) NOT NULL,
	`description` TEXT,
	`rules` JSON DEFAULT NULL,
	`benefits` JSON DEFAULT NULL,
	`date` DATE DEFAULT NULL,
	`time` VARCHAR(100) DEFAULT NULL,
	`speaker1` VARCHAR(255) DEFAULT NULL,
	`speaker2` VARCHAR(255) DEFAULT NULL,
	`speaker3` VARCHAR(255) DEFAULT NULL,
	`moderator` VARCHAR(255) DEFAULT NULL,
	`location` VARCHAR(255) DEFAULT NULL,
	`quota` INT UNSIGNED NOT NULL DEFAULT 0,
	`registered` INT UNSIGNED NOT NULL DEFAULT 0,
	`poster` VARCHAR(255) DEFAULT NULL,
	`event_type` ENUM('free','paid') NOT NULL DEFAULT 'free',
	`price` INT UNSIGNED NOT NULL DEFAULT 0,
	`available` TINYINT(1) NOT NULL DEFAULT 1,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `idx_events_category` (`category`),
	INDEX `idx_events_date` (`date`),
	INDEX `idx_events_event_type` (`event_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Registrations table
CREATE TABLE `registrations` (
	`id` VARCHAR(50) NOT NULL,
	`event_id` VARCHAR(50) NOT NULL,
	`name` VARCHAR(255) NOT NULL,
	`whatsapp` VARCHAR(25) NOT NULL,
	`institution` VARCHAR(255) DEFAULT NULL,
	`proof_image` VARCHAR(255) DEFAULT NULL,
	`file_name` VARCHAR(255) DEFAULT NULL,
	`registered_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `idx_reg_event` (`event_id`),
	CONSTRAINT `fk_registrations_event` FOREIGN KEY (`event_id`) REFERENCES `events`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Documentations table (past events / gallery)
CREATE TABLE `documentations` (
	`id` VARCHAR(50) NOT NULL,
	`event_id` VARCHAR(50) DEFAULT NULL,
	`category` ENUM('open-class','webinar','seminar','bootcamp') NOT NULL,
	`category_label` VARCHAR(100) DEFAULT NULL,
	`event_title` VARCHAR(255) DEFAULT NULL,
	`year` YEAR DEFAULT NULL,
	`images` JSON DEFAULT NULL,
	`description` TEXT DEFAULT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `idx_docs_year` (`year`),
	CONSTRAINT `fk_docs_event` FOREIGN KEY (`event_id`) REFERENCES `events`(`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Seed data (based on frontend mockData) -- events with event_type and price
INSERT INTO `events` (`id`,`category`,`category_label`,`title`,`description`,`rules`,`benefits`,`date`,`time`,`speaker1`,`speaker2`,`speaker3`,`moderator`,`location`,`quota`,`registered`,`poster`,`event_type`,`price`,`available`)
VALUES
('oc-1','open-class','COCONUT Open Class','Pengenalan Web Development dengan React',
 'Pelajari dasar-dasar pengembangan web modern menggunakan React.js. Materi mencakup komponen, state management, dan pembuatan UI interaktif.',
 JSON_ARRAY('Peserta wajib hadir 15 menit sebelum acara dimulai','Membawa laptop pribadi dengan Node.js terinstall','Mengikuti seluruh rangkaian acara hingga selesai','Menjaga ketertiban selama acara berlangsung'),
 JSON_ARRAY('E-Sertifikat kehadiran','Materi pembelajaran lengkap','Akses ke grup diskusi eksklusif','Hands-on project portfolio'),
 '2025-03-15','09:00 - 12:00 WIB','Ahmad Rizky, S.Kom','Dewi Sartika, M.Cs',NULL,'Budi Santoso','Lab Komputer Gedung A Lt.3',50,32,'/placeholder.svg','free',0,1),

('oc-2','open-class','COCONUT Open Class','UI/UX Design Fundamentals',
 'Workshop desain UI/UX untuk pemula. Belajar prinsip desain, wireframing, dan prototyping menggunakan Figma.',
 JSON_ARRAY('Peserta wajib membuat akun Figma sebelum acara','Hadir tepat waktu','Mengikuti seluruh sesi workshop'),
 JSON_ARRAY('E-Sertifikat','Template Figma gratis','Portfolio project','Networking dengan desainer profesional'),
 '2025-03-22','13:00 - 16:00 WIB','Sarah Amanda, S.Ds',NULL,NULL,'Rina Kartika','Aula Fakultas Ilmu Komputer',40,40,'/placeholder.svg','free',0,1),

('wb-1','webinar','Webinar','Karir di Dunia IT: Peluang dan Tantangan',
 'Webinar online membahas berbagai jalur karir di industri teknologi informasi, tips memulai karir, dan skill yang dibutuhkan.',
 JSON_ARRAY('Peserta wajib join Zoom 10 menit sebelum acara','Menjaga mikrofon dalam keadaan mute','Pertanyaan diajukan melalui fitur Q&A'),
 JSON_ARRAY('E-Sertifikat','Rekaman webinar','Slide presentasi','Kesempatan tanya jawab langsung'),
 '2025-04-05','19:00 - 21:00 WIB','Dr. Hendra Wijaya','Fitri Rahmawati, M.T','Dimas Prasetyo',NULL,'Online via Zoom',200,87,'/placeholder.svg','free',0,1),

('sm-1','seminar','Seminar','Artificial Intelligence & Masa Depan Teknologi',
 'Seminar nasional tentang perkembangan AI, machine learning, dan dampaknya terhadap berbagai industri di Indonesia.',
 JSON_ARRAY('Peserta wajib melakukan registrasi ulang di tempat','Berpakaian rapi dan sopan','Dilarang membawa makanan ke dalam ruangan','Mengikuti seluruh rangkaian acara'),
 JSON_ARRAY('Sertifikat seminar nasional','Goodie bag eksklusif','Networking session','Lunch & coffee break','Materi presentasi digital'),
 '2025-04-20','08:00 - 15:00 WIB','Prof. Agus Setiawan, Ph.D','Maria Chen, MSc','Robert Tanaka','Lisa Permata, S.Kom','Auditorium Universitas, Gedung Utama',300,156,'/placeholder.svg','paid',150000,1),

('bc-1','bootcamp','Bootcamp','Full-Stack Developer Bootcamp',
 'Bootcamp intensif selama 3 hari untuk mempelajari pengembangan aplikasi full-stack dengan teknologi modern.',
 JSON_ARRAY('Peserta wajib hadir selama 3 hari penuh','Membawa laptop dengan spesifikasi minimum yang ditentukan','Menyelesaikan pre-assignment sebelum bootcamp','Bekerja dalam tim selama bootcamp'),
 JSON_ARRAY('Sertifikat kompetensi','Project portfolio lengkap','Mentoring 1-on-1','Akses materi seumur hidup','Rekomendasi kerja'),
 '2025-05-10','08:00 - 17:00 WIB (3 hari)','Yusuf Hakim, S.T','Nina Saraswati, M.Kom',NULL,'Andi Firmansyah','Co-working Space TechHub, Lantai 5',30,28,'/placeholder.svg','paid',500000,1),

('wb-2','webinar','Webinar','Cybersecurity Awareness',
 'Webinar tentang keamanan siber untuk mahasiswa dan umum.',
 JSON_ARRAY('Hadir tepat waktu','Menjaga mikrofon mute'),
 JSON_ARRAY('E-Sertifikat','Materi presentasi'),
 '2025-06-01','14:00 - 16:00 WIB','TBA',NULL,NULL,'TBA','Online via Zoom',150,0,'/placeholder.svg','free',0,0)
;

-- Seed data: documentations
INSERT INTO `documentations` (`id`,`event_id`,`category`,`category_label`,`event_title`,`year`,`images`,`description`)
VALUES
('doc-1',NULL,'open-class','COCONUT Open Class','Belajar Python dari Nol',2024,
 JSON_ARRAY('/placeholder.svg','/placeholder.svg','/placeholder.svg'),
 'Dokumentasi kegiatan Open Class Python yang diikuti oleh 45 peserta.'),

('doc-2',NULL,'seminar','Seminar','Seminar Teknologi Cloud Computing',2024,
 JSON_ARRAY('/placeholder.svg','/placeholder.svg'),
 'Seminar nasional dengan 250+ peserta membahas cloud computing.'),

('doc-3',NULL,'webinar','Webinar','Webinar Data Science for Beginners',2023,
 JSON_ARRAY('/placeholder.svg','/placeholder.svg','/placeholder.svg','/placeholder.svg'),
 'Webinar online yang diikuti oleh 180 peserta dari berbagai universitas.'),

('doc-4',NULL,'bootcamp','Bootcamp','Mobile App Development Bootcamp',2023,
 JSON_ARRAY('/placeholder.svg','/placeholder.svg'),
 'Bootcamp 3 hari pengembangan aplikasi mobile dengan Flutter.');


-- Seed data: registrants
INSERT INTO `registrations` (`id`,`event_id`,`name`,`whatsapp`,`institution`,`proof_image`,`registered_at`)
VALUES
('reg-1','oc-1','Muhammad Faisal','081234567890','Universitas Indonesia','/placeholder.svg','2025-02-20 10:30:00'),
('reg-2','oc-1','Siti Nurhaliza','082345678901','Institut Teknologi Bandung','/placeholder.svg','2025-02-21 14:15:00'),
('reg-3','wb-1','Andi Pratama','083456789012','Universitas Gadjah Mada','/placeholder.svg','2025-03-01 09:00:00'),
('reg-4','sm-1','Rina Susanti','084567890123','SMA Negeri 1 Jakarta','/placeholder.svg','2025-03-05 16:45:00')
;

-- Useful indexes
CREATE INDEX `idx_events_title` ON `events` (`title`(100));
CREATE INDEX `idx_regs_name` ON `registrations` (`name`(100));

-- Re-enable foreign key checks after schema creation and seeding
SET FOREIGN_KEY_CHECKS = 1;

-- End of migration

-- Posters table (previous activities / gallery)
CREATE TABLE IF NOT EXISTS `posters` (
	`id` VARCHAR(50) NOT NULL,
	`title` VARCHAR(255) NOT NULL,
	`type` VARCHAR(100) DEFAULT NULL,
	`image` VARCHAR(255) DEFAULT NULL,
	`date` VARCHAR(100) DEFAULT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;