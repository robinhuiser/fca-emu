--
-- Cloud API.
-- Prepared SQL queries for 'Attribute' definition.
--


--
-- SELECT template for table `Attribute`
--
SELECT `name`, `value` FROM `Attribute` WHERE 1;

--
-- INSERT template for table `Attribute`
--
INSERT INTO `Attribute`(`name`, `value`) VALUES (?, ?);

--
-- UPDATE template for table `Attribute`
--
UPDATE `Attribute` SET `name` = ?, `value` = ? WHERE 1;

--
-- DELETE template for table `Attribute`
--
DELETE FROM `Attribute` WHERE 0;

