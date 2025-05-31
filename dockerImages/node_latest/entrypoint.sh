#!/usr/bin/env sh

iptables -I INPUT 1 -m conntrack --ctstate ESTABLISHED,RELATED -j DROP

#RUN iptables -I INPUT 1 -i lo -j ACCEPT
#
#RUN iptables -A OUTPUT -p udp --dport 53 -j ACCEPT
#
#RUN iptables -A OUTPUT -p tcp -d npmjs.org --dport 80 -j ACCEPT
#RUN iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
#
#RUN iptables -P INPUT DROP
#RUN iptables -P OUTPUT DROP