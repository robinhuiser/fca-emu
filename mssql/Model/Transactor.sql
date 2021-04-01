--
-- Cloud API.
-- Prepared SQL queries for 'Transactor' definition.
--


--
-- SELECT template for table `Transactor`
--
SELECT `type`, `id`, `fullname`, `firstname`, `lastname`, `dateOfBirth`, `addresses`, `contactPoints`, `preferences`, `taxInformation`, `URI`, `swift`, `bankCode`, `bankName`, `branchCode`, `accountNumber` FROM `Transactor` WHERE 1;

--
-- INSERT template for table `Transactor`
--
INSERT INTO `Transactor`(`type`, `id`, `fullname`, `firstname`, `lastname`, `dateOfBirth`, `addresses`, `contactPoints`, `preferences`, `taxInformation`, `URI`, `swift`, `bankCode`, `bankName`, `branchCode`, `accountNumber`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Transactor`
--
UPDATE `Transactor` SET `type` = ?, `id` = ?, `fullname` = ?, `firstname` = ?, `lastname` = ?, `dateOfBirth` = ?, `addresses` = ?, `contactPoints` = ?, `preferences` = ?, `taxInformation` = ?, `URI` = ?, `swift` = ?, `bankCode` = ?, `bankName` = ?, `branchCode` = ?, `accountNumber` = ? WHERE 1;

--
-- DELETE template for table `Transactor`
--
DELETE FROM `Transactor` WHERE 0;

