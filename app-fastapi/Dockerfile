FROM python:3.9
ENV PIP_DISABLE_PIP_VERSION_CHECK=on
RUN pip install poetry

WORKDIR /app

COPY poetry.lock pyproject.toml main.py /app/
RUN poetry config virtualenvs.create false
RUN poetry install --no-interaction

EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--workers", "8"]
