--
-- Cloud API.
-- Prepared SQL queries for 'ExchangeTransactionResult' definition.
--


--
-- SELECT template for table `ExchangeTransactionResult`
--
SELECT `status`, `reason`, `exchangeTransactionId` FROM `ExchangeTransactionResult` WHERE 1;

--
-- INSERT template for table `ExchangeTransactionResult`
--
INSERT INTO `ExchangeTransactionResult`(`status`, `reason`, `exchangeTransactionId`) VALUES (?, ?, ?);

--
-- UPDATE template for table `ExchangeTransactionResult`
--
UPDATE `ExchangeTransactionResult` SET `status` = ?, `reason` = ?, `exchangeTransactionId` = ? WHERE 1;

--
-- DELETE template for table `ExchangeTransactionResult`
--
DELETE FROM `ExchangeTransactionResult` WHERE 0;

