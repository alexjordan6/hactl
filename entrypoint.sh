#!/bin/bash
# Start HAProxy in the background
haproxy -f /usr/local/etc/haproxy/haproxy.cfg &
# Keep the container running
tail -f /dev/null
