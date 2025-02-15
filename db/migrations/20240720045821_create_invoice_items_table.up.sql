CREATE TABLE invoice_items (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  invoice_id BIGINT NOT NULL,
  item_id BIGINT NOT NULL,
  item_name VARCHAR(150),
  quantity INTEGER NOT NULL DEFAULT 0,
  unit_price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE invoice_items
ADD CONSTRAINT fk_invoice
FOREIGN KEY (invoice_id) REFERENCES invoices(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

CREATE INDEX idx_invoice_id ON invoice_items(invoice_id);

ALTER TABLE invoice_items
ADD CONSTRAINT fk_item
FOREIGN KEY (item_id) REFERENCES items(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

CREATE INDEX idx_item_id ON invoice_items(item_id);
