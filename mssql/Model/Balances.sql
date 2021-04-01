--
-- Cloud API.
-- Prepared SQL queries for 'Balances' definition.
--


--
-- SELECT template for table `Balances`
--
SELECT `availableBalance`, `currentBalance`, `cashBalance`, `minimumBalance`, `outstandingBalance`, `principalBalance`, `accruedInterest`, `investmentBalance`, `payOffAmount`, `ledgerBalance`, `URI` FROM `Balances` WHERE 1;

--
-- INSERT template for table `Balances`
--
INSERT INTO `Balances`(`availableBalance`, `currentBalance`, `cashBalance`, `minimumBalance`, `outstandingBalance`, `principalBalance`, `accruedInterest`, `investmentBalance`, `payOffAmount`, `ledgerBalance`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Balances`
--
UPDATE `Balances` SET `availableBalance` = ?, `currentBalance` = ?, `cashBalance` = ?, `minimumBalance` = ?, `outstandingBalance` = ?, `principalBalance` = ?, `accruedInterest` = ?, `investmentBalance` = ?, `payOffAmount` = ?, `ledgerBalance` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Balances`
--
DELETE FROM `Balances` WHERE 0;

