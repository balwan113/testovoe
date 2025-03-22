#!/bin/sh
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME"; do
  echo "ожидаем бд"
  sleep 1
done

echo "база готова"
/usr/local/bin/migrate -path /root/db/migrations -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up || {
  echo "ошибка при миграции"
  exit 1
}

echo "миграции успешно применены"
/root/main


