create database test;

use test

create table UserInfo(
    ID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Username char(255) NOT NULL,
    Email    char(255) NOT NULL,
    Password char(255) NOT NULL,
    PRIMARY KEY (ID)
);