--
-- Cloud API.
-- Prepared SQL queries for 'Fee' definition.
--


--
-- SELECT template for table `Fee`
--
SELECT `amount`, `currencyCode`, `owernship` FROM `Fee` WHERE 1;

--
-- INSERT template for table `Fee`
--
INSERT INTO `Fee`(`amount`, `currencyCode`, `owernship`) VALUES (?, ?, ?);

--
-- UPDATE template for table `Fee`
--
UPDATE `Fee` SET `amount` = ?, `currencyCode` = ?, `owernship` = ? WHERE 1;

--
-- DELETE template for table `Fee`
--
DELETE FROM `Fee` WHERE 0;

