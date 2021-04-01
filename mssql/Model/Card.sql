--
-- Cloud API.
-- Prepared SQL queries for 'Card' definition.
--


--
-- SELECT template for table `Card`
--
SELECT `type`, `id`, `number`, `startDate`, `expiryDate`, `holderName`, `network`, `status`, `URI` FROM `Card` WHERE 1;

--
-- INSERT template for table `Card`
--
INSERT INTO `Card`(`type`, `id`, `number`, `startDate`, `expiryDate`, `holderName`, `network`, `status`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Card`
--
UPDATE `Card` SET `type` = ?, `id` = ?, `number` = ?, `startDate` = ?, `expiryDate` = ?, `holderName` = ?, `network` = ?, `status` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Card`
--
DELETE FROM `Card` WHERE 0;

