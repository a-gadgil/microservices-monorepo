from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="Order Management Service")

class Order(BaseModel):
    item_id: str
    quantity: int
    user_id: str

@app.get("/health")
def health_check():
    return {"status": "UP", "service": "order-service"}

@app.post("/api/orders")
def create_order(order: Order):
    return {
        "order_id": "ORD-998877",
        "status": "CREATED",
        "item_id": order.item_id,
        "quantity": order.quantity,
        "user_id": order.user_id,
        "message": "Order placed successfully!"
    }