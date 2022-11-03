#!/bin/bash  
  
list="rootfs usr data data2"  
for i in $list;  
do  
echo $i is appoint ;  
done

for i in f1 f2 f3 ;
do
echo $i is appoint ;
done


for i in $* ;
do
echo $i is input chart\! ;
done


for i in `ls`;
do
echo $i is file name\! ;
done

arr=( 'rootfs' 'usr' 'data' )

for item in ${arr[@]}
do
    echo $item in arr\! ;
done
