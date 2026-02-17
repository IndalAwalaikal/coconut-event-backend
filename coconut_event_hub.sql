-- MariaDB dump 10.19  Distrib 10.4.32-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: coconut_event_hub
-- ------------------------------------------------------
-- Server version	10.4.32-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admins`
--

DROP TABLE IF EXISTS `admins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `role` varchar(50) DEFAULT 'admin',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admins`
--

LOCK TABLES `admins` WRITE;
/*!40000 ALTER TABLE `admins` DISABLE KEYS */;
INSERT INTO `admins` VALUES (1,'admin','$2a$12$ODn6WQi1NiHzNzxYc8S/iOud.TrGPcrcqHxfl2WEjCXVRhLmqUozW','admin','admin','2026-02-07 12:43:27','2026-02-07 12:43:27');
/*!40000 ALTER TABLE `admins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `documentations`
--

DROP TABLE IF EXISTS `documentations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `documentations` (
  `id` varchar(50) NOT NULL,
  `event_id` varchar(50) DEFAULT NULL,
  `category` enum('open-class','webinar','seminar','bootcamp') NOT NULL,
  `category_label` varchar(100) DEFAULT NULL,
  `event_title` varchar(255) DEFAULT NULL,
  `year` int(10) unsigned DEFAULT NULL,
  `images` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`images`)),
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `idx_docs_year` (`year`),
  KEY `fk_docs_event` (`event_id`),
  CONSTRAINT `fk_docs_event` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `documentations`
--

LOCK TABLES `documentations` WRITE;
/*!40000 ALTER TABLE `documentations` DISABLE KEYS */;
INSERT INTO `documentations` VALUES ('075700c0-dca0-430c-9296-8a6d7300b721',NULL,'open-class','COCONUT Open Class','Batch 6',2025,'[\"/storage/documentations/3b52fd0e-3f67-43a7-948d-f40417b9e148.JPG\",\"/storage/documentations/1cdcdd3f-e0fd-4f8f-9348-892cf5f307ce.JPG\",\"/storage/documentations/ca5b36f1-0754-4366-93d6-cfd04b385af0.JPG\",\"/storage/documentations/c4797de9-0bbb-4cc8-8d47-2abd3c7893fe.JPG\",\"/storage/documentations/d8df74e2-ce4f-46a8-bfd2-557019a5a20d.JPG\"]','Programming 101: Mengenal dan Belajar Lebih dalam Konsep Pemrograman','2026-02-08 10:39:16','2026-02-08 10:39:16'),('0e688fed-986a-486c-a40a-4a1688f254b2',NULL,'open-class','COCONUT Open Class','Batch 1',2023,'[\"/storage/documentations/9b023913-7e8b-406f-9432-9ddbb8a3d3db.png\",\"/storage/documentations/8c9e1350-774a-4ba1-bf71-81855ded54fb.png\",\"/storage/documentations/36b496ce-dc12-414a-84bc-bcd1b1c131cd.png\",\"/storage/documentations/f0a02b32-513b-4c78-850f-1ce6311a8c50.png\"]','Introduce Your Self to Information Technology','2026-02-08 10:31:32','2026-02-08 10:31:32'),('186ab5b6-a0aa-4bb5-9f42-ea526d66776e',NULL,'open-class','COCONUT Open Class','Batch 4',2024,'[\"/storage/documentations/a6259b79-9a98-4347-b25f-b86503cdf42e.JPG\",\"/storage/documentations/487f3a3f-67c7-482c-af39-7c370cbf4e89.JPG\",\"/storage/documentations/f3532026-d779-475a-bda7-1073336cdf8b.JPG\",\"/storage/documentations/d1a16cde-d092-4350-ac4c-65ce7bd578ba.JPG\"]','Introduction to Sveltekit: The Frontend Framework of the Future','2026-02-08 10:37:34','2026-02-08 10:37:34'),('398d1b57-3b74-4ef2-940f-1dd6c91e244c',NULL,'open-class','COCONUT Open Class','Batch 8',2025,'[\"/storage/documentations/465961b1-874d-401a-a6df-216fbc41aac1.JPG\",\"/storage/documentations/d78a2023-6125-4ad1-800b-d82ea8f0c73b.JPG\",\"/storage/documentations/fb404347-62fd-44c7-b0a6-4aa9652f8a06.JPG\",\"/storage/documentations/6efdc10b-6d41-4be3-97ee-86b988b140fe.JPG\"]','Go REST, Go Fast: Membangun REST API dengan Golang','2026-02-08 10:41:20','2026-02-08 10:41:20'),('4ad9de0e-981b-4c6b-9668-33fa8a0d31fa',NULL,'open-class','COCONUT Open Class','Batch 5',2024,'[\"/storage/documentations/88d46d95-1fe8-414d-82ef-4a5c1079d497.JPG\",\"/storage/documentations/2f06a336-6139-4cde-9762-9d7805daa3ec.JPG\",\"/storage/documentations/ca7050b9-228a-45d5-9c6f-5b77b6e92fee.JPG\",\"/storage/documentations/5ae52a8f-ff4a-4f93-90e7-6aab896321d2.JPG\",\"/storage/documentations/d7e479d3-8bbf-4eb5-a70f-cfc568d0bf81.JPG\"]','Web Development With Laravel : Pengantar HTML, CSS, Dan JavaScript & Dasar PHP dan Framework Laravel','2026-02-08 10:38:30','2026-02-08 10:50:10'),('55623f89-790e-467f-afa0-305082193349',NULL,'seminar','Seminar','BLOCKCHAIN',2024,'[\"/storage/documentations/67bddbe7-656a-4583-b067-ab4717bb91e8.JPG\",\"/storage/documentations/5d40ed24-e85e-4e3f-afbb-3c88378ac18d.JPG\",\"/storage/documentations/5f0974bf-bcbe-43b1-9e67-728a40e02fce.JPG\",\"/storage/documentations/4ce54b4d-03d4-452e-8cd8-013b38d4ba2a.JPG\",\"/storage/documentations/97ef64b3-bd61-4541-b845-c6f97fd56d78.JPG\",\"/storage/documentations/b6734346-2d02-4105-90bd-1ef8b070f2c8.JPG\",\"/storage/documentations/b7b27da0-ed67-46fe-9ade-6ab9d6860d07.JPG\",\"/storage/documentations/46cebf28-d1e2-4254-8134-1badc73da6ed.JPG\",\"/storage/documentations/a5d8002f-5967-488d-b1c2-5c1aae3baf30.JPG\"]','BLOCKCHAIN DEMYSTIFIED: The Technology Behind the Digital Revolution','2026-02-08 10:29:18','2026-02-08 10:29:18'),('72f05453-3195-47a7-a68c-2520b178a918',NULL,'open-class','COCONUT Open Class','Batch 7',2025,'[\"/storage/documentations/3d9b91cd-c59c-42c4-a7d5-99a9e77dd415.JPG\",\"/storage/documentations/01f4d77e-a130-4673-8007-828fd4c51c9e.JPG\",\"/storage/documentations/261727b1-a3f5-46c8-b753-1205db6925dd.JPG\",\"/storage/documentations/5ae1a083-5341-47fd-b4a3-8b89c5caa081.JPG\",\"/storage/documentations/481472ce-23c5-481e-9093-9553e68a6c4e.JPG\"]','Level UP Your Skills: Bangun Portofolio Interaktif dengan Next.js dari Nol','2026-02-08 10:40:26','2026-02-08 10:40:26'),('7c0e5025-1a09-4626-b832-90e58dda3668',NULL,'open-class','COCONUT Open Class','Batch 9',2026,'[\"/storage/documentations/a05ff9db-cb00-4903-a7e2-ecb24450133c.JPG\",\"/storage/documentations/427f901b-3846-4305-a66b-51efb996e07d.JPG\",\"/storage/documentations/fa6b0b2b-4fe1-48ea-9088-a68ae01b5141.JPG\",\"/storage/documentations/9c5ae6e7-c68e-4a48-8a40-822286cd954b.JPG\",\"/storage/documentations/c99ed1e9-aa00-4015-b751-d1eeee38180c.JPG\"]','Beyond SQL Injection: Modern Web Exploitation in API & AI Era','2026-02-08 11:16:08','2026-02-08 11:16:08'),('a3be205f-8360-4490-9324-09420a791587',NULL,'seminar','Seminar','IOT',2025,'[\"/storage/documentations/3486b46a-7bf0-42af-913a-2543cf830505.JPG\",\"/storage/documentations/d138e757-dd1b-4d81-b178-b7249aca343c.JPG\",\"/storage/documentations/ecf36fb2-3cba-45fb-91f1-7525399be0d5.JPG\",\"/storage/documentations/e36f8bc4-e839-421a-a1d9-a62ff6925ee7.JPG\",\"/storage/documentations/3fb59221-41fa-4b5d-a2f5-b37f90f3bf0d.JPG\",\"/storage/documentations/92860daf-dc23-450a-8e37-54613ce5b464.JPG\",\"/storage/documentations/4a9b90a8-c224-4076-b7b6-3b5663a52f81.JPG\",\"/storage/documentations/a5301bb1-1e87-4e07-9633-6a56cfe44fe1.JPG\"]','IOT-Based Network Management : Solusi Pintar dalam Pengelolaan Infrastruktur Jaringan','2026-02-08 10:30:09','2026-02-08 10:50:43'),('ada5c872-e7a8-4b75-98c5-0457618264e7',NULL,'seminar','Seminar','AI',2023,'[\"/storage/documentations/fe07bb1d-91db-4240-988b-2ea148ab3fc2.JPG\",\"/storage/documentations/778a8796-551a-4bde-b8df-799649c4aa8c.JPG\",\"/storage/documentations/f7eaf2b7-a61b-464e-a5b5-686176a2af8a.JPG\",\"/storage/documentations/72544bd8-4d94-4436-a27f-f2e3f55b959a.JPG\",\"/storage/documentations/0067b80c-7de0-498d-acf1-9f9b7841b291.JPG\",\"/storage/documentations/490379c7-8849-43e3-b921-20a976c41395.JPG\",\"/storage/documentations/cfbe7efd-8a31-4732-9708-abbaa7625f5e.JPG\",\"/storage/documentations/6c87cdc4-33b5-4782-8e91-fbc08d0febb2.JPG\",\"/storage/documentations/e092a4c5-d541-499e-9ba8-cd1bb528916e.JPG\"]','Introduction to Machine Learning : Basic Concepts, Types, and Learning Process','2026-02-08 10:27:09','2026-02-08 10:51:29'),('cc3b304f-dc9e-43ac-b34e-8e69d6ca0c44',NULL,'open-class','COCONUT Open Class','Batch 2',2024,'[\"/storage/documentations/ea7a778c-b576-42e4-96a4-13b07c51bdf6.png\",\"/storage/documentations/82ffb22f-102f-4957-a9fa-bb8873fc7386.png\",\"/storage/documentations/1691203a-a84c-4755-92f5-a0d03154c1ad.png\",\"/storage/documentations/efc8ed44-35a0-4895-94cf-74c0b92b5d5f.png\"]','CRUD: Belajar Mengolah Data Menggunakan Golang dan MYSQL','2026-02-08 10:34:31','2026-02-08 10:34:31');
/*!40000 ALTER TABLE `documentations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events` (
  `id` varchar(50) NOT NULL,
  `category` enum('open-class','webinar','seminar','bootcamp') NOT NULL,
  `category_label` varchar(100) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `rules` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin CHECK (json_valid(`rules`)),
  `benefits` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin CHECK (json_valid(`benefits`)),
  `date` date DEFAULT NULL,
  `time` varchar(100) DEFAULT NULL,
  `speaker1` varchar(255) DEFAULT NULL,
  `speaker2` varchar(255) DEFAULT NULL,
  `speaker3` varchar(255) DEFAULT NULL,
  `moderator` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `quota` int(10) unsigned NOT NULL DEFAULT 0,
  `registered` int(10) unsigned NOT NULL DEFAULT 0,
  `poster` varchar(255) DEFAULT NULL,
  `event_type` enum('free','paid') NOT NULL DEFAULT 'free',
  `price` int(10) unsigned NOT NULL DEFAULT 0,
  `available` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `idx_events_category` (`category`),
  KEY `idx_events_date` (`date`),
  KEY `idx_events_title` (`title`(100)),
  KEY `idx_events_event_type` (`event_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES ('0d959556-48cf-4353-bd0e-9cc2d4c08497','bootcamp','Bootcamp','Man Behind The Hat','Mini Bootcamp MAN BEHIND THE HAT merupakan program pelatihan intensif yang dirancang untuk membekali peserta dengan pemahaman mendalam mengenai Offensive Security (Offsec) dengan fokus utama pada Offensive OPSEC (Operational Security). Kegiatan ini menekankan bahwa keberhasilan sebuah operasi penetrasi tidak hanya terletak pada kemampuan mengeksploitasi sistem, tetapi juga pada bagaimana seorang attacker mampu menjaga anonimitas, meminimalkan jejak digital, dan menghindari deteksi forensik.\r\n\r\nBootcamp ini dilaksanakan dalam tiga pertemuan terstruktur yang membentuk satu siklus serangan lengkap (attack lifecycle), mulai dari tahap reconnaissance hingga post-exploitation dan penyusunan narasi serangan. Setiap sesi dirancang berbasis praktik (hands-on lab) menggunakan lingkungan terisolasi berbasis virtual machine, sehingga peserta dapat mensimulasikan skenario serangan secara realistis namun tetap aman dan terkendali.\r\n\r\nPada pertemuan pertama, peserta akan mempelajari teknik reconnaissance dan attack surface mapping, termasuk pendekatan active dan passive reconnaissance, service enumeration, serta penyusunan recon report untuk memetakan potensi celah yang dapat dieksploitasi. Peserta tidak hanya diajarkan teknik, tetapi juga dibangun mindset analitis untuk memahami bagaimana dan mengapa suatu sistem rentan diserang.\r\n\r\nPertemuan kedua berfokus pada eksploitasi aplikasi web sebagai initial access vector. Peserta akan melakukan vulnerability assessment, mengeksplorasi teknik injection lanjutan, abuse terhadap mekanisme autentikasi, hingga memperoleh initial shell pada Ubuntu Server. Tahap ini menekankan pemahaman bagaimana sebuah aplikasi web dapat menjadi pintu masuk kompromi sistem secara menyeluruh.\r\n\r\nPada pertemuan ketiga, peserta akan mendalami fase post-exploitation, termasuk local enumeration, privilege escalation, persistence technique, serta pemahaman struktur infrastruktur target. Selain itu, peserta akan dilatih menyusun attack narrative dan reporting structure untuk menjelaskan bagaimana sistem gagal diamankan dari perspektif teknis dan strategis.\r\n\r\nMelalui bootcamp ini, peserta diharapkan mampu:\r\n\r\nMemahami dan menerapkan siklus serangan secara utuh.\r\n\r\nMengidentifikasi dan mengeksploitasi kerentanan secara sistematis.\r\n\r\nMenjaga anonimitas dan mengelola jejak digital dalam skenario simulasi.\r\n\r\nMenjelaskan proses kompromi sistem secara teknis dan terstruktur.\r\n\r\nMini Bootcamp ini ditujukan bagi mahasiswa Teknik Informatika dan penggiat keamanan siber yang ingin meningkatkan kompetensi teknis dalam skenario Offsec realistis dengan pendekatan profesional dan bertanggung jawab.','[\"Laptop dengan RAM minimal 8GB untuk menjalankan VM dengan stabil\",\"Penyimpanan kosong minimal 20GB untuk virtual disk\",\"Software VirtualBox atau VMware Workstation versi terbaru\",\"Image Kali Linux atau Parrot Security OS dan Ubuntu Server yang siap digunakan\",\"Koneksi internet stabil minimal 10 Mbps\"]','[\"E-Sertifikat\",\"Materi \\u0026 Modul\",\"Full Praktik\",\"Networking \\u0026 Grup Diskusi\",\"Rekaman Ulang\",\"Free Konsultasi\"]','2026-03-06','13.00 - Selesai','Abdul Rahman Wahid','','','Alya','Online via Zoom',50,0,'/storage/posters/65d5f24e-c2bb-4d7d-ab9a-9ad7103457fa.jpeg','paid',30000,1,'2026-02-14 05:23:38','2026-02-14 05:24:31'),('f9a46677-6c4e-4365-8fc8-7d22602366f4','bootcamp','Bootcamp','Hacking Modern Systems Ethically: Advanced Web, API, and Authentication Security Bootcamp for Cloud-Native and AI-Integrated Platforms','Hacking Modern Systems Ethically: Advanced Web, API, and Authentication Security Bootcamp for Cloud-Native and AI-Integrated Platforms adalah bootcamp intensif yang dirancang untuk membekali peserta dengan pemahaman dan keterampilan praktis dalam mengamankan aplikasi modern berbasis web, API, dan sistem autentikasi yang digunakan pada arsitektur cloud-native dan platform yang terintegrasi dengan AI.\r\n\r\nMelalui pendekatan hands-on lab dan studi kasus dunia nyata, peserta akan mempelajari teknik identifikasi celah keamanan pada aplikasi web modern, eksploitasi kerentanan API, serta analisis kesalahan implementasi autentikasi dan otorisasi seperti OAuth, JWT, dan session management. Bootcamp ini juga membahas bagaimana arsitektur microservices, cloud deployment, dan integrasi AI memperluas attack surface dan menuntut strategi keamanan yang lebih adaptif.\r\n\r\nKegiatan ini diselenggarakan oleh COCONUT Computer Club sebagai wadah pembelajaran praktis bagi mahasiswa dan pemula tingkat menengah yang ingin membangun kompetensi di bidang application security, ethical hacking, dan secure system design, sekaligus mempersiapkan peserta menghadapi tantangan keamanan pada sistem modern di dunia industri.','[\"Peserta wajib mengikuti kegiatan dari awal hingga akhir sesi.\",\"Hadir tepat waktu\",\"Berpakaian Rapi dan Sopan\",\"Wajib Oncamera\"]','[\"E-Sertifikat\",\"Modul\",\"Relasi\"]','2026-02-16','14:00 - 16:00 WIB','Nawat S.Kom., M.T., ','Saudah S.Kom','','Fajrul Bangsat','Online via Zoom',50,3,'/storage/posters/1ee6773f-0dc9-48e9-9da8-e8e5a3a75a50.jpg','free',0,1,'2026-02-08 16:24:18','2026-02-14 15:55:26');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posters`
--

DROP TABLE IF EXISTS `posters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posters` (
  `id` varchar(50) NOT NULL,
  `title` varchar(255) NOT NULL,
  `type` varchar(100) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `date` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posters`
--

LOCK TABLES `posters` WRITE;
/*!40000 ALTER TABLE `posters` DISABLE KEYS */;
INSERT INTO `posters` VALUES ('022c93d8-cad0-45b9-9337-60b0d60527ce','Introduction to Sveltekit: The Frontend Framework of the Future','Open Class','/storage/posters/c3d800a3-5003-4256-ae20-41a95b7ab9af.jpg','11 Oktober 2024','2026-02-08 02:05:16','2026-02-08 02:05:16'),('0b5756af-8925-4ea9-9247-82619c892b1a','Pengenalan Sistem Operasi dan Praktik Perintah Dasar Linux','Open Class','/storage/posters/76f64df2-5211-42f9-9382-32470fe0b8da.jpg','10 Agustus 2024','2026-02-08 02:04:24','2026-02-08 02:04:24'),('148b3001-13a6-4687-ac6d-cc6f1c1ead30','The Art of Prompting Crafting Better AI Responses','Webinar','/storage/posters/92af4885-a076-4e92-b61b-929d8b9376fe.jpeg','12 April 2025','2026-02-08 01:52:04','2026-02-08 01:52:04'),('3427485c-eac2-49de-a4e7-35951534b4a7','CRUD: Belajar Mengolah Data Menggunakan Golang dan MYSQL','Open Class','/storage/posters/eb362c6a-dc0f-4b57-b2a3-ebdcae57f6a2.png','10 Mei 2024','2026-02-08 02:03:16','2026-02-08 02:03:16'),('52784580-6c7b-4af4-aa30-6ff7cc8fc8b5','BLOCKCHAIN DEMYSTIFIED : The Technology Behind the Digital Revolution','Seminar','/storage/posters/1e20c916-e30f-4b1a-bdb5-ef33681a5667.jpeg','12 Oktober 2024','2026-02-08 01:56:56','2026-02-08 01:56:56'),('6236cb1f-07bb-45d3-976e-30430ec08ff6','Introduction to Machine Learning : Basic Concepts, Types, and Learning Process','Seminar','/storage/posters/8446b3ba-36f6-4a9a-9c23-de1985332981.jpeg','03 November 2023','2026-02-08 01:54:44','2026-02-08 01:54:44'),('6680d1ab-2be8-49c6-a1e3-2a56697120db','AI & Future Of Humanity : Penerapan Natural Language Processing (NLP) Dalam Mendiskripsikan Ulasan Pengguna','Webinar','/storage/posters/6001cc3f-c856-4b3d-8b16-cd7e33ce1c7b.png','4 Februari 2024','2026-02-08 01:46:12','2026-02-08 01:46:12'),('762ac7ab-9d03-4b5f-9780-4fdb4eb3f1be','THE DISPLAY IS MAGIC: Boostrap, Framework CSS & Javascript','Open Class','/storage/posters/64069131-cf9a-41a0-b83e-5cf704fa2da3.png','12 Januari 2024','2026-02-08 02:01:54','2026-02-08 02:01:54'),('7cd9e6bd-5d89-40ed-8897-beeac2c582d6','Blockchain 101 : Teknologi dasar di Balik Revolusi Digital & Web 3 : Evolusi Internet Berbasis Blockchain','Webinar','/storage/posters/036a684b-a819-403b-92d8-9f3ba1197373.jpg','22 & 23 Maret 2025','2026-02-08 01:50:20','2026-02-08 01:50:20'),('7e921eec-95a0-4c84-b73c-f463fda9c0a8','IOT-Based Network Management : Solusi Pintar dalam Pengelolaan Infrastruktur Jaringan','Seminar','/storage/posters/0594f038-9818-45d9-a2bb-01e4e10c7d59.jpg','24 Oktober 2025','2026-02-08 01:58:41','2026-02-08 01:58:41'),('89be13c6-28a8-4357-9da5-300ac426900d','Go REST, Go Fast: Membangun REST API dengan Golang','Open Class','/storage/posters/d705b684-8d6e-402b-b912-ec517bc30161.png','17 Oktober 2025','2026-02-08 02:09:24','2026-02-08 02:09:24'),('a12a5927-df7e-4363-bdba-0cec3266fcb3','Level UP Your Skills: Bangun Portofolio Interaktif dengan Next.js dari Nol','Open Class','/storage/posters/9a18b913-36b6-4d49-aa8a-2f3eef65252b.png','13 Juni 2025','2026-02-08 02:08:52','2026-02-08 02:08:52'),('a8da05a9-e123-4aba-b99d-167f8d734230','Web Development With Laravel : Pengantar HTML, CSS, Dan JavaScript & Dasar PHP dan Framework Laravel','Open Class','/storage/posters/8ea46dde-ffe6-4726-998f-c39c3110ccee.jpg','28 Desember 2024','2026-02-08 02:07:27','2026-02-08 02:07:27'),('f0f2c3e4-0d66-4ff8-8315-d0d79ec87b0f','Programming 101: Mengenal dan Belajar Lebih dalam Konsep Pemrograman','Open Class','/storage/posters/8d8b9c39-8cbe-4bd8-92d0-4ca11a72e5ae.jpg','02 Maret 2025','2026-02-08 02:08:19','2026-02-08 02:08:19'),('f67098da-3560-4e5e-b690-c00dbac5c48d','Introduce Your Self to Information Technology','Open Class','/storage/posters/d3a46bff-5fe1-4e2e-ac76-5dada332c8b1.png','12 & 13 Agustus 2023','2026-02-08 02:00:42','2026-02-08 02:00:42'),('fc47ce78-45c7-4279-9f67-122f8ac9d80d','Beyond SQL Injection: Modern Web Exploitation in API & AI Era','Open Class','/storage/posters/2817723d-badc-4ae7-bd0c-5ae391a23966.jpeg','01 Februari 2026','2026-02-08 02:23:03','2026-02-08 02:23:03');
/*!40000 ALTER TABLE `posters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `registrations`
--

DROP TABLE IF EXISTS `registrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `registrations` (
  `id` varchar(50) NOT NULL,
  `event_id` varchar(50) NOT NULL,
  `name` varchar(255) NOT NULL,
  `whatsapp` varchar(25) NOT NULL,
  `institution` varchar(255) DEFAULT NULL,
  `proof_image` varchar(255) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `registered_at` datetime NOT NULL DEFAULT current_timestamp(),
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `idx_reg_event` (`event_id`),
  KEY `idx_regs_name` (`name`(100)),
  CONSTRAINT `fk_registrations_event` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `registrations`
--

LOCK TABLES `registrations` WRITE;
/*!40000 ALTER TABLE `registrations` DISABLE KEYS */;
INSERT INTO `registrations` VALUES ('0e727515-dd97-4115-85aa-164883c72928','f9a46677-6c4e-4365-8fc8-7d22602366f4','kjbuyv','09-0uy75645321`','lkjhgyft','/storage/registrations/e3d3cd1a-1b8d-42cd-a27e-bb13d1ce0fc2.pdf','15 Jan [DRAFT] Jadwal Perkuliahan JTIK Semester Genap 25-26.pdf','2026-02-14 15:55:26','2026-02-14 15:55:26','2026-02-14 15:55:26'),('12a85d30-15d2-4e88-aaa8-9a415ebc4961','f9a46677-6c4e-4365-8fc8-7d22602366f4','morgan','0239427236288','unhas','/storage/registrations/f62a2f41-dca4-413f-9d8b-b71690b55b1d.png','Screenshot from 2025-07-29 03-05-19.png','2026-02-14 04:51:29','2026-02-14 04:51:29','2026-02-14 04:51:29'),('b622e5ab-921c-4572-80b2-c6b28b3ab9b0','f9a46677-6c4e-4365-8fc8-7d22602366f4','nawat','0821737137718','unm','/storage/registrations/b923a2f2-d499-4e6e-a4ef-510de41d0ba8.png','4K-Blue-Linux-Space.png','2026-02-08 16:27:20','2026-02-08 16:27:20','2026-02-08 16:27:20');
/*!40000 ALTER TABLE `registrations` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-02-17 20:03:24
