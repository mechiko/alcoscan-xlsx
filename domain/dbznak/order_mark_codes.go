package dbznak

import "database/sql"

type OrderMarkCodes struct {
	Id                  int64          `db:"id,omitempty"`
	CreateDate          string         `db:"create_date"`
	Gtin                string         `db:"gtin"`
	Quantity            int            `db:"quantity"`
	SerialNumberType    string         `db:"serial_number_type"`
	TemplateId          int            `db:"template_id"`
	CisType             string         `db:"cis_type"`
	ContactPerson       string         `db:"contact_person"`
	ReleaseMethodType   string         `db:"release_method_type"`
	CreateMethodType    string         `db:"create_method_type"`
	ProductionOrderId   string         `db:"production_order_id"`
	ProductName         string         `db:"product_name"`
	ProductCapacity     string         `db:"product_capacity"`
	ProductShelfLife    string         `db:"product_shelf_life"`
	ProductTemplate     string         `db:"product_template"`
	Comment             string         `db:"comment"`
	Version             string         `db:"version"`
	State               string         `db:"state"`
	Status              string         `db:"status"`
	OrderId             string         `db:"order_id"`
	Archive             int            `db:"archive"`
	Json                sql.NullString `db:"json"`
	PaymentType         string         `db:"payment_type"`
	ServiceProviderId   string         `db:"service_provider_id"`
	ServiceProviderName string         `db:"service_provider_name"`
	ServiceProviderRole string         `db:"service_provider_role"`
}
