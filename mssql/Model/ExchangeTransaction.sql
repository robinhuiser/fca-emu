--
-- Cloud API.
-- Prepared SQL queries for 'ExchangeTransaction' definition.
--


--
-- SELECT template for table `ExchangeTransaction`
--
SELECT `id`, `type`, `debtor`, `creditor`, `intermediateInstitutions`, `remittanceInformation`, `isBatch`, `priority`, `executionDate`, `isRecurring`, `recurringSchedule`, `amount`, `fee`, `URI` FROM `ExchangeTransaction` WHERE 1;

--
-- INSERT template for table `ExchangeTransaction`
--
INSERT INTO `ExchangeTransaction`(`id`, `type`, `debtor`, `creditor`, `intermediateInstitutions`, `remittanceInformation`, `isBatch`, `priority`, `executionDate`, `isRecurring`, `recurringSchedule`, `amount`, `fee`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `ExchangeTransaction`
--
UPDATE `ExchangeTransaction` SET `id` = ?, `type` = ?, `debtor` = ?, `creditor` = ?, `intermediateInstitutions` = ?, `remittanceInformation` = ?, `isBatch` = ?, `priority` = ?, `executionDate` = ?, `isRecurring` = ?, `recurringSchedule` = ?, `amount` = ?, `fee` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `ExchangeTransaction`
--
DELETE FROM `ExchangeTransaction` WHERE 0;

