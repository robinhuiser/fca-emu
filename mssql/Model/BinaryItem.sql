--
-- Cloud API.
-- Prepared SQL queries for 'BinaryItem' definition.
--


--
-- SELECT template for table `BinaryItem`
--
SELECT `format`, `length`, `attributes`, `itemId`, `content`, `URI` FROM `BinaryItem` WHERE 1;

--
-- INSERT template for table `BinaryItem`
--
INSERT INTO `BinaryItem`(`format`, `length`, `attributes`, `itemId`, `content`, `URI`) VALUES (?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `BinaryItem`
--
UPDATE `BinaryItem` SET `format` = ?, `length` = ?, `attributes` = ?, `itemId` = ?, `content` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `BinaryItem`
--
DELETE FROM `BinaryItem` WHERE 0;

