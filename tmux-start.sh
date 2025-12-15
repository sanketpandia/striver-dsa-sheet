#!/bin/bash

# Create a new tmux session named 'dev' with the first window (this will be window 0)
tmux new-session -d -s dev

# Create window 1 for nvim
tmux new-window -t dev:1 -n editor
tmux send-keys -t dev:1 'nvim ./' C-m

# Create window 2 for terminal
tmux new-window -t dev:2 -n terminal

# Kill window 0
tmux kill-window -t dev:0

# Attach to the session, starting at window 1
tmux attach-session -t dev
