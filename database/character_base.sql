-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 04, 2021 at 01:52 PM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gw_emu`
--

-- --------------------------------------------------------

--
-- Table structure for table `character_base`
--

CREATE TABLE `character_base` (
  `id` int(4) UNSIGNED NOT NULL,
  `account_id` int(32) NOT NULL,
  `name` char(32) COLLATE utf8_unicode_ci NOT NULL,
  `gender` enum('female','male') COLLATE utf8_unicode_ci NOT NULL DEFAULT 'male',
  `GM` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `camp` tinyint(1) UNSIGNED NOT NULL,
  `profession` tinyint(1) UNSIGNED NOT NULL,
  `fighter_job_lv` tinyint(1) UNSIGNED NOT NULL,
  `scholar_job_lv` tinyint(1) UNSIGNED NOT NULL,
  `fighter_job_exp` int(4) UNSIGNED NOT NULL,
  `scholar_job_exp` int(4) UNSIGNED NOT NULL,
  `curHP` int(2) UNSIGNED NOT NULL DEFAULT 150,
  `curMP` int(2) UNSIGNED NOT NULL,
  `status` tinyint(1) NOT NULL,
  `belief` tinyint(1) UNSIGNED NOT NULL,
  `prestige` smallint(2) UNSIGNED NOT NULL,
  `earl_rank` tinyint(1) UNSIGNED NOT NULL,
  `consortia` smallint(2) UNSIGNED ZEROFILL NOT NULL DEFAULT 00,
  `consortia_job` tinyint(1) UNSIGNED ZEROFILL NOT NULL DEFAULT 0 COMMENT '6: President, 5: Vice President, 4: Director, 3: Elite 2: Member, 1: Student, 0: No',
  `consortia_contribute` int(4) UNSIGNED NOT NULL,
  `store_num` int(2) UNSIGNED NOT NULL DEFAULT 10,
  `bag_num` int(2) UNSIGNED NOT NULL DEFAULT 1,
  `hair_style` tinyint(4) UNSIGNED NOT NULL,
  `face_shap` tinyint(4) UNSIGNED DEFAULT NULL,
  `Map` tinyint(1) UNSIGNED NOT NULL DEFAULT 1,
  `Pos_X` float(10,5) NOT NULL DEFAULT 0.00000,
  `Pos_Z` float(10,5) NOT NULL DEFAULT 0.00000,
  `Money` int(4) UNSIGNED NOT NULL DEFAULT 0,
  `Stone` int(4) UNSIGNED NOT NULL DEFAULT 0,
  `SkillPoint` int(4) UNSIGNED ZEROFILL NOT NULL DEFAULT 0000,
  `SkillExp` int(4) UNSIGNED ZEROFILL NOT NULL DEFAULT 0000,
  `MaxHP` int(4) UNSIGNED NOT NULL DEFAULT 0,
  `MaxMP` int(4) UNSIGNED NOT NULL DEFAULT 0,
  `Register_time` datetime NOT NULL,
  `LastLogin_time` datetime NOT NULL,
  `mutetime` int(4) UNSIGNED ZEROFILL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `character_base`
--

INSERT INTO `character_base` (`id`, `account_id`, `name`, `gender`, `GM`, `camp`, `profession`, `fighter_job_lv`, `scholar_job_lv`, `fighter_job_exp`, `scholar_job_exp`, `curHP`, `curMP`, `status`, `belief`, `prestige`, `earl_rank`, `consortia`, `consortia_job`, `consortia_contribute`, `store_num`, `bag_num`, `hair_style`, `face_shap`, `Map`, `Pos_X`, `Pos_Z`, `Money`, `Stone`, `SkillPoint`, `SkillExp`, `MaxHP`, `MaxMP`, `Register_time`, `LastLogin_time`, `mutetime`) VALUES
(1, 1, 'sush1', 'male', 1, 0, 0, 1, 0, 0, 0, 123456789, 177, 0, 0, 0, 0, 00, 0, 0, 10, 1, 0, 0, 1, -108.00000, -122.00000, 0, 0, 0000, 0000, 1368, 177, '2008-04-28 15:19:04', '0000-00-00 00:00:00', 0000);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `character_base`
--
ALTER TABLE `character_base`
  ADD PRIMARY KEY (`id`,`name`),
  ADD UNIQUE KEY `RoleName` (`name`),
  ADD KEY `Accounts` (`account_id`),
  ADD KEY `index` (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `character_base`
--
ALTER TABLE `character_base`
  MODIFY `id` int(4) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `character_base`
--
ALTER TABLE `character_base`
  ADD CONSTRAINT `character_base_ibfk_1` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
