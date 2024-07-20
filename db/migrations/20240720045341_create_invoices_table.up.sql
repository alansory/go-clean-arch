CREATE TABLE invoices (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  invoice_number VARCHAR(255) NOT NULL,
  invoice_subject VARCHAR(150),
  issue_date DATE NOT NULL,
  due_date DATE NOT NULL,
  status ENUM('paid', 'unpaid') NOT NULL DEFAULT 'unpaid',
  customer_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE invoices
ADD CONSTRAINT fk_customer
FOREIGN KEY (customer_id) REFERENCES customers(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

CREATE INDEX idx_customer_id ON invoices(customer_id);