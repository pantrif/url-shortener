CREATE DATABASE IF NOT EXISTS shortener;
use shortener;
CREATE TABLE `shortened_urls` (
  `id` int(12) unsigned NOT NULL auto_increment,
  `long_url` varchar(255) NOT NULL,
  `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY  (`id`),
  UNIQUE KEY `long` (`long_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;