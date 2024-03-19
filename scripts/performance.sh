echo $(ps --no-headers -p $1 -o %cpu,%mem; ps --no-headers --ppid $1 -o %cpu,%mem)

