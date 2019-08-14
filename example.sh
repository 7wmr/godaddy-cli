#!/bin/bash

# This script is used to check and update your GoDaddy DNS server to the IP address of your current internet connection.
# Special thanks to mfox for his ps script
# https://github.com/markafox/GoDaddy_Powershell_DDNS
#
# First go to GoDaddy developer site to create a developer account and get your key and secret
#
# https://developer.godaddy.com/getstarted
# Be aware that there are 2 types of key and secret - one for the test server and one for the production server
# Get a key and secret for the production server
# 
#Update the first 4 variables with your information
 
domain="8lr.uk"   # your domain
name="dev"     # name of A record to update
key="e42XDsA1Rngo_CWdAjp6PwePfv3CgEq3D6R"      # key for godaddy developer API
secret="HRiA7fAEHnmukiFfS2KSYv"   # secret for godaddy developer API

headers="Authorization: sso-key $key:$secret"

 echo $headers

result=$(curl -s -X GET -H "$headers" \
 "https://api.godaddy.com/v1/domains/$domain/records/A/$name")

dnsIp=$(echo $result | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b")
 echo "dnsIp:" $dnsIp

# Get public ip address there are several websites that can do this.
ret=$(curl -s GET "http://ipinfo.io/json")
currentIp=$(echo $ret | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b")
 echo "currentIp:" $currentIp

if [ $dnsIp != $currentIp ]; then
	echo "Ips are not equal"
	request='[{"data":"'$currentIp'","ttl":600}]'
	echo $request
	nresult=$(curl -i -s -X PUT \
 -H "$headers" \
 -H "Content-Type: application/json" \
 -d $request "https://api.godaddy.com/v1/domains/$domain/records/A/$name")
	echo $nresult
fi
