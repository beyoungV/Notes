1,程序连接数据库会有连接泄漏的情况，需要及时释放连接

2,Go sql包中的Query和QueryRow两个方法的连接不会自动释放连接，只有在遍历完结果或者调用close方法才会关闭连接

3,Go sql中的Ping和Exec方法在调用结束以后就会自动释放连接

4,忽略了函数的某个返回值不代表这个值就不存在了，如果该返回值需要close才会释放资源，直接忽略了就会导致资源的泄漏。

5,有close方法的变量，在使用后要及时调用该方法，释放资源
