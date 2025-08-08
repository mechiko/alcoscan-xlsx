package dbznak

type OrderMarkCodesSerialNumbers struct {
	Id               int64  `db:"id,omitempty"`
	IdOrderMarkCodes int64  `db:"id_order_mark_codes"`
	Gtin             string `db:"gtin"`
	SerialNumber     string `db:"serial_number"`
	Code             string `db:"code"`
	BlockId          string `db:"block_id"`
	Status           string `db:"status"`
}
