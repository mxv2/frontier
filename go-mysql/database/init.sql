CREATE TABLE users
(
    id   INT(5) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(50)     NOT NULL,
    age  INT(2) UNSIGNED NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 9331
  DEFAULT CHARSET = utf8;

INSERT INTO users(name, age)
VALUES ('Simon', 25),
       ('Andrew', 20),
       ('James', 25),
       ('John', 31),
       ('Phillip', 26),
       ('Thaddeus', 25),
       ('Bartholomew', 24),
       ('James the Less', 22),
       ('Matthew', 23),
       ('Simon the Canaanite', 30),
       ('Judas Iscariot', 27);
