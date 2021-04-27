import http from 'k6/http';
import { check, group, sleep, fail } from 'k6';

export let options = {
  vus: 1, // 1 user looping
  duration: '5s', // for 5 seconds

  thresholds: {
    http_req_duration: ['p(99)<400'], // 99% of requests must complete below 200ms
  },
};

const BASE_URL = `http://localhost:${__ENV.EMULATOR_LISTEN_PORT}/cloud`;
const ACCESS_TOKEN = `${__ENV.FCA_SECRET}`;
const ENTITY_ID = `${__ENV.ENTITY_ID}`;
const MASK = false
const ENHANCE = false
const LIMIT = 20

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
  getEntityAccountsListJSON.accounts.forEach(function (arrayItem){

    // Get account info
    let getAccountRes = http.get(`${BASE_URL}/v1/account/${arrayItem.id}?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountJSON = JSON.parse(getAccountRes.body);
    check(getAccountRes, {
      "GetAccount: status is 200": (r) => r.status === 200,
      "GetAccount: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetAccount: has valid account ID": (r) => getAccountJSON.id === `${arrayItem.id}`,
    });

    // Get account details
    let getAccountDetailsRes = http.get(`${BASE_URL}/v1/account/${arrayItem.id}/details?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountDetailsJSON = JSON.parse(getAccountDetailsRes.body);
    check(getAccountDetailsRes, {
      "GetAccountDetails: status is 200": (r) => r.status === 200,
      "GetAccountDetails: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetAccountDetails: has valid account ID": (r) => getAccountDetailsJSON.id === `${arrayItem.id}`,
    });

    // Get account balances
    let getAccountBalancesRes = http.get(`${BASE_URL}/v1/account/${arrayItem.id}/balances?mask=${MASK}&enhance=${ENHANCE}`, authHeaders);
    let getAccountBalancesJSON = JSON.parse(getAccountBalancesRes.body);
    check(getAccountBalancesRes, {
      "GetAccountBalances: status is 200": (r) => r.status === 200,
      "GetAccountBalances: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetAccountBalances: has available balance": (r) => isFloat(getAccountBalancesJSON.availableBalance),
      "GetAccountBalances: has current balance": (r) => isFloat(getAccountBalancesJSON.currentBalance),
    });

    // Get transactions for account
    let getAccountTransactionsRes = http.get(`${BASE_URL}/v1/account/${arrayItem.id}/transactions?mask=${MASK}&enhance=${ENHANCE}&limit=${LIMIT}`, authHeaders);
    let getAccountTransactionsJSON = JSON.parse(getAccountTransactionsRes.body);
    check(getAccountTransactionsRes, {
      "GetAccountTransactions: status is 200": (r) => r.status === 200,
      "GetAccountTransactions: http version is 1.1": (r) => r.proto === "HTTP/1.1",
      "GetEntityAccountsList: has one or more transactions": (r) => getAccountTransactionsJSON.totalItems > 0,
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
