--
-- Cloud API.
-- Prepared SQL queries for 'SecurityInformation' definition.
--


--
-- SELECT template for table `SecurityInformation`
--
SELECT `username`, `token`, `lastLoginDate` FROM `SecurityInformation` WHERE 1;

--
-- INSERT template for table `SecurityInformation`
--
INSERT INTO `SecurityInformation`(`username`, `token`, `lastLoginDate`) VALUES (?, ?, ?);

--
-- UPDATE template for table `SecurityInformation`
--
UPDATE `SecurityInformation` SET `username` = ?, `token` = ?, `lastLoginDate` = ? WHERE 1;

--
-- DELETE template for table `SecurityInformation`
--
DELETE FROM `SecurityInformation` WHERE 0;

