# CREATE USER TABLE
CREATE TABLE IF NOT EXISTS user (
    id INT(11) NOT NULL auto_increment,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255) ,
    PRIMARY KEY (id)
);
# DELETE FROM user;
# INSERT INTO user (id, last_name, first_name, address, city)
# VALUES (1, "Maruejol", "Cedric", "rue des fesses", "fesseland");