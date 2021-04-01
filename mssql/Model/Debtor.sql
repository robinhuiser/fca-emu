--
-- Cloud API.
-- Prepared SQL queries for 'Debtor' definition.
--


--
-- SELECT template for table `Debtor`
--
SELECT `debtorType` FROM `Debtor` WHERE 1;

--
-- INSERT template for table `Debtor`
--
INSERT INTO `Debtor`(`debtorType`) VALUES (?);

--
-- UPDATE template for table `Debtor`
--
UPDATE `Debtor` SET `debtorType` = ? WHERE 1;

--
-- DELETE template for table `Debtor`
--
DELETE FROM `Debtor` WHERE 0;

