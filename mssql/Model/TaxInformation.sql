--
-- Cloud API.
-- Prepared SQL queries for 'TaxInformation' definition.
--


--
-- SELECT template for table `TaxInformation`
--
SELECT `type`, `taxId` FROM `TaxInformation` WHERE 1;

--
-- INSERT template for table `TaxInformation`
--
INSERT INTO `TaxInformation`(`type`, `taxId`) VALUES (?, ?);

--
-- UPDATE template for table `TaxInformation`
--
UPDATE `TaxInformation` SET `type` = ?, `taxId` = ? WHERE 1;

--
-- DELETE template for table `TaxInformation`
--
DELETE FROM `TaxInformation` WHERE 0;

