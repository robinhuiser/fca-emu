--
-- Cloud API.
-- Prepared SQL queries for 'EntityDebtor' definition.
--


--
-- SELECT template for table `EntityDebtor`
--
SELECT `debtorType`, `type`, `id`, `name`, `taxInformation`, `active`, `dateCreated`, `securityInformation`, `URI` FROM `EntityDebtor` WHERE 1;

--
-- INSERT template for table `EntityDebtor`
--
INSERT INTO `EntityDebtor`(`debtorType`, `type`, `id`, `name`, `taxInformation`, `active`, `dateCreated`, `securityInformation`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `EntityDebtor`
--
UPDATE `EntityDebtor` SET `debtorType` = ?, `type` = ?, `id` = ?, `name` = ?, `taxInformation` = ?, `active` = ?, `dateCreated` = ?, `securityInformation` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `EntityDebtor`
--
DELETE FROM `EntityDebtor` WHERE 0;

