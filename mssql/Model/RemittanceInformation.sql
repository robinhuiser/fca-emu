--
-- Cloud API.
-- Prepared SQL queries for 'RemittanceInformation' definition.
--


--
-- SELECT template for table `RemittanceInformation`
--
SELECT `type`, `issuer`, `reference` FROM `RemittanceInformation` WHERE 1;

--
-- INSERT template for table `RemittanceInformation`
--
INSERT INTO `RemittanceInformation`(`type`, `issuer`, `reference`) VALUES (?, ?, ?);

--
-- UPDATE template for table `RemittanceInformation`
--
UPDATE `RemittanceInformation` SET `type` = ?, `issuer` = ?, `reference` = ? WHERE 1;

--
-- DELETE template for table `RemittanceInformation`
--
DELETE FROM `RemittanceInformation` WHERE 0;

