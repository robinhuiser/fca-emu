--
-- Cloud API.
-- Prepared SQL queries for 'EntityProfile' definition.
--


--
-- SELECT template for table `EntityProfile`
--
SELECT `type`, `id`, `fullname`, `firstname`, `lastname`, `dateOfBirth`, `addresses`, `contactPoints`, `preferences`, `taxInformation`, `URI` FROM `EntityProfile` WHERE 1;

--
-- INSERT template for table `EntityProfile`
--
INSERT INTO `EntityProfile`(`type`, `id`, `fullname`, `firstname`, `lastname`, `dateOfBirth`, `addresses`, `contactPoints`, `preferences`, `taxInformation`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `EntityProfile`
--
UPDATE `EntityProfile` SET `type` = ?, `id` = ?, `fullname` = ?, `firstname` = ?, `lastname` = ?, `dateOfBirth` = ?, `addresses` = ?, `contactPoints` = ?, `preferences` = ?, `taxInformation` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `EntityProfile`
--
DELETE FROM `EntityProfile` WHERE 0;

