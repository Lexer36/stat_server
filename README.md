Микросервис на golang для сбора статистики

- В качестве хранилища использован Postgres
- Миграции применяются посредством библиотеки https://github.com/golang-migrate/migrate/
> go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest