package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestQueries_CreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  200,
		Currency: "USD",
	}

	account, _ := testQueries.CreateAccount(context.Background(), arg)
	log.Print(account)

	require.NotNil(t, account)

}
