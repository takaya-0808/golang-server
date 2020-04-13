create database test2;

use test2

create table userdata(
    ID MEDIUMINT NOT NULL AUTO_INCREMENT,
    username char(255) NOT NULL,
    email    char(255) NOT NULL,
    password char(255) NOT NULL,
    PRIMARY KEY (ID)
);%               