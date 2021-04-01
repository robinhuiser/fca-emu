--
-- Cloud API.
-- Prepared SQL queries for 'ErrorResponse' definition.
--


--
-- SELECT template for table `ErrorResponse`
--
SELECT `status`, `messages` FROM `ErrorResponse` WHERE 1;

--
-- INSERT template for table `ErrorResponse`
--
INSERT INTO `ErrorResponse`(`status`, `messages`) VALUES (?, ?);

--
-- UPDATE template for table `ErrorResponse`
--
UPDATE `ErrorResponse` SET `status` = ?, `messages` = ? WHERE 1;

--
-- DELETE template for table `ErrorResponse`
--
DELETE FROM `ErrorResponse` WHERE 0;

