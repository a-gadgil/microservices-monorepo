const express = require('express');
const axios = require('axios');
const app = express();
app.use(express.json());

const USER_SERVICE_URL = process.env.USER_SERVICE_URL || 'http://user-service-svc:8000';
const DISPATCH_SERVICE_URL = process.env.DISPATCH_SERVICE_URL || 'http://dispatch-service-svc:8080';
const BILLING_SERVICE_URL = process.env.BILLING_SERVICE_URL || 'http://billing-service-svc:8080';

app.get('/health', (req, res) => {
    res.json({ status: 'UP', service: 'api-gateway' });
});

// Proxy to User Service
app.all('/users*', async (req, res) => {
    try {
        const response = await axios({
            method: req.method,
            url: `${USER_SERVICE_URL}${req.originalUrl}`,
            data: req.body
        });
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({ error: 'Error proxying to User Service' });
    }
});

// Proxy to Dispatch Service
app.all('/dispatch*', async (req, res) => {
    try {
        const response = await axios({
            method: req.method,
            url: `${DISPATCH_SERVICE_URL}${req.originalUrl}`,
            data: req.body
        });
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({ error: 'Error proxying to Dispatch Service' });
    }
});

// Proxy to Billing Service
app.all('/billing*', async (req, res) => {
    try {
        const response = await axios({
            method: req.method,
            url: `${BILLING_SERVICE_URL}${req.originalUrl}`,
            data: req.body
        });
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({ error: 'Error proxying to Billing Service' });
    }
});

app.listen(8080, '0.0.0.0', () => {
    console.log('API Gateway running on port 8080');
});