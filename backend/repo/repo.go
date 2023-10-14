package repo

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"math/big"
	_ "net/http"
)

func InsertService(ctx *gin.Context, service models.Service) error {
	const query = `
INSERT INTO bank.bank_services VALUES ($1, $2);
`
	if _, err := db.ExecContext(ctx, query, service.ID, service.Description); err != nil {
		return err
	}

	return nil
}

func InsertClient(ctx context.Context, client models.Client) error {
	const query = `
INSERT INTO bank.clients (id, latitude, longitude) VALUES ($1, $2, $3);
`

	lat, _, err := big.ParseFloat(client.LocationLat, 10, 1, big.ToNearestEven)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	lon, _, err := big.ParseFloat(client.LocationLon, 10, 1, big.ToNearestEven)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	if _, err = db.ExecContext(ctx, query, client.ID, lat, lon); err != nil {
		return err
	}

	return nil
}

func GetClientLocationByID(ctx context.Context, id uuid.UUID) (string, string, error) {
	const query = `
SELECT latitude, longitude FROM bank.clients WHERE id = $1;`

	var (
		lat, long *BigFloatScanner
	)

	if err := db.QueryRowContext(ctx, query, id).Scan(&lat, &long); err != nil {
		return "", "", fmt.Errorf("scan: %w", err)
	}

	return lat.BigFloat().String(), long.BigFloat().String(), nil
}

func GetBanks(ctx context.Context) ([]models.Bank, error) {
	const query = `
SELECT id, salepointname, latitude, longitude FROM bank.banks;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Bank

	for rows.Next() {
		var bank models.Bank
		if err = rows.Scan(&bank.ID, &bank.Name, &bank.Latitude, &bank.Longitude); err != nil {
			return nil, err
		}

		res = append(res, bank)
	}

	if len(res) == 0 {
		return nil, sql.ErrNoRows
	}

	return res, nil
}

func GetBank(id string) (*models.Bank, error) {
	query := `SELECT id, salepointname, latitude, longitude FROM bank.banks WHERE id = $1`
	var res *models.Bank
	err := models.DB.QueryRow(query, id).Scan(&res)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
