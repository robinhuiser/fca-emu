--
-- Cloud API.
-- Prepared SQL queries for 'ExchangeTransactionsList' definition.
--


--
-- SELECT template for table `ExchangeTransactionsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `exchanges` FROM `ExchangeTransactionsList` WHERE 1;

--
-- INSERT template for table `ExchangeTransactionsList`
--
INSERT INTO `ExchangeTransactionsList`(`status`, `messages`, `totalItems`, `nextCursor`, `exchanges`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `ExchangeTransactionsList`
--
UPDATE `ExchangeTransactionsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `exchanges` = ? WHERE 1;

--
-- DELETE template for table `ExchangeTransactionsList`
--
DELETE FROM `ExchangeTransactionsList` WHERE 0;

