#!/bin/bash

usage()
{
        echo  "Usage: focus.sh [options]";
        echo  "###################################"
        echo  "  build     Build and start focus service"
        echo  "  start     Start focus service"
        echo  "  stop      Stop focus service"
        echo  "  restart   Restart focus service"
        echo  "  rebuild   stop,delete focus service,after then Build"
        echo  "  ps        Show all services status"
        echo  "  help      Show usage"
        echo  "  Note: If no option ,default to Help. "
        echo  ""
        echo  "  "
        echo  "Example: "
        echo  "  \"start.sh build\" "
        echo  "  \"start.sh start\" "
        echo  "  \"start.sh stop\" "
        echo  "###################################"
}

build()
{
  echo  "#########build start... "
  echo  " "
  docker-compose -f docker/docker-compose.yml up -d --build;
}

start()
{
  echo  "#########start start... "
  echo  " "
  docker-compose -f docker/docker-compose.yml start;
}

stop()
{
  echo  "#########stop start... "
  echo  " "
  docker-compose -f docker/docker-compose.yml stop;
}

restart()
{
  echo  "#########restart start... "
  echo  " "
  stop;
  start;
}

rebuild()
{
  echo  "#########rebuild start... "
  echo  " "
  stop;
  docker-compose -f docker/docker-compose.yml rm -f;
  rm -rf docker/mysql/*
  rm -rf logs/*
  build
}

ps()
{
  echo  "#########ps start... "
  echo  " "
  docker-compose -f docker/docker-compose.yml ps;
}

if [ $# -gt 0 ];
then
   case $1 in
   ps)
       ps
       ;;
   start)
       start
       ;;
   stop)
       stop
       ;;
   build)
       build
       ;;
   rebuild)
       rebuild
       ;;
   restart)
       restart
       ;;
   help)
       usage
       ;;
   *)
       usage
       ;;
   esac

   echo  " "
   echo  "#########end"
else
  usage;
fi