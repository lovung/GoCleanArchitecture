#!/bin/bash

RED=`tput setaf 1`
GRN=`tput setaf 2`
PUR=`tput setaf 13`
RESET=`tput sgr0`
echo "${GRN}Listing all packages${RESET}"
PKG_LIST+="$(go list ./... | grep -v /vendor/ | grep -v /mock | grep -v migrations) "
for i in $PKG_LIST
do
    echo $i
done
echo "----------------"
EXIT_CODE=0
# echo "${GRN}Golint...${RESET}"
# go list ./... | grep -v /vendor/ | grep -v migrations | xargs -L1 golint -set_exit_status -v
# EXIT_CODE+=$?
echo "${GRN}Golangci-lint...${RESET}"
golangci-lint run --timeout 5m0s -v ./...
EXIT_CODE+=$?
if [ $EXIT_CODE != 0 ]
then
    echo "${GRN}Success${RESET}"
else
    echo "${RED}Failed${RESET}"
fi
exit $EXIT_CODE

