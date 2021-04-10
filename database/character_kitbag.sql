-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 09, 2021 at 03:05 PM
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
-- Database: `godswar`
--

-- --------------------------------------------------------

--
-- Table structure for table `character_kitbag`
--

CREATE TABLE `character_kitbag` (
  `user_id` int(4) UNSIGNED NOT NULL,
  `kitbag_1` varchar(4000) DEFAULT NULL,
  `kitbag_2` varchar(4000) DEFAULT NULL,
  `kitbag_3` varchar(4000) DEFAULT NULL,
  `kitbag_4` varchar(4000) DEFAULT NULL,
  `storage` varchar(4000) DEFAULT NULL,
  `equip` varchar(2000) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=gb2312;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `character_kitbag`
--
ALTER TABLE `character_kitbag`
  ADD PRIMARY KEY (`user_id`),
  ADD KEY `index` (`user_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
