package keeper_test

import (
	"fmt"

	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/collection"
)

func (s *KeeperTestSuite) TestSendCoins() {
	testCases := map[string]struct {
		tokenID string
		amount sdk.Int
		valid  bool
	}{
		"valid send (fungible token)": {
			tokenID: s.ftClassID + fmt.Sprintf("%08x", 0),
			amount: s.balance,
			valid: true,
		},
		"valid send (non-fungible token)": {
			tokenID: s.nftClassID + fmt.Sprintf("%08x", 3),
			amount: sdk.OneInt(),
			valid: true,
		},
		"insufficient tokens": {
			tokenID: s.ftClassID + fmt.Sprintf("%08x", 0),
			amount: s.balance.Add(sdk.OneInt()),
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			customerBalance := s.keeper.GetBalance(ctx, s.contractID, s.customer, tc.tokenID)
			operatorBalance := s.keeper.GetBalance(ctx, s.contractID, s.operator, tc.tokenID)

			coin := collection.Coin{TokenId: tc.tokenID, Amount: tc.amount}
			err := s.keeper.SendCoins(ctx, s.contractID, s.customer, s.operator, []collection.Coin{coin})
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			newCustomerBalance := s.keeper.GetBalance(ctx, s.contractID, s.customer, tc.tokenID)
			newOperatorBalance := s.keeper.GetBalance(ctx, s.contractID, s.operator, tc.tokenID)
			s.Require().True(customerBalance.Sub(tc.amount).Equal(newCustomerBalance))
			s.Require().True(operatorBalance.Add(tc.amount).Equal(newOperatorBalance))
		})
	}
}

func (s *KeeperTestSuite) TestAuthorizeOperator() {
	// make sure the dummy contract does not exist
	dummyContractID := "fee1dead"
	_, err := s.keeper.GetContract(s.ctx, dummyContractID)
	s.Require().Error(err)

	contractDescriptions := map[string]string{
		s.contractID: "valid",
		dummyContractID: "not-exists",
	}
	userDescriptions := map[sdk.AccAddress]string{
		s.vendor: "vendor",
		s.operator: "operator",
		s.customer: "customer",
	}
	for id, idDesc := range contractDescriptions {
		for operator, operatorDesc := range userDescriptions {
			for from, fromDesc := range userDescriptions {
				name := fmt.Sprintf("ContractID: %s, Operator: %s, From: %s", idDesc, operatorDesc, fromDesc)
				s.Run(name, func() {
					ctx, _ := s.ctx.CacheContext()

					_, idErr := s.keeper.GetContract(ctx, id)
					_, authErr := s.keeper.GetAuthorization(ctx, id, from, operator)
					err := s.keeper.AuthorizeOperator(ctx, id, from, operator)
					if idErr == nil && authErr != nil {
						s.Require().NoError(err)
						_, authErr = s.keeper.GetAuthorization(ctx, id, from, operator)
						s.Require().NoError(authErr)
					} else {
						s.Require().Error(err)
					}
				})
			}
		}
	}
}

func (s *KeeperTestSuite) TestRevokeOperator() {
	// make sure the dummy contract does not exist
	dummyContractID := "fee1dead"
	_, err := s.keeper.GetContract(s.ctx, dummyContractID)
	s.Require().Error(err)

	contractDescriptions := map[string]string{
		s.contractID: "valid",
		dummyContractID: "not-exists",
	}
	userDescriptions := map[sdk.AccAddress]string{
		s.vendor: "vendor",
		s.operator: "operator",
		s.customer: "customer",
	}
	for id, idDesc := range contractDescriptions {
		for operator, operatorDesc := range userDescriptions {
			for from, fromDesc := range userDescriptions {
				name := fmt.Sprintf("ContractID: %s, Operator: %s, From: %s", idDesc, operatorDesc, fromDesc)
				s.Run(name, func() {
					ctx, _ := s.ctx.CacheContext()

					_, idErr := s.keeper.GetContract(ctx, id)
					_, authErr := s.keeper.GetAuthorization(ctx, id, from, operator)
					err := s.keeper.RevokeOperator(ctx, id, from, operator)
					if idErr == nil && authErr == nil {
						s.Require().NoError(err)
						_, authErr = s.keeper.GetAuthorization(ctx, id, from, operator)
						s.Require().Error(authErr)
					} else {
						s.Require().Error(err)
					}
				})
			}
		}
	}
}
