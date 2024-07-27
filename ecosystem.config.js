module.exports = {
    apps: [
        {
            name: "hotel-reservation",
            script: "nodemon",
            args: [
                "--watch", "*.go",
                "--exec", "sh -c 'lsof -ti tcp:3000 | xargs kill -9 || true && make run'",
                "--ext", "go",
                "--restartable", "rs"
            ],
            interpreter: "/opt/homebrew/bin/bash", // Specify the shell interpreter
            watch: false,
            restart_delay: 1000,
            env: {
                NODE_ENV: "development"
            },
            exec_mode: "fork",
            kill_timeout: 3000,
            wait_ready: true,
            listen_timeout: 3000,
            shutdown_with_message: true,
            stop_signal: "SIGINT",
        }
    ]
};