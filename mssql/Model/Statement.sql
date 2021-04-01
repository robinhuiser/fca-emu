--
-- Cloud API.
-- Prepared SQL queries for 'Statement' definition.
--


--
-- SELECT template for table `Statement`
--
SELECT `id`, `accountId`, `name`, `attributes`, `dateCreated`, `dateRevised`, `items` FROM `Statement` WHERE 1;

--
-- INSERT template for table `Statement`
--
INSERT INTO `Statement`(`id`, `accountId`, `name`, `attributes`, `dateCreated`, `dateRevised`, `items`) VALUES (?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Statement`
--
UPDATE `Statement` SET `id` = ?, `accountId` = ?, `name` = ?, `attributes` = ?, `dateCreated` = ?, `dateRevised` = ?, `items` = ? WHERE 1;

--
-- DELETE template for table `Statement`
--
DELETE FROM `Statement` WHERE 0;

