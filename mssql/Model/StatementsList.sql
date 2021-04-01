--
-- Cloud API.
-- Prepared SQL queries for 'StatementsList' definition.
--


--
-- SELECT template for table `StatementsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `statements` FROM `StatementsList` WHERE 1;

--
-- INSERT template for table `StatementsList`
--
INSERT INTO `StatementsList`(`status`, `messages`, `totalItems`, `nextCursor`, `statements`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `StatementsList`
--
UPDATE `StatementsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `statements` = ? WHERE 1;

--
-- DELETE template for table `StatementsList`
--
DELETE FROM `StatementsList` WHERE 0;

