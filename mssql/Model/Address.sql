--
-- Cloud API.
-- Prepared SQL queries for 'Address' definition.
--


--
-- SELECT template for table `Address`
--
SELECT `type`, `line1`, `line2`, `line3`, `city`, `state`, `postalCode`, `country`, `primary` FROM `Address` WHERE 1;

--
-- INSERT template for table `Address`
--
INSERT INTO `Address`(`type`, `line1`, `line2`, `line3`, `city`, `state`, `postalCode`, `country`, `primary`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Address`
--
UPDATE `Address` SET `type` = ?, `line1` = ?, `line2` = ?, `line3` = ?, `city` = ?, `state` = ?, `postalCode` = ?, `country` = ?, `primary` = ? WHERE 1;

--
-- DELETE template for table `Address`
--
DELETE FROM `Address` WHERE 0;

