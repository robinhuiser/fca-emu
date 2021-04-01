--
-- Cloud API.
-- Prepared SQL queries for 'EntityList' definition.
--


--
-- SELECT template for table `EntityList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `entities` FROM `EntityList` WHERE 1;

--
-- INSERT template for table `EntityList`
--
INSERT INTO `EntityList`(`status`, `messages`, `totalItems`, `nextCursor`, `entities`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `EntityList`
--
UPDATE `EntityList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `entities` = ? WHERE 1;

--
-- DELETE template for table `EntityList`
--
DELETE FROM `EntityList` WHERE 0;

