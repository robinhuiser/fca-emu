--
-- Cloud API.
-- Prepared SQL queries for 'ContactPoint' definition.
--


--
-- SELECT template for table `ContactPoint`
--
SELECT `type`, `name`, `prefix`, `value`, `suffix` FROM `ContactPoint` WHERE 1;

--
-- INSERT template for table `ContactPoint`
--
INSERT INTO `ContactPoint`(`type`, `name`, `prefix`, `value`, `suffix`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `ContactPoint`
--
UPDATE `ContactPoint` SET `type` = ?, `name` = ?, `prefix` = ?, `value` = ?, `suffix` = ? WHERE 1;

--
-- DELETE template for table `ContactPoint`
--
DELETE FROM `ContactPoint` WHERE 0;

