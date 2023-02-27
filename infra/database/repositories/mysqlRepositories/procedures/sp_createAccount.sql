DROP PROCEDURE IF EXISTS `sp_createAccount`;

DELIMITER $$
CREATE PROCEDURE `sp_createAccount` (
    IN `in_holder` INT(11),
    IN `in_agency` VARCHAR(150),
    IN `in_number` VARCHAR(150)
)
BEGIN
    INSERT INTO `account`(`holder`,`agency`,`number`, `activated`, `blocked`) VALUES (`in_holder`, `in_agency`, `in_number`, true, true);

    SET @li := (SELECT LAST_INSERT_ID()); 

    INSERT INTO `account_statement`(`account`,`transaction_type`) VALUES (@li, 1);
END $$
DELIMITER ;
