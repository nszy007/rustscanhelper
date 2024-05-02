# rustscanhelper
将rustscan的-g方式输出结果以IP:PORT每行方式输出  
## 使用方法1
对于扫描目标不多的情况，可以  
rustscan -a xxx.xxx.xxx.xxx -r 1-65535 -g | ./rustscanhelper  
```
rustscan -a 192.168.31.1 -g | ./rustscanhelper
192.168.31.1:53
192.168.31.1:80
192.168.31.1:443
192.168.31.1:784
192.168.31.1:8098
192.168.31.1:8080
192.168.31.1:8999
192.168.31.1:16971
192.168.31.1:54171
```
## 使用方法2
对于扫描目标较多的情况，建议rustscan先用>将结果输出到指定文件，再用rustscanhelper以-i参数读取方式使用  
```
(base) ➜  rustscanhelper rustscan -a 192.168.31.193 -g > target.txt                    
(base) ➜  rustscanhelper ./rustscanhelper -i target.txt     
192.168.31.193:5000
192.168.31.193:7000
192.168.31.193:49156
```
同时可以使用-o参数将结果输出到指定文件
```
./rustscanhelper -h           
Usage of ./rustscanhelper:
  -i string
        解析rustscan -g输出结果的文件路径
  -o string
        输出结果文件路径
```
