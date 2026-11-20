# Stage 1: Build isolated virtual environment
FROM python:3.13-slim AS builder

WORKDIR /build
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

COPY requirements.txt .
RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

# Stage 2: Minimal non-root runtime image
FROM python:3.13-slim

RUN groupadd --system --gid 10001 pyapp && \
    useradd --system --uid 10001 --gid pyapp --no-create-home pyapp

WORKDIR /app
COPY --from=builder /opt/venv /opt/venv
COPY app/ ./app/

ENV PATH="/opt/venv/bin:$PATH" \
    PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    APP_HOST=0.0.0.0 \
    APP_PORT=8000 \
    WEB_CONCURRENCY=2

USER pyapp:pyapp
EXPOSE 8000

CMD ["sh", "-c", "exec gunicorn app.main:app --worker-class uvicorn.workers.UvicornWorker --workers ${WEB_CONCURRENCY} --bind ${APP_HOST}:${APP_PORT} --access-logfile - --error-logfile -"]
