<h1 align="center">
  Flashcard Maker
</h1>

<h2 align="center">
  Inspiration from https://quizlet.com
</h2>

## 🚀 Quick start

1.  **Config.**
     + Create .env file 
          ```shell
          cp .env-example .env
          ```
       Edit .env
    
     + Generating VAPID Keys: https://github.com/SherClockHolmes/webpush-go
    
2.  **Start.**

    In root directory: 
     + Start server: 
         ```shell
          make dev-server
          ```
     + Create new terminal, start scheduler(cronjob):
          ```shell
          cd scheduler
          go run main.go
          ```
3.  **Frontend Source.**

    https://github.com/khacsonit/flashcard_maker
