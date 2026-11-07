const express = require('express');
const app = express();
const PORT = process.env.PORT || 8080;

app.use(express.json());

// Root and health endpoints
app.get('/', (req, res) => {
  jsonResponse = {
    message: "Welcome to Quick-Commerce API Gateway",
    services: {
      orders: "/api/orders",
      inventory: "/api/inventory",
      dispatch: "/api/dispatch"
    }
  };
  res.json(jsonResponse);
});

app.get('/health', (req, res) => {
  res.status(200).json({ status: "UP", service: "api-gateway" });
});

// Stub routing placeholders for downstream microservices
app.all('/api/*', (req, res) => {
  res.json({
    routedPath: req.path,
    message: "Gateway proxy route placeholder — downstream service integration pending."
  });
});

app.listen(PORT, () => {
  console.log(`API Gateway running on port ${PORT}`);
});