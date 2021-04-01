--
-- Cloud API.
-- Prepared SQL queries for 'RoutingNumber' definition.
--


--
-- SELECT template for table `RoutingNumber`
--
SELECT `type`, `number` FROM `RoutingNumber` WHERE 1;

--
-- INSERT template for table `RoutingNumber`
--
INSERT INTO `RoutingNumber`(`type`, `number`) VALUES (?, ?);

--
-- UPDATE template for table `RoutingNumber`
--
UPDATE `RoutingNumber` SET `type` = ?, `number` = ? WHERE 1;

--
-- DELETE template for table `RoutingNumber`
--
DELETE FROM `RoutingNumber` WHERE 0;

