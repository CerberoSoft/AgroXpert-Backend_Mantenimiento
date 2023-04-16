package models

type HistoricHarvest struct {
	Harvest         Harvest         `json:"harvest"`
	Estimates       []EstimateModel `json:"estimates"`
	FinalProduction FinalProduction `json:"finalProduction"`
}
