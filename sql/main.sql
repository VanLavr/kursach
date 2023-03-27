CREATE DATABASE crud1;

USE crud1;

DROP TABLE products;
SELECT * FROM products;
CREATE TABLE products (
	id INT PRIMARY KEY,
    NameOfProduct VARCHAR(100),
    Category VARCHAR(100),
    Price FLOAT
);

INSERT INTO products(id, NameOfProduct, Category, Price)
VALUES
	(1, "Tshirt", "Clothes", 12.5),
    (2, "Ring", "Fingeys", 3.99),
    (3, "Tshirt", "Clothes", 10.99),
    (4, "Pants", "Clothes", 5.11),
    (5, "Trinket", "Fingeys", 0.35);
