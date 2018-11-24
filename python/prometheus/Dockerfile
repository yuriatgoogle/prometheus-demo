FROM python:2

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD [ "uwsgi", "--http", "127.0.0.1:8000", "--wsgi-file", "prometheus-flask-dispatcher.py", "--callable", "app_dispatch" ]