CREATE INDEX "omcsn_code" ON "order_mark_codes_serial_numbers" (
	"code"	ASC
);
CREATE INDEX "omac_code" ON "order_mark_aggregation_codes" (
	"code"	ASC
);
CREATE INDEX "oma_" ON "order_mark_aggregation" (
	"unit_serial_number"	ASC
);