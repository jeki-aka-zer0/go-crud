USE go_crud;

DROP TABLE IF EXISTS `engineers`;
CREATE TABLE `engineers`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `first_name` varchar(255) NOT NULL,
    `last_name`  varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

LOCK
TABLES `engineers` WRITE;
INSERT INTO `engineers` (`first_name`, `last_name`)
VALUES ('Barbara', 'Liskov'),
       ('Linus', 'Torvalds'),
       ('Robert', 'Martin'),
       ('Martin', 'Fowler'),
       ('Tim', 'Berners-Lee');
UNLOCK
TABLES;