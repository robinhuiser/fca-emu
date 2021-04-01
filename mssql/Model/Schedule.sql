--
-- Cloud API.
-- Prepared SQL queries for 'Schedule' definition.
--


--
-- SELECT template for table `Schedule`
--
SELECT `strategy`, `frequency`, `dayOn`, `startDateTime`, `endDateTime`, `nextDateTime`, `repeatCount`, `isEveryTime` FROM `Schedule` WHERE 1;

--
-- INSERT template for table `Schedule`
--
INSERT INTO `Schedule`(`strategy`, `frequency`, `dayOn`, `startDateTime`, `endDateTime`, `nextDateTime`, `repeatCount`, `isEveryTime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Schedule`
--
UPDATE `Schedule` SET `strategy` = ?, `frequency` = ?, `dayOn` = ?, `startDateTime` = ?, `endDateTime` = ?, `nextDateTime` = ?, `repeatCount` = ?, `isEveryTime` = ? WHERE 1;

--
-- DELETE template for table `Schedule`
--
DELETE FROM `Schedule` WHERE 0;

