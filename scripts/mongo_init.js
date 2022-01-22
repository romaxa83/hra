use orders

db.orders.stats()

db.orders.createIndex({ name: 1, description: 1 });
db.orders.createIndex({ '$**': 'text' });

db.orders.getIndexes();