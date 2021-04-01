--
-- Cloud API.
-- Prepared SQL queries for 'BaseErrorProperties' definition.
--


--
-- SELECT template for table `BaseErrorProperties`
--
SELECT `status`, `messages` FROM `BaseErrorProperties` WHERE 1;

--
-- INSERT template for table `BaseErrorProperties`
--
INSERT INTO `BaseErrorProperties`(`status`, `messages`) VALUES (?, ?);

--
-- UPDATE template for table `BaseErrorProperties`
--
UPDATE `BaseErrorProperties` SET `status` = ?, `messages` = ? WHERE 1;

--
-- DELETE template for table `BaseErrorProperties`
--
DELETE FROM `BaseErrorProperties` WHERE 0;

