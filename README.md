<h1 align="center">
  <br>
  The Amazing Ledger
</h1>

![Main](https://github.com/stone-co/the-amazing-ledger/workflows/Main/badge.svg)

The Amazing Ledger is a double-entry accounting system used to record transactions. The entries are processed, validated and persisted to a database and their original data is **never** modified.

# Concepts

The Amazing Ledger records a journal that keeps track of all credits and debits to and from accounts (an account being anywhere money can go, not necessarily a bank account).

Since it is a double-entry accounting system, there must always be a debit from at least one account for every credit made to another account, therefore all **entries** must balance to zero within a **transaction**.

For example, 

| Account          | Amount |
|------------------|--------|
| Satoshi Nakamoto | $ -50  |
| Hal Finney       | $ 50   |

The system supports the following account types:
- **equity**: Represents your net worth. This will be used to start the ledger and its values are naturally negative.
- **asset**: Represents the money you have, where money sits. The values of theses accounts are naturally positive.
- **liability**: Represents the money you owe and is a synonym to debt. The values of theses accounts are naturally negative.
- **expense**: Represents the money you spent, where money goes. The values of theses accounts are naturally positive.
- **income**: Represents the money you have earned, where money comes from. The values of theses accounts are naturally negative.

# Usage

A Web API is used to issue commands and query for data in the ledger.