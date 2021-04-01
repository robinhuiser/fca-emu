--
-- Cloud API.
-- Prepared SQL queries for 'EntityCreditor' definition.
--


--
-- SELECT template for table `EntityCreditor`
--
SELECT `creditorType`, `target`, `type`, `id`, `name`, `taxInformation`, `active`, `dateCreated`, `securityInformation`, `URI` FROM `EntityCreditor` WHERE 1;

--
-- INSERT template for table `EntityCreditor`
--
INSERT INTO `EntityCreditor`(`creditorType`, `target`, `type`, `id`, `name`, `taxInformation`, `active`, `dateCreated`, `securityInformation`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `EntityCreditor`
--
UPDATE `EntityCreditor` SET `creditorType` = ?, `target` = ?, `type` = ?, `id` = ?, `name` = ?, `taxInformation` = ?, `active` = ?, `dateCreated` = ?, `securityInformation` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `EntityCreditor`
--
DELETE FROM `EntityCreditor` WHERE 0;

