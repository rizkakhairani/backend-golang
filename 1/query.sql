CREATE DATABASE db_backend_golang;

USE db_backend_golang;

CREATE TABLE USER (
	ID INT NOT NULL AUTO_INCREMENT,
    UserName VARCHAR(30) NOT NULL,
    Parent INT NOT NULL,
    PRIMARY KEY (ID)
);

INSERT INTO USER (UserName, Parent) VALUES 
	('Ali', 2),
    ('Budi', 0),
    ('Cecep', 1);

SELECT USER.ID AS ID, USER.UserName AS UserName, 
	PARENT.UserName AS ParentUserName
FROM USER
LEFT JOIN USER PARENT 
ON USER.Parent = PARENT.ID;