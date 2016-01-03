#!/bin/sh
sugoi clearFirewall
sugoi redirectPort 80 5000
sugoi redirectPort 443 5001
sugoi blockPort 3000