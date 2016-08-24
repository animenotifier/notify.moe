#!/bin/sh

# Firewall
sugoi clearFirewall
#sugoi redirectPort 80 5000
#sugoi redirectPort 443 5001
sugoi blockPort 3000

# Use swap only when needed
sudo sysctl vm.swappiness=5

# TCP settings
sudo sysctl net.ipv4.ip_local_port_range="15000 61000"
sudo sysctl net.ipv4.tcp_fin_timeout=15
sudo sysctl net.ipv4.tcp_tw_recycle=1
sudo sysctl net.ipv4.tcp_tw_reuse=1
sudo sysctl net.ipv4.tcp_max_syn_backlog=40000
sudo sysctl net.ipv4.tcp_sack=1
sudo sysctl net.ipv4.tcp_window_scaling=1
sudo sysctl net.ipv4.tcp_keepalive_intvl=30
sudo sysctl net.ipv4.tcp_moderate_rcvbuf=1

# net.core
sudo sysctl net.core.somaxconn=1024
#sudo sysctl net.core.wmem_default=8388608
#sudo sysctl net.core.rmem_default=8388608

# Max files
ulimit -n 2048
