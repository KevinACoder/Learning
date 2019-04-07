# go to home dir
cd ~
# use color to diff dir and file
ls -F
# show details as well as hide file
ls -a
# list files of sub dir
ls -R
ls -FR
# show details
ls -l
# (file globbing) 
# ? match just one character, * match one or more characters
ls -l man?s*
ls -l do[cm]ker
ls -l do[c-m]ker
ls -l do[!m]ker #exclude of
ls -l --time=atime do[!m]ker

cp src/data dst
cp src/data dst/data
ls -R
# ask user wheather to overwrite before cp
cp -i src/data dst/data
# cp the whole folder
cp -R src/ dst2
ls -lR

# build symbolic file link
#  contents do not need to be the same
ln -s src/data sl_data
ls -l *data
# the two files share the same inode
# sl is just a ptr to src file
ls -i sl_data src/data

# hard link is a independent virtual file
# with src file info
# hl is a different file with diff inode

ln src/data hl_data
ls -li hl_data sl_data src/data

# move do not change inode and time stamp
mv dst/data dst/data2
rm -i sl_data

# create parent dir if not exist
mkdir -p l1/l2/l3
touch l1/l2/l3/data
rmdir l1/l2/l3
rm -ri l1/l2/l3

# check type of file
file data
file /bin/ls
# add line number for showing content
cat -n data
# line number only for in-use lines
cat -b data
# show content by pages
more file_navi.sh
less file_navi.sh
# show last ten lines
tail file_navi.sh
tail -n 2 file_navi.sh
# enable showing file to update in real time
tail -n 2 -f file_navi.sh
head -n 9 file_navi.sh