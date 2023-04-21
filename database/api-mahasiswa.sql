-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 21, 2023 at 07:24 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `api-mahasiswa`
--

-- --------------------------------------------------------

--
-- Table structure for table `hobi`
--

CREATE TABLE `hobi` (
  `id` int(11) NOT NULL,
  `nama_hobi` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `hobi`
--

INSERT INTO `hobi` (`id`, `nama_hobi`) VALUES
(1, 'membaca'),
(2, 'bernyanyi'),
(3, 'sepak bola'),
(4, 'bermain game'),
(5, 'memancing');

-- --------------------------------------------------------

--
-- Table structure for table `jurusan`
--

CREATE TABLE `jurusan` (
  `id` int(11) NOT NULL,
  `nama_jurusan` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `jurusan`
--

INSERT INTO `jurusan` (`id`, `nama_jurusan`) VALUES
(1, 'Teknik Mesin'),
(2, 'Teknik Informatika'),
(3, 'Peternakan'),
(4, 'Kedokteran'),
(5, 'Pertanian'),
(6, 'Teknik Sipil');

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `id` int(11) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `usia` int(2) UNSIGNED NOT NULL,
  `gender` enum('0','1') NOT NULL,
  `tanggal_registrasi` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`id`, `nama`, `usia`, `gender`, `tanggal_registrasi`) VALUES
(1, 'Dionysius Yudha Riadi', 20, '1', '2020-01-02 15:04:05'),
(2, 'Budi Alexander', 21, '1', '2020-01-02 15:04:05'),
(3, 'Himuro Sakura', 19, '0', '2023-04-21 07:17:50'),
(4, 'Hinata Himawari', 19, '0', '2023-04-21 07:18:10'),
(5, 'Tenten Uzumaki', 20, '0', '2023-04-21 07:18:10');

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa_hobi`
--

CREATE TABLE `mahasiswa_hobi` (
  `id` int(11) NOT NULL,
  `id_mahasiswa` int(11) NOT NULL,
  `id_hobi` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `mahasiswa_hobi`
--

INSERT INTO `mahasiswa_hobi` (`id`, `id_mahasiswa`, `id_hobi`) VALUES
(1, 1, 4),
(2, 4, 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `hobi`
--
ALTER TABLE `hobi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jurusan`
--
ALTER TABLE `jurusan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mahasiswa_hobi`
--
ALTER TABLE `mahasiswa_hobi`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_mahasiswa_hobi__hobi` (`id_hobi`),
  ADD KEY `fk_mahasiswa_hobi__mahasiswa` (`id_mahasiswa`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `hobi`
--
ALTER TABLE `hobi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `jurusan`
--
ALTER TABLE `jurusan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `mahasiswa_hobi`
--
ALTER TABLE `mahasiswa_hobi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `mahasiswa_hobi`
--
ALTER TABLE `mahasiswa_hobi`
  ADD CONSTRAINT `fk_mahasiswa_hobi__hobi` FOREIGN KEY (`id_hobi`) REFERENCES `hobi` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_mahasiswa_hobi__mahasiswa` FOREIGN KEY (`id_mahasiswa`) REFERENCES `mahasiswa` (`id`) ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
