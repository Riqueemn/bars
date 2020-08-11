package bars

import (
	"errors"

	"github.com/autotrend/entities"
)

func BarsClose(bars []entities.Bar, period int) ([]float64, error) {

	closes := make([]float64, 0)
	if len(bars) >= period {
		for _, bar := range bars {
			closes = append(closes, bar.Close)
		}

		return closes, nil
	}

	return closes, errors.New("Erro no BarsClose")

}

func BarsOpen(bars []entities.Bar, period int) ([]float64, error) {

	opens := make([]float64, 0)
	if len(bars) >= period {
		for _, bar := range bars {
			opens = append(opens, bar.Open)
		}

		return opens, nil
	}

	return opens, errors.New("Erro no BarsOpen")

}

func BarsLow(bars []entities.Bar, period int) ([]float64, error) {

	lows := make([]float64, 0)
	if len(bars) >= period {
		for _, bar := range bars {
			lows = append(lows, bar.Low)
		}

		return lows, nil
	}

	return lows, errors.New("Erro no BarsLow")

}

func BarsHigh(bars []entities.Bar, period int) ([]float64, error) {

	highs := make([]float64, 0)
	if len(bars) >= period {
		for _, bar := range bars {
			highs = append(highs, bar.High)
		}

		return highs, nil
	}

	return highs, errors.New("Erro no BarsHigh")

}
