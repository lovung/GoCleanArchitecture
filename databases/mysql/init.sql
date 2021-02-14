-- create the databases
CREATE DATABASE IF NOT EXISTS gocleanarchitecture;

-- create the users for each database
CREATE USER admin IDENTIFIED BY 'gocleanarchitecture12345';
-- GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON gocleanarchitecture TO admin;
GRANT ALL PRIVILEGES ON DATABASE gocleanarchitecture TO admin;

