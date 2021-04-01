--
-- Cloud API.
-- Prepared SQL queries for 'ProductsList' definition.
--


--
-- SELECT template for table `ProductsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `products` FROM `ProductsList` WHERE 1;

--
-- INSERT template for table `ProductsList`
--
INSERT INTO `ProductsList`(`status`, `messages`, `totalItems`, `nextCursor`, `products`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `ProductsList`
--
UPDATE `ProductsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `products` = ? WHERE 1;

--
-- DELETE template for table `ProductsList`
--
DELETE FROM `ProductsList` WHERE 0;

