#!/bin/sh

APP_NAME=DataReportUpdate
APP_PATH=/home/centos/data-report-update
APP_PID=$APP_NAME\.pid
ENV_NAME=test
LOG_FILE=logs/data-report.log
SLEEP_SECONDS=3

is_exist() {
  pid=$(ps -ef | grep $APP_NAME | grep -v grep | awk '{print $2}')

  if [ -z "${pid}" ]; then
    return 0
  else
    return 1
  fi
}

start() {
  is_exist
  if [ $? -eq "1" ]; then
    echo "$APP_NAME is already running pid is ${pid}"
  else
    nohup $APP_PATH/$APP_NAME -env=$ENV_NAME >>$APP_PATH/nohup.out 2>&1 &
    echo $! >$APP_PATH/$APP_PID
    echo "start $APP_NAME successfully, pid is $! "
    tail -20 $APP_PATH/$LOG_FILE
  fi
}

stop() {
  # is_exist
  pidf=$(cat $APP_PATH/$APP_PID)
  # echo "$pidf"
  echo "pid = $pidf begin kill $pidf"
  kill $pidf
  rm -rf $APP_PATH/$APP_PID
  sleep $SLEEP_SECONDS
  # 判断服务进程是否存在
  is_exist
  if [ $? -eq "1" ]; then
    echo "pid = $pid begin kill -9 $pid"
    kill -9 $pid
    sleep $SLEEP_SECONDS
    echo "$APP_NAME process stopped！"
  else
    echo "$APP_NAME is not running！"
  fi
}

status() {
  is_exist
  if [ $? -eq "1" ]; then
    echo "$APP_NAME is running，pid is ${pid}"
  else
    echo "$APP_NAME is not running！"
  fi
}

restart() {
  stop
  start
}

usage() {
  echo "Usage: sh run.sh [ start | stop | restart | status ]"
  exit 1
}

case "$1" in
'start')
  start
  ;;
'stop')
  stop
  ;;
'restart')
  restart
  ;;
'status')
  status
  ;;
*)
  usage
  ;;
esac
exit 0
