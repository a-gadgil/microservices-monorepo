from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="User Service", version="1.0.0")

class User(BaseModel):
    name: str
    email: str
    role: str  # rider or driver

@app.get("/health")
def health_check():
    return {"status": "UP", "service": "user-service"}

@app.post("/users")
def register_user(user: User):
    return {
        "message": f"User {user.name} registered successfully as {user.role}",
        "data": user
    }