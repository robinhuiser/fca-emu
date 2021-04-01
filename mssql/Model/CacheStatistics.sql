--
-- Cloud API.
-- Prepared SQL queries for 'CacheStatistics' definition.
--


--
-- SELECT template for table `CacheStatistics`
--
SELECT `finiteType`, `finiteReference`, `cacheSize`, `cacheCount` FROM `CacheStatistics` WHERE 1;

--
-- INSERT template for table `CacheStatistics`
--
INSERT INTO `CacheStatistics`(`finiteType`, `finiteReference`, `cacheSize`, `cacheCount`) VALUES (?, ?, ?, ?);

--
-- UPDATE template for table `CacheStatistics`
--
UPDATE `CacheStatistics` SET `finiteType` = ?, `finiteReference` = ?, `cacheSize` = ?, `cacheCount` = ? WHERE 1;

--
-- DELETE template for table `CacheStatistics`
--
DELETE FROM `CacheStatistics` WHERE 0;

