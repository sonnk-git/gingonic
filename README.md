<h1 align="center">
  Flashcard Maker
</h1>

<h2 align="center">
  Inspiration from https://quizlet.com
</h2>

## 🚀 Quick start

1.  **Config.**

    Edit .env
    
    Generating VAPID Keys: https://github.com/SherClockHolmes/webpush-go
    
2.  **Start.**

    In root directory: 
     + Start server: 
         ```shell
          make dev-server
          ```
     + Start scheduler:
          ```shell
          cd scheduler
          go run main.go
          ```
