--
-- Cloud API.
-- Prepared SQL queries for 'AccountCreditor' definition.
--


--
-- SELECT template for table `AccountCreditor`
--
SELECT `creditorType`, `target`, `id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `creditInformation`, `loanInterest`, `creditInterest`, `preferences`, `bank`, `owners`, `product`, `URI` FROM `AccountCreditor` WHERE 1;

--
-- INSERT template for table `AccountCreditor`
--
INSERT INTO `AccountCreditor`(`creditorType`, `target`, `id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `creditInformation`, `loanInterest`, `creditInterest`, `preferences`, `bank`, `owners`, `product`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `AccountCreditor`
--
UPDATE `AccountCreditor` SET `creditorType` = ?, `target` = ?, `id` = ?, `type` = ?, `number` = ?, `parentId` = ?, `entities` = ?, `name` = ?, `title` = ?, `iban` = ?, `routingnumbers` = ?, `balances` = ?, `dateCreated` = ?, `dateOpened` = ?, `dateLastUpdated` = ?, `dateClosed` = ?, `currencyCode` = ?, `status` = ?, `source` = ?, `interestReporting` = ?, `creditInformation` = ?, `loanInterest` = ?, `creditInterest` = ?, `preferences` = ?, `bank` = ?, `owners` = ?, `product` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `AccountCreditor`
--
DELETE FROM `AccountCreditor` WHERE 0;

