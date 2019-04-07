ls -lF /bin/bash /bin/tcsh /bin/dash /bin/csh /bin/sh
# run a set of cmd in a sub-shell
(pwd ; ls ; pwd; ls; echo $BASH_SUBSHELL)
# run in current shell
{pwd ; ls ; pwd; ls; echo $BASH_SUBSHELL}
(pwd ; (echo $BASH_SUBSHELL))

# run cmd in back ground
sleep 3000&
ps -f
# check back ground jobs
jobs
(sleep 2 ; echo $BASH_SUBSHELL ; sleep 2)&

# run cmd as co-process
coproc sleep 10
coproc my_job { sleep 10; }

# check if a cmd is of built-in
type -a cd ps pwd
# check recent cmds
history