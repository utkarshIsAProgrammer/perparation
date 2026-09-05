-- show all databases
SHOW DATABASES;

/* DROP DATABASE PC; */

-- create a database

CREATE DATABASE MACHINES;
CREATE DATABASE COMPUTER;

-- use a database
USE MACHINES;

-- show selected database
SELECT DATABASE();

-- create a table
CREATE TABLE GAMES (
    name VARCHAR(50),
    ratings INT
)

-- show the tables in the database
SHOW TABLES;

-- show the columns of the table
SHOW COLUMNS FROM GAMES;

-- describe the table structure
DESC GAMES;

-- default values for the table
CREATE TABLE Best_Games (
    name VARCHAR(50) DEFAULT 'Minecraft',
    ratings INT 
)

-- show the columns of the table
SHOW COLUMNS FROM Best_Games;

-- describe the table structure
DESC Best_Games;

-- delete a table
DROP TABLE GAMES;

-- data types
CREATE TABLE Data_Types (
    int_col INT,
    big_int_col BIGINT,
    double_distance DOUBLE,
    decimal_col DECIMAL(10,2),
    varchar_col VARCHAR(255),
    text_col TEXT,
    char_txt CHAR(3),
    date_col DATE,
    datetime_col DATETIME,
    timestamp_col TIMESTAMP,
    boolean_col BOOLEAN,
    json_col JSON,
    enum_role ENUM('admin', 'user', 'tester', 'moderator')
);

-- show columns from the table
SHOW COLUMNS FROM Data_Types;

-- insert data into table
INSERT INTO Data_Types (
    int_col,
    big_int_col,
    double_distance,
    decimal_col,
    varchar_col,
    text_col,
    char_txt,
    date_col,
    datetime_col,
    timestamp_col,
    boolean_col,
    json_col,
    enum_role
)
VALUES (
    5,
    5797539847530945028,
    45345.56,
    9.25,
    'Indiedev!',
    'This is a whole text.',
    'Hey',
    '2026-09-05',
    '2026-09-05 18:13:00',
    '2026-09-05 18:13:00',
    TRUE,
    '{"name": "Indiedev", "level": 10}',
    'admin'
);

-- show entire table
SELECT * FROM Data_Types;

-- select specific property from the table
SELECT int_col FROM Data_Types;

-- select multiple properties from the table
SELECT int_col, varchar_col FROM Data_Types;
