import http from 'k6/http';
import { check, group, sleep, fail } from 'k6';

export let options = {
  vus: 1, // 1 user looping
  duration: '5s', // for 5 seconds

  thresholds: {
    // 95% of requests must finish within 50ms, 99% within 80, and 99.9% within 200ms.
    http_req_duration: ['p(95) < 50', 'p(99) < 80', 'p(99.9) < 200'],
    // The error rate must be lower than 1%.
    http_req_failed: ['rate<0.01'], 
  },
};

const BASE_URL = `http://localhost:${__ENV.EMULATOR_LISTEN_PORT}/cloud`;
const ACCESS_TOKEN = `${__ENV.FCA_SECRET}`;
const ENTITY_ID = `${__ENV.ENTITY_ID}`;
const MASK = false
const ENHANCE = false
const LIMIT = 20
const INLINE = true

export default () => {
  let authHeaders = {
    headers: {
      Accept: 'application/json',
      'X-TOKEN': `${ACCESS_TOKEN}`,
      'X-TRACE-ID': uuidv4(),
    },
  };

  // Get products
  let getProductsRes = http.get(`${BASE_URL}/v1/products?productType=CHECKING&limit=${LIMIT}&enhance=${ENHANCE}`, authHeaders);
  let getProductsJSON = JSON.parse(getProductsRes.body);
  check(getProductsRes, {
    "GetProducts: status is 200": (r) => r.status === 200,
    "GetProducts: http version is 1.1": (r) => r.proto === "HTTP/1.1",
    "GetProducts: has valid totalItems": (r) => isInt(getProductsJSON.totalItems),
  });

  // Get entity
  let getEntityRes = http.get(`${BASE_URL}/v1/entity/${ENTITY_ID}?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
  let getEntityJSON = JSON.parse(getEntityRes.body);
  check(getEntityRes, {
    "GetEntity: status is 200": (r) => r.status === 200,
    "GetEntity: http version is 1.1": (r) => r.proto === "HTTP/1.1",
    "GetEntity: has valid entity ID": (r) => getEntityJSON.id === ENTITY_ID,
  });

  // Get entity profile
  let getEntityProfileRes = http.get(`${BASE_URL}/v1/entity/${ENTITY_ID}/profile?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
  let getEntityProfileJSON = JSON.parse(getEntityProfileRes.body);
  check(getEntityProfileRes, {
    "GetEntityProfile: status is 200": (r) => r.status === 200,
    "GetEntityProfile: http version is 1.1": (r) => r.proto === "HTTP/1.1",
    "GetEntityProfile: has valid entity ID": (r) => getEntityProfileJSON.id === ENTITY_ID,
  });

  // Get accounts for entity
  let getEntityAccountsListRes = http.get(`${BASE_URL}/v1/entity/${ENTITY_ID}/accounts?mask=${MASK}&limit=${LIMIT}&enhance=${ENHANCE}`, authHeaders);
  let getEntityAccountsListJSON = JSON.parse(getEntityAccountsListRes.body);
  check(getEntityAccountsListRes, {
    "GetEntityAccountsList: status is 200": (r) => r.status === 200,
    "GetEntityAccountsList: http version is 1.1": (r) => r.proto === "HTTP/1.1",
    "GetEntityAccountsList: has one or more accounts": (r) => getEntityAccountsListJSON.totalItems > 0,
  });

  // For each account...
  getEntityAccountsListJSON.accounts.forEach(function (accountItem){

    // Get account info
    let getAccountRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountJSON = JSON.parse(getAccountRes.body);
    check(getAccountRes, {
      "GetAccount: status is 200": (r) => r.status === 200,
      "GetAccount: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetAccount: has valid account ID": (r) => getAccountJSON.id === `${accountItem.id}`,
    });

    // Get account details
    let getAccountDetailsRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/details?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountDetailsJSON = JSON.parse(getAccountDetailsRes.body);
    check(getAccountDetailsRes, {
      "GetAccountDetails: status is 200": (r) => r.status === 200,
      "GetAccountDetails: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetAccountDetails: has valid account ID": (r) => getAccountDetailsJSON.id === `${accountItem.id}`,
    });

    // Get account balances
    let getAccountBalancesRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/balances?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountBalancesJSON = JSON.parse(getAccountBalancesRes.body);
    check(getAccountBalancesRes, {
      "GetAccountBalances: status is 200": (r) => r.status === 200,
      "GetAccountBalances: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetAccountBalances: has available balance": (r) => isFloat(getAccountBalancesJSON.availableBalance),
      "GetAccountBalances: has current balance": (r) => isFloat(getAccountBalancesJSON.currentBalance),
    });

    // Get account cards
    let getAccountCardsRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/cards?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountCardsJSON = JSON.parse(getAccountCardsRes.body);
    check(getAccountCardsRes, {
      "GetAccountCards: status is 200 or 404": (r) => r.status === 200 || r.status === 404,
      "GetAccountCards: http version is 1.1": (r) => r.proto === "HTTP/1.1",
    });
    if (getAccountCardsRes.status === 200) {
      check(getAccountCardsRes, {
        "GetAccountCards: has one or more cards": (r) => getAccountCardsJSON.totalItems > 0,
      });
    }
    
    // Get transactions for account
    let getAccountTransactionsRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/transactions?mask=${MASK}&enhance=${ENHANCE}&limit=${LIMIT}`, authHeaders);
    let getAccountTransactionsJSON = JSON.parse(getAccountTransactionsRes.body);
    check(getAccountTransactionsRes, {
      "GetAccountTransactions: status is 200": (r) => r.status === 200,
      "GetAccountTransactions: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetEntityAccountsList: has one or more transactions": (r) => getAccountTransactionsJSON.totalItems > 0,
    });

    getAccountTransactionsJSON.transactions.forEach(function (transactionItem){
      // Get transaction details for each transaction
      let getAccountTransactionRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/transaction/${transactionItem.id}?mask=${MASK}&enhance=${ENHANCE}&inline=${INLINE}`, authHeaders);
      let getAccountTransactionJSON = JSON.parse(getAccountTransactionRes.body);
      check(getAccountTransactionRes, {
        "GetAccountTransaction: status is 200": (r) => r.status === 200,
        "GetAccountTransaction: http version is 1.1": (r) => r.proto === "HTTP/1.1",
        "GetAccountTransaction: has valid account ID": (r) => getAccountTransactionJSON.accountId === `${accountItem.id}`,
        "GetAccountTransaction: has valid entity ID": (r) => getAccountTransactionJSON.entityId === ENTITY_ID,
        "GetAccountTransaction: has valid ID": (r) => getAccountTransactionJSON.id === `${transactionItem.id}`,
      });

      // Get images for each transaction
      let getAccountTransactionImagesRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/transaction/${transactionItem.id}/images?mask=${MASK}&enhance=${ENHANCE}&inline=${INLINE}`, authHeaders);
      let getAccountTransactionImagesJSON = JSON.parse(getAccountTransactionImagesRes.body);
      if (getAccountTransactionImagesJSON.totalItems > 0) {
        check(getAccountTransactionImagesRes, {
          "GetAccountTransactionImages: status is 200": (r) => r.status === 200,
          "GetAccountTransactionImages: http version is 1.1": (r) => r.proto === "HTTP/1.1",
          "GetAccountTransactionImages: has one or more images": (r) => getAccountTransactionImagesJSON.binaries.length > 0,
        });

        // Get (first) image based upon account_id and image_id
        let imageId = getAccountTransactionImagesJSON.binaries[0].itemId
        let getAccountTransactionImageRes = http.get(`${BASE_URL}/v1/account/${accountItem.id}/image/${imageId}`, authHeaders);
        let getAccountTransactionImageJSON = JSON.parse(getAccountTransactionImageRes.body);
        check(getAccountTransactionImageRes, {
          "GetAccountTransactionImage: status is 200": (r) => r.status === 200,
          "GetAccountTransactionImage: http version is 1.1": (r) => r.proto === "HTTP/1.1",
        });
      }
    });
  })
};

// Supporting functions
function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

function isFloat(n){
  return Number(n) === n && n % 1 !== 0;
}

function isInt(n) {
  return n % 1 === 0;
}
