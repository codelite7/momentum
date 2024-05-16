FROM python:3.12-bookworm

RUN pip install poetry

COPY . .

RUN poetry install

ENTRYPOINT ["poetry", "run", "python", "-m", "api_langchain.main"]