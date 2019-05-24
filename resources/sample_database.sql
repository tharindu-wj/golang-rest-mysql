-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: May 24, 2019 at 02:55 PM
-- Server version: 5.7.26-0ubuntu0.18.04.1
-- PHP Version: 7.2.17-0ubuntu0.18.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_rest_mysql`
--

-- --------------------------------------------------------

--
-- Table structure for table `companies`
--

CREATE TABLE `companies` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `location` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`id`, `name`, `location`, `created_at`) VALUES
(1, 'Company 1', 'Location 1', '2019-05-23 10:48:40'),
(2, 'Company 2', 'Location 2', '2019-05-23 10:48:40'),
(3, 'Company 3', 'Location 3', '2019-05-23 10:48:40'),
(4, 'Company 4', 'Location 4', '2019-05-23 10:48:40');

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `body` longtext NOT NULL,
  `user` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `title`, `body`, `user`, `created_at`) VALUES
(1, 'What is Lorem Ipsum?', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.\r\n', 1, '2019-05-20 03:25:21'),
(4, 'Where can I get some?', 'There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don\'t look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn\'t anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.', 1, '2019-05-20 03:28:00'),
(5, 'test title', 'body', 2, '2019-05-23 13:10:24'),
(6, 'test title', 'body', 2, '2019-05-23 13:57:09'),
(7, 'test title', 'body', 2, '2019-05-23 14:03:45'),
(8, 'test title', 'body', 2, '2019-05-23 14:04:44'),
(9, 'test title', 'body', 2, '2019-05-23 14:05:15'),
(10, 'test title', 'body', 2, '2019-05-23 14:17:12'),
(11, 'test title', 'body', 2, '2019-05-23 14:18:20'),
(12, 'test title', 'body', 2, '2019-05-23 14:18:58'),
(13, 'asdf1234', 'body 1234', 1, '2019-05-23 14:32:14'),
(14, 'asdf1234', 'body 1234', 1, '2019-05-23 14:36:11'),
(15, '', '', 2, '2019-05-23 14:48:53'),
(16, '', '', 2, '2019-05-23 14:53:15'),
(17, '', '', 2, '2019-05-23 14:54:07'),
(18, 'qwerty', '', 2, '2019-05-23 14:56:37'),
(19, 'qwerty', '', 1, '2019-05-23 14:59:04'),
(20, 'qwerty', 'asdf12345', 1, '2019-05-23 14:59:27'),
(21, 'qwerty', 'asdf12345', 1, '2019-05-23 15:00:18'),
(22, 'qwerty', 'asdf12345', 1, '2019-05-23 15:13:29'),
(23, 'qwerty1234', 'asdf12345', 1, '2019-05-23 15:13:59'),
(24, 'qwerty1234', 'asdf12345', 1, '2019-05-23 15:14:20'),
(25, 'qwerty1234', 'asdf12345', 1, '2019-05-24 09:03:25'),
(26, 'Blog Post', 'asdf12345', 1, '2019-05-24 09:04:28');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `company` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `company`, `created_at`) VALUES
(1, 'test user', 'asdf@gmail.com', 1, '2019-05-19 18:30:39'),
(2, 'Test 2', 'test@gmail.com', 1, '2019-05-20 04:34:06'),
(4, 'New User', 'asdf3@gmail.com', 1, '2019-05-23 15:19:53');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `companies`
--
ALTER TABLE `companies`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user` (`user`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `company` (`company`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `companies`
--
ALTER TABLE `companies`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `posts`
--
ALTER TABLE `posts`
  ADD CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`user`) REFERENCES `users` (`id`);

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
