--
-- Cloud API.
-- Prepared SQL queries for 'AccountsList' definition.
--


--
-- SELECT template for table `AccountsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `accounts` FROM `AccountsList` WHERE 1;

--
-- INSERT template for table `AccountsList`
--
INSERT INTO `AccountsList`(`status`, `messages`, `totalItems`, `nextCursor`, `accounts`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `AccountsList`
--
UPDATE `AccountsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `accounts` = ? WHERE 1;

--
-- DELETE template for table `AccountsList`
--
DELETE FROM `AccountsList` WHERE 0;

