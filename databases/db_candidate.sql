/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 10.1.32-MariaDB : Database - db_candidate
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_candidate` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `db_candidate`;

/*Table structure for table `t_candidate` */

DROP TABLE IF EXISTS `t_candidate`;

CREATE TABLE `t_candidate` (
  `candidate_id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `phone_number` varchar(20) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `dob` varchar(10) DEFAULT NULL,
  `pob` varchar(255) DEFAULT NULL,
  `gender` varchar(1) DEFAULT NULL,
  `year_exp` varchar(15) DEFAULT NULL,
  `last_salary` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`candidate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;

/*Data for the table `t_candidate` */

insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (1,'email1@example.com','1234567890','John Doe','1990-01-01','City1','M','< 1 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (2,'email2@example.com','9876543210','Jane Smith','1995-02-15','City2','F','2 - 3 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (3,'email3@example.com','5555555555','Alice Johnson','1998-05-20','City3','F','3 - 4 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (4,'email4@example.com','9999999999','Bob Williams','1987-09-10','City4','M','> 5 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (5,'email5@example.com','1111111111','Emily Davis','1993-07-08','City5','F','> 5 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (6,'email6@example.com','7777777777','Michael Brown','1985-04-25','City6','M','3 - 4 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (7,'email7@example.com','3333333333','Olivia Lee','2000-03-12','City7','F','2 - 3 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (8,'email8@example.com','4444444444','William Jones','1992-11-30','City8','M','< 1 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (9,'email9@example.com','6666666666','Sophia Miller','1996-12-18','City9','F','< 1 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (10,'email10@example.com','8888888888','James Watson','1989-08-05','City10','M','2 - 3 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (11,'email11@example.com','9999999999','Kyle Wilson','1989-08-05','City10','M','3 - 4 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (12,'email12@example.com','1010101010','Triya Waymen','1989-08-05','City10','M','> 5 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (13,'email13@example.com','1111111111','Plyio Iosas','1989-08-05','City10','M','3 - 4 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (14,'email14@example.com','1212121212','James Wilson','1989-08-05','City10','M','< 1 years','Rp4.000.000');
insert  into `t_candidate`(`candidate_id`,`email`,`phone_number`,`full_name`,`dob`,`pob`,`gender`,`year_exp`,`last_salary`) values (15,'email15@example.com','1313131313','James Iopdf','1989-08-05','City10','M','< 1 years','Rp4.000.000');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
