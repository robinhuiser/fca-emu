--
-- Cloud API.
-- Prepared SQL queries for 'Relationship' definition.
--


--
-- SELECT template for table `Relationship`
--
SELECT `type`, `id`, `fullname`, `firstname`, `lastname`, `dateOfBirth`, `addresses`, `contactPoints`, `preferences`, `taxInformation`, `URI`, `relation` FROM `Relationship` WHERE 1;

--
-- INSERT template for table `Relationship`
--
INSERT INTO `Relationship`(`type`, `id`, `fullname`, `firstname`, `lastname`, `dateOfBirth`, `addresses`, `contactPoints`, `preferences`, `taxInformation`, `URI`, `relation`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Relationship`
--
UPDATE `Relationship` SET `type` = ?, `id` = ?, `fullname` = ?, `firstname` = ?, `lastname` = ?, `dateOfBirth` = ?, `addresses` = ?, `contactPoints` = ?, `preferences` = ?, `taxInformation` = ?, `URI` = ?, `relation` = ? WHERE 1;

--
-- DELETE template for table `Relationship`
--
DELETE FROM `Relationship` WHERE 0;

