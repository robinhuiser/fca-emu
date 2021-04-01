--
-- Cloud API.
-- Prepared SQL queries for 'Account' definition.
--


--
-- SELECT template for table `Account`
--
SELECT `id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `creditInformation`, `loanInterest`, `creditInterest`, `preferences`, `bank`, `owners`, `product`, `URI` FROM `Account` WHERE 1;

--
-- INSERT template for table `Account`
--
INSERT INTO `Account`(`id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `creditInformation`, `loanInterest`, `creditInterest`, `preferences`, `bank`, `owners`, `product`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `Account`
--
UPDATE `Account` SET `id` = ?, `type` = ?, `number` = ?, `parentId` = ?, `entities` = ?, `name` = ?, `title` = ?, `iban` = ?, `routingnumbers` = ?, `balances` = ?, `dateCreated` = ?, `dateOpened` = ?, `dateLastUpdated` = ?, `dateClosed` = ?, `currencyCode` = ?, `status` = ?, `source` = ?, `interestReporting` = ?, `creditInformation` = ?, `loanInterest` = ?, `creditInterest` = ?, `preferences` = ?, `bank` = ?, `owners` = ?, `product` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `Account`
--
DELETE FROM `Account` WHERE 0;

