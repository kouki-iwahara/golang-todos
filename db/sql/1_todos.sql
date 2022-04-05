DROP DATABASE IF EXISTS golang_todos_db;
CREATE DATABASE golang_todos_db;
USE golang_todos_db;
DROP TABLE IF EXISTS todos;
 
CREATE TABLE todos (
id         INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
content    VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
)DEFAULT CHARACTER SET=utf8mb4;


INSERT INTO todos (content) VALUES ("GOの学習");
INSERT INTO todos (content) VALUES ("Linuxの学習");
INSERT INTO todos (content) VALUES ("本屋に行く");