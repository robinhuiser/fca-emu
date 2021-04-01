--
-- Cloud API.
-- Prepared SQL queries for 'CreditInformation' definition.
--


--
-- SELECT template for table `CreditInformation`
--
SELECT `creditLimit`, `creditAvailable`, `creditLimitExpiryDate`, `loanTermsUnit`, `loanTerms`, `principalDueAmount`, `interestDueAmount`, `paymentDueAmount`, `paymentDueDate`, `paymentDueDay`, `lastPaymentAmount`, `lastPaymentDate`, `isRevolving`, `maturityDate`, `maturityAmount`, `initialAmount`, `interestEarnedYTD`, `interestPaidYTD`, `interestEarnedCTD`, `interestPaidCTD`, `interestEarnedPriorYear`, `interestPaidPriorYear` FROM `CreditInformation` WHERE 1;

--
-- INSERT template for table `CreditInformation`
--
INSERT INTO `CreditInformation`(`creditLimit`, `creditAvailable`, `creditLimitExpiryDate`, `loanTermsUnit`, `loanTerms`, `principalDueAmount`, `interestDueAmount`, `paymentDueAmount`, `paymentDueDate`, `paymentDueDay`, `lastPaymentAmount`, `lastPaymentDate`, `isRevolving`, `maturityDate`, `maturityAmount`, `initialAmount`, `interestEarnedYTD`, `interestPaidYTD`, `interestEarnedCTD`, `interestPaidCTD`, `interestEarnedPriorYear`, `interestPaidPriorYear`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

--
-- UPDATE template for table `CreditInformation`
--
UPDATE `CreditInformation` SET `creditLimit` = ?, `creditAvailable` = ?, `creditLimitExpiryDate` = ?, `loanTermsUnit` = ?, `loanTerms` = ?, `principalDueAmount` = ?, `interestDueAmount` = ?, `paymentDueAmount` = ?, `paymentDueDate` = ?, `paymentDueDay` = ?, `lastPaymentAmount` = ?, `lastPaymentDate` = ?, `isRevolving` = ?, `maturityDate` = ?, `maturityAmount` = ?, `initialAmount` = ?, `interestEarnedYTD` = ?, `interestPaidYTD` = ?, `interestEarnedCTD` = ?, `interestPaidCTD` = ?, `interestEarnedPriorYear` = ?, `interestPaidPriorYear` = ? WHERE 1;

--
-- DELETE template for table `CreditInformation`
--
DELETE FROM `CreditInformation` WHERE 0;

