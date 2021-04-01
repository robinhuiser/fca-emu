--
-- Cloud API.
-- Prepared SQL queries for 'CacheStatisticsList' definition.
--


--
-- SELECT template for table `CacheStatisticsList`
--
SELECT `status`, `messages`, `totalItems`, `nextCursor`, `caches` FROM `CacheStatisticsList` WHERE 1;

--
-- INSERT template for table `CacheStatisticsList`
--
INSERT INTO `CacheStatisticsList`(`status`, `messages`, `totalItems`, `nextCursor`, `caches`) VALUES (?, ?, ?, ?, ?);

--
-- UPDATE template for table `CacheStatisticsList`
--
UPDATE `CacheStatisticsList` SET `status` = ?, `messages` = ?, `totalItems` = ?, `nextCursor` = ?, `caches` = ? WHERE 1;

--
-- DELETE template for table `CacheStatisticsList`
--
DELETE FROM `CacheStatisticsList` WHERE 0;

