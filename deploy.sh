#!/bin/bash
pack
go build
scp notify.moe eduard@arn:~/beta/notify.moe.new
ssh eduard@arn 'cd beta; killall notify.moe; rm notify.moe; mv notify.moe.new notify.moe; ./notify.moe &'