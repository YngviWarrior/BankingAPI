USE dock;

DROP TABLE IF EXISTS `holder`;
CREATE TABLE IF NOT EXISTS `holder`(
    `holder` BIGINT(20) AUTO_INCREMENT,
    `full_name` VARCHAR(150) NOT NULL,
    `cpf` VARCHAR(150) NOT NULL UNIQUE,
    `verified` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`holder`)
);

DROP TABLE IF EXISTS `account`;
CREATE TABLE IF NOT EXISTS `account`(
    `account` BIGINT(20) AUTO_INCREMENT,
    `holder` BIGINT(20) NOT NULL,
    `agency` VARCHAR(150) NOT NULL,
    `number` VARCHAR(150) NOT NULL UNIQUE,
    `balance` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `activated` TINYINT(1) NOT NULL DEFAULT 0,
    `blocked` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`account`),
    FOREIGN KEY (`holder`) REFERENCES `holder`(`holder`)
);

DROP TABLE IF EXISTS `transaction_type`;
CREATE TABLE IF NOT EXISTS `transaction_type`(
    `transaction_type` TINYINT(11) AUTO_INCREMENT,
    `description` VARCHAR(150) NOT NULL DEFAULT 0,
    PRIMARY KEY (`transaction_type`)
);

INSERT INTO `transaction_type` (`description`) VALUES ('openning account'),
('deposit'),
('withdraw'),
('transfer');

DROP TABLE IF EXISTS `account_statement`;
CREATE TABLE IF NOT EXISTS `account_statement`(
    `account_statement` BIGINT(20) AUTO_INCREMENT,
    `account` BIGINT(20) NOT NULL,
    `transaction_type` TINYINT(11) NOT NULL,    
    `previous_balance` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `current_balance` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `registered_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`account_statement`),
    FOREIGN KEY (`account`) REFERENCES `account`(`account`),
    FOREIGN KEY (`transaction_type`) REFERENCES `transaction_type`(`transaction_type`)
);

