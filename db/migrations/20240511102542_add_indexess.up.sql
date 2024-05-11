CREATE INDEX idx_users_id ON users(id);
CREATE INDEX idx_users_phone_number ON users(phone_number);
CREATE INDEX idx_products_id ON products(id);
CREATE INDEX idx_customer_id ON customers(id);
CREATE INDEX idx_customer_phone_number ON customers(phone_number);
CREATE INDEX idx_orders_id ON orders(id);
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_order_details_order_id ON order_details(order_id);
CREATE INDEX idx_order_details_product_id ON order_details(product_id);