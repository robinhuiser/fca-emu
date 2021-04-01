--
-- Cloud API.
-- Prepared SQL queries for 'BaseListProperties' definition.
--


--
-- SELECT template for table `BaseListProperties`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor` FROM `BaseListProperties` WHERE 1;

--
-- INSERT template for table `BaseListProperties`
--
INSERT INTO `BaseListProperties`(`status`, `messages`, `totalItems`, `nextCursor`) VALUES (?, ?, ?, ?);

--
-- UPDATE template for table `BaseListProperties`
--
UPDATE `BaseListProperties` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ? WHERE 1;

--
-- DELETE template for table `BaseListProperties`
--
DELETE FROM `BaseListProperties` WHERE 0;

