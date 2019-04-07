# display processes belongs to current
# console and current user
ps
ps -A
# display all process with full format
ps -ef
ps -l
# process tag | process status | Ni: nice
F S   UID   PID  PPID  C PRI  NI ADDR SZ WCHAN  TTY          TIME CMD
0 S  1000  1008  1005  0  80   0 -  1553 wait   pts/0    00:00:02 bash
0 R  1000  2380  1008  0  80   0 -  1818 -      pts/0    00:00:00 ps

# trace parent and child process
ps --forest
# display process info in real-time manner
top

# kill process by process name
killall http*
# kill process by PID
kill 3048

# show all mounted device
mount
# mount usb stick, mount sdb1 to
#  disk
mount -t vfat /dev/sdb1 /media/disk

# check disk info, mounted dev info
df
# use MB to count memory
df -h

# sort by line on alphbtic order
sort data
# sort by order of number
sort -n nums > s_nums
# sort by time stamp
sort -M dates > s_dates
sort -M --reverse dates > s_dates
# split by ':' and use col 3 as sort key
sort -t ':' -k 3 -n /etc/passwd
# check whihc files of current dir take more space 
du -sh * | sort -nr

# search 'sort' in currend dir with line num show
grep -n sort ./*
# check how many matches
grep -c sort ./*

# zip all matched files
gzip dat*
# unzip all matched files
gzip -d dat*
# zip the folder src into src.tar
tar -cvf src.tar src/
# list all content in the tar
tar -tf src.tar
tar -xvf src.tar