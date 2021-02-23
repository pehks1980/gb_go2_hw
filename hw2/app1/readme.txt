
на PC ноутбуке goland в папке с go файлом:
GOOS=linux GOARCH=arm CC_FOR_TARGET=arm-linux-gnueabi-gcc go build .


Testing on my rasp pi4 (linux/ubuntu)

(env) user@ubuntu200:~/golang$ uname -a
Linux ubuntu200 5.4.0-1025-raspi #28-Ubuntu SMP PREEMPT Wed Dec 9 17:13:54 UTC 2020 armv7l armv7l armv7l GNU/Linux
(env) user@ubuntu200:~/golang$ 

(env) user@ubuntu200:~/golang$ ./app1 out.txt
Делитель не должен быть по идее = 0!  ОШИБКА делитель = 0!
error: divide by zero. time: 2021-02-23 01:28:05.791274865 +0000 UTC m=+0.001252810

22 bytes written
done
result= 0
(env) user@ubuntu200:~/golang$ ./app1
Делитель не должен быть по идее = 0!  ОШИБКА делитель = 0!
error: divide by zero. time: 2021-02-23 01:28:20.917250428 +0000 UTC m=+0.000887760

2021/02/23 01:28:20 Ошибка файла open /home/user/golang/hw1/app1/out.txt: no such file or directory
result= 0

(env) user@ubuntu200:~/golang$ less -S out.txt
Делитель=3.5000
Делимое=0.0000
Частное=0.0000
out.txt (END)

