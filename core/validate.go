package core

import (
	"errors"
	"fmt"
	"math/big"
)

var zeroAddress Address

// Validate checks a BetTerms object for the invariants that must hold before a
// bet can be hashed and created. These are conservative, explicit checks — we
// err toward rejecting ambiguous terms rather than letting them on-chain.
func (t BetTerms) Validate() error {
	var errs []error

	if t.YesAgent == zeroAddress {
		errs = append(errs, errors.New("yesAgent is the zero address"))
	}
	if t.NoAgent == zeroAddress {
		errs = append(errs, errors.New("noAgent is the zero address"))
	}
	if t.YesAgent == t.NoAgent && t.YesAgent != zeroAddress {
		errs = append(errs, errors.New("yesAgent and noAgent must be different"))
	}
	if t.CollateralToken == zeroAddress {
		errs = append(errs, errors.New("collateralToken is the zero address"))
	}
	if !positive(t.YesStake) {
		errs = append(errs, errors.New("yesStake must be > 0"))
	}
	if !positive(t.NoStake) {
		errs = append(errs, errors.New("noStake must be > 0"))
	}
	if t.Statement == "" {
		errs = append(errs, errors.New("statement is empty"))
	}
	if t.PrimarySource == "" {
		errs = append(errs, errors.New("primarySource is empty"))
	}
	if t.EventTime == 0 {
		errs = append(errs, errors.New("eventTime is unset"))
	}
	if t.ClaimDeadline == 0 {
		errs = append(errs, errors.New("claimDeadline is unset"))
	}
	// The claim deadline must come at or after the event it settles.
	if t.EventTime != 0 && t.ClaimDeadline != 0 && t.ClaimDeadline < t.EventTime {
		errs = append(errs, fmt.Errorf("claimDeadline (%d) precedes eventTime (%d)", t.ClaimDeadline, t.EventTime))
	}
	if t.ChallengeWindow == 0 {
		errs = append(errs, errors.New("challengeWindow is zero (no window to contest a claim)"))
	}
	if t.Nonce == nil {
		errs = append(errs, errors.New("nonce is nil (needed to disambiguate otherwise-identical terms)"))
	}

	return errors.Join(errs...)
}

func positive(x *big.Int) bool { return x != nil && x.Sign() > 0 }
