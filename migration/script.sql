SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS activity_levels;
DROP TABLE IF EXISTS health_goals;

SET FOREIGN_KEY_CHECKS = 1;

-- Create the activity_levels table
CREATE TABLE IF NOT EXISTS activity_levels (
    al_id VARCHAR(36) PRIMARY KEY,
    al_type BIGINT NOT NULL,
    al_desc VARCHAR(255) NOT NULL,
    al_value FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create the health_goals table
CREATE TABLE IF NOT EXISTS health_goals (
    hg_id VARCHAR(36) PRIMARY KEY,
    hg_type BIGINT NOT NULL,
    hg_desc VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    gender BIGINT NOT NULL,
    telp VARCHAR(255) NOT NULL,
    profpic VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL,
    place VARCHAR(255) NOT NULL,
    height FLOAT NOT NULL,
    weight FLOAT NOT NULL,
    weight_goal FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    hg_id VARCHAR(36),
    al_id VARCHAR(36)
);

DROP TABLE IF EXISTS tokens;

CREATE TABLE tokens (
    id VARCHAR(36) PRIMARY KEY,
    userId VARCHAR(36) NOT NULL,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    INDEX idx_tokens_deleted_at (deleted_at),
    CONSTRAINT fk_tokens_user FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

DESCRIBE tokens;

CREATE TABLE Stores (
    STORE_ID VARCHAR(255) PRIMARY KEY,
    STORE_NAME VARCHAR(255) NOT NULL,
    STORE_USERNAME VARCHAR(255) NOT NULL,
    STORE_ADDRESS VARCHAR(255) NOT NULL,
    STORE_CONTACT VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    USER_ID VARCHAR(255),
    FOREIGN KEY (USER_ID) REFERENCES Users(ID)
);

-- Add foreign key constraints
ALTER TABLE users
ADD CONSTRAINT FK_users_health_goals
FOREIGN KEY (hg_id) REFERENCES health_goals(hg_id);

ALTER TABLE users
DROP FOREIGN KEY FK_users_activitylevel;

ALTER TABLE users
DROP FOREIGN KEY FK_users_health_goals;

ALTER TABLE users
ADD CONSTRAINT FK_users_activity_levels
FOREIGN KEY (al_id) REFERENCES activity_levels(al_id);

SELECT constraint_name, table_name, constraint_type
FROM information_schema.table_constraints
WHERE table_schema = 'nutriomatic'
ORDER BY table_name, constraint_type;

SELECT CONCAT('ALTER TABLE ', table_name, ' DROP FOREIGN KEY ', constraint_name, ';') AS sql_command
FROM information_schema.table_constraints
WHERE constraint_type = 'FOREIGN KEY'
AND table_schema = 'nutriomatic';


ALTER TABLE activity_levels DROP FOREIGN KEY fk_users_activitylevel;
ALTER TABLE health_goals DROP FOREIGN KEY fk_users_healthgoal;
ALTER TABLE tokens DROP FOREIGN KEY fk_tokens_user;
ALTER TABLE users DROP FOREIGN KEY FK_users_activity_levels;
ALTER TABLE users DROP FOREIGN KEY FK_users_health_goals;