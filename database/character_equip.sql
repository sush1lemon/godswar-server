-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 14, 2021 at 03:47 PM
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
-- Table structure for table `character_equip`
--

CREATE TABLE `character_equip` (
  `user_id` int(4) UNSIGNED NOT NULL,
  `body_part_id` tinyint(1) NOT NULL,
  `prop_id` smallint(2) NOT NULL,
  `type1` tinyint(1) DEFAULT NULL,
  `quality1` tinyint(1) DEFAULT NULL,
  `value1` float(4,0) DEFAULT NULL,
  `type2` tinyint(1) DEFAULT NULL,
  `quality2` tinyint(1) DEFAULT NULL,
  `value2` float(4,0) DEFAULT NULL,
  `type3` tinyint(1) DEFAULT NULL,
  `quality3` tinyint(1) DEFAULT NULL,
  `value3` float(4,0) DEFAULT NULL,
  `type4` tinyint(1) DEFAULT NULL,
  `quality4` tinyint(1) DEFAULT NULL,
  `value4` float(4,0) DEFAULT NULL,
  `type5` tinyint(1) DEFAULT NULL,
  `quality5` tinyint(1) DEFAULT NULL,
  `value5` float(4,0) DEFAULT NULL,
  `isbind` tinyint(1) UNSIGNED DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=gb2312;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `character_equip`
--
ALTER TABLE `character_equip`
  ADD PRIMARY KEY (`user_id`,`body_part_id`),
  ADD KEY `index` (`user_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
