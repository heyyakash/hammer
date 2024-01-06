# Hammer - A Simple HTTP Load Testing Tool 🛠️🚀

## Introduction
Hammer is a simple HTTP load testing tool designed to help you test the performance of your web applications. With a few configurations, you can unleash the power of goroutines and measure the response time of your server under various loads. Let's dive into how you can wield the mighty Hammer! ⚒️💻

## Configuration Options 🛠️
- **Target URL (-url):** The URL you want to benchmark (default: http://localhost:8000) 🌐
- **Total Requests (-r):** Total number of requests to perform (default: 100) 📈
- **Concurrency (-c):** Total number of goroutines to run concurrently (default: 1) 🏃
- **Timeout (-t):** Request timeout in seconds (default: 20) ⏰

## Getting Started 🔨
1. Clone the Hammer repository.
2. Run `go build` to build the executable.
3. Run the executable with your desired configuration.

```bash
./hammer -url=http://your-target-url.com -r=500 -c=10 -t=30
```
# Explanation of the Code 💬

1. Command Line Flags: Accepts various configurations via command line flags.
2. Goroutines: Spawns goroutines to perform concurrent requests.
3. Response Time Channel: Uses a channel to collect response times from goroutines.
4. Calculate Results: After all goroutines finish, it calculates and displays the results.

# Future Plans
1. Add load-testing for POST, UPDATE routes as well (Currently only supporting GET request)

Now you're ready to swing the Hammer and unleash the power of load testing! 💪🚀
