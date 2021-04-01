--
-- INSERT template for table `Account`
--
-- INSERT INTO `Account`(`id`, `type`, `number`, `parentId`, `entities`, `name`, `title`, `iban`, `routingnumbers`, `balances`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `bank`, `owners`, `product`, `URI`) 
-- VALUES ("123456-DDA", "DDA", "xxxxxxx1213", "12345ABC", null, "Matt Dollar", "401k Plan", "GB29NWBK60161331926819", null, null, "YYYY-MM-DDThh:mm:ss", "YYYY-MM-DDThh:mm:ss", "YYYY-MM-DDThh:mm:ss", "YYYY-MM-DDThh:mm:ss", "USD", "OPEN", "core", 1, null, null, null, null, null, null, null, "http://host/images/1234.jpg");

INSERT INTO `Account`(`id`, `type`, `number`, `parentId`, `name`, `title`, `iban`, `dateCreated`, `dateOpened`, `dateLastUpdated`, `dateClosed`, `currencyCode`, `status`, `source`, `interestReporting`, `URI`) 
VALUES ("e14f44ea-5380-4368-9631-45c1d6a3bf6c", "DDA", "xxxxxxx1213", "97208803-274d-4d82-be2b-482651314934", "Matt Dollar", "401k Plan", "GB29NWBK60161331926819", "2020-02-18T16:44:01", "2020-02-18T16:44:01", "2020-02-18T16:44:01", "2020-02-18T16:44:01", "USD", "OPEN", "core", 1, "http://host/images/1234.jpg");
