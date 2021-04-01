--
-- Cloud API.
-- Prepared SQL queries for 'Entity' definition.
--


--
-- SELECT template for table `Entity`
--
SELECT `type`, `id`, `name`, `taxInformation`, `active`, `dateCreated`, `securityInformation`, `URI` FROM `Entity` WHERE 1;

--
-- INSERT template for table `Entity`
--
INSERT INTO `Entity`(`type`, `id`, `name`, `taxInformation`, `active`, `dateCreated`, `securityInformation`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Entity`
--
UPDATE `Entity` SET `type` = ?, `id` = ?, `name` = ?, `taxInformation` = ?, `active` = ?, `dateCreated` = ?, `securityInformation` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Entity`
--
DELETE FROM `Entity` WHERE 0;

