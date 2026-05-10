CREATE DATABASE IF NOT EXISTS sharing_vision;

USE sharing_vision;

CREATE TABLE IF NOT EXISTS posts (
    id           INT           NOT NULL AUTO_INCREMENT,
    title        VARCHAR(200)  NOT NULL,
    content      TEXT          NOT NULL,
    category     VARCHAR(100)  NOT NULL,
    created_date TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status       VARCHAR(100)  NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
