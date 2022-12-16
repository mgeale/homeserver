INSERT INTO users (name, email, hashed_password, created) VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2018-12-23 17:25:22'
);

INSERT INTO balances (name, balance, balanceaud, pricebookid, productid, created) VALUES (
    'BAL-0022',
    100.89,
    1000.01,
    3333,
    2222,
    '2018-12-23 17:25:22'
);

INSERT INTO transactions (name, amount, date, type, created) VALUES (
    'name',
    100,
    '2018-12-23 17:25:22',
    'Repayment',
    '2018-12-23 17:25:22'
);

INSERT INTO tokens (hash, user_id, expiry, scope) VALUES (
    0xB141DDEDC0A61707090303ADDDEA841D5284FFF0A80C56A6BB7394EB26957596,
    1,
    '2030-01-01 00:00:01',
    'authentication'
);

INSERT INTO users_permissions VALUES (
    1,
    1
);