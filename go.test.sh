#!/usr/bin/env bash

set -e
echo "mode: set" > cover.out
echo "" > coverage.txt

pkg_array=("usecase" "repository" "controller")

test_flag=false

for d in $(go list ./... | grep -v vendor); do
    test_flag=false
    for pkg in ${pkg_array[@]}; do
        if [[ $d =~ $pkg ]]; then
            test_flag=true
        fi
    done
    if [[ $test_flag == true ]]; then
        go test -cover -coverprofile=profile.out "$d"
        if [ -f profile.out ]; then
            # カバレッジ
            if [[ $(cat profile.out|grep -v 'mode') != "" ]]; then
                cat profile.out|grep -v 'mode' >>cover.out

                go tool cover -func=profile.out |
                tail -n1                        |
                awk -F ' ' '{print "'$d'",$3}'  >>coverage.txt
            else
                echo "$d 0%"                    >>coverage.txt
            fi
            rm profile.out
        fi
    fi
done

go tool cover -func=cover.out                   |
tail -n1                                        |
awk -F ' ' '{print "total",$3}'                 >>coverage.txt