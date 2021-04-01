--
-- Cloud API.
-- Prepared SQL queries for 'Bank' definition.
--


--
-- SELECT template for table `Bank`
--
SELECT `swift`, `bankCode`, `bankName`, `branchCode`, `URI` FROM `Bank` WHERE 1;

--
-- INSERT template for table `Bank`
--
INSERT INTO `Bank`(`swift`, `bankCode`, `bankName`, `branchCode`, `URI`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `Bank`
--
UPDATE `Bank` SET `swift` = ?, `bankCode` = ?, `bankName` = ?, `branchCode` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Bank`
--
DELETE FROM `Bank` WHERE 0;

