--
-- Cloud API.
-- Prepared SQL queries for 'Transaction' definition.
--


--
-- SELECT template for table `Transaction`
--
SELECT `entityId`, `accountId`, `id`, `sequenceInDay`, `status`, `executedAmount`, `executedCurrencyCode`, `exchangeRate`, `originatingAmount`, `originatingCurrencyCode`, `direction`, `runningBalance`, `createdDate`, `postedDate`, `executedDate`, `updatedDate`, `description`, `memo`, `group`, `type`, `mainCategory`, `subCategory`, `checkNumber`, `images`, `latitude`, `longitude`, `merchantCode`, `reversal`, `reversalFor`, `reversed`, `reversedBy`, `transactor`, `URI` FROM `Transaction` WHERE 1;

--
-- INSERT template for table `Transaction`
--
INSERT INTO `Transaction`(`entityId`, `accountId`, `id`, `sequenceInDay`, `status`, `executedAmount`, `executedCurrencyCode`, `exchangeRate`, `originatingAmount`, `originatingCurrencyCode`, `direction`, `runningBalance`, `createdDate`, `postedDate`, `executedDate`, `updatedDate`, `description`, `memo`, `group`, `type`, `mainCategory`, `subCategory`, `checkNumber`, `images`, `latitude`, `longitude`, `merchantCode`, `reversal`, `reversalFor`, `reversed`, `reversedBy`, `transactor`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Transaction`
--
UPDATE `Transaction` SET `entityId` = ?, `accountId` = ?, `id` = ?, `sequenceInDay` = ?, `status` = ?, `executedAmount` = ?, `executedCurrencyCode` = ?, `exchangeRate` = ?, `originatingAmount` = ?, `originatingCurrencyCode` = ?, `direction` = ?, `runningBalance` = ?, `createdDate` = ?, `postedDate` = ?, `executedDate` = ?, `updatedDate` = ?, `description` = ?, `memo` = ?, `group` = ?, `type` = ?, `mainCategory` = ?, `subCategory` = ?, `checkNumber` = ?, `images` = ?, `latitude` = ?, `longitude` = ?, `merchantCode` = ?, `reversal` = ?, `reversalFor` = ?, `reversed` = ?, `reversedBy` = ?, `transactor` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Transaction`
--
DELETE FROM `Transaction` WHERE 0;

