package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/line/ostracon/libs/log"
	ostproto "github.com/line/ostracon/proto/ostracon/types"
	"github.com/line/tm-db/v2/memdb"
	"github.com/stretchr/testify/suite"

	"github.com/line/lfb-sdk/codec"
	"github.com/line/lfb-sdk/simapp"
	"github.com/line/lfb-sdk/store"
	sdk "github.com/line/lfb-sdk/types"
	"github.com/line/lfb-sdk/x/params/types"
)

type SubspaceTestSuite struct {
	suite.Suite

	cdc   codec.BinaryMarshaler
	amino *codec.LegacyAmino
	ctx   sdk.Context
	ss    *types.Subspace
}

func (suite *SubspaceTestSuite) SetupTest() {
	db := memdb.NewDB()

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	suite.NoError(ms.LoadLatestVersion())

	encCfg := simapp.MakeTestEncodingConfig()
	ss := types.NewSubspace(encCfg.Marshaler, encCfg.Amino, key, "testsubspace")

	suite.cdc = encCfg.Marshaler
	suite.amino = encCfg.Amino
	suite.ctx = sdk.NewContext(ms, ostproto.Header{}, false, log.NewNopLogger())
	suite.ss = ss.WithKeyTable(paramKeyTable())
}

func (suite *SubspaceTestSuite) TestKeyTable() {
	suite.Require().True(suite.ss.HasKeyTable())
	suite.Require().Panics(func() {
		suite.ss.WithKeyTable(paramKeyTable())
	})
	suite.Require().NotPanics(func() {
		ss := types.NewSubspace(suite.cdc, suite.amino, key, "testsubspace2")
		ss = ss.WithKeyTable(paramKeyTable())
	})
}

func (suite *SubspaceTestSuite) TestGetSet() {
	var v time.Duration
	t := time.Hour * 48

	suite.Require().Panics(func() {
		suite.ss.Get(suite.ctx, keyUnbondingTime, &v)
	})
	suite.Require().NotEqual(t, v)
	suite.Require().NotPanics(func() {
		suite.ss.Set(suite.ctx, keyUnbondingTime, t)
	})
	suite.Require().NotPanics(func() {
		suite.ss.Get(suite.ctx, keyUnbondingTime, &v)
	})
	suite.Require().Equal(t, v)
}

func (suite *SubspaceTestSuite) TestGetIfExists() {
	var v time.Duration

	suite.Require().NotPanics(func() {
		suite.ss.GetIfExists(suite.ctx, keyUnbondingTime, &v)
	})
	suite.Require().Equal(time.Duration(0), v)
}

func (suite *SubspaceTestSuite) TestGetRaw() {
	t := time.Hour * 48

	suite.Require().NotPanics(func() {
		suite.ss.Set(suite.ctx, keyUnbondingTime, t)
	})
	suite.Require().NotPanics(func() {
		res := suite.ss.GetRaw(suite.ctx, keyUnbondingTime)
		suite.Require().Equal("2231373238303030303030303030303022", fmt.Sprintf("%X", res))
	})
}

func (suite *SubspaceTestSuite) TestHas() {
	t := time.Hour * 48

	suite.Require().False(suite.ss.Has(suite.ctx, keyUnbondingTime))
	suite.Require().NotPanics(func() {
		suite.ss.Set(suite.ctx, keyUnbondingTime, t)
	})
	suite.Require().True(suite.ss.Has(suite.ctx, keyUnbondingTime))
}

func (suite *SubspaceTestSuite) TestUpdate() {
	suite.Require().Panics(func() {
		suite.ss.Update(suite.ctx, []byte("invalid_key"), nil) // nolint:errcheck
	})

	t := time.Hour * 48
	suite.Require().NotPanics(func() {
		suite.ss.Set(suite.ctx, keyUnbondingTime, t)
	})

	bad := time.Minute * 5

	bz, err := suite.amino.MarshalJSON(bad)
	suite.Require().NoError(err)
	suite.Require().Error(suite.ss.Update(suite.ctx, keyUnbondingTime, bz))

	good := time.Hour * 360
	bz, err = suite.amino.MarshalJSON(good)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.ss.Update(suite.ctx, keyUnbondingTime, bz))

	var v time.Duration

	suite.Require().NotPanics(func() {
		suite.ss.Get(suite.ctx, keyUnbondingTime, &v)
	})
	suite.Require().Equal(good, v)
}

func (suite *SubspaceTestSuite) TestGetParamSet() {
	a := params{
		UnbondingTime: time.Hour * 48,
		MaxValidators: 100,
		BondDenom:     "stake",
	}
	suite.Require().NotPanics(func() {
		suite.ss.Set(suite.ctx, keyUnbondingTime, a.UnbondingTime)
		suite.ss.Set(suite.ctx, keyMaxValidators, a.MaxValidators)
		suite.ss.Set(suite.ctx, keyBondDenom, a.BondDenom)
	})

	b := params{}
	suite.Require().NotPanics(func() {
		suite.ss.GetParamSet(suite.ctx, &b)
	})
	suite.Require().Equal(a.UnbondingTime, b.UnbondingTime)
	suite.Require().Equal(a.MaxValidators, b.MaxValidators)
	suite.Require().Equal(a.BondDenom, b.BondDenom)
}

func (suite *SubspaceTestSuite) TestSetParamSet() {
	testCases := []struct {
		name string
		ps   types.ParamSet
	}{
		{"invalid unbonding time", &params{time.Hour * 1, 100, "stake"}},
		{"invalid bond denom", &params{time.Hour * 48, 100, ""}},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.Require().Panics(func() {
				suite.ss.SetParamSet(suite.ctx, tc.ps)
			})
		})
	}

	a := params{
		UnbondingTime: time.Hour * 48,
		MaxValidators: 100,
		BondDenom:     "stake",
	}
	suite.Require().NotPanics(func() {
		suite.ss.SetParamSet(suite.ctx, &a)
	})

	b := params{}
	suite.Require().NotPanics(func() {
		suite.ss.GetParamSet(suite.ctx, &b)
	})
	suite.Require().Equal(a.UnbondingTime, b.UnbondingTime)
	suite.Require().Equal(a.MaxValidators, b.MaxValidators)
	suite.Require().Equal(a.BondDenom, b.BondDenom)
}

func (suite *SubspaceTestSuite) TestParamSetCache() {
	a := params{
		UnbondingTime: time.Hour * 48,
		MaxValidators: 100,
		BondDenom:     "stake",
	}
	b := params{}
	// confirm cache is empty
	suite.Require().False(suite.ss.GetCachedValueForTesting(keyUnbondingTime, &b.UnbondingTime))
	suite.Require().False(suite.ss.GetCachedValueForTesting(keyMaxValidators, &b.MaxValidators))
	suite.Require().False(suite.ss.GetCachedValueForTesting(keyBondDenom, &b.BondDenom))
	suite.Require().False(suite.ss.HasCacheForTesting(keyUnbondingTime))
	suite.Require().False(suite.ss.HasCacheForTesting(keyMaxValidators))
	suite.Require().False(suite.ss.HasCacheForTesting(keyBondDenom))
	suite.Require().NotEqual(a, b)

	suite.Require().NotPanics(func() {
		suite.ss.SetParamSet(suite.ctx, &a)
	})
	// confirm cached
	suite.Require().True(suite.ss.GetCachedValueForTesting(keyUnbondingTime, &b.UnbondingTime))
	suite.Require().True(suite.ss.GetCachedValueForTesting(keyMaxValidators, &b.MaxValidators))
	suite.Require().True(suite.ss.GetCachedValueForTesting(keyBondDenom, &b.BondDenom))
	suite.Require().Equal(a, b)

	// update local variable
	a.UnbondingTime = time.Hour * 24
	a.MaxValidators = 50
	a.BondDenom = "link"

	// confirm the cache is not updated
	c := params{}
	suite.Require().True(suite.ss.GetCachedValueForTesting(keyUnbondingTime, &c.UnbondingTime))
	suite.Require().True(suite.ss.GetCachedValueForTesting(keyMaxValidators, &c.MaxValidators))
	suite.Require().True(suite.ss.GetCachedValueForTesting(keyBondDenom, &c.BondDenom))
	suite.Require().Equal(b, c) // cached value is not updated

	// update only cache
	suite.ss.SetCacheForTesting(keyUnbondingTime, &a.UnbondingTime)
	suite.ss.SetCacheForTesting(keyMaxValidators, &a.MaxValidators)
	suite.ss.SetCacheForTesting(keyBondDenom, &a.BondDenom)

	// ensure GetParamSet to get value not from db but from cache
	suite.Require().NotPanics(func() {
		suite.ss.GetParamSet(suite.ctx, &c)
	})
	suite.Require().Equal(a, c)
}

func (suite *SubspaceTestSuite) TestName() {
	suite.Require().Equal("testsubspace", suite.ss.Name())
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(SubspaceTestSuite))
}
