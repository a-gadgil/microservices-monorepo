CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

INSERT INTO items (name)
SELECT * FROM (VALUES
    ('Containerized Python Slim'),
    ('Gunicorn + Uvicorn Non-Root'),
    ('Docker Internal Bridge Network')
) AS v(name)
WHERE NOT EXISTS (SELECT 1 FROM items);
