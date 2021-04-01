--
-- Cloud API.
-- Prepared SQL queries for 'SearchFilter' definition.
--


--
-- SELECT template for table `SearchFilter`
--
SELECT `field`, `query`, `operator` FROM `SearchFilter` WHERE 1;

--
-- INSERT template for table `SearchFilter`
--
INSERT INTO `SearchFilter`(`field`, `query`, `operator`) VALUES (?, ?, ?);

--
-- UPDATE template for table `SearchFilter`
--
UPDATE `SearchFilter` SET `field` = ?, `query` = ?, `operator` = ? WHERE 1;

--
-- DELETE template for table `SearchFilter`
--
DELETE FROM `SearchFilter` WHERE 0;

