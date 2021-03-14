package model

import "sync"

type LocalStock struct {
	Stock     int
	SaleCount int
	mux       sync.Mutex
}

func NewLocalStock(stock, SaleCount int) LocalStock {
	return LocalStock{
		Stock:     stock,
		SaleCount: SaleCount,
		mux:       sync.Mutex{},
	}
}

// LocalDeductionStock 本地扣库存
func (spike *LocalStock) DeductionStock() (isOk bool) {
	spike.mux.Lock()
	defer spike.mux.Unlock()
	if spike.SaleCount >= spike.Stock {
		return
	}
	spike.SaleCount++
	isOk = true
	return
}

func (spike *LocalStock) IncrStock() {
	spike.mux.Lock()
	defer spike.mux.Unlock()
	spike.SaleCount++
	return
}

