package repo

import (
	"database/sql"
	"fmt"
	"github.com/bashkirian/gin-service/models"
	"golang.org/x/net/context"
)

func InsertService(ctx context.Context, service models.Service) error {
	const query = `
INSERT INTO bank.bank_services VALUES ($1, $2);
`

	if _, err := DB.ExecContext(ctx, query, service.ID, service.Description); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func GetBanks(ctx context.Context) ([]models.Bank, error) {
	const query = `
SELECT id, salepointname, latitude, longitude FROM bank.banks;`

	rows, err := DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w")
	}
	defer rows.Close()

	var res []models.Bank

	for rows.Next() {
		var bank models.Bank
		if err = rows.Scan(&bank.ID, &bank.Latitude, &bank.Longitude); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		res = append(res, bank)
	}

	if len(res) == 0 {
		return nil, sql.ErrNoRows
	}

	return res, nil
}
