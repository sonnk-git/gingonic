<h1 align="center">
  Flashcard Maker
</h1>

<h2 align="center">
  Inspiration from https://quizlet.com
</h2>

## Requirement

- Golang version: 1.18.9 ^
- Postgres DB is installed
- Frontend Source

  https://github.com/sonnk-git/flashcard_maker

## Quick start

1. **Config**
    + In root directory

         ```shell
         go mod tidy
         ```

    + Create .env file

         ```shell
         cp .env-example .env
         ```

    + Edit .env

      + Generating VAPID Keys: https://github.com/SherClockHolmes/webpush-go (We have VAPID_PUBLIC_KEY and
        VAPID_PRIVATE_KEY, you must put it into .env)

      + Sample .env file

         ```shell
         MODE=debug
         DB_HOST=localhost
         DB_USER=postgres
         DB_PASS=password
         DB_NAME=gingonic
         DB_PORT=5432
         TIMEZONE=Asia/Ho_Chi_Minh
         APP_SECRET_KEY=secret
         VAPID_PUBLIC_KEY=BLPvDei9pAtoZweEjYW7J5tKfJobcWzuj8mhxSWlRckIa6tW5lHeur7xZUGGh65AURT-F3cNoCgq_a-38EaoLVg
         VAPID_PRIVATE_KEY=cwPV5XDjvmQUQJbjjkslTUb7m1BraQDdH82j_Agogk6
         ```

2. **Start.**

   In root directory:
    + Start server:
         ```shell
         make dev-server
         ```
    + Create new terminal, start scheduler(cronjob):
         ```shell
         make scheduler
         ```
