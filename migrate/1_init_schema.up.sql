# CREATE USER TABLE
CREATE TABLE toto (
    id INT(11) NOT NULL auto_increment,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255) ,
    PRIMARY KEY (id)
);
INSERT INTO toto VALUES (1, "Maruejol", "Cedric", "rue des fesses", "fesseland");