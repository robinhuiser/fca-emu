--
-- Cloud API.
-- Prepared SQL queries for 'Creditor' definition.
--


--
-- SELECT template for table `Creditor`
--
SELECT `creditorType`, `target` FROM `Creditor` WHERE 1;

--
-- INSERT template for table `Creditor`
--
INSERT INTO `Creditor`(`creditorType`, `target`) VALUES (?, ?);

--
-- UPDATE template for table `Creditor`
--
UPDATE `Creditor` SET `creditorType` = ?, `target` = ? WHERE 1;

--
-- DELETE template for table `Creditor`
--
DELETE FROM `Creditor` WHERE 0;

