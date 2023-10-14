package repo

import (
	"database/sql"
	_ "fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	_ "net/http"
)

func InsertService(ctx *gin.Context, service models.Service) error {
	const query = `
INSERT INTO bank.bank_services VALUES ($1, $2);
`
	if _, err := models.DB.ExecContext(ctx, query, service.ID, service.Description); err != nil {
		return err
	}

	return nil
}

func GetBanks(ctx context.Context) ([]models.Bank, error) {
	const query = `
SELECT id, salepointname, latitude, longitude FROM bank.banks;`

	rows, err := models.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Bank

	for rows.Next() {
		var bank models.Bank
		if err = rows.Scan(&bank.ID, &bank.Latitude, &bank.Longitude); err != nil {
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
	query := `SELECT id, salepointname, latitude, longitude FROM banks WHERE id = $1`
	var res *models.Bank
	err := models.DB.QueryRow(query, id).Scan(&res) 
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}