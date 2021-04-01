--
-- Cloud API.
-- Prepared SQL queries for 'Interest' definition.
--


--
-- SELECT template for table `Interest`
--
SELECT `rate`, `termsUnit`, `terms` FROM `Interest` WHERE 1;

--
-- INSERT template for table `Interest`
--
INSERT INTO `Interest`(`rate`, `termsUnit`, `terms`) VALUES (?, ?, ?);

--
-- UPDATE template for table `Interest`
--
UPDATE `Interest` SET `rate` = ?, `termsUnit` = ?, `terms` = ? WHERE 1;

--
-- DELETE template for table `Interest`
--
DELETE FROM `Interest` WHERE 0;

