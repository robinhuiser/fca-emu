--
-- Cloud API.
-- Prepared SQL queries for 'Product' definition.
--


--
-- SELECT template for table `Product`
--
SELECT `id`, `number`, `type`, `typeName`, `subType`, `subTypeName`, `URI` FROM `Product` WHERE 1;

--
-- INSERT template for table `Product`
--
INSERT INTO `Product`(`id`, `number`, `type`, `typeName`, `subType`, `subTypeName`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Product`
--
UPDATE `Product` SET `id` = ?, `number` = ?, `type` = ?, `typeName` = ?, `subType` = ?, `subTypeName` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Product`
--
DELETE FROM `Product` WHERE 0;

