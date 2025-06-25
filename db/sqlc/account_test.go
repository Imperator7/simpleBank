package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Imperator7/simpleBank.git/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// Check the correctness of the returned account
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}	

func TestGetAccount(t *testing.T) {
	// Create a new account
 arg := CreateRandomAccount(t)

 account, err := testQueries.GetAccount(context.Background(), arg.ID)

 require.NoError(t, err)
 require.NotEmpty(t, account)

 require.Equal(t, arg.ID, account.ID)
 require.Equal(t, arg.Owner, account.Owner)
 require.Equal(t, arg.Balance, account.Balance)
 require.Equal(t, arg.Currency, account.Currency)
 require.WithinDuration(t, arg.CreatedAt, account.CreatedAt, time.Second*2)
}

func TestUpdateAccount (t *testing.T) {
	account := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account.ID,
		Balance: util.RandomMoney(),
	}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second*2)

}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	accountCheck, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountCheck)
}

func TestListAccounts(t *testing.T){
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
  }

 arg := ListAccountsParams{
	Limit: 5,
	Offset: 5,
 }

 accounts, err := testQueries.ListAccounts(context.Background(), arg)
 require.NoError(t, err)
 require.Len(t, accounts, 5)

 for _, account := range accounts {
	require.NotEmpty(t, account)
 }
}