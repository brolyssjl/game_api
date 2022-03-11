-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: gamedb
-- ------------------------------------------------------
-- Server version	8.0.28-0ubuntu0.20.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `game_states`
--

DROP TABLE IF EXISTS `game_states`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game_states` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `games_played` mediumint unsigned NOT NULL DEFAULT '0',
  `score` mediumint unsigned NOT NULL DEFAULT '0',
  `user_id` varchar(75) NOT NULL,
  PRIMARY KEY (`id`,`user_id`),
  KEY `fk_game_states_user_idx` (`user_id`),
  CONSTRAINT `fk_game_states_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game_states`
--

LOCK TABLES `game_states` WRITE;
/*!40000 ALTER TABLE `game_states` DISABLE KEYS */;
INSERT INTO `game_states` VALUES (1,5,10,'fc8a8b83-a764-4655-a918-5cb3b957bf4e');
/*!40000 ALTER TABLE `game_states` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_has_friends`
--

DROP TABLE IF EXISTS `user_has_friends`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_has_friends` (
  `user_id` varchar(75) NOT NULL,
  `friend_id` varchar(75) NOT NULL,
  PRIMARY KEY (`user_id`,`friend_id`),
  KEY `fk_User_has_friend_User2_idx` (`friend_id`),
  KEY `fk_User_has_friend_User1_idx` (`user_id`),
  CONSTRAINT `fk_User_has_friend_User1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_User_has_friend_User2` FOREIGN KEY (`friend_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_has_friends`
--

LOCK TABLES `user_has_friends` WRITE;
/*!40000 ALTER TABLE `user_has_friends` DISABLE KEYS */;
INSERT INTO `user_has_friends` VALUES ('fc8a8b83-a764-4655-a918-5cb3b957bf4e','18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1'),('fc8a8b83-a764-4655-a918-5cb3b957bf4e','2d18862b-b9c3-40f5-803e-5e100a520249'),('fc8a8b83-a764-4655-a918-5cb3b957bf4e','f9a9af78-6681-4d7d-8ae7-fc41e7a24d08');
/*!40000 ALTER TABLE `user_has_friends` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(75) NOT NULL,
  `name` varchar(25) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='Users table';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1','George'),('2d18862b-b9c3-40f5-803e-5e100a520249','Doom Slayer'),('52a47226-19bf-4404-bf74-616e4dc04459','Carlitos'),('f9a9af78-6681-4d7d-8ae7-fc41e7a24d08','Kirito'),('fc8a8b83-a764-4655-a918-5cb3b957bf4e','Jonatan');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-11  0:05:51
