/* SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO"; */
/* SET AUTOCOMMIT = 0; */
/* START TRANSACTION; */
/* SET time_zone = "+00:00"; */

-- --------------------------------------------------------

--
-- Table structure for table `Account` generated from model 'Account'
-- Account object
--

CREATE TABLE IF NOT EXISTS `Account` (
  `id` TEXT NOT NULL COMMENT 'Unique composite identifier of the account in the system of record',
  `type` TEXT DEFAULT NULL COMMENT 'The type of account as enumerated by the system of record',
  `number` TEXT DEFAULT NULL COMMENT 'The account number as represented in the system of record (masked if mask&#x3D;true)',
  `parentId` TEXT DEFAULT NULL COMMENT 'Unique identifier of the parent account, in the case of a sub account',
  `entities` JSON DEFAULT NULL COMMENT 'List of entities this account is associated with',
  `name` TEXT DEFAULT NULL COMMENT 'The name of the primary account holder',
  `title` TEXT DEFAULT NULL COMMENT 'The given title of an account',
  `iban` TEXT DEFAULT NULL COMMENT 'The international bank account number',
  `routingnumbers` JSON DEFAULT NULL COMMENT 'The routing numbers of this account',
  `balances` TEXT DEFAULT NULL,
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was created',
  `dateOpened` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was opened or activated',
  `dateLastUpdated` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was last updated',
  `dateClosed` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was closed or deactivated',
  `currencyCode` TEXT DEFAULT NULL COMMENT 'The currency code of the account',
  `status` TEXT DEFAULT NULL COMMENT 'The current status of the account',
  `source` TEXT DEFAULT NULL COMMENT 'The source of the account, where the account is located',
  `interestReporting` TINYINT(1) DEFAULT NULL COMMENT 'Indicates if interest is reportable on this account.',
  `creditInformation` TEXT DEFAULT NULL,
  `loanInterest` TEXT DEFAULT NULL,
  `creditInterest` TEXT DEFAULT NULL,
  `preferences` JSON DEFAULT NULL COMMENT 'Unique preferences of the account',
  `bank` TEXT DEFAULT NULL,
  `owners` JSON DEFAULT NULL COMMENT 'List of owners of the account',
  `product` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Account object';

--
-- Table structure for table `AccountCreditor` generated from model 'AccountCreditor'
-- An Account That is Credited
--

CREATE TABLE IF NOT EXISTS `AccountCreditor` (
  `creditorType` TEXT NOT NULL,
  `target` ENUM('PAYMENT', 'PRINCIPAL', 'INTEREST') DEFAULT NULL COMMENT 'When crediting a loan-style account, where to target incoming funds.',
  `id` TEXT NOT NULL COMMENT 'Unique composite identifier of the account in the system of record',
  `type` TEXT DEFAULT NULL COMMENT 'The type of account as enumerated by the system of record',
  `number` TEXT DEFAULT NULL COMMENT 'The account number as represented in the system of record (masked if mask&#x3D;true)',
  `parentId` TEXT DEFAULT NULL COMMENT 'Unique identifier of the parent account, in the case of a sub account',
  `entities` JSON DEFAULT NULL COMMENT 'List of entities this account is associated with',
  `name` TEXT DEFAULT NULL COMMENT 'The name of the primary account holder',
  `title` TEXT DEFAULT NULL COMMENT 'The given title of an account',
  `iban` TEXT DEFAULT NULL COMMENT 'The international bank account number',
  `routingnumbers` JSON DEFAULT NULL COMMENT 'The routing numbers of this account',
  `balances` TEXT DEFAULT NULL,
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was created',
  `dateOpened` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was opened or activated',
  `dateLastUpdated` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was last updated',
  `dateClosed` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was closed or deactivated',
  `currencyCode` TEXT DEFAULT NULL COMMENT 'The currency code of the account',
  `status` TEXT DEFAULT NULL COMMENT 'The current status of the account',
  `source` TEXT DEFAULT NULL COMMENT 'The source of the account, where the account is located',
  `interestReporting` TINYINT(1) DEFAULT NULL COMMENT 'Indicates if interest is reportable on this account.',
  `creditInformation` TEXT DEFAULT NULL,
  `loanInterest` TEXT DEFAULT NULL,
  `creditInterest` TEXT DEFAULT NULL,
  `preferences` JSON DEFAULT NULL COMMENT 'Unique preferences of the account',
  `bank` TEXT DEFAULT NULL,
  `owners` JSON DEFAULT NULL COMMENT 'List of owners of the account',
  `product` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='An Account That is Credited';

--
-- Table structure for table `AccountDebtor` generated from model 'AccountDebtor'
-- An Account That is Debited
--

CREATE TABLE IF NOT EXISTS `AccountDebtor` (
  `debtorType` TEXT NOT NULL,
  `id` TEXT NOT NULL COMMENT 'Unique composite identifier of the account in the system of record',
  `type` TEXT DEFAULT NULL COMMENT 'The type of account as enumerated by the system of record',
  `number` TEXT DEFAULT NULL COMMENT 'The account number as represented in the system of record (masked if mask&#x3D;true)',
  `parentId` TEXT DEFAULT NULL COMMENT 'Unique identifier of the parent account, in the case of a sub account',
  `entities` JSON DEFAULT NULL COMMENT 'List of entities this account is associated with',
  `name` TEXT DEFAULT NULL COMMENT 'The name of the primary account holder',
  `title` TEXT DEFAULT NULL COMMENT 'The given title of an account',
  `iban` TEXT DEFAULT NULL COMMENT 'The international bank account number',
  `routingnumbers` JSON DEFAULT NULL COMMENT 'The routing numbers of this account',
  `balances` TEXT DEFAULT NULL,
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was created',
  `dateOpened` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was opened or activated',
  `dateLastUpdated` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was last updated',
  `dateClosed` TEXT DEFAULT NULL COMMENT 'ISO 6801 date when the account was closed or deactivated',
  `currencyCode` TEXT DEFAULT NULL COMMENT 'The currency code of the account',
  `status` TEXT DEFAULT NULL COMMENT 'The current status of the account',
  `source` TEXT DEFAULT NULL COMMENT 'The source of the account, where the account is located',
  `interestReporting` TINYINT(1) DEFAULT NULL COMMENT 'Indicates if interest is reportable on this account.',
  `creditInformation` TEXT DEFAULT NULL,
  `loanInterest` TEXT DEFAULT NULL,
  `creditInterest` TEXT DEFAULT NULL,
  `preferences` JSON DEFAULT NULL COMMENT 'Unique preferences of the account',
  `bank` TEXT DEFAULT NULL,
  `owners` JSON DEFAULT NULL COMMENT 'List of owners of the account',
  `product` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='An Account That is Debited';

--
-- Table structure for table `AccountsList` generated from model 'AccountsList'
--

CREATE TABLE IF NOT EXISTS `AccountsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `accounts` JSON DEFAULT NULL COMMENT 'The list of accounts'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `AccountsList_allOf` generated from model 'AccountsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `AccountsList_allOf` (
  `accounts` JSON DEFAULT NULL COMMENT 'The list of accounts'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Address` generated from model 'Address'
-- Address object
--

CREATE TABLE IF NOT EXISTS `Address` (
  `type` ENUM('BUSINESS', 'RESIDENTIAL', 'POBOX') DEFAULT NULL,
  `line1` TEXT DEFAULT NULL COMMENT 'Address line',
  `line2` TEXT DEFAULT NULL COMMENT 'Address line',
  `line3` TEXT DEFAULT NULL COMMENT 'Address line',
  `city` TEXT DEFAULT NULL COMMENT 'Address City or Town',
  `state` TEXT DEFAULT NULL COMMENT 'State or Region',
  `postalCode` TEXT DEFAULT NULL COMMENT 'Postal code or Zip code',
  `country` TEXT DEFAULT NULL COMMENT 'Country',
  `primary` TINYINT(1) DEFAULT NULL COMMENT 'Primary Address'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Address object';

--
-- Table structure for table `Attribute` generated from model 'Attribute'
-- Name value pair property
--

CREATE TABLE IF NOT EXISTS `Attribute` (
  `name` TEXT DEFAULT NULL COMMENT 'Attribute name',
  `value` TEXT DEFAULT NULL COMMENT 'Attribute value'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Name value pair property';

--
-- Table structure for table `Balances` generated from model 'Balances'
-- Account Balances object
--

CREATE TABLE IF NOT EXISTS `Balances` (
  `availableBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The available balance is your current balance less any outstanding holds or debits that have not yet posted to your account.',
  `currentBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The current balance is the total amount of funds in your account.',
  `cashBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of money in an account.  It is calculated by adding the initial deposit to all subsequent deposits and then subtracting all disbursements.',
  `minimumBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The minimum amount of funds a account may have before indicated as overdrawn.',
  `outstandingBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The total amount of the loan including interest/fees. The outstanding balance usually implies that extra interest/fees has been accrued.',
  `principalBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The total loan amount without any added interest/fees',
  `accruedInterest` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The accumulated interest charges that have been recognized in the books of accounts but have yet to be paid',
  `investmentBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amounts credited to the Investment, less the amounts debited from the Investment and the amounts redeemed by such Investor from time to time',
  `payOffAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The total amount needed to pay off the loan.',
  `ledgerBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'A ledger balance is calculated at the end of each business day by a bank and includes all debits and credits.',
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Account Balances object';

--
-- Table structure for table `Bank` generated from model 'Bank'
-- Bank object
--

CREATE TABLE IF NOT EXISTS `Bank` (
  `swift` TEXT DEFAULT NULL COMMENT 'Standard format of Business Identifier Codes (bic), unique identification code',
  `bankCode` TEXT DEFAULT NULL COMMENT 'A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.',
  `bankName` TEXT DEFAULT NULL COMMENT 'The financial institution name',
  `branchCode` TEXT DEFAULT NULL COMMENT 'The branch code of a bank branch helps in distinguishing one branch from another',
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Bank object';

--
-- Table structure for table `BaseErrorProperties` generated from model 'BaseErrorProperties'
-- Base properties for all lists object
--

CREATE TABLE IF NOT EXISTS `BaseErrorProperties` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Base properties for all lists object';

--
-- Table structure for table `BaseListProperties` generated from model 'BaseListProperties'
-- Base properties for all lists object
--

CREATE TABLE IF NOT EXISTS `BaseListProperties` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Base properties for all lists object';

--
-- Table structure for table `BinaryItem` generated from model 'BinaryItem'
-- Binary Item Object
--

CREATE TABLE IF NOT EXISTS `BinaryItem` (
  `format` ENUM('application/pdf', 'image/jpeg', 'image/png', 'image/svg', 'image/gif', 'text/html', 'text/plain') DEFAULT NULL COMMENT 'The mime type of the binary content.',
  `length` INT DEFAULT NULL COMMENT 'The length of the binary content in bytes.',
  `attributes` JSON DEFAULT NULL COMMENT 'List of attributes for a binary item, e.g description, creation time...',
  `itemId` TEXT DEFAULT NULL COMMENT 'A base64 encoded identifier used by the connectors to locate the binary item',
  `content` MEDIUMBLOB DEFAULT NULL COMMENT 'The optionally inlined base64 encoded content of the item.',
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Binary Item Object';

--
-- Table structure for table `BinaryItemList` generated from model 'BinaryItemList'
--

CREATE TABLE IF NOT EXISTS `BinaryItemList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `binaries` JSON DEFAULT NULL COMMENT 'The list of binary items'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `BinaryItemList_allOf` generated from model 'BinaryItemListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `BinaryItemList_allOf` (
  `binaries` JSON DEFAULT NULL COMMENT 'The list of binary items'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `CacheReference` generated from model 'CacheReference'
-- Reference to items in cache used to refresh/evict/renew the cache
--

CREATE TABLE IF NOT EXISTS `CacheReference` (
  `attributes` JSON DEFAULT NULL COMMENT 'List of attributes to locate and reference items in cache'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Reference to items in cache used to refresh/evict/renew the cache';

--
-- Table structure for table `CacheStatistics` generated from model 'CacheStatistics'
-- A summary of the caches for a given finite item type
--

CREATE TABLE IF NOT EXISTS `CacheStatistics` (
  `finiteType` TEXT DEFAULT NULL,
  `finiteReference` JSON DEFAULT NULL COMMENT 'List of attributes identifying the cache referenced',
  `cacheSize` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The memory usage of the cache, in megabytes',
  `cacheCount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The number of items in the cache'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='A summary of the caches for a given finite item type';

--
-- Table structure for table `CacheStatisticsList` generated from model 'CacheStatisticsList'
--

CREATE TABLE IF NOT EXISTS `CacheStatisticsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `caches` JSON DEFAULT NULL COMMENT 'The list of cache statistics'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `CacheStatisticsList_allOf` generated from model 'CacheStatisticsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `CacheStatisticsList_allOf` (
  `caches` JSON DEFAULT NULL COMMENT 'The list of cache statistics'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Card` generated from model 'Card'
-- Card object
--

CREATE TABLE IF NOT EXISTS `Card` (
  `type` ENUM('CREDIT', 'DEBIT', 'LOYALTY') DEFAULT NULL COMMENT 'The type of card',
  `id` TEXT DEFAULT NULL COMMENT 'Unique card identifier in the system of record',
  `number` TEXT DEFAULT NULL COMMENT 'Card number (masked if masked&#x3D;true)',
  `startDate` TEXT DEFAULT NULL COMMENT 'ISO 6801 date of when the card is started/activated',
  `expiryDate` TEXT DEFAULT NULL COMMENT 'ISO 6801 date of when the card expires',
  `holderName` TEXT DEFAULT NULL COMMENT 'The name of the cardholder',
  `network` TEXT DEFAULT NULL COMMENT 'The card network program',
  `status` TEXT DEFAULT NULL COMMENT 'The status of the card',
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Card object';

--
-- Table structure for table `CardsList` generated from model 'CardsList'
--

CREATE TABLE IF NOT EXISTS `CardsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `cards` JSON DEFAULT NULL COMMENT 'The list of cards'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `CardsList_allOf` generated from model 'CardsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `CardsList_allOf` (
  `cards` JSON DEFAULT NULL COMMENT 'The list of cards'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ContactPoint` generated from model 'ContactPoint'
-- Contact Point object
--

CREATE TABLE IF NOT EXISTS `ContactPoint` (
  `type` ENUM('SMS', 'EMAIL', 'VOICE') DEFAULT NULL,
  `name` TEXT DEFAULT NULL COMMENT 'Name of the contact point',
  `prefix` TEXT DEFAULT NULL COMMENT 'Text preceeding contact point',
  `value` TEXT DEFAULT NULL COMMENT 'Contact point value',
  `suffix` TEXT DEFAULT NULL COMMENT 'Additional information'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Contact Point object';

--
-- Table structure for table `CreditInformation` generated from model 'CreditInformation'
-- Account Credit object
--

CREATE TABLE IF NOT EXISTS `CreditInformation` (
  `creditLimit` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The maximum credit limit amount',
  `creditAvailable` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of credit available based on limit',
  `creditLimitExpiryDate` TEXT DEFAULT NULL COMMENT 'ISO 8601 date of when the credit limit expires',
  `loanTermsUnit` TEXT DEFAULT NULL COMMENT 'The measure of terms',
  `loanTerms` INT DEFAULT NULL COMMENT 'The number of terms agreed',
  `principalDueAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of principal within the next due payment',
  `interestDueAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of interest within the next due payment',
  `paymentDueAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of the next due payment',
  `paymentDueDate` TEXT DEFAULT NULL COMMENT 'ISO 8601 date of the next due payment',
  `paymentDueDay` INT DEFAULT NULL COMMENT 'Day of the month on which payments are due',
  `lastPaymentAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of the last payment',
  `lastPaymentDate` TEXT DEFAULT NULL COMMENT 'ISO 8601 date of the last payment',
  `isRevolving` TINYINT(1) DEFAULT NULL COMMENT 'Indicates Revolving Credit Line',
  `maturityDate` TEXT DEFAULT NULL COMMENT 'ISO 8601 formatted maturity date',
  `maturityAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount due at maturity',
  `initialAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The original amount of the credit line or loan.',
  `interestEarnedYTD` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of earning interest, Year To Date.',
  `interestPaidYTD` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of interest paid, Year To Date.',
  `interestEarnedCTD` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of earning interest, Calendar To Date.',
  `interestPaidCTD` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of interest paid, Calendar To Date.',
  `interestEarnedPriorYear` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of earning interest, Prior Year.',
  `interestPaidPriorYear` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of interest paid, Prior Year.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Account Credit object';

--
-- Table structure for table `Creditor` generated from model 'Creditor'
-- Account or Entity that receives an exchange
--

CREATE TABLE IF NOT EXISTS `Creditor` (
  `creditorType` TEXT NOT NULL,
  `target` ENUM('PAYMENT', 'PRINCIPAL', 'INTEREST') DEFAULT NULL COMMENT 'When crediting a loan-style account, where to target incoming funds.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Account or Entity that receives an exchange';

--
-- Table structure for table `Debtor` generated from model 'Debtor'
-- Account or Entity that originates an exchange
--

CREATE TABLE IF NOT EXISTS `Debtor` (
  `debtorType` TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Account or Entity that originates an exchange';

--
-- Table structure for table `Entity` generated from model 'Entity'
-- Entity Object
--

CREATE TABLE IF NOT EXISTS `Entity` (
  `type` ENUM('PERSON', 'BUSINESS', 'SYSTEM') DEFAULT NULL COMMENT 'The type of profile',
  `id` TEXT NOT NULL COMMENT 'Unique identifier of the entity',
  `name` TEXT DEFAULT NULL COMMENT 'Name of the entity',
  `taxInformation` JSON DEFAULT NULL COMMENT 'Entity tax information',
  `active` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether entity is active',
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 8601 Date entity got created',
  `securityInformation` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Entity Object';

--
-- Table structure for table `EntityCreditor` generated from model 'EntityCreditor'
-- An Entity That is Credited
--

CREATE TABLE IF NOT EXISTS `EntityCreditor` (
  `creditorType` TEXT NOT NULL,
  `target` ENUM('PAYMENT', 'PRINCIPAL', 'INTEREST') DEFAULT NULL COMMENT 'When crediting a loan-style account, where to target incoming funds.',
  `type` ENUM('PERSON', 'BUSINESS', 'SYSTEM') DEFAULT NULL COMMENT 'The type of profile',
  `id` TEXT NOT NULL COMMENT 'Unique identifier of the entity',
  `name` TEXT DEFAULT NULL COMMENT 'Name of the entity',
  `taxInformation` JSON DEFAULT NULL COMMENT 'Entity tax information',
  `active` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether entity is active',
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 8601 Date entity got created',
  `securityInformation` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='An Entity That is Credited';

--
-- Table structure for table `EntityDebtor` generated from model 'EntityDebtor'
-- An Entity That is Debited
--

CREATE TABLE IF NOT EXISTS `EntityDebtor` (
  `debtorType` TEXT NOT NULL,
  `type` ENUM('PERSON', 'BUSINESS', 'SYSTEM') DEFAULT NULL COMMENT 'The type of profile',
  `id` TEXT NOT NULL COMMENT 'Unique identifier of the entity',
  `name` TEXT DEFAULT NULL COMMENT 'Name of the entity',
  `taxInformation` JSON DEFAULT NULL COMMENT 'Entity tax information',
  `active` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether entity is active',
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 8601 Date entity got created',
  `securityInformation` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='An Entity That is Debited';

--
-- Table structure for table `EntityList` generated from model 'EntityList'
--

CREATE TABLE IF NOT EXISTS `EntityList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `entities` JSON DEFAULT NULL COMMENT 'The list of entities'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `EntityList_allOf` generated from model 'EntityListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `EntityList_allOf` (
  `entities` JSON DEFAULT NULL COMMENT 'The list of entities'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `EntityProfile` generated from model 'EntityProfile'
-- Entity Profile Information
--

CREATE TABLE IF NOT EXISTS `EntityProfile` (
  `type` ENUM('PERSON', 'BUSINESS', 'SYSTEM') DEFAULT NULL COMMENT 'The type of profile',
  `id` TEXT NOT NULL COMMENT 'Unique identifier of the entity',
  `fullname` TEXT DEFAULT NULL COMMENT 'The fullname of a the entity',
  `firstname` TEXT DEFAULT NULL COMMENT 'The legal firstname of the person',
  `lastname` TEXT DEFAULT NULL COMMENT 'The legal lastname/surname of the person',
  `dateOfBirth` TEXT DEFAULT NULL COMMENT 'ISO 8601 Date of Birth',
  `addresses` JSON DEFAULT NULL COMMENT 'Entity addresses',
  `contactPoints` JSON DEFAULT NULL COMMENT 'Entity contact points',
  `preferences` JSON DEFAULT NULL COMMENT 'Entity preferences',
  `taxInformation` JSON DEFAULT NULL COMMENT 'Entity tax information',
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Entity Profile Information';

--
-- Table structure for table `ErrorResponse` generated from model 'ErrorResponse'
--

CREATE TABLE IF NOT EXISTS `ErrorResponse` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ExchangeTransaction` generated from model 'ExchangeTransaction'
-- Exchange Transaction
--

CREATE TABLE IF NOT EXISTS `ExchangeTransaction` (
  `id` TEXT DEFAULT NULL COMMENT 'Unique identifier for this exchange transaction',
  `type` ENUM('ACH', 'WIRE', 'P2P', 'A2A') DEFAULT NULL COMMENT 'The type of exchange or money movement',
  `debtor` TEXT DEFAULT NULL,
  `creditor` TEXT DEFAULT NULL,
  `intermediateInstitutions` JSON DEFAULT NULL,
  `remittanceInformation` TEXT DEFAULT NULL,
  `isBatch` TINYINT(1) DEFAULT NULL COMMENT 'Indicate if this transaction is part of a batch',
  `priority` TEXT DEFAULT NULL COMMENT 'Indicates the priority of this transaction',
  `executionDate` TEXT DEFAULT NULL COMMENT 'ISO 8601 date of execution',
  `isRecurring` TINYINT(1) DEFAULT NULL COMMENT 'Indicates if this is a recurring transaction',
  `recurringSchedule` TEXT DEFAULT NULL,
  `amount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of the transaction',
  `fee` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Exchange Transaction';

--
-- Table structure for table `ExchangeTransactionResult` generated from model 'ExchangeTransactionResult'
-- The resulting information from attempting to perform an Exchange Transaction
--

CREATE TABLE IF NOT EXISTS `ExchangeTransactionResult` (
  `status` TEXT DEFAULT NULL COMMENT 'The status of the attempt',
  `reason` TEXT DEFAULT NULL COMMENT 'A readable reason for the status provided.',
  `exchangeTransactionId` TEXT DEFAULT NULL COMMENT 'The newly created exchange transaction identifier'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='The resulting information from attempting to perform an Exchange Transaction';

--
-- Table structure for table `ExchangeTransactionsList` generated from model 'ExchangeTransactionsList'
--

CREATE TABLE IF NOT EXISTS `ExchangeTransactionsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `exchanges` JSON DEFAULT NULL COMMENT 'The list of exchange transactions'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ExchangeTransactionsList_allOf` generated from model 'ExchangeTransactionsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `ExchangeTransactionsList_allOf` (
  `exchanges` JSON DEFAULT NULL COMMENT 'The list of exchange transactions'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Fee` generated from model 'Fee'
-- Fee object
--

CREATE TABLE IF NOT EXISTS `Fee` (
  `amount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Fee amount',
  `currencyCode` TEXT DEFAULT NULL COMMENT 'Fee currency code',
  `owernship` ENUM('ORIGINATOR', 'SHARED', 'RECIPIENT') DEFAULT NULL COMMENT 'Indicates who is responsible for the fee'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Fee object';

--
-- Table structure for table `FiniteURI` generated from model 'FiniteURI'
-- Unique Resource Identifier in the API
--

CREATE TABLE IF NOT EXISTS `FiniteURI` (
  `URL` TEXT DEFAULT NULL COMMENT 'The URL to the URI'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Unique Resource Identifier in the API';

--
-- Table structure for table `Interest` generated from model 'Interest'
-- Interest object
--

CREATE TABLE IF NOT EXISTS `Interest` (
  `rate` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The effective interest rate % of a loan or saving account',
  `termsUnit` TEXT DEFAULT NULL COMMENT 'The measure of terms',
  `terms` INT DEFAULT NULL COMMENT 'The number of terms agreed'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Interest object';

--
-- Table structure for table `IntermediateInstitution` generated from model 'IntermediateInstitution'
-- Intermediate institution for ACH or Wire exchanges
--

CREATE TABLE IF NOT EXISTS `IntermediateInstitution` (
  `swift` TEXT DEFAULT NULL COMMENT 'Standard format of Business Identifier Codes (bic), unique identification code',
  `bankCode` TEXT DEFAULT NULL COMMENT 'A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.',
  `bankName` TEXT DEFAULT NULL COMMENT 'The financial institution name',
  `branchCode` TEXT DEFAULT NULL COMMENT 'The branch code of a bank branch helps in distinguishing one branch from another',
  `URI` TEXT DEFAULT NULL,
  `message` TEXT DEFAULT NULL COMMENT 'Message to the intermediate institution'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Intermediate institution for ACH or Wire exchanges';

--
-- Table structure for table `IntermediateInstitution_allOf` generated from model 'IntermediateInstitutionUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `IntermediateInstitution_allOf` (
  `message` TEXT DEFAULT NULL COMMENT 'Message to the intermediate institution'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Product` generated from model 'Product'
-- Product Object
--

CREATE TABLE IF NOT EXISTS `Product` (
  `id` TEXT DEFAULT NULL COMMENT 'Unique identifier of the product',
  `number` TEXT DEFAULT NULL COMMENT 'Number used to identify the product',
  `type` TEXT DEFAULT NULL COMMENT 'Product type',
  `typeName` TEXT DEFAULT NULL COMMENT 'Product type friendly name',
  `subType` TEXT DEFAULT NULL COMMENT 'Product sub type',
  `subTypeName` TEXT DEFAULT NULL COMMENT 'Product sub type friendly name',
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Product Object';

--
-- Table structure for table `ProductsList` generated from model 'ProductsList'
--

CREATE TABLE IF NOT EXISTS `ProductsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `products` JSON DEFAULT NULL COMMENT 'The list of products'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ProductsList_allOf` generated from model 'ProductsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `ProductsList_allOf` (
  `products` JSON DEFAULT NULL COMMENT 'The list of products'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Relationship` generated from model 'Relationship'
-- Relationship object
--

CREATE TABLE IF NOT EXISTS `Relationship` (
  `type` ENUM('PERSON', 'BUSINESS', 'SYSTEM') DEFAULT NULL COMMENT 'The type of profile',
  `id` TEXT NOT NULL COMMENT 'Unique identifier of the entity',
  `fullname` TEXT DEFAULT NULL COMMENT 'The fullname of a the entity',
  `firstname` TEXT DEFAULT NULL COMMENT 'The legal firstname of the person',
  `lastname` TEXT DEFAULT NULL COMMENT 'The legal lastname/surname of the person',
  `dateOfBirth` TEXT DEFAULT NULL COMMENT 'ISO 8601 Date of Birth',
  `addresses` JSON DEFAULT NULL COMMENT 'Entity addresses',
  `contactPoints` JSON DEFAULT NULL COMMENT 'Entity contact points',
  `preferences` JSON DEFAULT NULL COMMENT 'Entity preferences',
  `taxInformation` JSON DEFAULT NULL COMMENT 'Entity tax information',
  `URI` TEXT DEFAULT NULL,
  `relation` TEXT DEFAULT NULL COMMENT 'The type of relationship between entities'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Relationship object';

--
-- Table structure for table `Relationship_allOf` generated from model 'RelationshipUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `Relationship_allOf` (
  `relation` TEXT DEFAULT NULL COMMENT 'The type of relationship between entities'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `RemittanceInformation` generated from model 'RemittanceInformation'
-- Remittance Information
--

CREATE TABLE IF NOT EXISTS `RemittanceInformation` (
  `type` TEXT DEFAULT NULL COMMENT 'Type of remittance information',
  `issuer` TEXT DEFAULT NULL COMMENT 'Issuer information',
  `reference` TEXT DEFAULT NULL COMMENT 'Reference information/content'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Remittance Information';

--
-- Table structure for table `RoutingNumber` generated from model 'RoutingNumber'
-- Routing Number object
--

CREATE TABLE IF NOT EXISTS `RoutingNumber` (
  `type` TEXT DEFAULT NULL COMMENT 'The type of routing number',
  `number` TEXT DEFAULT NULL COMMENT 'The routing number value'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Routing Number object';

--
-- Table structure for table `Schedule` generated from model 'Schedule'
-- Schedule object
--

CREATE TABLE IF NOT EXISTS `Schedule` (
  `strategy` ENUM('NONE', 'BEFORE', 'AFTER') DEFAULT NULL COMMENT 'The execution strategy of the schedule',
  `frequency` TEXT DEFAULT NULL COMMENT 'Indicates how frequencly this schedule should execute',
  `dayOn` TEXT DEFAULT NULL COMMENT 'The day of the month on which a execution should take',
  `startDateTime` TEXT DEFAULT NULL COMMENT 'ISO 6801 formatted date for the schedule to start',
  `endDateTime` TEXT DEFAULT NULL COMMENT 'ISO 6801 formatted date for the schedule to complete',
  `nextDateTime` TEXT DEFAULT NULL COMMENT 'ISO 6801 formatted date for the next scheduled execution',
  `repeatCount` INT DEFAULT NULL COMMENT 'The number of times this schedule should execute',
  `isEveryTime` TINYINT(1) DEFAULT NULL COMMENT 'Indicates if the execution is even or odd'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Schedule object';

--
-- Table structure for table `SearchFilter` generated from model 'SearchFilter'
-- Search filter entry
--

CREATE TABLE IF NOT EXISTS `SearchFilter` (
  `field` TEXT DEFAULT NULL COMMENT 'The name of the field of element to query',
  `query` TEXT DEFAULT NULL COMMENT 'The value to query for',
  `operator` ENUM('EQUAL', 'LIKE', 'GREATER', 'LESS', 'NOT_LIKE', 'NOT_EQUAL', 'IN', 'NOT_IN') DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Search filter entry';

--
-- Table structure for table `SecurityInformation` generated from model 'SecurityInformation'
-- Security Information
--

CREATE TABLE IF NOT EXISTS `SecurityInformation` (
  `username` TEXT DEFAULT NULL COMMENT 'Unique username',
  `token` TEXT DEFAULT NULL COMMENT 'Unique token to identify entity for subsequent requests.  E.g. OAuth access_token',
  `lastLoginDate` TEXT DEFAULT NULL COMMENT 'ISO 8601 date the entity last logged in'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Security Information';

--
-- Table structure for table `Statement` generated from model 'Statement'
-- Account Statement
--

CREATE TABLE IF NOT EXISTS `Statement` (
  `id` TEXT DEFAULT NULL COMMENT 'The ID of this Statement.',
  `accountId` TEXT DEFAULT NULL COMMENT 'The ID of the Account for this Statement.',
  `name` TEXT DEFAULT NULL COMMENT 'The name of the document for this Statement.',
  `attributes` JSON DEFAULT NULL COMMENT 'List of attributes for a binary item, e.g description, creation time...',
  `dateCreated` TEXT DEFAULT NULL COMMENT 'ISO 8601 date of creation',
  `dateRevised` TEXT DEFAULT NULL COMMENT 'ISO 8601 date of revision',
  `items` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Account Statement';

--
-- Table structure for table `StatementsList` generated from model 'StatementsList'
--

CREATE TABLE IF NOT EXISTS `StatementsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `statements` JSON DEFAULT NULL COMMENT 'The list of statements'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `StatementsList_allOf` generated from model 'StatementsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `StatementsList_allOf` (
  `statements` JSON DEFAULT NULL COMMENT 'The list of statements'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TaxInformation` generated from model 'TaxInformation'
-- Tax information
--

CREATE TABLE IF NOT EXISTS `TaxInformation` (
  `type` ENUM('SSN', 'EIN', 'ITIN', 'ATIN', 'OTHER') DEFAULT NULL,
  `taxId` TEXT DEFAULT NULL COMMENT 'Tax identification value'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Tax information';

--
-- Table structure for table `Transaction` generated from model 'Transaction'
-- Transaction object
--

CREATE TABLE IF NOT EXISTS `Transaction` (
  `entityId` TEXT DEFAULT NULL COMMENT 'Unique identifier of the entity',
  `accountId` TEXT DEFAULT NULL COMMENT 'Unique identifier of the account in the system of record',
  `id` TEXT DEFAULT NULL COMMENT 'Unique identifier of the transaction in the system of record',
  `sequenceInDay` INT DEFAULT NULL COMMENT 'An integer sequence that can be used for ordering transactions within one day.',
  `status` TEXT DEFAULT NULL,
  `executedAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The amount of the transaction when it was executed and booked to the account.',
  `executedCurrencyCode` TEXT DEFAULT NULL COMMENT 'The currency code of the executed amount',
  `exchangeRate` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The exchange rate used to convert the originating amount to the executing amount',
  `originatingAmount` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The originating amount of the transaction when it was originally submitted to the account.',
  `originatingCurrencyCode` TEXT DEFAULT NULL COMMENT 'The currency code of the originating transaction amount',
  `direction` ENUM('DEBIT', 'CREDIT') DEFAULT NULL,
  `runningBalance` DECIMAL(20, 9) DEFAULT NULL COMMENT 'The balance of the account at the moment this transaction was executed',
  `createdDate` TEXT DEFAULT NULL COMMENT 'ISO 6801 date of when the transaction was created in the system of record',
  `postedDate` TEXT DEFAULT NULL COMMENT 'ISO 6801 date of when the transaction was posted to the account',
  `executedDate` TEXT DEFAULT NULL COMMENT 'ISO 6801 date of when the transaction was executed and booked to the account',
  `updatedDate` TEXT DEFAULT NULL COMMENT 'ISO 6801 date of when the transaction was updated on the account',
  `description` TEXT DEFAULT NULL COMMENT 'Unique identifier of the transaction in the system of record',
  `memo` TEXT DEFAULT NULL COMMENT 'Additional information about the transaction',
  `group` TEXT DEFAULT NULL COMMENT 'The grouping of this transaction in the system of record',
  `type` TEXT DEFAULT NULL COMMENT 'The type of this transaction in the system of record',
  `mainCategory` TEXT DEFAULT NULL COMMENT 'The main category of the transaction',
  `subCategory` TEXT DEFAULT NULL COMMENT 'The sub category of the transaction',
  `checkNumber` TEXT DEFAULT NULL COMMENT 'The check number associated with the transaction in case it was a check',
  `images` TEXT DEFAULT NULL,
  `latitude` TEXT DEFAULT NULL COMMENT 'The latitude coordinate where the transaction took place',
  `longitude` TEXT DEFAULT NULL COMMENT 'The longitude coordinate where the transaction took place',
  `merchantCode` TEXT DEFAULT NULL COMMENT 'The code used to identify the merchant for the transaction',
  `reversal` TINYINT(1) DEFAULT NULL COMMENT 'true if this transaction is a reversal, false otherwise.',
  `reversalFor` TEXT DEFAULT NULL COMMENT 'The transaction identifier this transaction is reversing',
  `reversed` TINYINT(1) DEFAULT NULL COMMENT 'true if this transaction has been reversed, false otherwise.',
  `reversedBy` TEXT DEFAULT NULL COMMENT 'The transaction identifier this transaction was reversed by',
  `transactor` TEXT DEFAULT NULL,
  `URI` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Transaction object';

--
-- Table structure for table `TransactionsList` generated from model 'TransactionsList'
--

CREATE TABLE IF NOT EXISTS `TransactionsList` (
  `status` TINYINT(1) DEFAULT NULL COMMENT 'Indicates whether a operation was successful',
  `messages` JSON DEFAULT NULL COMMENT 'Potential information or errror messages',
  `totalItems` INT DEFAULT NULL COMMENT 'The total number of list items that exist for the list or query',
  `nextCursor` TEXT DEFAULT NULL COMMENT 'The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.',
  `transactions` JSON DEFAULT NULL COMMENT 'The list of transactions'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TransactionsList_allOf` generated from model 'TransactionsListUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `TransactionsList_allOf` (
  `transactions` JSON DEFAULT NULL COMMENT 'The list of transactions'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Transactor` generated from model 'Transactor'
-- Person or Entity that performed a transaction
--

CREATE TABLE IF NOT EXISTS `Transactor` (
  `type` ENUM('PERSON', 'BUSINESS', 'SYSTEM') DEFAULT NULL COMMENT 'The type of profile',
  `id` TEXT NOT NULL COMMENT 'Unique identifier of the entity',
  `fullname` TEXT DEFAULT NULL COMMENT 'The fullname of a the entity',
  `firstname` TEXT DEFAULT NULL COMMENT 'The legal firstname of the person',
  `lastname` TEXT DEFAULT NULL COMMENT 'The legal lastname/surname of the person',
  `dateOfBirth` TEXT DEFAULT NULL COMMENT 'ISO 8601 Date of Birth',
  `addresses` JSON DEFAULT NULL COMMENT 'Entity addresses',
  `contactPoints` JSON DEFAULT NULL COMMENT 'Entity contact points',
  `preferences` JSON DEFAULT NULL COMMENT 'Entity preferences',
  `taxInformation` JSON DEFAULT NULL COMMENT 'Entity tax information',
  `URI` TEXT DEFAULT NULL,
  `swift` TEXT DEFAULT NULL COMMENT 'Standard format of Business Identifier Codes (bic), unique identification code',
  `bankCode` TEXT DEFAULT NULL COMMENT 'A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.',
  `bankName` TEXT DEFAULT NULL COMMENT 'The financial institution name',
  `branchCode` TEXT DEFAULT NULL COMMENT 'The branch code of a bank branch helps in distinguishing one branch from another',
  `accountNumber` TEXT DEFAULT NULL COMMENT 'The account number of the transactor performing the transaction'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Person or Entity that performed a transaction';

--
-- Table structure for table `Transactor_allOf` generated from model 'TransactorUnderscoreallOf'
--

CREATE TABLE IF NOT EXISTS `Transactor_allOf` (
  `accountNumber` TEXT DEFAULT NULL COMMENT 'The account number of the transactor performing the transaction'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


