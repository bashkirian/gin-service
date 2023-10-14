package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
)

type BigFloatScanner struct {
	float *big.Float
}

var _ sql.Scanner = (*BigFloatScanner)(nil)

func (s *BigFloatScanner) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		s.float = nil
		return nil

	case string:
		if len(src) == 0 {
			return errors.New("scanned string is empty")
		}

		s.float = big.NewFloat(0).SetPrec(64)
		if err := s.float.UnmarshalText([]byte(src)); err != nil {
			return fmt.Errorf("unmarshal string: %w", err)
		}

		return nil

	case float64:
		return errors.New("will not scan type float64 into *big.Float")

	case int64:
		s.float = big.NewFloat(float64(src))
		return nil

	default:
		return fmt.Errorf("unable to scan type %T into *big.Float", src)
	}
}

func (s *BigFloatScanner) BigFloat() *big.Float {
	return s.float
}
