--
-- Cloud API.
-- Prepared SQL queries for 'IntermediateInstitution' definition.
--


--
-- SELECT template for table `IntermediateInstitution`
--
SELECT `swift`, `bankCode`, `bankName`, `branchCode`, `URI`, `message` FROM `IntermediateInstitution` WHERE 1;

--
-- INSERT template for table `IntermediateInstitution`
--
INSERT INTO `IntermediateInstitution`(`swift`, `bankCode`, `bankName`, `branchCode`, `URI`, `message`) VALUES (?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `IntermediateInstitution`
--
UPDATE `IntermediateInstitution` SET `swift` = ?, `bankCode` = ?, `bankName` = ?, `branchCode` = ?, `URI` = ?, `message` = ? WHERE 1;

--
-- DELETE template for table `IntermediateInstitution`
--
DELETE FROM `IntermediateInstitution` WHERE 0;

