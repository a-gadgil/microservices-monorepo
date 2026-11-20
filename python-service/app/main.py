import json
import os
from contextlib import asynccontextmanager

import psycopg
import redis
from fastapi import FastAPI, HTTPException, Response

DB_DSN = os.getenv("DB_DSN", "postgresql://appuser:apppass@127.0.0.1:5432/appdb")
REDIS_ADDR = os.getenv("REDIS_ADDR", "127.0.0.1:6379")
CACHE_TTL = int(os.getenv("CACHE_TTL", "30"))

redis_host, redis_port = REDIS_ADDR.split(":")
rdb = redis.Redis(host=redis_host, port=int(redis_port), decode_responses=True, socket_timeout=2)


@asynccontextmanager
async def lifespan(app: FastAPI):
    yield
    rdb.close()


app = FastAPI(title="3-Tier Python Service", lifespan=lifespan)


@app.get("/healthz")
def healthz():
    return {"status": "alive"}


@app.get("/readyz")
def readyz():
    try:
        rdb.ping()
        with psycopg.connect(DB_DSN, connect_timeout=2) as conn:
            with conn.cursor() as cur:
                cur.execute("SELECT 1;")
    except Exception as exc:
        raise HTTPException(status_code=503, detail=f"not ready: {exc}")
    return {"status": "ready"}


@app.get("/items")
def list_items(response: Response):
    cache_key = "items:all"
    try:
        cached = rdb.get(cache_key)
        if cached:
            response.headers["X-Cache"] = "HIT"
            return json.loads(cached)
    except redis.RedisError:
        pass

    with psycopg.connect(DB_DSN, connect_timeout=3) as conn:
        with conn.cursor() as cur:
            cur.execute("SELECT id, name FROM items ORDER BY id;")
            rows = [{"id": r[0], "name": r[1]} for r in cur.fetchall()]

    try:
        rdb.setex(cache_key, CACHE_TTL, json.dumps(rows))
    except redis.RedisError:
        pass

    response.headers["X-Cache"] = "MISS"
    return rows
