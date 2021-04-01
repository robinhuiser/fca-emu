--
-- Cloud API.
-- Prepared SQL queries for 'BinaryItemList' definition.
--


--
-- SELECT template for table `BinaryItemList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `binaries` FROM `BinaryItemList` WHERE 1;

--
-- INSERT template for table `BinaryItemList`
--
INSERT INTO `BinaryItemList`(`status`, `messages`, `totalItems`, `nextCursor`, `binaries`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `BinaryItemList`
--
UPDATE `BinaryItemList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `binaries` = ? WHERE 1;

--
-- DELETE template for table `BinaryItemList`
--
DELETE FROM `BinaryItemList` WHERE 0;

