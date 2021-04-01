--
-- Cloud API.
-- Prepared SQL queries for 'TransactionsList' definition.
--


--
-- SELECT template for table `TransactionsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `transactions` FROM `TransactionsList` WHERE 1;

--
-- INSERT template for table `TransactionsList`
--
INSERT INTO `TransactionsList`(`status`, `messages`, `totalItems`, `nextCursor`, `transactions`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `TransactionsList`
--
UPDATE `TransactionsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `transactions` = ? WHERE 1;

--
-- DELETE template for table `TransactionsList`
--
DELETE FROM `TransactionsList` WHERE 0;

