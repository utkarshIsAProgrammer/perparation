USE MACHINES;
SELECT DATABASE()

CREATE TABLE Grocery (
    items VARCHAR(200) PRIMARY KEY,
    quantity INT
)

INSERT INTO Grocery(items, quantity)
VALUES ("Biscuit", 2),("Kurkure", 4),("Banana", 12), ("Drinks", 1)

SELECT * FROM Grocery;