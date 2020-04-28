ssh -nNT -L $(pwd)/docker.sock:/var/run/docker.sock root@51.15.244.251
rm $(pwd)/docker.sock
