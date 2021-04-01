--
-- Cloud API.
-- Prepared SQL queries for 'AccountDebtor' definition.
--


--
-- SELECT template for table `AccountDebtor`
--
SELECT `debtorType`, `id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `creditInformation`, `loanInterest`, `creditInterest`, `preferences`, `bank`, `owners`, `product`, `URI` FROM `AccountDebtor` WHERE 1;

--
-- INSERT template for table `AccountDebtor`
--
INSERT INTO `AccountDebtor`(`debtorType`, `id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `creditInformation`, `loanInterest`, `creditInterest`, `preferences`, `bank`, `owners`, `product`, `URI`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `AccountDebtor`
--
UPDATE `AccountDebtor` SET `debtorType` = ?, `id` = ?, `type` = ?, `number` = ?, `parentId` = ?, `entities` = ?, `name` = ?, `title` = ?, `iban` = ?, `routingnumbers` = ?, `balances` = ?, `dateCreated` = ?, `dateOpened` = ?, `dateLastUpdated` = ?, `dateClosed` = ?, `currencyCode` = ?, `status` = ?, `source` = ?, `interestReporting` = ?, `creditInformation` = ?, `loanInterest` = ?, `creditInterest` = ?, `preferences` = ?, `bank` = ?, `owners` = ?, `product` = ?, `URI` = ? WHERE 1;

--
-- DELETE template for table `AccountDebtor`
--
DELETE FROM `AccountDebtor` WHERE 0;

