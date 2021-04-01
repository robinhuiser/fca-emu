--
-- Cloud API.
-- Prepared SQL queries for 'CardsList' definition.
--


--
-- SELECT template for table `CardsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `cards` FROM `CardsList` WHERE 1;

--
-- INSERT template for table `CardsList`
--
INSERT INTO `CardsList`(`status`, `messages`, `totalItems`, `nextCursor`, `cards`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `CardsList`
--
UPDATE `CardsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `cards` = ? WHERE 1;

--
-- DELETE template for table `CardsList`
--
DELETE FROM `CardsList` WHERE 0;

